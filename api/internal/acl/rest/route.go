package rest

import (
	"github.com/locksmithhq/locksmith/api/internal/acl/di"
	adapterMiddleware "github.com/locksmithhq/locksmith/api/internal/adapter/middleware"
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
		r.Use(adapterMiddleware.NewBasicAuthMiddleware())
		r.Get("/acl/permissions/user/{user}/domain/{domain}", di.NewGetPermissionsForUserInDomainHandler().Execute)
	})
}
