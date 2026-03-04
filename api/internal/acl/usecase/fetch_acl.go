package usecase

import (
	"context"

	"github.com/booscaaa/go-paginate/v3/paginate"
	"github.com/booscaaa/locksmith/api/internal/acl/contract"
	"github.com/booscaaa/locksmith/api/internal/acl/domain"
)

type fetchAclUseCase struct {
	fetchAclRepository contract.FetchAclRepository
}

// Execute implements contract.FetchAclUseCase.
func (u *fetchAclUseCase) Execute(ctx context.Context, filter paginate.PaginationParams) ([]domain.Acl, error) {
	return u.fetchAclRepository.Execute(ctx, filter)
}

func NewFetchAclUseCase(fetchAclRepository contract.FetchAclRepository) contract.FetchAclUseCase {
	return &fetchAclUseCase{fetchAclRepository: fetchAclRepository}
}
