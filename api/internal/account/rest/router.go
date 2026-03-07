package rest

import (
	"github.com/locksmithhq/locksmith/api/internal/account/di"
	adapterMiddleware "github.com/locksmithhq/locksmith/api/internal/adapter/middleware"
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
		r.With(locksmith.AclMiddleware("domain:locksmith", "module:accounts", "action:read:all")).Get("/projects/{project_id}/accounts/{id}", di.NewGetAccountByProjectIDAndIDHandler().Execute)
	})

	router.Group(func(r chi.Router) {
		r.Use(adapterMiddleware.BasicAuthMiddleware)

		r.Post("/accounts", di.NewCreateAccountHandler().Execute)
		r.Get("/accounts/{id}", di.NewGetAccountByProjectIDAndIDHandler().Execute)
		r.Put("/accounts/{account_id}", di.NewUpdateAccountHandler().Execute)
	})

	// No middleware: JWT from login's change_password_jwt is verified inside the usecase
	router.Post("/accounts/change-password", di.NewChangePasswordHandler().Execute)

}
