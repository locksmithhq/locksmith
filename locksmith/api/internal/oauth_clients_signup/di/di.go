package di

import (
	"github.com/locksmithhq/locksmith/api/internal/adapter/database"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients_signup/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients_signup/handler"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients_signup/repository"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients_signup/usecase"
)

func NewGetSignupByClientIDHandler() contract.GetSignupByClientIDHandler {
	return handler.NewGetSignupByClientIDHandler(
		usecase.NewGetSignupByClientIDUseCase(
			repository.NewGetSignupByClientIDRepository(
				database.GetConnection(),
			),
		),
	)
}

func NewCreateSignupByClientIDHandler() contract.CreateSignupByClientIDHandler {
	return handler.NewCreateSignupByClientIDHandler(
		usecase.NewCreateSignupByClientIDUseCase(
			repository.NewCreateSignupByClientIDRepository(
				database.GetConnection(),
			),
		),
	)
}

func NewUpdateSignupByClientIDHandler() contract.UpdateSignupByClientIDHandler {
	return handler.NewUpdateSignupByClientIDHandler(
		usecase.NewUpdateSignupByClientIDUseCase(
			repository.NewUpdateSignupByClientIDRepository(
				database.GetConnection(),
			),
		),
	)
}
