package usecase

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"time"

	"github.com/locksmithhq/locksmith/api/internal/core/types/database"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/contract"
)

type logoutUseCase struct {
	getRefreshTokenByHash contract.GetRefreshTokenByHashRepository
	updateRefreshToken    contract.UpdateRefreshTokenRepository
	revokeUserSession     contract.RevokeUserSessionRepository
}

func (u *logoutUseCase) Execute(ctx context.Context, rawRefreshToken string) error {
	hasher := sha256.New()
	hasher.Write([]byte(rawRefreshToken))
	tokenHash := hex.EncodeToString(hasher.Sum(nil))

	storedToken, err := u.getRefreshTokenByHash.Execute(ctx, tokenHash)
	if err != nil {
		if errors.Is(err, stackerror.ErrNotFound) {
			return nil // token not found — cookie was already invalid, nothing to revoke
		}
		return err
	}

	if !storedToken.Revoked {
		now := time.Now().Format(time.RFC3339)
		storedToken.Revoked = true
		storedToken.RevokedAt = database.ParseNull(now)
		storedToken.RevokedReason = database.ParseNull("user_logout")
		storedToken.LastUsedAt = database.ParseNull(now)

		if err := u.updateRefreshToken.Execute(ctx, storedToken); err != nil {
			return err
		}
	}

	if !storedToken.SessionID.IsNull() {
		_ = u.revokeUserSession.Execute(ctx, storedToken.SessionID.String())
	}

	return nil
}

func NewLogoutUseCase(
	getRefreshTokenByHash contract.GetRefreshTokenByHashRepository,
	updateRefreshToken contract.UpdateRefreshTokenRepository,
	revokeUserSession contract.RevokeUserSessionRepository,
) contract.LogoutUseCase {
	return &logoutUseCase{
		getRefreshTokenByHash: getRefreshTokenByHash,
		updateRefreshToken:    updateRefreshToken,
		revokeUserSession:     revokeUserSession,
	}
}
