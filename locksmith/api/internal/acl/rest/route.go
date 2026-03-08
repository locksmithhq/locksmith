package rest

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/locksmithhq/locksmith/api/internal/acl/di"
	"github.com/locksmithhq/locksmith/api/internal/adapter/database"
	"github.com/go-chi/chi/v5"
	"github.com/locksmithhq/locksmith-go"
)

func InitializeAclRouter(router chi.Router) {
	router.Group(func(r chi.Router) {
		r.Use(locksmith.AuthMiddlewareCookie("LOCKSMITHACCESSTOKEN"))

		r.With(locksmith.AclMiddleware("domain:locksmith", "module:acl", "action:read:all")).Get("/acl", di.NewFetchAclHandler().Execute)
		r.With(locksmith.AclMiddleware("domain:locksmith", "module:acl", "action:create:role")).Post("/acl/role", di.NewCreateRoleHandler().Execute)
		r.With(locksmith.AclMiddleware("domain:locksmith", "module:acl", "action:create:module")).Post("/acl/module", di.NewCreateModuleHandler().Execute)
		r.With(locksmith.AclMiddleware("domain:locksmith", "module:acl", "action:create:action")).Post("/acl/action", di.NewCreateActionHandler().Execute)
		r.With(locksmith.AclMiddleware("domain:locksmith", "module:acl", "action:read:all")).Get("/acl/roles", di.NewFetchRolesHandler().Execute)
		r.With(locksmith.AclMiddleware("domain:locksmith", "module:acl", "action:read:all")).Get("/acl/modules", di.NewFetchModulesHandler().Execute)
		r.With(locksmith.AclMiddleware("domain:locksmith", "module:acl", "action:read:all")).Get("/acl/actions", di.NewFetchActionsHandler().Execute)
		r.With(locksmith.AclMiddleware("domain:locksmith", "module:acl", "action:read:all")).Get("/acl/projects/{projectId}", di.NewFetchProjectAclHandler().Execute)
		r.With(locksmith.AclMiddleware("domain:locksmith", "module:acl", "action:create:all")).Post("/acl/projects/{projectId}", di.NewCreateProjectAclHandler().Execute)
	})

	router.Group(func(r chi.Router) {
		r.Use(BasicAuthMiddleware)
		r.Get("/acl/permissions/user/{user}/domain/{domain}", di.NewGetPermissionsForUserInDomainHandler().Execute)
	})
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
