package stackerror

import (
	"encoding/json"
	"errors"
	"net/http"
)

func HttpResponse(w http.ResponseWriter, caller string, e error) {
	var usecaseError UseCaseError
	if errors.As(e, &usecaseError) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(usecaseError.statusCode)
		json.NewEncoder(w).Encode(struct {
			Message string `json:"message"`
			Local   string `json:"local"`
			Err   string `json:"err"`
			Trace   string `json:"trace"`
		}{
			Message: usecaseError.message,
			Local:   "HANDLER: " + caller + " -> " + usecaseError.local,
			Err:   usecaseError.err.Error(),
			Trace:   usecaseError.trace,
		})
		return
	}

	var repositoryError RepositoryError
	if errors.As(e, &repositoryError) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(repositoryError.statusCode)
		json.NewEncoder(w).Encode(struct {
			Message string `json:"message"`
			Local   string `json:"local"`
			Err   string `json:"err"`
			File    string `json:"file"`
		}{
			Message: repositoryError.message,
			Local:   "HANDLER: " + caller + " -> " + repositoryError.local,
			Err:   repositoryError.err.Error(),
			File:    repositoryError.trace,
		})
		return
	}

	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(struct {
		Message string `json:"message"`
		Local   string `json:"local"`
		Err   string `json:"err"`
		Trace   string `json:"trace"`
	}{
		Message: e.Error(),
		Local:   "HANDLER: " + caller,
		Err:   e.Error(),
		Trace:   getStack(),
	})
}
