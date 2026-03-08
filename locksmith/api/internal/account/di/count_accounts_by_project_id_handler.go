package di

import (
	"github.com/locksmithhq/locksmith/api/internal/account/contract"
	"github.com/locksmithhq/locksmith/api/internal/account/handler"
	"github.com/locksmithhq/locksmith/api/internal/account/repository"
	"github.com/locksmithhq/locksmith/api/internal/account/usecase"
	"github.com/locksmithhq/locksmith/api/internal/adapter/database"
)

func NewCountAccountsByProjectIDHandler() contract.CountAccountsByProjectIDHandler {
	return handler.NewCountAccountsByProjectIDHandler(
		usecase.NewCountAccountsByProjectIDUseCase(
			repository.NewCountAccountsByProjectIDRepository(
				database.GetConnection(),
			),
		),
	)
}
