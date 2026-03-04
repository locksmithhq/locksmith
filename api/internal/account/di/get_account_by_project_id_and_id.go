package di

import (
	"github.com/locksmithhq/locksmith/api/internal/account/contract"
	"github.com/locksmithhq/locksmith/api/internal/account/handler"
	"github.com/locksmithhq/locksmith/api/internal/account/repository"
	"github.com/locksmithhq/locksmith/api/internal/account/usecase"
	"github.com/locksmithhq/locksmith/api/internal/adapter/database"
)

func NewGetAccountByProjectIDAndIDHandler() contract.GetAccountByProjectIDAndIDHandler {
	return handler.NewGetAccountByProjectIDAndIDHandler(
		usecase.NewGetAccountByProjectIDAndIDUseCase(
			repository.NewGetAccountByProjectIDAndIDRepository(database.GetConnection()),
		),
	)
}
