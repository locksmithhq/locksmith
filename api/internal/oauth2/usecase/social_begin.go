package usecase

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/domain"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/types/input"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/types/output"
)

type socialBeginUseCase struct {
	getClientByClientIDRepository    contract.GetClientByClientIDRepository
	getSocialProviderByClientRepository contract.GetSocialProviderByClientRepository
	createSocialStateRepository      contract.CreateSocialStateRepository
}

func (u *socialBeginUseCase) Execute(ctx context.Context, in input.SocialBegin) (output.SocialBegin, error) {
	client, err := u.getClientByClientIDRepository.Execute(ctx, in.ClientID)
	if err != nil {
		if errors.Is(err, stackerror.ErrNotFound) {
			return output.SocialBegin{}, stackerror.NewUseCaseError(
				"SocialBeginUseCase",
				err,
				stackerror.WithMessage("client id is not valid"),
				stackerror.WithStatusCode(http.StatusBadRequest),
			)
		}
		return output.SocialBegin{}, stackerror.NewUseCaseError("SocialBeginUseCase", err)
	}

	socialProvider, err := u.getSocialProviderByClientRepository.Execute(ctx, client.ID, in.Provider)
	if err != nil {
		if errors.Is(err, stackerror.ErrNotFound) {
			return output.SocialBegin{}, stackerror.NewUseCaseError(
				"SocialBeginUseCase",
				err,
				stackerror.WithMessage("social provider not configured"),
				stackerror.WithStatusCode(http.StatusForbidden),
			)
		}
		return output.SocialBegin{}, stackerror.NewUseCaseError("SocialBeginUseCase", err)
	}

	if !socialProvider.Enabled {
		return output.SocialBegin{}, stackerror.NewUseCaseError(
			"SocialBeginUseCase",
			errors.New("social provider is disabled"),
			stackerror.WithMessage("social provider is disabled"),
			stackerror.WithStatusCode(http.StatusForbidden),
		)
	}

	if client.RequirePKCE && in.CodeChallenge == "" {
		return output.SocialBegin{}, stackerror.NewUseCaseError(
			"SocialBeginUseCase",
			fmt.Errorf("pkce required for this client"),
			stackerror.WithMessage("code_challenge is required for this client"),
			stackerror.WithStatusCode(http.StatusBadRequest),
		)
	}

	if in.CodeChallenge != "" && in.CodeChallengeMethod != "S256" {
		return output.SocialBegin{}, stackerror.NewUseCaseError(
			"SocialBeginUseCase",
			fmt.Errorf("unsupported code_challenge_method: %q", in.CodeChallengeMethod),
			stackerror.WithMessage("only S256 is accepted as code_challenge_method"),
			stackerror.WithStatusCode(http.StatusBadRequest),
		)
	}

	baseURL := os.Getenv("LOCKSMITH_BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:4000"
	}

	callbackURL := fmt.Sprintf("%s/api/oauth2/social/%s/callback", baseURL, in.Provider)
	scopes := strings.Split(socialProvider.Scopes, " ")

	gothProvider, err := newGothProvider(in.Provider, socialProvider.ClientKey, socialProvider.ClientSecret, callbackURL, scopes...)
	if err != nil {
		return output.SocialBegin{}, err
	}

	nonceBytes := make([]byte, 32)
	if _, err = rand.Read(nonceBytes); err != nil {
		return output.SocialBegin{}, stackerror.NewUseCaseError(
			"SocialBeginUseCase",
			err,
			stackerror.WithStatusCode(http.StatusInternalServerError),
		)
	}
	nonce := base64.RawURLEncoding.EncodeToString(nonceBytes)

	sess, err := gothProvider.BeginAuth(nonce)
	if err != nil {
		return output.SocialBegin{}, stackerror.NewUseCaseError("SocialBeginUseCase", err)
	}

	authURL, err := sess.GetAuthURL()
	if err != nil {
		return output.SocialBegin{}, stackerror.NewUseCaseError("SocialBeginUseCase", err)
	}

	socialState := domain.SocialState{
		Nonce:               nonce,
		ClientID:            in.ClientID,
		RedirectURI:         in.RedirectURI,
		State:               in.State,
		CodeChallenge:       in.CodeChallenge,
		CodeChallengeMethod: in.CodeChallengeMethod,
		ExpiresAt:           time.Now().Add(10 * time.Minute).Format(time.RFC3339),
	}

	if err = u.createSocialStateRepository.Execute(ctx, socialState); err != nil {
		return output.SocialBegin{}, stackerror.NewUseCaseError("SocialBeginUseCase", err)
	}

	return output.SocialBegin{AuthURL: authURL}, nil
}

func NewSocialBeginUseCase(
	getClientByClientIDRepository contract.GetClientByClientIDRepository,
	getSocialProviderByClientRepository contract.GetSocialProviderByClientRepository,
	createSocialStateRepository contract.CreateSocialStateRepository,
) contract.SocialBeginUseCase {
	return &socialBeginUseCase{
		getClientByClientIDRepository:       getClientByClientIDRepository,
		getSocialProviderByClientRepository: getSocialProviderByClientRepository,
		createSocialStateRepository:         createSocialStateRepository,
	}
}
