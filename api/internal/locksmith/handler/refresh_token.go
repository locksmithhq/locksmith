package handler

import (
	"net/http"

	"github.com/locksmithhq/locksmith/api/internal/locksmith/contract"
	oauth2Contract "github.com/locksmithhq/locksmith/api/internal/oauth2/contract"
)

type refreshTokenHandler struct {
	generateRefreshTokenUseCase oauth2Contract.GenerateRefreshTokenUseCase
}

// Execute implements contract.RefreshTokenHandler.
func (h *refreshTokenHandler) Execute(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("LOCKSMITHREFRESHTOKEN")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	token, err := h.generateRefreshTokenUseCase.Execute(r.Context(), cookie.Value)
	if err != nil {
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

	setCookies(w, r, token.AccessToken, token.RefreshToken, token.ExpiresIn)
}

func NewRefreshTokenHandler(generateRefreshTokenUseCase oauth2Contract.GenerateRefreshTokenUseCase) contract.RefreshTokenHandler {
	return &refreshTokenHandler{
		generateRefreshTokenUseCase: generateRefreshTokenUseCase,
	}
}
