package usecase

import (
	"context"
	"errors"
	"net/http"

	"github.com/booscaaa/locksmith/api/internal/account/contract"
	"github.com/booscaaa/locksmith/api/internal/account/types/input"
	"github.com/booscaaa/locksmith/api/internal/core/types/stackerror"
	oauth2Contract "github.com/booscaaa/locksmith/api/internal/oauth2/contract"
	"github.com/locksmithhq/locksmith-go"
)

type changePasswordUseCase struct {
	updateAccountPasswordRepository contract.UpdateAccountPasswordRepository
	getClientByClientIDRepository   oauth2Contract.GetClientByClientIDRepository
}

func (u *changePasswordUseCase) Execute(ctx context.Context, in input.ChangePassword) error {
	client, err := u.getClientByClientIDRepository.Execute(ctx, in.ClientID)
	if err != nil {
		return stackerror.NewUseCaseError(
			"ChangePasswordUseCase",
			err,
		)
	}

	claims, valid := locksmith.VerifyTokenWithClientSecret(in.Jwt, client.ClientSecret+"-renew")
	if !valid {
		return stackerror.NewUseCaseError(
			"ChangePasswordUseCase",
			errors.New("invalid token"),
			stackerror.WithStatusCode(http.StatusUnauthorized),
		)
	}

	if err := u.updateAccountPasswordRepository.Execute(ctx, claims.Data["sub"].(string), in.Password); err != nil {
		return stackerror.NewUseCaseError(
			"ChangePasswordUseCase",
			err,
		)
	}

	return nil
}

func NewChangePasswordUseCase(
	updateAccountPasswordRepository contract.UpdateAccountPasswordRepository,
	getClientByClientIDRepository oauth2Contract.GetClientByClientIDRepository,
) contract.ChangePasswordUseCase {
	return &changePasswordUseCase{
		updateAccountPasswordRepository: updateAccountPasswordRepository,
		getClientByClientIDRepository:   getClientByClientIDRepository,
	}
}
