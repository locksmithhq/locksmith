package di

import (
	"github.com/locksmithhq/locksmith/api/internal/acl/contract"
	"github.com/locksmithhq/locksmith/api/internal/acl/handler"
	"github.com/locksmithhq/locksmith/api/internal/acl/repository"
	"github.com/locksmithhq/locksmith/api/internal/acl/usecase"
	"github.com/locksmithhq/locksmith/api/internal/adapter/database"
)

func NewFetchRolesHandler() contract.FetchRolesHandler {
	return handler.NewFetchRolesHandler(
		usecase.NewFetchRolesUseCase(
			repository.NewFetchRolesRepository(
				database.GetConnection(),
			),
		),
	)
}
