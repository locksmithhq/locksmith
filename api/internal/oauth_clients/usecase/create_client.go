package usecase

import (
	"context"
	"errors"

	"github.com/booscaaa/locksmith/api/internal/core/types/stackerror"
	"github.com/booscaaa/locksmith/api/internal/oauth_clients/contract"
	"github.com/booscaaa/locksmith/api/internal/oauth_clients/types/input"
	"github.com/booscaaa/locksmith/api/internal/oauth_clients/types/output"
	"github.com/google/uuid"
)

type createClientUseCase struct {
	createClientRepository                    contract.CreateClientRepository
	getClientByProjectIDAndClientIDRepository contract.GetClientByProjectIDAndClientIDRepository
}

// Execute implements contract.CreateClientUseCase.
func (u *createClientUseCase) Execute(ctx context.Context, in input.Client) (output.Client, error) {
	client, err := u.getClientByProjectIDAndClientIDRepository.Execute(ctx, in.ProjectID, in.ClientID)
	if err != nil && !errors.Is(err, stackerror.ErrNotFound) {
		return output.Client{}, stackerror.NewUseCaseError("CreateClientUseCase", err)
	}

	if client.ID != "" {
		return output.Client{}, nil
	}

	if in.ClientID == "" {
		in.ClientID = uuid.New().String()
	}

	if in.ClientSecret == "" {
		in.ClientSecret = uuid.New().String()
	}

	client, err = u.createClientRepository.Execute(ctx, in)
	if err != nil {
		return output.Client{}, stackerror.NewUseCaseError("CreateClientUseCase", err)
	}

	return output.NewClientFromDomain(client), nil
}

func NewCreateClientUseCase(
	createClientRepository contract.CreateClientRepository,
	getClientByProjectIDAndClientIDRepository contract.GetClientByProjectIDAndClientIDRepository,
) contract.CreateClientUseCase {
	return &createClientUseCase{
		createClientRepository:                    createClientRepository,
		getClientByProjectIDAndClientIDRepository: getClientByProjectIDAndClientIDRepository,
	}
}
