package usecase

import (
	"context"

	"github.com/booscaaa/locksmith/api/internal/oauth_clients_signup/contract"
	"github.com/booscaaa/locksmith/api/internal/oauth_clients_signup/domain"
)

type getSignupByClientIDUseCase struct {
	getSignupByClientIDRepository contract.GetSignupByClientIDRepository
}

// Execute implements contract.GetSignupByClientIDUseCase.
func (u *getSignupByClientIDUseCase) Execute(ctx context.Context, clientID string) (domain.Signup, error) {
	signup, err := u.getSignupByClientIDRepository.Execute(ctx, clientID)
	if err != nil {
		return domain.Signup{}, err
	}

	return signup, nil
}

func NewGetSignupByClientIDUseCase(
	getSignupByClientIDRepository contract.GetSignupByClientIDRepository,
) contract.GetSignupByClientIDUseCase {
	return &getSignupByClientIDUseCase{
		getSignupByClientIDRepository: getSignupByClientIDRepository,
	}
}
