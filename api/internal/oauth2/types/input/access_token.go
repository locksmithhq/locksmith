package input

type AccessToken struct {
	Code            string `json:"code"`
	ClientID        string `json:"client_id"`
	ClientSecret    string `json:"client_secret"`
	GrantType       string `json:"grant_type"`
	DeviceName      string `json:"device_name"`
	CodeVerifier    string `json:"code_verifier"`
	DeviceID        string
	IPAddress       string
	UserAgent       string
	DeviceType      string
	Browser         string
	BrowserVersion  string
	OS              string
	OSVersion       string
	LocationCountry string
	LocationRegion  string
	LocationCity    string
}
