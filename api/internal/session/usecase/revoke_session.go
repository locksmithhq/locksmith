package usecase

import (
	"context"

	"github.com/locksmithhq/locksmith/api/internal/session/contract"
)

type revokeSessionUseCase struct {
	revokeSessionRepository contract.RevokeSessionRepository
}

func (u *revokeSessionUseCase) Execute(ctx context.Context, projectID, sessionID string) error {
	return u.revokeSessionRepository.Execute(ctx, projectID, sessionID)
}

func NewRevokeSessionUseCase(revokeSessionRepository contract.RevokeSessionRepository) contract.RevokeSessionUseCase {
	return &revokeSessionUseCase{revokeSessionRepository: revokeSessionRepository}
}
