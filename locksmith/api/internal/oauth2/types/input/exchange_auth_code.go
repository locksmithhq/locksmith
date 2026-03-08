package input

type ExchangeAuthCode struct {
	Code            string `json:"code"`
	ClientID        string `json:"client_id"`
	RedirectURI     string `json:"redirect_uri"`
	IPAddress       string `json:"ip_address"`
	UserAgent       string `json:"user_agent"`
	DeviceName      string `json:"device_name"`
	DeviceType      string `json:"device_type"`
	Browser         string `json:"browser"`
	BrowserVersion  string `json:"browser_version"`
	OS              string `json:"os"`
	OSVersion       string `json:"os_version"`
	LocationCountry string `json:"location_country"`
	LocationRegion  string `json:"location_region"`
	LocationCity    string `json:"location_city"`
	RefreshToken    string
	CodeVerifier    string
	DeviceID        string
}
