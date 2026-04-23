package handler

import (
	"net/http"

	"github.com/locksmithhq/locksmith/api/internal/locksmith/contract"
	oauth2Contract "github.com/locksmithhq/locksmith/api/internal/oauth2/contract"
)

type logoutHandler struct {
	logoutUseCase oauth2Contract.LogoutUseCase
}

// Execute implements contract.LogoutHandler.
func (h *logoutHandler) Execute(w http.ResponseWriter, r *http.Request) {
	if cookie, err := r.Cookie("LOCKSMITHREFRESHTOKEN"); err == nil && cookie.Value != "" {
		_ = h.logoutUseCase.Execute(r.Context(), cookie.Value)
	}

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

func NewLogoutHandler(logoutUseCase oauth2Contract.LogoutUseCase) contract.LogoutHandler {
	return &logoutHandler{logoutUseCase: logoutUseCase}
}
