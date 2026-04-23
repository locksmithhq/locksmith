package handler

import (
	"net"
	"net/http"
	"net/url"
	"strings"
)

const RefreshTokenMaxAge = 15 * 24 * 60 * 60 // 15 days

func setCookies(w http.ResponseWriter, r *http.Request, accessToken string, refreshToken string, expiresIn int) {
	domain := getDomainFromRequest(r)
	http.SetCookie(w, &http.Cookie{
		Name:     "LOCKSMITHACCESSTOKEN",
		Value:    accessToken,
		Path:     "/",
		Domain:   domain,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   expiresIn,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "LOCKSMITHREFRESHTOKEN",
		Value:    refreshToken,
		Path:     "/",
		Domain:   domain,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   RefreshTokenMaxAge,
	})
}

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
