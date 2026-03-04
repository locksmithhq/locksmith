package di

import (
	"github.com/booscaaa/locksmith/api/internal/acl/contract"
	"github.com/booscaaa/locksmith/api/internal/acl/handler"
	"github.com/booscaaa/locksmith/api/internal/acl/repository"
	"github.com/booscaaa/locksmith/api/internal/acl/usecase"
	"github.com/booscaaa/locksmith/api/internal/adapter/database"
)

func NewFetchActionsHandler() contract.FetchActionsHandler {
	return handler.NewFetchActionsHandler(
		usecase.NewFetchActionsUseCase(
			repository.NewFetchActionsRepository(
				database.GetConnection(),
			),
		),
	)
}
