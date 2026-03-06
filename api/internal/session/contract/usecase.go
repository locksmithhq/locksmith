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

type FetchSessionsByAccountIDUseCase interface {
	Execute(ctx context.Context, projectID, accountID string, page, limit int) ([]output.Session, error)
}

type CountSessionsByAccountIDUseCase interface {
	Execute(ctx context.Context, projectID, accountID string) (int64, error)
}

type RevokeSessionUseCase interface {
	Execute(ctx context.Context, projectID, sessionID string) error
}

type FetchRefreshTokensByAccountIDUseCase interface {
	Execute(ctx context.Context, projectID, accountID, sessionID string, page, limit int) ([]output.RefreshToken, error)
}

type CountRefreshTokensByAccountIDUseCase interface {
	Execute(ctx context.Context, projectID, accountID, sessionID string) (int64, error)
}
