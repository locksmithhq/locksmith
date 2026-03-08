package di

import (
	"github.com/locksmithhq/locksmith/api/internal/adapter/database"
	"github.com/locksmithhq/locksmith/api/internal/project/contract"
	"github.com/locksmithhq/locksmith/api/internal/project/handler"
	"github.com/locksmithhq/locksmith/api/internal/project/repository"
	"github.com/locksmithhq/locksmith/api/internal/project/usecase"
)

func NewUpdateProjectHandler() contract.UpdateProjectHandler {
	conn := database.GetConnection()
	return handler.NewUpdateProjectHandler(
		usecase.NewUpdateProjectUseCase(
			repository.NewUpdateProjectRepository(conn),
		),
	)
}
