package contract

import "net/http"

type FetchAclHandler interface {
	Execute(http.ResponseWriter, *http.Request)
}

type GetPermissionsForUserInDomainHandler interface {
	Execute(http.ResponseWriter, *http.Request)
}

type CreateRoleHandler interface {
	Execute(http.ResponseWriter, *http.Request)
}

type CreateModuleHandler interface {
	Execute(http.ResponseWriter, *http.Request)
}

type CreateActionHandler interface {
	Execute(http.ResponseWriter, *http.Request)
}

type FetchRolesHandler interface {
	Execute(http.ResponseWriter, *http.Request)
}

type FetchModulesHandler interface {
	Execute(http.ResponseWriter, *http.Request)
}

type FetchActionsHandler interface {
	Execute(http.ResponseWriter, *http.Request)
}

type FetchProjectAclHandler interface {
	Execute(http.ResponseWriter, *http.Request)
}

type CreateProjectAclHandler interface {
	Execute(http.ResponseWriter, *http.Request)
}
