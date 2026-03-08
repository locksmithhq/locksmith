package input

import (
	"net/http"

	"github.com/booscaaa/initializers/validator"
	"github.com/locksmithhq/locksmith/api/internal/project/domain"
)

type Project struct {
	Name        string `json:"name" validate:"required,min=3,max=100"`
	Description string `json:"description" validate:"required,min=3,max=255"`
	Domain      string `json:"domain" validate:"required,min=3,max=255"`
}

func (in Project) ToProjectDomain() domain.Project {
	return domain.NewProject(
		domain.WithName(in.Name),
		domain.WithDescription(in.Description),
		domain.WithDomain(in.Domain),
	)
}

func (p Project) ValidateHttp(w http.ResponseWriter) error {
	if err := validator.ValidateStruct(&p); err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	return nil
}
