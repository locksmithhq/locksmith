package di

import (
	"github.com/locksmithhq/locksmith/api/internal/acl/contract"
	"github.com/locksmithhq/locksmith/api/internal/acl/handler"
	"github.com/locksmithhq/locksmith/api/internal/acl/repository"
	"github.com/locksmithhq/locksmith/api/internal/acl/usecase"
	"github.com/locksmithhq/locksmith/api/internal/adapter/database"
)

func NewCreateModuleHandler() contract.CreateModuleHandler {
	return handler.NewCreateModuleHandler(
		usecase.NewCreateModuleUseCase(
			repository.NewCreateModuleRepository(database.GetConnection()),
		),
	)
}
