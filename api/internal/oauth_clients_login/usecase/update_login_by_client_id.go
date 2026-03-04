package usecase

import (
	"context"

	"github.com/locksmithhq/locksmith/api/internal/oauth_clients_login/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients_login/types/input"
)

type updateLoginByClientIDUseCase struct {
	updateLoginByClientIDRepository contract.UpdateLoginByClientIDRepository
}

// Execute implements contract.UpdateLoginByClientIDUseCase.
func (u *updateLoginByClientIDUseCase) Execute(ctx context.Context, clientID string, in input.Login) error {
	loginDomain := in.ToLoginDomain()
	loginDomain.ClientID = clientID
	return u.updateLoginByClientIDRepository.Execute(ctx, loginDomain)
}

func NewUpdateLoginByClientIDUseCase(updateLoginByClientIDRepository contract.UpdateLoginByClientIDRepository) contract.UpdateLoginByClientIDUseCase {
	return &updateLoginByClientIDUseCase{
		updateLoginByClientIDRepository: updateLoginByClientIDRepository,
	}
}
