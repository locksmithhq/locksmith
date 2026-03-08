package rest

import (
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients/di"
	"github.com/go-chi/chi/v5"
	"github.com/locksmithhq/locksmith-go"
)

func InitializeOauthClientsRouter(router chi.Router) {
	router.Get("/oauth2/resolve-domain", di.NewResolveCustomDomainHandler().Execute)

	router.Group(func(r chi.Router) {
		r.Use(locksmith.AuthMiddlewareCookie("LOCKSMITHACCESSTOKEN"))
		r.With(locksmith.AclMiddleware("domain:locksmith", "module:oauth_clients", "action:read:all")).Get("/projects/{project_id}/clients", di.NewFetchClientsByProjectIDHandler().Execute)
		r.With(locksmith.AclMiddleware("domain:locksmith", "module:oauth_clients", "action:read:one")).Get("/projects/{project_id}/clients/{id}", di.NewGetClientByIDAndProjectIDHandler().Execute)
		r.With(locksmith.AclMiddleware("domain:locksmith", "module:oauth_clients", "action:create:one")).Post("/projects/{project_id}/clients", di.NewCreateClientHandler().Execute)
		r.With(locksmith.AclMiddleware("domain:locksmith", "module:oauth_clients", "action:update:one")).Put("/projects/{project_id}/clients/{id}", di.NewUpdateClientHandler().Execute)
	})
}
