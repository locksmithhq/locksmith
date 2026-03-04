package di

import (
	"github.com/locksmithhq/locksmith/api/internal/account/contract"
	"github.com/locksmithhq/locksmith/api/internal/account/handler"
	"github.com/locksmithhq/locksmith/api/internal/account/repository"
	"github.com/locksmithhq/locksmith/api/internal/account/usecase"
	aclRepository "github.com/locksmithhq/locksmith/api/internal/acl/repository"
	"github.com/locksmithhq/locksmith/api/internal/adapter/database"
)

func NewCreateAccountHandler() contract.CreateAccountHandler {
	return handler.NewCreateAccountHandler(
		usecase.NewCreateAccountUseCase(
			repository.NewCreateAccountRepository(database.GetConnection()),
			repository.NewGetAccountByEmailAndProjectIDRepository(database.GetConnection()),
			aclRepository.NewGetProjectDomainByProjectIDRepository(database.GetConnection()),
		),
	)
}
