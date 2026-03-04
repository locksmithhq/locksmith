package usecase

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"

	"github.com/booscaaa/locksmith/api/internal/core/types/database"
	"github.com/booscaaa/locksmith/api/internal/core/types/stackerror"
	"github.com/booscaaa/locksmith/api/internal/oauth2/contract"
	"github.com/booscaaa/locksmith/api/internal/oauth2/domain"
	"github.com/booscaaa/locksmith/api/internal/oauth2/types/output"
	"github.com/google/uuid"
	"github.com/locksmithhq/locksmith-go"
)

type generateRefreshTokenUseCase struct {
	getRefreshTokenByHashRepository contract.GetRefreshTokenByHashRepository
	updateRefreshTokenRepository    contract.UpdateRefreshTokenRepository
	createRefreshTokenRepository    contract.CreateRefreshTokenRepository
	getClientByClientIDRepository   contract.GetClientByClientIDRepository
	getClientByIDRepository         contract.GetClientByIDRepository
}

// Execute implements contract.GenerateRefreshTokenUseCase.
func (u *generateRefreshTokenUseCase) Execute(ctx context.Context, refreshToken string) (output.AccessToken, error) {
	// 1. Hash the incoming refresh token
	hasher := sha256.New()
	hasher.Write([]byte(refreshToken))
	tokenHash := hex.EncodeToString(hasher.Sum(nil))

	// 2. Fetch the refresh token from DB
	storedToken, err := u.getRefreshTokenByHashRepository.Execute(ctx, tokenHash)
	if err != nil {
		return output.AccessToken{}, fmt.Errorf("invalid refresh token")
	}

	// 3. Validate token
	if storedToken.Revoked {
		// TODO: Security Alert! Attempt to use revoked token.
		// Ideally, we should revoke the entire chain here.
		return output.AccessToken{}, fmt.Errorf("refresh token revoked")
	}

	expiresAt, err := time.Parse(time.RFC3339, storedToken.ExpiresAt)
	if err != nil {
		return output.AccessToken{}, fmt.Errorf("invalid expiration time")
	}

	if time.Now().After(expiresAt) {
		return output.AccessToken{}, fmt.Errorf("refresh token expired")
	}

	// 4. Rotation: Revoke the old token
	revokedReason := "rotation"
	revokedAt := time.Now().Format(time.RFC3339)
	storedToken.Revoked = true
	storedToken.RevokedAt = database.ParseNull(revokedAt)
	storedToken.RevokedReason = database.ParseNull(revokedReason)
	storedToken.LastUsedAt = database.ParseNull(revokedAt) // Update last used

	err = u.updateRefreshTokenRepository.Execute(ctx, storedToken)
	if err != nil {
		return output.AccessToken{}, fmt.Errorf("failed to revoke old refresh token")
	}

	// 5. Create a new Refresh Token
	newRefreshTokenString := uuid.New().String()
	hasher.Reset()
	hasher.Write([]byte(newRefreshTokenString))
	newTokenHash := hex.EncodeToString(hasher.Sum(nil))

	newExpiresAt := time.Now().Add(30 * 24 * time.Hour).Format(time.RFC3339)
	newRotationCount := storedToken.RotationCount + 1

	newRefreshToken := domain.RefreshToken{
		TokenHash:     newTokenHash,
		SessionID:     storedToken.SessionID,
		AccountID:     storedToken.AccountID,
		ClientID:      storedToken.ClientID,
		RotationCount: newRotationCount,
		ParentTokenID: database.ParseNull(storedToken.ID),
		ExpiresAt:     newExpiresAt,
	}

	_, err = u.createRefreshTokenRepository.Execute(ctx, newRefreshToken)
	if err != nil {
		return output.AccessToken{}, fmt.Errorf("failed to create new refresh token")
	}

	// 6. Fetch Client (needed for claims)
	client, err := u.getClientByIDRepository.Execute(ctx, storedToken.ClientID)
	if err != nil {
		return output.AccessToken{}, stackerror.NewUseCaseError(
			"GenerateAccessTokenUseCase",
			err,
			stackerror.WithMessage("invalid client_id"),
			stackerror.WithStatusCode(http.StatusBadRequest),
		)
	}

	claims := locksmith.NewTokenClaims(storedToken.AccountID, client.ClientID, client.Domain)
	accessToken := locksmith.GetSignToken(claims, 5*time.Minute, client.ClientSecret)

	return output.AccessToken{
		AccessToken:  accessToken,
		RefreshToken: newRefreshTokenString,
		TokenType:    "Bearer",
		ExpiresIn:    int((5 * time.Minute).Seconds()),
	}, nil
}

func NewGenerateRefreshTokenUseCase(
	getRefreshTokenByHashRepository contract.GetRefreshTokenByHashRepository,
	updateRefreshTokenRepository contract.UpdateRefreshTokenRepository,
	createRefreshTokenRepository contract.CreateRefreshTokenRepository,
	getClientByClientIDRepository contract.GetClientByClientIDRepository,
	getClientByIDRepository contract.GetClientByIDRepository,
) contract.GenerateRefreshTokenUseCase {
	return &generateRefreshTokenUseCase{
		getRefreshTokenByHashRepository: getRefreshTokenByHashRepository,
		updateRefreshTokenRepository:    updateRefreshTokenRepository,
		createRefreshTokenRepository:    createRefreshTokenRepository,
		getClientByClientIDRepository:   getClientByClientIDRepository,
		getClientByIDRepository:         getClientByIDRepository,
	}
}
