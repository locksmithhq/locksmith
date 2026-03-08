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

type FetchSessionsByAccountIDRepository interface {
	Execute(ctx context.Context, projectID, accountID string, page, limit int) ([]domain.Session, error)
}

type CountSessionsByAccountIDRepository interface {
	Execute(ctx context.Context, projectID, accountID string) (int64, error)
}

type RevokeSessionRepository interface {
	Execute(ctx context.Context, projectID, sessionID string) error
}

type FetchRefreshTokensByAccountIDRepository interface {
	Execute(ctx context.Context, projectID, accountID, sessionID string, page, limit int) ([]domain.RefreshToken, error)
}

type CountRefreshTokensByAccountIDRepository interface {
	Execute(ctx context.Context, projectID, accountID, sessionID string) (int64, error)
}
