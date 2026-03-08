package usecase

import (
	"context"
	"fmt"

	"github.com/locksmithhq/locksmith/api/internal/account/contract"
	"github.com/locksmithhq/locksmith/api/internal/account/types/input"
	"github.com/locksmithhq/locksmith/api/internal/account/types/output"
	aclContract "github.com/locksmithhq/locksmith/api/internal/acl/contract"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith-go"
)

type createAccountUseCase struct {
	createAccountRepository                 contract.CreateAccountRepository
	getAccountByEmailAndProjectIDRepository contract.GetAccountByEmailAndProjectIDRepository
	getProjectDomainByProjectIDRepository   aclContract.GetProjectDomainByProjectIDRepository
}

// Execute implements contract.CreateAccountUseCase.
func (u *createAccountUseCase) Execute(ctx context.Context, in input.Account) (output.Account, error) {
	account, err := u.createAccountRepository.Execute(ctx, in.ToDomain())
	if err != nil {
		return output.Account{}, stackerror.NewUseCaseError("CreateAccountUseCase", err)
	}

	domain, err := u.getProjectDomainByProjectIDRepository.Execute(ctx, in.ProjectID)
	if err != nil {
		return output.Account{}, stackerror.NewUseCaseError("CreateAccountUseCase", err)
	}

	if _, err := locksmith.AddRoleForUser(
		fmt.Sprintf("user:%s", account.ID),
		in.RoleName,
		domain,
	); err != nil {
		return output.Account{}, stackerror.NewUseCaseError("CreateAccountUseCase", err)
	}

	return output.NewAccountFromDomain(account), nil
}

func NewCreateAccountUseCase(
	createAccountRepository contract.CreateAccountRepository,
	getAccountByEmailAndProjectIDRepository contract.GetAccountByEmailAndProjectIDRepository,
	getProjectDomainByProjectIDRepository aclContract.GetProjectDomainByProjectIDRepository,
) contract.CreateAccountUseCase {
	return &createAccountUseCase{
		createAccountRepository:                 createAccountRepository,
		getAccountByEmailAndProjectIDRepository: getAccountByEmailAndProjectIDRepository,
		getProjectDomainByProjectIDRepository:   getProjectDomainByProjectIDRepository,
	}
}
