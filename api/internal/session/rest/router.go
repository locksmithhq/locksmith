package rest

import (
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
	})
}
