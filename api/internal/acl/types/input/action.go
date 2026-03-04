package input

import (
	"net/http"

	"github.com/booscaaa/initializers/validator"
	"github.com/booscaaa/locksmith/api/internal/acl/domain"
)

type Action struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

func (in Action) ToActionDomain() domain.Action {
	return domain.NewAction(
		domain.WithActionTitle(in.Title),
	)
}

func (in Action) ValidateHttp(w http.ResponseWriter) error {
	if err := validator.ValidateStruct(&in); err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	return nil
}
