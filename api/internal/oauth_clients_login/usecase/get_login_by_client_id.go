package usecase

import (
	"context"

	"github.com/booscaaa/locksmith/api/internal/core/types/stackerror"
	"github.com/booscaaa/locksmith/api/internal/oauth_clients_login/contract"
	"github.com/booscaaa/locksmith/api/internal/oauth_clients_login/domain"
)

type getLoginByClientIDUseCase struct {
	getLoginByClientIDRepository contract.GetLoginByClientIDRepository
}

// Execute implements contract.GetLoginByClientIDUseCase.
func (u *getLoginByClientIDUseCase) Execute(ctx context.Context, clientID string) (domain.Login, error) {
	login, err := u.getLoginByClientIDRepository.Execute(ctx, clientID)
	if err != nil {
		return domain.Login{}, stackerror.NewUseCaseError("GetLoginByClientIDUseCase", err)
	}

	return login, nil
}

func NewGetLoginByClientIDUseCase(
	getLoginByClientIDRepository contract.GetLoginByClientIDRepository,
) contract.GetLoginByClientIDUseCase {
	return &getLoginByClientIDUseCase{
		getLoginByClientIDRepository: getLoginByClientIDRepository,
	}
}
