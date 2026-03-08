package usecase

import (
	"context"

	"github.com/locksmithhq/locksmith/api/internal/account/contract"
	"github.com/locksmithhq/locksmith/api/internal/account/types/output"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
)

type getAccountByProjectIDAndIDUseCase struct {
	getAccountByProjectIDAndIDRepository contract.GetAccountByProjectIDAndIDRepository
}

func (u *getAccountByProjectIDAndIDUseCase) Execute(ctx context.Context, projectID string, id string) (output.Account, error) {
	account, err := u.getAccountByProjectIDAndIDRepository.Execute(ctx, projectID, id)
	if err != nil {
		return output.Account{}, stackerror.NewUseCaseError(
			"GetAccountByProjectIDAndIDUseCase",
			err,
		)
	}

	return output.NewAccountFromDomain(account), nil
}

func NewGetAccountByProjectIDAndIDUseCase(
	getAccountByProjectIDAndIDRepository contract.GetAccountByProjectIDAndIDRepository,
) contract.GetAccountByProjectIDAndIDUseCase {
	return &getAccountByProjectIDAndIDUseCase{
		getAccountByProjectIDAndIDRepository: getAccountByProjectIDAndIDRepository,
	}
}
