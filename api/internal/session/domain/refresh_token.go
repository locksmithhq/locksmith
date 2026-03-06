package domain

type RefreshToken struct {
	ID            string `json:"id" db:"id"`
	SessionID     string `json:"session_id" db:"session_id"`
	AccountID     string `json:"account_id" db:"account_id"`
	ClientID      string `json:"client_id" db:"client_id"`
	RotationCount int    `json:"rotation_count" db:"rotation_count"`
	ParentTokenID string `json:"parent_token_id" db:"parent_token_id"`
	ExpiresAt     string `json:"expires_at" db:"expires_at"`
	Revoked       bool   `json:"revoked" db:"revoked"`
	RevokedAt     string `json:"revoked_at" db:"revoked_at"`
	RevokedReason string `json:"revoked_reason" db:"revoked_reason"`
	LastUsedAt    string `json:"last_used_at" db:"last_used_at"`
	CreatedAt     string `json:"created_at" db:"created_at"`
	ClientName    string `json:"client_name" db:"client_name"`
}
