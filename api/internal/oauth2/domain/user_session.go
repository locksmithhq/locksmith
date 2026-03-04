package domain

import "github.com/booscaaa/locksmith/api/internal/core/types/database"

type UserSession struct {
	ID              string        `json:"id" db:"id"`
	AccountID       string        `json:"account_id" db:"account_id"`
	ClientID        string        `json:"client_id" db:"client_id"`
	JTI             string        `json:"jti" db:"jti"`
	IPAddress       database.Null `json:"ip_address" db:"ip_address"`
	UserAgent       database.Null `json:"user_agent" db:"user_agent"`
	DeviceName      database.Null `json:"device_name" db:"device_name"`
	DeviceID        database.Null `json:"device_id" db:"device_id"`
	DeviceType      database.Null `json:"device_type" db:"device_type"`
	Browser         database.Null `json:"browser" db:"browser"`
	BrowserVersion  database.Null `json:"browser_version" db:"browser_version"`
	OS              database.Null `json:"os" db:"os"`
	OSVersion       database.Null `json:"os_version" db:"os_version"`
	LocationCountry database.Null `json:"location_country" db:"location_country"`
	LocationRegion  database.Null `json:"location_region" db:"location_region"`
	LocationCity    database.Null `json:"location_city" db:"location_city"`
	ExpiresAt       string        `json:"expires_at" db:"expires_at"`
	Revoked         bool          `json:"revoked" db:"revoked"`
	RevokedAt       database.Null `json:"revoked_at" db:"revoked_at"`
	RevokedReason   database.Null `json:"revoked_reason" db:"revoked_reason"`
	LastActivity    database.Null `json:"last_activity" db:"last_activity"`
	CreatedAt       string        `json:"created_at" db:"created_at"`
}
