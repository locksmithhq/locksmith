package contract

import (
	"context"

	"github.com/locksmithhq/locksmith/api/internal/session/domain"
)

type FetchSessionsByProjectIDRepository interface {
	Execute(ctx context.Context, projectID string, page, limit int, search string) ([]domain.Session, error)
}

type CountSessionsByProjectIDRepository interface {
	Execute(ctx context.Context, projectID string, search string) (int64, error)
}
