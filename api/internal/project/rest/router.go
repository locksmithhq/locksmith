package rest

import (
	"github.com/locksmithhq/locksmith/api/internal/project/di"
	"github.com/go-chi/chi/v5"
	"github.com/locksmithhq/locksmith-go"
)

func InitializeProjectRouter(router chi.Router) {
	router.Group(func(r chi.Router) {
		r.Use(locksmith.AuthMiddlewareCookie("LOCKSMITHACCESSTOKEN"))
		r.With(locksmith.AclMiddleware("domain:locksmith", "module:project", "action:read:all")).Get("/projects", di.NewFetchProjectsHandler().Execute)
		r.With(locksmith.AclMiddleware("domain:locksmith", "module:project", "action:read:one")).Get("/projects/{id}", di.NewGetProjectByIDHandler().Execute)
		r.With(locksmith.AclMiddleware("domain:locksmith", "module:project", "action:create:one")).Post("/projects", di.NewCreateProjectHandler().Execute)
		r.With(locksmith.AclMiddleware("domain:locksmith", "module:project", "action:update:one")).Put("/projects/{id}", di.NewUpdateProjectHandler().Execute)
		r.With(locksmith.AclMiddleware("domain:locksmith", "module:project", "action:delete:one")).Delete("/projects/{id}", di.NewDeleteProjectHandler().Execute)
	})
}
