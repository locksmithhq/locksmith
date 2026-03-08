package di

import (
	"github.com/locksmithhq/locksmith/api/internal/adapter/database"
	"github.com/locksmithhq/locksmith/api/internal/project/contract"
	"github.com/locksmithhq/locksmith/api/internal/project/handler"
	"github.com/locksmithhq/locksmith/api/internal/project/repository"
	"github.com/locksmithhq/locksmith/api/internal/project/usecase"
)

func NewGetProjectByIDHandler() contract.GetProjectByIDHandler {
	return handler.NewGetProjectByIDHandler(
		usecase.NewGetProjectByIDUseCase(
			repository.NewGetProjectByIDRepository(
				database.GetConnection(),
			),
		),
	)
}
