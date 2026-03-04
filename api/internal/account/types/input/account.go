package input

import (
	"net/http"

	"github.com/booscaaa/initializers/validator"
	"github.com/booscaaa/locksmith/api/internal/account/domain"
)

type Account struct {
	ProjectID          string `json:"project_id" validate:"required"`
	Name               string `json:"name" validate:"required,min=2"`
	Email              string `json:"email" validate:"required,email"`
	Username           string `json:"username" validate:"required,min=2"`
	Password           string `json:"password" validate:"required,min=8"`
	RoleName           string `json:"role_name" validate:"required,min=2"`
	MustChangePassword bool   `json:"must_change_password"`
}

func (a Account) ToDomain() domain.Account {
	return domain.NewAccount(
		domain.WithProjectID(a.ProjectID),
		domain.WithName(a.Name),
		domain.WithEmail(a.Email),
		domain.WithUsername(a.Username),
		domain.WithPassword(a.Password),
		domain.WithRoleName(a.RoleName),
		domain.WithMustChangePassword(a.MustChangePassword),
	)
}

func (a Account) ValidateHttp(w http.ResponseWriter) error {
	if err := validator.ValidateStruct(&a); err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	return nil
}
