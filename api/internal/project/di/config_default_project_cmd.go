package di

import (
	"github.com/booscaaa/locksmith/api/internal/adapter/database"
	"github.com/booscaaa/locksmith/api/internal/project/cmd"
	"github.com/booscaaa/locksmith/api/internal/project/contract"
	"github.com/booscaaa/locksmith/api/internal/project/repository"
	"github.com/booscaaa/locksmith/api/internal/project/usecase"
)

func NewConfigDefaultProjectCMD() contract.ConfigDefaultProjectCMD {
	return cmd.NewConfigDefaultProject(
		usecase.NewCreateProjectUseCase(
			repository.NewCreateProjectRepository(
				database.GetConnection(),
			),
			repository.NewGetProjectByNameRepository(
				database.GetConnection(),
			),
		),
	)
}
