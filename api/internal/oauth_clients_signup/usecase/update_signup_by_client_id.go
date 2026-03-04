package usecase

import (
	"context"

	"github.com/booscaaa/locksmith/api/internal/oauth_clients_signup/contract"
	"github.com/booscaaa/locksmith/api/internal/oauth_clients_signup/types/input"
)

type updateSignupByClientIDUseCase struct {
	updateSignupByClientIDRepository contract.UpdateSignupByClientIDRepository
}

// Execute implements contract.UpdateSignupByClientIDUseCase.
func (u *updateSignupByClientIDUseCase) Execute(ctx context.Context, clientID string, in input.Signup) error {
	signupDomain := in.ToSignupDomain()
	signupDomain.ClientID = clientID
	return u.updateSignupByClientIDRepository.Execute(ctx, signupDomain)
}

func NewUpdateSignupByClientIDUseCase(updateSignupByClientIDRepository contract.UpdateSignupByClientIDRepository) contract.UpdateSignupByClientIDUseCase {
	return &updateSignupByClientIDUseCase{
		updateSignupByClientIDRepository: updateSignupByClientIDRepository,
	}
}
