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
)

type socialCallbackUseCase struct {
	getSocialStateByNonceRepository      contract.GetSocialStateByNonceRepository
	deleteSocialStateRepository          contract.DeleteSocialStateRepository
	getClientByClientIDRepository        contract.GetClientByClientIDRepository
	getSocialProviderByClientRepository  contract.GetSocialProviderByClientRepository
	getAccountBySocialProviderRepository contract.GetAccountBySocialProviderRepository
	getAccountByEmailAndProjectRepository contract.GetAccountByEmailAndProjectRepository
	createAccountSocialProviderRepository contract.CreateAccountSocialProviderRepository
	createAccountRepository              contract.CreateAccountRepository
	createAuthCodeRepository             contract.CreateAuthCodeRepository
	getSignupByClientIDRepository        contract.GetSignupByClientIDRepository
}

func (u *socialCallbackUseCase) Execute(ctx context.Context, in input.SocialCallback, w http.ResponseWriter, r *http.Request) error {
	socialState, err := u.getSocialStateByNonceRepository.Execute(ctx, in.Nonce)
	if err != nil {
		if errors.Is(err, stackerror.ErrNotFound) {
			return stackerror.NewUseCaseError(
				"SocialCallbackUseCase",
				err,
				stackerror.WithMessage("invalid or expired social state"),
				stackerror.WithStatusCode(http.StatusBadRequest),
			)
		}
		return stackerror.NewUseCaseError("SocialCallbackUseCase", err)
	}

	client, err := u.getClientByClientIDRepository.Execute(ctx, socialState.ClientID)
	if err != nil {
		return stackerror.NewUseCaseError("SocialCallbackUseCase", err)
	}

	socialProvider, err := u.getSocialProviderByClientRepository.Execute(ctx, client.ID, in.Provider)
	if err != nil {
		return stackerror.NewUseCaseError("SocialCallbackUseCase", err)
	}

	baseURL := os.Getenv("LOCKSMITH_BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:4000"
	}

	callbackURL := fmt.Sprintf("%s/api/oauth2/social/%s/callback", baseURL, in.Provider)
	scopes := strings.Split(socialProvider.Scopes, " ")

	gothProvider, err := newGothProvider(in.Provider, socialProvider.ClientKey, socialProvider.ClientSecret, callbackURL, scopes...)
	if err != nil {
		return err
	}

	sess, err := newGothSession(in.Provider)
	if err != nil {
		return err
	}
	_, err = sess.Authorize(gothProvider, r.URL.Query())
	if err != nil {
		return stackerror.NewUseCaseError(
			"SocialCallbackUseCase",
			err,
			stackerror.WithMessage("failed to exchange authorization code"),
			stackerror.WithStatusCode(http.StatusBadRequest),
		)
	}

	gothUser, err := gothProvider.FetchUser(sess)
	if err != nil {
		return stackerror.NewUseCaseError(
			"SocialCallbackUseCase",
			err,
			stackerror.WithMessage("failed to fetch user from provider"),
			stackerror.WithStatusCode(http.StatusBadRequest),
		)
	}

	account, err := u.getAccountBySocialProviderRepository.Execute(ctx, in.Provider, gothUser.UserID)
	if err != nil && !errors.Is(err, stackerror.ErrNotFound) {
		return stackerror.NewUseCaseError("SocialCallbackUseCase", err)
	}

	if errors.Is(err, stackerror.ErrNotFound) {
		account, err = u.getAccountByEmailAndProjectRepository.Execute(ctx, gothUser.Email, client.ProjectID)
		if err != nil && !errors.Is(err, stackerror.ErrNotFound) {
			return stackerror.NewUseCaseError("SocialCallbackUseCase", err)
		}

		if errors.Is(err, stackerror.ErrNotFound) {
			signup, signupErr := u.getSignupByClientIDRepository.Execute(ctx, client.ID)
			if signupErr != nil || !signup.Enabled {
				return stackerror.NewUseCaseError(
					"SocialCallbackUseCase",
					errors.New("signup is disabled"),
					stackerror.WithMessage("registration is disabled for this client"),
					stackerror.WithStatusCode(http.StatusForbidden),
				)
			}

			name := gothUser.Name
			if name == "" {
				name = gothUser.Email
			}

			account, err = u.createAccountRepository.Execute(ctx, domain.Account{
				ProjectID: client.ProjectID,
				Name:      name,
				Email:     gothUser.Email,
				Username:  gothUser.Email,
				RoleName:  signup.DefaultRoleName,
			})
			if err != nil {
				return stackerror.NewUseCaseError("SocialCallbackUseCase", err)
			}
		}

		if linkErr := u.createAccountSocialProviderRepository.Execute(ctx, domain.AccountSocialProvider{
			AccountID:      account.ID,
			Provider:       in.Provider,
			ProviderUserID: gothUser.UserID,
			Email:          gothUser.Email,
		}); linkErr != nil {
			return stackerror.NewUseCaseError("SocialCallbackUseCase", linkErr)
		}
	}

	if delErr := u.deleteSocialStateRepository.Execute(ctx, in.Nonce); delErr != nil {
		return stackerror.NewUseCaseError("SocialCallbackUseCase", delErr)
	}

	codeBytes := make([]byte, 32)
	if _, err = rand.Read(codeBytes); err != nil {
		return stackerror.NewUseCaseError(
			"SocialCallbackUseCase",
			err,
			stackerror.WithStatusCode(http.StatusInternalServerError),
		)
	}
	code := base64.URLEncoding.EncodeToString(codeBytes)

	authCode := domain.NewAuthCode(
		"",
		code,
		client.ID,
		account.ID,
		socialState.RedirectURI,
		socialState.CodeChallenge,
		socialState.CodeChallengeMethod,
		time.Now().Add(5*time.Minute).Format(time.RFC3339),
		false,
		time.Now().Format(time.RFC3339),
	)

	authCode, err = u.createAuthCodeRepository.Execute(ctx, authCode)
	if err != nil {
		return stackerror.NewUseCaseError("SocialCallbackUseCase", err)
	}

	redirectTo := fmt.Sprintf("%s?code=%s&state=%s", socialState.RedirectURI, authCode.Code, socialState.State)
	http.Redirect(w, r, redirectTo, http.StatusFound)
	return nil
}

