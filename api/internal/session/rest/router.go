package rest

import (
	adapterMiddleware "github.com/locksmithhq/locksmith/api/internal/adapter/middleware"
	"github.com/locksmithhq/locksmith/api/internal/session/di"
	"github.com/go-chi/chi/v5"
	"github.com/locksmithhq/locksmith-go"
)

func InitializeSessionRouter(router chi.Router) {
	router.Group(func(r chi.Router) {
		r.Use(locksmith.AuthMiddlewareCookie("LOCKSMITHACCESSTOKEN"))

		r.With(locksmith.AclMiddleware("domain:locksmith", "module:sessions", "action:read:all")).
			Get("/projects/{project_id}/sessions", di.NewFetchSessionsByProjectIDHandler().Execute)
		r.With(locksmith.AclMiddleware("domain:locksmith", "module:sessions", "action:read:all")).
			Get("/projects/{project_id}/sessions/count", di.NewCountSessionsByProjectIDHandler().Execute)

		r.With(locksmith.AclMiddleware("domain:locksmith", "module:sessions", "action:read:all")).
			Get("/projects/{project_id}/accounts/{account_id}/sessions", di.NewFetchSessionsByAccountIDHandler().Execute)
		r.With(locksmith.AclMiddleware("domain:locksmith", "module:sessions", "action:read:all")).
			Get("/projects/{project_id}/accounts/{account_id}/sessions/count", di.NewCountSessionsByAccountIDHandler().Execute)

		r.With(locksmith.AclMiddleware("domain:locksmith", "module:sessions", "action:revoke")).
			Delete("/projects/{project_id}/sessions/{session_id}", di.NewRevokeSessionHandler().Execute)

		r.With(locksmith.AclMiddleware("domain:locksmith", "module:sessions", "action:read:all")).
			Get("/projects/{project_id}/accounts/{account_id}/refresh-tokens", di.NewFetchRefreshTokensByAccountIDHandler().Execute)
		r.With(locksmith.AclMiddleware("domain:locksmith", "module:sessions", "action:read:all")).
			Get("/projects/{project_id}/accounts/{account_id}/refresh-tokens/count", di.NewCountRefreshTokensByAccountIDHandler().Execute)
	})

	router.Group(func(r chi.Router) {
		r.Use(adapterMiddleware.NewBasicAuthMiddleware())

		r.Get("/accounts/{account_id}/sessions", di.NewFetchSessionsByAccountIDHandler().Execute)
		r.Get("/accounts/{account_id}/sessions/count", di.NewCountSessionsByAccountIDHandler().Execute)
		r.Delete("/accounts/{account_id}/sessions/{session_id}", di.NewRevokeSessionHandler().Execute)
	})
}
