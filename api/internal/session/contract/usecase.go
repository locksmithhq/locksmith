package contract

import (
	"context"

	"github.com/locksmithhq/locksmith/api/internal/session/types/output"
)

type FetchSessionsByProjectIDUseCase interface {
	Execute(ctx context.Context, projectID string, page, limit int, search string) ([]output.Session, error)
}

type CountSessionsByProjectIDUseCase interface {
	Execute(ctx context.Context, projectID string, search string) (int64, error)
}
