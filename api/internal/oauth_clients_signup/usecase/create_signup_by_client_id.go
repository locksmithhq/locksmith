package usecase

import (
	"context"

	"github.com/locksmithhq/locksmith/api/internal/oauth_clients_signup/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients_signup/types/input"
)

type createSignupByClientIDUseCase struct {
	createSignupByClientIDRepository contract.CreateSignupByClientIDRepository
}

// Execute implements contract.CreateSignupByClientIDUseCase.
func (u *createSignupByClientIDUseCase) Execute(ctx context.Context, clientID string, in input.Signup) error {
	signupDomain := in.ToSignupDomain()
	signupDomain.ClientID = clientID
	return u.createSignupByClientIDRepository.Execute(ctx, signupDomain)
}

func NewCreateSignupByClientIDUseCase(createSignupByClientIDRepository contract.CreateSignupByClientIDRepository) contract.CreateSignupByClientIDUseCase {
	return &createSignupByClientIDUseCase{
		createSignupByClientIDRepository: createSignupByClientIDRepository,
	}
}
