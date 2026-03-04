package contract

import "net/http"

type FetchProjectsHandler interface {
	Execute(http.ResponseWriter, *http.Request)
}

type GetProjectByIDHandler interface {
	Execute(http.ResponseWriter, *http.Request)
}

type CreateProjectHandler interface {
	Execute(http.ResponseWriter, *http.Request)
}

type UpdateProjectHandler interface {
	Execute(http.ResponseWriter, *http.Request)
}

type DeleteProjectHandler interface {
	Execute(http.ResponseWriter, *http.Request)
}
