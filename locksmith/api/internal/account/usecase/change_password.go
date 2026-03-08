package usecase

import (
	"context"
	"errors"
	"net/http"

	"github.com/locksmithhq/locksmith-go"
	"github.com/locksmithhq/locksmith/api/internal/account/contract"
	"github.com/locksmithhq/locksmith/api/internal/account/types/input"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	oauth2Contract "github.com/locksmithhq/locksmith/api/internal/oauth2/contract"
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

	if err := u.updateAccountPasswordRepository.Execute(ctx, claims.Sub, in.Password); err != nil {
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
