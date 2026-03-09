package handler

import (
	"net/http"

	"github.com/locksmithhq/locksmith-go"
	"github.com/locksmithhq/locksmith/api/internal/locksmith/contract"
)

type statusHandler struct{}

// Execute implements contract.StatusHandler.
func (h *statusHandler) Execute(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("LOCKSMITHACCESSTOKEN")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if _, valid := locksmith.VerifyToken(cookie.Value); !valid {
		domain := getDomainFromRequest(r)
		http.SetCookie(w, &http.Cookie{
			Name:     "LOCKSMITHACCESSTOKEN",
			Value:    "",
			Path:     "/",
			Domain:   domain,
			MaxAge:   -1,
			Secure:   true,
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
		})
		http.SetCookie(w, &http.Cookie{
			Name:     "LOCKSMITHREFRESHTOKEN",
			Value:    "",
			Path:     "/",
			Domain:   domain,
			MaxAge:   -1,
			Secure:   true,
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
		})
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func NewStatusHandler() contract.StatusHandler {
	return &statusHandler{}
}
