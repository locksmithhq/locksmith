package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/locksmithhq/locksmith-go"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
)

func InitializeLocksmithRouter(router chi.Router) {
	redirectUrl := os.Getenv("LOCKSMITH_BASE_URL")

	router.Get("/locksmith/callback", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		if code == "" {
			stackerror.HttpResponse(w, "CallbackHandler", fmt.Errorf("code is required"))
			return
		}

		// Read PKCE code_verifier from cookie set by the frontend before the redirect
		codeVerifier := ""
		if cv, err := r.Cookie("pkce_cv"); err == nil {
			codeVerifier = cv.Value
		}

		type tokenRequest struct {
			Code         string `json:"code"`
			ClientID     string `json:"client_id"`
			ClientSecret string `json:"client_secret"`
			GrantType    string `json:"grant_type"`
			CodeVerifier string `json:"code_verifier,omitempty"`
		}

		reqBody := tokenRequest{
			Code:         code,
			ClientID:     os.Getenv("LOCKSMITH_APP_CLIENT_ID"),
			ClientSecret: os.Getenv("LOCKSMITH_APP_CLIENT_SECRET"),
			GrantType:    "authorization_code",
			CodeVerifier: codeVerifier,
		}

		jsonBody, err := json.Marshal(reqBody)
		if err != nil {
			stackerror.HttpResponse(w, "CallbackHandler", err)
			return
		}

		baseURL := os.Getenv("LOCKSMITH_BASE_URL")
		resp, err := http.Post(baseURL+"/api/oauth2/access-token", "application/json", bytes.NewReader(jsonBody))
		if err != nil {
			stackerror.HttpResponse(w, "CallbackHandler", err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			var apiErr locksmith.ApiError
			json.NewDecoder(resp.Body).Decode(&apiErr)
			stackerror.HttpResponse(w, "CallbackHandler", apiErr)
			return
		}

		var token struct {
			AccessToken  string `json:"access_token"`
			RefreshToken string `json:"refresh_token"`
			ExpiresIn    int    `json:"expires_in"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&token); err != nil {
			stackerror.HttpResponse(w, "CallbackHandler", err)
			return
		}

		// Clear the PKCE code_verifier cookie
		http.SetCookie(w, &http.Cookie{
			Name:   "pkce_cv",
			Value:  "",
			Path:   "/",
			MaxAge: -1,
		})

		setCookies(w, r, token.AccessToken, token.RefreshToken, token.ExpiresIn)
		http.Redirect(w, r, redirectUrl, http.StatusFound)
	})

	router.Get("/locksmith/status", func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("LOCKSMITHACCESSTOKEN")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		accessToken := cookie.Value

		if _, valid := locksmith.VerifyToken(accessToken); !valid {
			http.SetCookie(w, &http.Cookie{
				Name:   "LOCKSMITHACCESSTOKEN",
				Value:  "",
				MaxAge: -1,
				Path:   "/",
			})

			http.SetCookie(w, &http.Cookie{
				Name:   "LOCKSMITHREFRESHTOKEN",
				Value:  "",
				MaxAge: -1,
				Path:   "/",
			})

			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusOK)
	})

	router.Post("/locksmith/r", func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("LOCKSMITHREFRESHTOKEN")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		token, err := locksmith.GenerateRefreshToken(r.Context(), locksmith.NewRefreshAccessTokenInput(
			cookie.Value,
		))
		if err != nil {
			stackerror.HttpResponse(w, "CallbackHandler", err)
			return
		}

		setCookies(w, r, token.AccessToken, token.RefreshToken, token.ExpiresIn)
	})
}

const RefreshTokenMaxAge = 15 * 24 * 60 * 60 // 15 days

func getDomainFromOrigin(r *http.Request) string {
	origin := r.Header.Get("Origin")
	if origin == "" {
		return ""
	}

	parsedURL, err := url.Parse(origin)
	if err != nil || parsedURL.Host == "" {
		return ""
	}

	return getDomainFromHost(parsedURL.Host)
}

func getDomainFromHost(host string) string {
	hostWithoutPort := strings.Split(host, ":")[0]

	if net.ParseIP(hostWithoutPort) != nil {
		return ".localhost"
	}

	parts := strings.Split(hostWithoutPort, ".")

	switch len(parts) {
	case 1:
		return "." + hostWithoutPort
	case 3:
		return "." + strings.Join(parts[len(parts)-2:], ".")
	default:
		return "." + strings.Join(parts[len(parts)-3:], ".")
	}
}

func setCookies(w http.ResponseWriter, r *http.Request, accessToken string, refreshToken string, expiresIn int) {
	domain := getDomainFromOrigin(r)
	http.SetCookie(w, &http.Cookie{
		Name:     "LOCKSMITHACCESSTOKEN",
		Value:    accessToken,
		Path:     "/",
		Domain:   domain,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   expiresIn,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "LOCKSMITHREFRESHTOKEN",
		Value:    refreshToken,
		Path:     "/",
		Domain:   domain,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   RefreshTokenMaxAge,
	})
}
