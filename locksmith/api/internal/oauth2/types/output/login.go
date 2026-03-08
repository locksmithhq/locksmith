package output

type Login struct {
	RedirectTo         string `json:"redirect_to,omitempty"`
	MustChangePassword bool   `json:"must_change_password"`
	ChangePasswordJWT  string `json:"change_password_jwt,omitempty"`
}
