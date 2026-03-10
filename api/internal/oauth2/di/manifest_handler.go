package di

import (
	"github.com/locksmithhq/locksmith/api/internal/adapter/database"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/handler"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/repository"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/usecase"
)

func NewPWAManifestHandler() contract.PWAManifestHandler {
	return handler.NewPWAManifestHandler(
		usecase.NewGetPWAManifestUseCase(
			repository.NewGetClientByClientIDRepository(database.GetConnection()),
			repository.NewGetLoginByClientIDRepository(database.GetConnection()),
		),
	)
}
