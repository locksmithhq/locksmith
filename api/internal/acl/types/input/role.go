package input

import (
	"net/http"

	"github.com/booscaaa/initializers/validator"
	"github.com/locksmithhq/locksmith/api/internal/acl/domain"
)

type Role struct {
	Id      string   `json:"id"`
	Title   string   `json:"title"`
	Modules []Module `json:"modules"`
}

func (in Role) ToRoleDomain() domain.Role {
	return domain.NewRole(
		domain.WithRoleTitle(in.Title),
	)
}

func (in Role) ValidateHttp(w http.ResponseWriter) error {
	if err := validator.ValidateStruct(&in); err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	return nil
}
