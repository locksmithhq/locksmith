package usecase

import (
	"context"

	"github.com/locksmithhq/locksmith/api/internal/oauth_clients_login/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients_login/types/input"
)

type createLoginByClientIDUseCase struct {
	createLoginByClientIDRepository contract.CreateLoginByClientIDRepository
}

// Execute implements contract.CreateLoginByClientIDUseCase.
func (u *createLoginByClientIDUseCase) Execute(ctx context.Context, clientID string, in input.Login) error {
	loginDomain := in.ToLoginDomain()
	loginDomain.ClientID = clientID
	return u.createLoginByClientIDRepository.Execute(ctx, loginDomain)
}

func NewCreateLoginByClientIDUseCase(createLoginByClientIDRepository contract.CreateLoginByClientIDRepository) contract.CreateLoginByClientIDUseCase {
	return &createLoginByClientIDUseCase{
		createLoginByClientIDRepository: createLoginByClientIDRepository,
	}
}
