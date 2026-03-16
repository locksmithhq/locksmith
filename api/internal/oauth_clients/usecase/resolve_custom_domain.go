package usecase

import (
	"context"

	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients/contract"
)

type resolveCustomDomainUseCase struct {
	getClientByCustomDomainRepository contract.GetClientByCustomDomainRepository
}

func (u *resolveCustomDomainUseCase) Execute(ctx context.Context, hostname string) (string, string, error) {
	clientID, redirectURI, err := u.getClientByCustomDomainRepository.Execute(ctx, hostname)
	if err != nil {
		return "", "", stackerror.NewUseCaseError("ResolveCustomDomainUseCase", err)
	}

	return clientID, redirectURI, nil
}

func NewResolveCustomDomainUseCase(
	getClientByCustomDomainRepository contract.GetClientByCustomDomainRepository,
) contract.ResolveCustomDomainUseCase {
	return &resolveCustomDomainUseCase{
		getClientByCustomDomainRepository: getClientByCustomDomainRepository,
	}
}
