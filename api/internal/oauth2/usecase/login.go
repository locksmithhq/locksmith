package usecase

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	contract "github.com/locksmithhq/locksmith/api/internal/oauth2/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/domain"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/types/input"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/types/output"
	"github.com/locksmithhq/locksmith-go"
)

type loginUseCase struct {
	getClientByClientIDRepository                   contract.GetClientByClientIDRepository
	getAccountByEmailPasswordAndProjectIDRepository contract.GetAccountByEmailPasswordAndProjectIDRepository
	createAuthCodeUseCase                           contract.CreateAuthCodeRepository
	getLoginByClientIDRepository                    contract.GetLoginByClientIDRepository
}

func (u *loginUseCase) Execute(ctx context.Context, in input.Login) (output.Login, error) {
	client, err := u.getClientByClientIDRepository.Execute(ctx, in.ClientID)
	if err != nil {
		if errors.Is(err, stackerror.ErrNotFound) {
			return output.Login{}, stackerror.NewUseCaseError(
				"LoginUseCase",
				err,
				stackerror.WithMessage("client id is not valid"),
				stackerror.WithStatusCode(http.StatusBadRequest),
			)
		}
		return output.Login{}, stackerror.NewUseCaseError(
			"LoginUseCase",
			err,
			stackerror.WithStatusCode(http.StatusInternalServerError),
		)
	}

	loginConfig, err := u.getLoginByClientIDRepository.Execute(ctx, client.ID)
	if err != nil && !errors.Is(err, stackerror.ErrNotFound) {
		return output.Login{}, stackerror.NewUseCaseError(
			"LoginUseCase",
			err,
		)
	}
	if !errors.Is(err, stackerror.ErrNotFound) && !loginConfig.Enabled {
		return output.Login{}, stackerror.NewUseCaseError(
			"LoginUseCase",
			errors.New("login is disabled"),
			stackerror.WithMessage("login is disabled for this client"),
			stackerror.WithStatusCode(http.StatusForbidden),
		)
	}

	account, err := u.getAccountByEmailPasswordAndProjectIDRepository.Execute(ctx, in.Email, in.Password, client.ProjectID)
	if err != nil {
		if errors.Is(err, stackerror.ErrNotFound) {
			return output.Login{}, stackerror.NewUseCaseError(
				"LoginUseCase",
				err,
				stackerror.WithMessage("account not found"),
			)
		}
		return output.Login{}, stackerror.NewUseCaseError(
			"LoginUseCase",
			err,
		)
	}

	if ok, err := locksmith.Enforce(
		fmt.Sprintf("user:%s", account.ID),
		client.Domain,
		"module:oauth2",
		"action:login",
	); err != nil || !ok {
		return output.Login{}, stackerror.NewUseCaseError(
			"LoginUseCase",
			errors.New("forbidden"),
			stackerror.WithMessage("user is not allowed to login in this domain"),
			stackerror.WithStatusCode(http.StatusForbidden),
		)
	}

	if account.MustChangePassword {
		claims := locksmith.NewTokenClaims(account.ID, client.ClientID, client.Domain)
		accessToken := locksmith.GetSignToken(claims, 1*time.Minute, client.ClientSecret+"-renew")

		return output.Login{
			MustChangePassword: true,
			ChangePasswordJWT:  accessToken,
		}, nil
	}

	// Generate random code
	b := make([]byte, 32)
	_, err = rand.Read(b)
	if err != nil {
		return output.Login{}, stackerror.NewUseCaseError(
			"LoginUseCase",
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
		in.CodeChallenge,
		in.CodeChallengeMethod,
		time.Now().Add(5*time.Minute).Format(time.RFC3339),
		false,
		time.Now().Format(time.RFC3339),
	)

	authCode, err = u.createAuthCodeUseCase.Execute(ctx, authCode)
	if err != nil {
		return output.Login{}, stackerror.NewUseCaseError(
			"LoginUseCase",
			err,
		)
	}

	redirectTo := fmt.Sprintf("%s?code=%s&state=%s", in.RedirectURI, authCode.Code, in.State)
	if in.DeviceID != "" {
		redirectTo += "&device_id=" + in.DeviceID
	}

	return output.Login{
		RedirectTo: redirectTo,
	}, nil
}

func NewLoginUseCase(
	getClientByClientIDRepository contract.GetClientByClientIDRepository,
	getAccountByEmailPasswordAndProjectIDRepository contract.GetAccountByEmailPasswordAndProjectIDRepository,
	createAuthCodeUseCase contract.CreateAuthCodeRepository,
	getLoginByClientIDRepository contract.GetLoginByClientIDRepository,
) contract.LoginUseCase {
	return &loginUseCase{
		getClientByClientIDRepository:                   getClientByClientIDRepository,
		getAccountByEmailPasswordAndProjectIDRepository: getAccountByEmailPasswordAndProjectIDRepository,
		createAuthCodeUseCase:                           createAuthCodeUseCase,
		getLoginByClientIDRepository:                    getLoginByClientIDRepository,
	}
}
