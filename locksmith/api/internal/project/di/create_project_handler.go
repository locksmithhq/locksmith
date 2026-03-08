package di

import (
	"github.com/locksmithhq/locksmith/api/internal/adapter/database"
	"github.com/locksmithhq/locksmith/api/internal/project/contract"
	"github.com/locksmithhq/locksmith/api/internal/project/handler"
	"github.com/locksmithhq/locksmith/api/internal/project/repository"
	"github.com/locksmithhq/locksmith/api/internal/project/usecase"
)

func NewCreateProjectHandler() contract.CreateProjectHandler {
	conn := database.GetConnection()
	return handler.NewCreateProjectHandler(
		usecase.NewCreateProjectUseCase(
			repository.NewCreateProjectRepository(conn),
			repository.NewGetProjectByNameRepository(conn),
		),
	)
}
