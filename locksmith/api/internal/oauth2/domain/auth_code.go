package domain

type AuthCode struct {
	ID                  string `json:"id" db:"id"`
	Code                string `json:"code" db:"code"`
	ClientID            string `json:"client_id" db:"client_id"`
	AccountID           string `json:"account_id" db:"account_id"`
	RedirectURI         string `json:"redirect_uri" db:"redirect_uri"`
	CodeChallenge       string `db:"code_challenge"`
	CodeChallengeMethod string `db:"code_challenge_method"`
	ExpiresAt           string `json:"expires_at" db:"expires_at"`
	Used                bool   `json:"used" db:"used"`
	CreatedAt           string `json:"created_at" db:"created_at"`
}

func NewAuthCode(
	id string,
	code string,
	clientID string,
	accountID string,
	redirectURI string,
	codeChallenge string,
	codeChallengeMethod string,
	expiresAt string,
	used bool,
	createdAt string,
) AuthCode {
	return AuthCode{
		ID:                  id,
		Code:                code,
		ClientID:            clientID,
		AccountID:           accountID,
		RedirectURI:         redirectURI,
		CodeChallenge:       codeChallenge,
		CodeChallengeMethod: codeChallengeMethod,
		ExpiresAt:           expiresAt,
		Used:                used,
		CreatedAt:           createdAt,
	}
}
