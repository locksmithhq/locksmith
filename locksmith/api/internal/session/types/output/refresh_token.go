package output

import "github.com/locksmithhq/locksmith/api/internal/session/domain"

type RefreshToken struct {
	ID            string `json:"id"`
	SessionID     string `json:"session_id"`
	AccountID     string `json:"account_id"`
	ClientID      string `json:"client_id"`
	RotationCount int    `json:"rotation_count"`
	ParentTokenID string `json:"parent_token_id"`
	ExpiresAt     string `json:"expires_at"`
	Revoked       bool   `json:"revoked"`
	RevokedAt     string `json:"revoked_at"`
	RevokedReason string `json:"revoked_reason"`
	LastUsedAt    string `json:"last_used_at"`
	CreatedAt     string `json:"created_at"`
	ClientName    string `json:"client_name"`
}

func NewRefreshTokenFromDomain(rt domain.RefreshToken) RefreshToken {
	return RefreshToken{
		ID:            rt.ID,
		SessionID:     rt.SessionID,
		AccountID:     rt.AccountID,
		ClientID:      rt.ClientID,
		RotationCount: rt.RotationCount,
		ParentTokenID: rt.ParentTokenID,
		ExpiresAt:     rt.ExpiresAt,
		Revoked:       rt.Revoked,
		RevokedAt:     rt.RevokedAt,
		RevokedReason: rt.RevokedReason,
		LastUsedAt:    rt.LastUsedAt,
		CreatedAt:     rt.CreatedAt,
		ClientName:    rt.ClientName,
	}
}

func NewRefreshTokensFromDomain(tokens []domain.RefreshToken) []RefreshToken {
	out := make([]RefreshToken, len(tokens))
	for i, rt := range tokens {
		out[i] = NewRefreshTokenFromDomain(rt)
	}
	return out
}
