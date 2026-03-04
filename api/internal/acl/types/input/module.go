package input

import (
	"net/http"

	"github.com/booscaaa/initializers/validator"
	"github.com/booscaaa/locksmith/api/internal/acl/domain"
)

type Module struct {
	Id      string   `json:"id"`
	Title   string   `json:"title"`
	Actions []Action `json:"actions"`
}

func (in Module) ToModuleDomain() domain.Module {
	return domain.NewModule(
		domain.WithModuleTitle(in.Title),
	)
}

func (in Module) ValidateHttp(w http.ResponseWriter) error {
	if err := validator.ValidateStruct(&in); err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	return nil
}
