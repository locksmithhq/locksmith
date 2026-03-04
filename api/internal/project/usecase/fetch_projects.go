package usecase

import (
	"context"

	"github.com/booscaaa/go-paginate/v3/paginate"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/project/contract"
	"github.com/locksmithhq/locksmith/api/internal/project/types/output"
)

type fetchProjectsUseCase struct {
	fetchProjectsRepository contract.FetchProjectsRepository
}

// Execute implements contract.FetchProjectsUseCase.
func (f *fetchProjectsUseCase) Execute(ctx context.Context, params paginate.PaginationParams) ([]output.Project, error) {
	projects, err := f.fetchProjectsRepository.Execute(ctx, params)
	if err != nil {
		return nil, stackerror.NewUseCaseError("FetchProjectsUseCase", err)
	}

	return output.NewProjects(projects), nil
}

func NewFetchProjectsUseCase(fetchProjectsRepository contract.FetchProjectsRepository) contract.FetchProjectsUseCase {
	return &fetchProjectsUseCase{fetchProjectsRepository: fetchProjectsRepository}
}
