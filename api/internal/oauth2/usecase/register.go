package usecase

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"strings"
	"time"

	accountContract "github.com/booscaaa/locksmith/api/internal/account/contract"
	accountInput "github.com/booscaaa/locksmith/api/internal/account/types/input"
	"github.com/booscaaa/locksmith/api/internal/core/types/stackerror"
	"github.com/booscaaa/locksmith/api/internal/oauth2/contract"
	"github.com/booscaaa/locksmith/api/internal/oauth2/domain"
	"github.com/booscaaa/locksmith/api/internal/oauth2/types/input"
	"github.com/booscaaa/locksmith/api/internal/oauth2/types/output"
)

type registerUseCase struct {
	getClientByClientIDRepository contract.GetClientByClientIDRepository
	getSignupByClientIDRepository contract.GetSignupByClientIDRepository
	createAccountUseCase          accountContract.CreateAccountUseCase
	createAuthCodeRepository      contract.CreateAuthCodeRepository
}

func (u *registerUseCase) Execute(ctx context.Context, in input.Register) (output.Login, error) {
	client, err := u.getClientByClientIDRepository.Execute(ctx, in.ClientID)
	if err != nil {
		return output.Login{}, stackerror.NewUseCaseError(
			"RegisterUseCase",
			err,
			stackerror.WithMessage("the client_id is not valid"),
			stackerror.WithStatusCode(http.StatusBadRequest),
		)
	}

	clientRedirectURIs := strings.Split(client.RedirectURIs, " ")
	if !slices.Contains(clientRedirectURIs, in.RedirectURI) {
		return output.Login{}, stackerror.NewUseCaseError(
			"RegisterUseCase",
			fmt.Errorf("invalid redirect_uri"),
			stackerror.WithMessage("the redirect_uri is not valid"),
			stackerror.WithStatusCode(http.StatusBadRequest),
		)
	}

	signup, err := u.getSignupByClientIDRepository.Execute(ctx, client.ID)
	if err != nil || !signup.Enabled {
		return output.Login{}, stackerror.NewUseCaseError(
			"RegisterUseCase",
			errors.New("signup is disabled"),
			stackerror.WithMessage("registration is disabled for this client"),
			stackerror.WithStatusCode(http.StatusForbidden),
		)
	}

	accountIn := accountInput.Account{
		ProjectID:          client.ProjectID,
		Name:               in.Name,
		Email:              in.Email,
		Username:           in.Email,
		Password:           in.Password,
		RoleName:           "user",
		MustChangePassword: false,
	}

	account, err := u.createAccountUseCase.Execute(ctx, accountIn)
	if err != nil {
		return output.Login{}, stackerror.NewUseCaseError(
			"RegisterUseCase",
			err,
			stackerror.WithStatusCode(http.StatusInternalServerError),
		)
	}

	b := make([]byte, 32)
	if _, err = rand.Read(b); err != nil {
		return output.Login{}, stackerror.NewUseCaseError(
			"RegisterUseCase",
			err,
			stackerror.WithStatusCode(http.StatusInternalServerError),
		)
	}
	code := base64.URLEncoding.EncodeToString(b)

	authCode := domain.NewAuthCode(
		"",
		code,
		client.ID,
		account.ID,
		client.RedirectURIs,
		"",
		"",
		time.Now().Add(5*time.Minute).Format(time.RFC3339),
		false,
		time.Now().Format(time.RFC3339),
	)

	authCode, err = u.createAuthCodeRepository.Execute(ctx, authCode)
	if err != nil {
		return output.Login{}, stackerror.NewUseCaseError(
			"RegisterUseCase",
			err,
		)
	}

	return output.Login{
		RedirectTo: fmt.Sprintf("%s?code=%s&state=%s", in.RedirectURI, authCode.Code, in.State),
	}, nil
}

func NewRegisterUseCase(
	getClientByClientIDRepository contract.GetClientByClientIDRepository,
	getSignupByClientIDRepository contract.GetSignupByClientIDRepository,
	createAccountUseCase accountContract.CreateAccountUseCase,
	createAuthCodeRepository contract.CreateAuthCodeRepository,
) contract.RegisterUseCase {
	return &registerUseCase{
		getClientByClientIDRepository: getClientByClientIDRepository,
		getSignupByClientIDRepository: getSignupByClientIDRepository,
		createAccountUseCase:          createAccountUseCase,
		createAuthCodeRepository:      createAuthCodeRepository,
	}
}
