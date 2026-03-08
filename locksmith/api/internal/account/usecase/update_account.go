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

type updateAccountUseCase struct {
	updateAccountRepository               contract.UpdateAccountRepository
	getProjectDomainByProjectIDRepository aclContract.GetProjectDomainByProjectIDRepository
}

// Execute implements contract.UpdateAccountUseCase.
func (u *updateAccountUseCase) Execute(ctx context.Context, in input.UpdateAccount) (output.Account, error) {
	account, err := u.updateAccountRepository.Execute(ctx, in.ToDomain())
	if err != nil {
		return output.Account{}, stackerror.NewUseCaseError("UpdateAccountUseCase", err)
	}

	domain, err := u.getProjectDomainByProjectIDRepository.Execute(ctx, in.ProjectID)
	if err != nil {
		return output.Account{}, stackerror.NewUseCaseError("UpdateAccountUseCase", err)
	}

	userId := fmt.Sprintf("user:%s", account.ID)

	// Remove old roles and add the new one
	if _, err := locksmith.DeleteRolesForUser(userId, domain); err != nil {
		return output.Account{}, stackerror.NewUseCaseError("UpdateAccountUseCase", err)
	}

	if _, err := locksmith.AddRoleForUser(
		userId,
		in.RoleName,
		domain,
	); err != nil {
		return output.Account{}, stackerror.NewUseCaseError("UpdateAccountUseCase", err)
	}

	return output.NewAccountFromDomain(account), nil
}

func NewUpdateAccountUseCase(
	updateAccountRepository contract.UpdateAccountRepository,
	getProjectDomainByProjectIDRepository aclContract.GetProjectDomainByProjectIDRepository,
) contract.UpdateAccountUseCase {
	return &updateAccountUseCase{
		updateAccountRepository:               updateAccountRepository,
		getProjectDomainByProjectIDRepository: getProjectDomainByProjectIDRepository,
	}
}
