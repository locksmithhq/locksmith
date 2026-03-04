package rest

import (
	"github.com/booscaaa/locksmith/api/internal/oauth_clients_signup/di"
	"github.com/go-chi/chi/v5"
	"github.com/locksmithhq/locksmith-go"
)

func InitializeOauthClientsSignupRouter(router chi.Router) {
	router.Group(func(r chi.Router) {
		r.Use(locksmith.AuthMiddlewareCookie("LOCKSMITHACCESSTOKEN"))
		r.With(locksmith.AclMiddleware("domain:locksmith", "oauth_clients_signup", "read:all")).Get("/projects/{project_id}/clients/{id}/signup", di.NewGetSignupByClientIDHandler().Execute)
		r.With(locksmith.AclMiddleware("domain:locksmith", "oauth_clients_signup", "create:one")).Post("/projects/{project_id}/clients/{id}/signup", di.NewCreateSignupByClientIDHandler().Execute)
		r.With(locksmith.AclMiddleware("domain:locksmith", "oauth_clients_signup", "update:one")).Put("/projects/{project_id}/clients/{id}/signup", di.NewUpdateSignupByClientIDHandler().Execute)
	})
}
