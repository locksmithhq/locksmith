package input

import (
	"net/http"

	"github.com/booscaaa/initializers/validator"
)

type ChangePassword struct {
	ClientID        string `json:"client_id" validate:"required"`
	Password        string `json:"password" validate:"required,min=8"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
	Jwt             string `json:"-" validate:"required"`
}

func (a ChangePassword) ValidateHttp(w http.ResponseWriter) error {
	if err := validator.ValidateStruct(&a); err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	return nil
}
