package rest

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/booscaaa/locksmith/api/internal/account/di"
	"github.com/booscaaa/locksmith/api/internal/adapter/database"
	"github.com/go-chi/chi/v5"
	"github.com/locksmithhq/locksmith-go"
)

func InitializeAccountsRouter(router chi.Router) {
	router.Group(func(r chi.Router) {
		r.Use(locksmith.AuthMiddlewareCookie("LOCKSMITHACCESSTOKEN"))

		r.With(locksmith.AclMiddleware("domain:locksmith", "module:accounts", "action:create:one")).Post("/projects/{project_id}/accounts", di.NewCreateAccountHandler().Execute)
		r.With(locksmith.AclMiddleware("domain:locksmith", "module:accounts", "action:update:one")).Put("/projects/{project_id}/accounts/{account_id}", di.NewUpdateAccountHandler().Execute)
		r.With(locksmith.AclMiddleware("domain:locksmith", "module:accounts", "action:read:all")).Get("/projects/{project_id}/accounts", di.NewFetchAccountsByProjectIDHandler().Execute)
		r.With(locksmith.AclMiddleware("domain:locksmith", "module:accounts", "action:read:all")).Get("/projects/{project_id}/accounts/count", di.NewCountAccountsByProjectIDHandler().Execute)
	})

	router.Group(func(r chi.Router) {
		r.Use(BasicAuthMiddleware)

		r.Post("/accounts", di.NewCreateAccountHandler().Execute)
		r.Get("/accounts/{id}", di.NewGetAccountByProjectIDAndIDHandler().Execute)
		r.Put("/accounts/{account_id}", di.NewUpdateAccountHandler().Execute)
	})

	// No middleware: JWT from login's change_password_jwt is verified inside the usecase
	router.Post("/accounts/change-password", di.NewChangePasswordHandler().Execute)

}

func BasicAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		basicAuth := r.Header.Get("Authorization")
		if basicAuth == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		basicAuth = strings.TrimPrefix(basicAuth, "Basic ")
		basicAuth = strings.TrimSpace(basicAuth)

		basiAuth, err := base64.StdEncoding.DecodeString(basicAuth)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		basics := strings.Split(string(basiAuth), ":")
		if len(basics) != 2 {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		var id string
		var projectID string
		if err := database.GetConnection().
			QueryRowContext(
				r.Context(),
				"SELECT id, project_id FROM oauth_clients WHERE client_id = $1 AND client_secret = $2",
				basics[0], basics[1],
			).Scan(&id, &projectID); err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if id == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		rctx := chi.RouteContext(r.Context())
		rctx.URLParams.Add("project_id", projectID)

		next.ServeHTTP(w, r)
	})
}
