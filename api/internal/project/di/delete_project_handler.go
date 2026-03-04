package di

import (
	"github.com/booscaaa/locksmith/api/internal/adapter/database"
	"github.com/booscaaa/locksmith/api/internal/project/contract"
	"github.com/booscaaa/locksmith/api/internal/project/handler"
	"github.com/booscaaa/locksmith/api/internal/project/repository"
	"github.com/booscaaa/locksmith/api/internal/project/usecase"
)

func NewDeleteProjectHandler() contract.DeleteProjectHandler {
	conn := database.GetConnection()
	return handler.NewDeleteProjectHandler(
		usecase.NewDeleteProjectUseCase(
			repository.NewDeleteProjectRepository(conn),
		),
	)
}
