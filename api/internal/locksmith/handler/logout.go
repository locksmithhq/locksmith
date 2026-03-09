package handler

import (
	"net/http"

	"github.com/locksmithhq/locksmith/api/internal/locksmith/contract"
)

type logoutHandler struct{}

// Execute implements contract.LogoutHandler.
func (h *logoutHandler) Execute(w http.ResponseWriter, r *http.Request) {
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
	w.WriteHeader(http.StatusNoContent)
}

func NewLogoutHandler() contract.LogoutHandler {
	return &logoutHandler{}
}
