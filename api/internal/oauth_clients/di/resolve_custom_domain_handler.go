package di

import (
	"github.com/locksmithhq/locksmith/api/internal/adapter/database"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients/handler"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients/repository"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients/usecase"
)

func NewResolveCustomDomainHandler() contract.ResolveCustomDomainHandler {
	conn := database.GetConnection()
	return handler.NewResolveCustomDomainHandler(
		usecase.NewResolveCustomDomainUseCase(
			repository.NewGetClientByCustomDomainRepository(conn),
		),
	)
}
