package usecase

import (
	"context"

	"github.com/booscaaa/locksmith/api/internal/core/types/stackerror"
	"github.com/booscaaa/locksmith/api/internal/oauth_clients/contract"
	"github.com/booscaaa/locksmith/api/internal/oauth_clients/types/output"
)

type getClientByIDAndProjectIDUseCase struct {
	getClientByIDAndProjectIDRepository contract.GetClientByIDAndProjectIDRepository
}

func (u *getClientByIDAndProjectIDUseCase) Execute(ctx context.Context, id string, projectID string) (output.Client, error) {
	client, err := u.getClientByIDAndProjectIDRepository.Execute(ctx, id, projectID)
	if err != nil {
		return output.Client{}, stackerror.NewUseCaseError("GetClientByIDAndProjectIDUseCase", err)
	}

	return output.NewClientFromDomain(client), nil
}

func NewGetClientByIDAndProjectIDUseCase(
	getClientByIDAndProjectIDRepository contract.GetClientByIDAndProjectIDRepository,
) contract.GetClientByIDAndProjectIDUseCase {
	return &getClientByIDAndProjectIDUseCase{
		getClientByIDAndProjectIDRepository: getClientByIDAndProjectIDRepository,
	}
}
