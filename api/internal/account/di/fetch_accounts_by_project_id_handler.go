package di

import (
	"github.com/booscaaa/locksmith/api/internal/account/contract"
	"github.com/booscaaa/locksmith/api/internal/account/handler"
	"github.com/booscaaa/locksmith/api/internal/account/repository"
	"github.com/booscaaa/locksmith/api/internal/account/usecase"
	"github.com/booscaaa/locksmith/api/internal/adapter/database"
)

func NewFetchAccountsByProjectIDHandler() contract.FetchAccountsByProjectIDHandler {
	return handler.NewFetchAccountsByProjectIDHandler(
		usecase.NewFetchAccountsByProjectIDUseCase(
			repository.NewFetchAccountsByProjectIDRepository(
				database.GetConnection(),
			),
		),
	)
}
