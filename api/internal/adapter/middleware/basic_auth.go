package middleware

import (
	"crypto/subtle"
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/locksmithhq/locksmith/api/internal/adapter/database"
	"github.com/locksmithhq/locksmith/api/internal/core/crypto"
	oauthClientContract "github.com/locksmithhq/locksmith/api/internal/oauth_clients/contract"
	oauthClientRepo "github.com/locksmithhq/locksmith/api/internal/oauth_clients/repository"
)

type basicAuthMiddleware struct {
	repo oauthClientContract.GetClientByClientIDRepository
}

func NewBasicAuthMiddleware() func(http.Handler) http.Handler {
	m := &basicAuthMiddleware{
		repo: oauthClientRepo.NewGetClientByClientIDRepository(database.GetConnection()),
	}
	return m.Handle
}

func (m *basicAuthMiddleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		basicAuth := r.Header.Get("Authorization")
		if basicAuth == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		basicAuth = strings.TrimPrefix(basicAuth, "Basic ")
		basicAuth = strings.TrimSpace(basicAuth)

		decoded, err := base64.StdEncoding.DecodeString(basicAuth)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		parts := strings.SplitN(string(decoded), ":", 2)
		if len(parts) != 2 {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		clientID, clientSecret := parts[0], parts[1]

		client, err := m.repo.Execute(r.Context(), clientID)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		decrypted, err := crypto.Decrypt(client.ClientSecret)
		if err != nil || subtle.ConstantTimeCompare([]byte(decrypted), []byte(clientSecret)) != 1 {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		rctx := chi.RouteContext(r.Context())
		rctx.URLParams.Add("project_id", client.ProjectID)

		next.ServeHTTP(w, r)
	})
}
