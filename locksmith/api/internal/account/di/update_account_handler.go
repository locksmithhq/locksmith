package di

import (
	"github.com/locksmithhq/locksmith/api/internal/account/contract"
	"github.com/locksmithhq/locksmith/api/internal/account/handler"
	"github.com/locksmithhq/locksmith/api/internal/account/repository"
	"github.com/locksmithhq/locksmith/api/internal/account/usecase"
	aclRepository "github.com/locksmithhq/locksmith/api/internal/acl/repository"
	"github.com/locksmithhq/locksmith/api/internal/adapter/database"
)

func NewUpdateAccountHandler() contract.UpdateAccountHandler {
	database := database.GetConnection()
	updateAccountRepository := repository.NewUpdateAccountRepository(database)
	getProjectDomainByProjectIDRepository := aclRepository.NewGetProjectDomainByProjectIDRepository(database)
	updateAccountUseCase := usecase.NewUpdateAccountUseCase(
		updateAccountRepository,
		getProjectDomainByProjectIDRepository,
	)
	return handler.NewUpdateAccountHandler(updateAccountUseCase)
}
