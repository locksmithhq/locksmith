package stackerror

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
)

func isProduction() bool {
	return os.Getenv("APP_ENV") == "production"
}

func HttpResponse(w http.ResponseWriter, caller string, e error) {
	var usecaseError UseCaseError
	if errors.As(e, &usecaseError) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(usecaseError.statusCode)

		if isProduction() {
			errorID := uuid.New().String()
			log.Printf("[ERROR] error_id=%s handler=%s local=%s err=%s trace=%s",
				errorID, caller, usecaseError.local, usecaseError.err.Error(), usecaseError.trace)
			json.NewEncoder(w).Encode(struct {
				Message string `json:"message"`
				ErrorID string `json:"error_id"`
			}{
				Message: usecaseError.message,
				ErrorID: errorID,
			})
			return
		}

		json.NewEncoder(w).Encode(struct {
			Message string `json:"message"`
			Local   string `json:"local"`
			Err     string `json:"err"`
			Trace   string `json:"trace"`
		}{
			Message: usecaseError.message,
			Local:   "HANDLER: " + caller + " -> " + usecaseError.local,
			Err:     usecaseError.err.Error(),
			Trace:   usecaseError.trace,
		})
		return
	}

	var repositoryError RepositoryError
	if errors.As(e, &repositoryError) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(repositoryError.statusCode)

		if isProduction() {
			errorID := uuid.New().String()
			log.Printf("[ERROR] error_id=%s handler=%s local=%s err=%s trace=%s",
				errorID, caller, repositoryError.local, repositoryError.err.Error(), repositoryError.trace)
			json.NewEncoder(w).Encode(struct {
				Message string `json:"message"`
				ErrorID string `json:"error_id"`
			}{
				Message: repositoryError.message,
				ErrorID: errorID,
			})
			return
		}

		json.NewEncoder(w).Encode(struct {
			Message string `json:"message"`
			Local   string `json:"local"`
			Err     string `json:"err"`
			File    string `json:"file"`
		}{
			Message: repositoryError.message,
			Local:   "HANDLER: " + caller + " -> " + repositoryError.local,
			Err:     repositoryError.err.Error(),
			File:    repositoryError.trace,
		})
		return
	}

	if isProduction() {
		errorID := uuid.New().String()
		log.Printf("[ERROR] error_id=%s handler=%s err=%s trace=%s",
			errorID, caller, e.Error(), getStack())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(struct {
			Message string `json:"message"`
			ErrorID string `json:"error_id"`
		}{
			Message: "Bad request",
			ErrorID: errorID,
		})
		return
	}

	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(struct {
		Message string `json:"message"`
		Local   string `json:"local"`
		Err     string `json:"err"`
		Trace   string `json:"trace"`
	}{
		Message: e.Error(),
		Local:   "HANDLER: " + caller,
		Err:     e.Error(),
		Trace:   getStack(),
	})
}
