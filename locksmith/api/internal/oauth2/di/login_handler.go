package di

import (
	"github.com/locksmithhq/locksmith/api/internal/adapter/database"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/handler"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/repository"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/usecase"
)

func NewLoginHandler() contract.LoginHandler {
	return handler.NewLoginHandler(
		usecase.NewLoginUseCase(
			repository.NewGetClientByClientIDRepository(database.GetConnection()),
			repository.NewGetAccountByEmailPasswordAndProjectIDRepository(database.GetConnection()),
			repository.NewCreateAuthCodeRepository(database.GetConnection()),
			repository.NewGetLoginByClientIDRepository(database.GetConnection()),
		),
	)
}
