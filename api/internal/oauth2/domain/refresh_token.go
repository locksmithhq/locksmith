package domain

import "github.com/booscaaa/locksmith/api/internal/core/types/database"

type RefreshToken struct {
	ID            string        `json:"id" db:"id"`
	TokenHash     string        `json:"token_hash" db:"token_hash"`
	SessionID     database.Null `json:"session_id" db:"session_id"`
	AccountID     string        `json:"account_id" db:"account_id"`
	ClientID      string        `json:"client_id" db:"client_id"`
	RotationCount int           `json:"rotation_count" db:"rotation_count"`
	ParentTokenID database.Null `json:"parent_token_id" db:"parent_token_id"`
	ExpiresAt     string        `json:"expires_at" db:"expires_at"`
	Revoked       bool          `json:"revoked" db:"revoked"`
	RevokedAt     database.Null `json:"revoked_at" db:"revoked_at"`
	RevokedReason database.Null `json:"revoked_reason" db:"revoked_reason"`
	LastUsedAt    database.Null `json:"last_used_at" db:"last_used_at"`
	CreatedAt     string        `json:"created_at" db:"created_at"`
}
