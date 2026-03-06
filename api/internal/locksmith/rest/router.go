package rest

import (
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

		token, err := locksmith.GenerateAccessToken(r.Context(), r, locksmith.AccessTokenInput{
			ClientID:     os.Getenv("LOCKSMITH_APP_CLIENT_ID"),
			ClientSecret: os.Getenv("LOCKSMITH_APP_CLIENT_SECRET"),
			GrantType:    "authorization_code",
			Code:         code,
			CodeVerifier: codeVerifier,
		})

		if err != nil {
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

	router.Post("/locksmith/logout", func(w http.ResponseWriter, r *http.Request) {
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
	})

	router.Get("/locksmith/status", func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("LOCKSMITHACCESSTOKEN")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		accessToken := cookie.Value

		if _, valid := locksmith.VerifyToken(accessToken); !valid {
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
	})

	router.Post("/locksmith/r", func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("LOCKSMITHREFRESHTOKEN")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		token, err := locksmith.GenerateRefreshToken(r.Context(), locksmith.RefreshAccessTokenInput{
			RefreshToken: cookie.Value,
		})
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
	})
}

const RefreshTokenMaxAge = 15 * 24 * 60 * 60 // 15 days

func getDomainFromRequest(r *http.Request) string {
	if domain := getDomainFromOrigin(r); domain != "" {
		return domain
	}
	return getDomainFromHost(r.Host)
}

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
	domain := getDomainFromRequest(r)
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