func NewSocialCallbackUseCase(
	getSocialStateByNonceRepository contract.GetSocialStateByNonceRepository,
	deleteSocialStateRepository contract.DeleteSocialStateRepository,
	getClientByClientIDRepository contract.GetClientByClientIDRepository,
	getSocialProviderByClientRepository contract.GetSocialProviderByClientRepository,
	getAccountBySocialProviderRepository contract.GetAccountBySocialProviderRepository,
	getAccountByEmailAndProjectRepository contract.GetAccountByEmailAndProjectRepository,
	createAccountSocialProviderRepository contract.CreateAccountSocialProviderRepository,
	createAccountRepository contract.CreateAccountRepository,
	createAuthCodeRepository contract.CreateAuthCodeRepository,
	getSignupByClientIDRepository contract.GetSignupByClientIDRepository,
) contract.SocialCallbackUseCase {
	return &socialCallbackUseCase{
		getSocialStateByNonceRepository:       getSocialStateByNonceRepository,
		deleteSocialStateRepository:           deleteSocialStateRepository,
		getClientByClientIDRepository:         getClientByClientIDRepository,
		getSocialProviderByClientRepository:   getSocialProviderByClientRepository,
		getAccountBySocialProviderRepository:  getAccountBySocialProviderRepository,
		getAccountByEmailAndProjectRepository: getAccountByEmailAndProjectRepository,
		createAccountSocialProviderRepository: createAccountSocialProviderRepository,
		createAccountRepository:               createAccountRepository,
		createAuthCodeRepository:              createAuthCodeRepository,
		getSignupByClientIDRepository:         getSignupByClientIDRepository,
	}
}
