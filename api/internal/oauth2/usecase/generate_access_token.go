package usecase

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"

	"github.com/locksmithhq/locksmith/api/internal/core/types/database"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/domain"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/types/input"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/types/output"
	"github.com/google/uuid"
	"github.com/locksmithhq/locksmith-go"
)

type generateAccessTokenUseCase struct {
	getAuthCodeByCodeRepository          contract.GetAuthCodeByCodeRepository
	updateAuthCodeRepository             contract.UpdateAuthCodeRepository
	getClientByClientIDRepository        contract.GetClientByClientIDRepository
	createUserSessionRepository          contract.CreateUserSessionRepository
	createRefreshTokenRepository         contract.CreateRefreshTokenRepository
	getUserSessionByDeviceRepository     contract.GetUserSessionByDeviceRepository
	updateUserSessionActivityRepository  contract.UpdateUserSessionActivityRepository
}

// Execute implements contract.GenerateAccessTokenUseCase.
func (u *generateAccessTokenUseCase) Execute(ctx context.Context, in input.AccessToken) (output.AccessToken, error) {
	authCode, err := u.getAuthCodeByCodeRepository.Execute(ctx, in.Code)
	if err != nil {
		return output.AccessToken{}, stackerror.NewUseCaseError(
			"GenerateAccessTokenUseCase",
			err,
			stackerror.WithMessage("invalid authorization code"),
			stackerror.WithStatusCode(http.StatusBadRequest),
		)
	}

	if authCode.Used {
		return output.AccessToken{}, stackerror.NewUseCaseError(
			"GenerateAccessTokenUseCase",
			fmt.Errorf("authorization code already used"),
			stackerror.WithMessage("authorization code already used"),
			stackerror.WithStatusCode(http.StatusBadRequest),
		)
	}

	expiresAt, err := time.Parse(time.RFC3339, authCode.ExpiresAt)
	if err != nil {
		return output.AccessToken{}, stackerror.NewUseCaseError(
			"GenerateAccessTokenUseCase",
			err,
			stackerror.WithMessage("invalid expiration time"),
			stackerror.WithStatusCode(http.StatusBadRequest),
		)
	}

	if time.Now().After(expiresAt) {
		return output.AccessToken{}, stackerror.NewUseCaseError(
			"GenerateAccessTokenUseCase",
			fmt.Errorf("authorization code expired"),
			stackerror.WithMessage("authorization code expired"),
			stackerror.WithStatusCode(http.StatusBadRequest),
		)
	}

	client, err := u.getClientByClientIDRepository.Execute(ctx, in.ClientID)
	if err != nil {
		return output.AccessToken{}, stackerror.NewUseCaseError(
			"GenerateAccessTokenUseCase",
			err,
			stackerror.WithMessage("invalid client_id"),
			stackerror.WithStatusCode(http.StatusBadRequest),
		)
	}

	if client.ID != authCode.ClientID {
		return output.AccessToken{}, stackerror.NewUseCaseError(
			"GenerateAccessTokenUseCase",
			fmt.Errorf("client_id mismatch"),
			stackerror.WithMessage("client_id mismatch"),
			stackerror.WithStatusCode(http.StatusBadRequest),
		)
	}

	// 5. Mark the auth code as used
	authCode.Used = true
	err = u.updateAuthCodeRepository.Execute(ctx, authCode)
	if err != nil {
		return output.AccessToken{}, stackerror.NewUseCaseError(
			"GenerateAccessTokenUseCase",
			err,
			stackerror.WithMessage("failed to update auth code"),
			stackerror.WithStatusCode(http.StatusInternalServerError),
		)
	}

	// 6. Validate PKCE
	if authCode.CodeChallenge != "" {
		if in.CodeVerifier == "" {
			return output.AccessToken{}, stackerror.NewUseCaseError(
				"GenerateAccessTokenUseCase",
				fmt.Errorf("code_verifier required"),
				stackerror.WithMessage("code_verifier required"),
				stackerror.WithStatusCode(http.StatusBadRequest),
			)
		}

		// Calculate hash of code_verifier
		sha256Verifier := sha256.Sum256([]byte(in.CodeVerifier))
		encodedVerifier := base64.RawURLEncoding.EncodeToString(sha256Verifier[:])

		if encodedVerifier != authCode.CodeChallenge {
			return output.AccessToken{}, stackerror.NewUseCaseError(
				"GenerateAccessTokenUseCase",
				fmt.Errorf("invalid code_verifier"),
				stackerror.WithMessage("invalid code_verifier"),
				stackerror.WithStatusCode(http.StatusBadRequest),
			)
		}
	}

	// 7. Create or Reuse User Session
	var userSession domain.UserSession

	// Check if session exists for device
	if in.DeviceID != "" {
		existingSession, err := u.getUserSessionByDeviceRepository.Execute(ctx, authCode.AccountID, client.ID, in.DeviceID)
		if err == nil {
			userSession = existingSession
		}
	}

	if userSession.ID != "" {
		_ = u.updateUserSessionActivityRepository.Execute(ctx, userSession.ID)
	}

	if userSession.ID == "" {
		jti := uuid.New().String()
		sessionExpiresAt := time.Now().Add(24 * time.Hour).Format(time.RFC3339) // 24 hours session

		newUserSession := domain.UserSession{
			AccountID:       authCode.AccountID,
			ClientID:        client.ID,
			JTI:             jti,
			IPAddress:       database.ParseNull(in.IPAddress),
			UserAgent:       database.ParseNull(in.UserAgent),
			DeviceName:      database.ParseNull(in.DeviceName),
			DeviceID:        database.ParseNull(in.DeviceID),
			DeviceType:      database.ParseNull(in.DeviceType),
			Browser:         database.ParseNull(in.Browser),
			BrowserVersion:  database.ParseNull(in.BrowserVersion),
			OS:              database.ParseNull(in.OS),
			OSVersion:       database.ParseNull(in.OSVersion),
			LocationCountry: database.ParseNull(in.LocationCountry),
			LocationRegion:  database.ParseNull(in.LocationRegion),
			LocationCity:    database.ParseNull(in.LocationCity),
			ExpiresAt:       sessionExpiresAt,
		}

		userSession, err = u.createUserSessionRepository.Execute(ctx, newUserSession)
		if err != nil {
			return output.AccessToken{}, stackerror.NewUseCaseError(
				"GenerateAccessTokenUseCase",
				err,
				stackerror.WithMessage("failed to create user session"),
				stackerror.WithStatusCode(http.StatusInternalServerError),
			)
		}
	}

	// 7. Generate and Save Refresh Token
	refreshTokenString := uuid.New().String()
	hasher := sha256.New()
	hasher.Write([]byte(refreshTokenString))
	tokenHash := hex.EncodeToString(hasher.Sum(nil))

	refreshTokenExpiresAt := time.Now().Add(30 * 24 * time.Hour).Format(time.RFC3339) // 30 days
	refreshToken := domain.RefreshToken{
		TokenHash: tokenHash,
		SessionID: database.ParseNull(userSession.ID),
		AccountID: authCode.AccountID,
		ClientID:  client.ID,
		ExpiresAt: refreshTokenExpiresAt,
	}

	_, err = u.createRefreshTokenRepository.Execute(ctx, refreshToken)
	if err != nil {
		return output.AccessToken{}, stackerror.NewUseCaseError(
			"GenerateAccessTokenUseCase",
			err,
			stackerror.WithMessage("failed to create refresh token"),
			stackerror.WithStatusCode(http.StatusInternalServerError),
		)
	}

	// 8. Generate JWT access token
	claims := locksmith.NewTokenClaims(authCode.AccountID, client.ClientID, client.Domain)
	accessToken := locksmith.GetSignToken(claims, 5*time.Minute, client.ClientSecret)

	// 10. Return tokens
	return output.AccessToken{
		AccessToken:  accessToken,
		RefreshToken: refreshTokenString,
		TokenType:    "Bearer",
		ExpiresIn:    int((5 * time.Minute).Seconds()),
	}, nil
}

func NewGenerateAccessTokenUseCase(
	getAuthCodeByCodeRepository contract.GetAuthCodeByCodeRepository,
	updateAuthCodeRepository contract.UpdateAuthCodeRepository,
	getClientByClientIDRepository contract.GetClientByClientIDRepository,
	createUserSessionRepository contract.CreateUserSessionRepository,
	createRefreshTokenRepository contract.CreateRefreshTokenRepository,
	getUserSessionByDeviceRepository contract.GetUserSessionByDeviceRepository,
	updateUserSessionActivityRepository contract.UpdateUserSessionActivityRepository,
) contract.GenerateAccessTokenUseCase {
	return &generateAccessTokenUseCase{
		getAuthCodeByCodeRepository:         getAuthCodeByCodeRepository,
		updateAuthCodeRepository:            updateAuthCodeRepository,
		getClientByClientIDRepository:       getClientByClientIDRepository,
		createUserSessionRepository:         createUserSessionRepository,
		createRefreshTokenRepository:        createRefreshTokenRepository,
		getUserSessionByDeviceRepository:    getUserSessionByDeviceRepository,
		updateUserSessionActivityRepository: updateUserSessionActivityRepository,
	}
}
