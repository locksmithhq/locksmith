package output

import "github.com/locksmithhq/locksmith/api/internal/session/domain"

type Session struct {
	ID              string `json:"id"`
	AccountID       string `json:"account_id"`
	ClientID        string `json:"client_id"`
	IPAddress       string `json:"ip_address"`
	DeviceType      string `json:"device_type"`
	Browser         string `json:"browser"`
	BrowserVersion  string `json:"browser_version"`
	OS              string `json:"os"`
	OSVersion       string `json:"os_version"`
	LocationCountry string `json:"location_country"`
	LocationRegion  string `json:"location_region"`
	LocationCity    string `json:"location_city"`
	Revoked         bool   `json:"revoked"`
	RevokedReason   string `json:"revoked_reason"`
	ExpiresAt       string `json:"expires_at"`
	LastActivity    string `json:"last_activity"`
	CreatedAt       string `json:"created_at"`
	AccountName     string `json:"account_name"`
	AccountEmail    string `json:"account_email"`
	ClientName      string `json:"client_name"`
}

func NewSessionFromDomain(s domain.Session) Session {
	return Session{
		ID:              s.ID,
		AccountID:       s.AccountID,
		ClientID:        s.ClientID,
		IPAddress:       s.IPAddress,
		DeviceType:      s.DeviceType,
		Browser:         s.Browser,
		BrowserVersion:  s.BrowserVersion,
		OS:              s.OS,
		OSVersion:       s.OSVersion,
		LocationCountry: s.LocationCountry,
		LocationRegion:  s.LocationRegion,
		LocationCity:    s.LocationCity,
		Revoked:         s.Revoked,
		RevokedReason:   s.RevokedReason,
		ExpiresAt:       s.ExpiresAt,
		LastActivity:    s.LastActivity,
		CreatedAt:       s.CreatedAt,
		AccountName:     s.AccountName,
		AccountEmail:    s.AccountEmail,
		ClientName:      s.ClientName,
	}
}

func NewSessionsFromDomain(sessions []domain.Session) []Session {
	out := make([]Session, len(sessions))
	for i, s := range sessions {
		out[i] = NewSessionFromDomain(s)
	}
	return out
}
