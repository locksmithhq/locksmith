package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/locksmithhq/locksmith-go"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/locksmith/contract"
	oauth2Contract "github.com/locksmithhq/locksmith/api/internal/oauth2/contract"
	oauth2Input "github.com/locksmithhq/locksmith/api/internal/oauth2/types/input"
)

type callbackHandler struct {
	generateAccessTokenUseCase oauth2Contract.GenerateAccessTokenUseCase
}

// Execute implements contract.CallbackHandler.
func (h *callbackHandler) Execute(w http.ResponseWriter, r *http.Request) {
	redirectURL := os.Getenv("LOCKSMITH_BASE_URL")

	code := r.URL.Query().Get("code")
	if code == "" {
		stackerror.HttpResponse(w, "CallbackHandler", fmt.Errorf("code is required"))
		return
	}

	codeVerifier := ""
	if cv, err := r.Cookie("pkce_cv"); err == nil {
		codeVerifier = cv.Value
	}

	fp := locksmith.Parse(r)
	deviceID := r.URL.Query().Get("device_id")
	if deviceID == "" {
		if cookie, err := r.Cookie("device_id"); err == nil {
			deviceID = cookie.Value
		}
	}

	token, err := h.generateAccessTokenUseCase.Execute(r.Context(), oauth2Input.AccessToken{
		ClientID:        os.Getenv("LOCKSMITH_APP_CLIENT_ID"),
		ClientSecret:    os.Getenv("LOCKSMITH_APP_CLIENT_SECRET"),
		GrantType:       "authorization_code",
		Code:            code,
		CodeVerifier:    codeVerifier,
		DeviceID:        deviceID,
		IPAddress:       fp.IPAddress,
		UserAgent:       fp.UserAgent,
		DeviceType:      fp.DeviceType,
		Browser:         fp.Browser,
		BrowserVersion:  fp.BrowserVersion,
		OS:              fp.OS,
		OSVersion:       fp.OSVersion,
		LocationCountry: fp.LocationCountry,
		LocationRegion:  fp.LocationRegion,
		LocationCity:    fp.LocationCity,
	})
	if err != nil {
		stackerror.HttpResponse(w, "CallbackHandler", err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "pkce_cv",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})

	setCookies(w, r, token.AccessToken, token.RefreshToken, token.ExpiresIn)
	http.Redirect(w, r, redirectURL, http.StatusFound)
}

func NewCallbackHandler(generateAccessTokenUseCase oauth2Contract.GenerateAccessTokenUseCase) contract.CallbackHandler {
	return &callbackHandler{
		generateAccessTokenUseCase: generateAccessTokenUseCase,
	}
}
