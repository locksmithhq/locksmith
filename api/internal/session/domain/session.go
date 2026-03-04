package domain

type Session struct {
	ID              string `json:"id" db:"id"`
	AccountID       string `json:"account_id" db:"account_id"`
	ClientID        string `json:"client_id" db:"client_id"`
	IPAddress       string `json:"ip_address" db:"ip_address"`
	DeviceType      string `json:"device_type" db:"device_type"`
	Browser         string `json:"browser" db:"browser"`
	BrowserVersion  string `json:"browser_version" db:"browser_version"`
	OS              string `json:"os" db:"os"`
	OSVersion       string `json:"os_version" db:"os_version"`
	LocationCountry string `json:"location_country" db:"location_country"`
	LocationCity    string `json:"location_city" db:"location_city"`
	Revoked         bool   `json:"revoked" db:"revoked"`
	RevokedReason   string `json:"revoked_reason" db:"revoked_reason"`
	ExpiresAt       string `json:"expires_at" db:"expires_at"`
	LastActivity    string `json:"last_activity" db:"last_activity"`
	CreatedAt       string `json:"created_at" db:"created_at"`
	AccountName     string `json:"account_name" db:"account_name"`
	AccountEmail    string `json:"account_email" db:"account_email"`
	ClientName      string `json:"client_name" db:"client_name"`
}
