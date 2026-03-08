package rest

import (
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients_login/di"
	"github.com/go-chi/chi/v5"
	"github.com/locksmithhq/locksmith-go"
)

func InitializeOauthClientsLoginRouter(router chi.Router) {
	router.Group(func(r chi.Router) {
		r.Use(locksmith.AuthMiddlewareCookie("LOCKSMITHACCESSTOKEN"))
		r.With(locksmith.AclMiddleware("domain:locksmith", "module:oauth_clients_login", "action:read:one")).Get("/projects/{project_id}/clients/{id}/login", di.NewGetLoginByClientIDHandler().Execute)
		r.With(locksmith.AclMiddleware("domain:locksmith", "module:oauth_clients_login", "action:create:one")).Post("/projects/{project_id}/clients/{id}/login", di.NewCreateLoginByClientIDHandler().Execute)
		r.With(locksmith.AclMiddleware("domain:locksmith", "module:oauth_clients_login", "action:update:one")).Put("/projects/{project_id}/clients/{id}/login", di.NewUpdateLoginByClientIDHandler().Execute)
	})
}
