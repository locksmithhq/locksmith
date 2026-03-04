package di

import (
	"github.com/booscaaa/locksmith/api/internal/acl/contract"
	"github.com/booscaaa/locksmith/api/internal/acl/handler"
	"github.com/booscaaa/locksmith/api/internal/acl/repository"
	"github.com/booscaaa/locksmith/api/internal/acl/usecase"
	"github.com/booscaaa/locksmith/api/internal/adapter/database"
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
