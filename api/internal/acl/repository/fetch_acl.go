package repository

import (
	"context"

	"github.com/booscaaa/go-paginate/v3/paginate"
	"github.com/booscaaa/initializers/postgres/types"
	"github.com/booscaaa/locksmith/api/internal/acl/contract"
	"github.com/booscaaa/locksmith/api/internal/acl/domain"
	"github.com/booscaaa/locksmith/api/internal/core/types/stackerror"
)

type fetchAclRepository struct {
	database types.Database
}

// FetchAll implements contract.FetchAclRepository.
func (r *fetchAclRepository) Execute(ctx context.Context, filter paginate.PaginationParams) ([]domain.Acl, error) {
	var roles []domain.Acl
	query, args, err := paginate.NewBuilder().
		Model(&domain.Acl{}).
		Table("locksmith_rules").
		Select("id", "p_type", "v0", "v1", "v2", "v3", "v4", "v5").
		FromStruct(filter).
		BuildSQL()
	if err != nil {
		return nil, stackerror.NewRepositoryError("FetchAll", err)
	}

	err = r.database.SelectContext(ctx, &roles, query, args...)
	if err != nil {
		return nil, stackerror.NewRepositoryError("FetchAll", err)
	}

	return roles, nil
}

func NewFetchAclRepository(database types.Database) contract.FetchAclRepository {
	return &fetchAclRepository{database: database}
}
