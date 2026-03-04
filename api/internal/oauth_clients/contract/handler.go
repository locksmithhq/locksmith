package contract

import "net/http"

type FetchClientsByProjectIDHandler interface {
	Execute(http.ResponseWriter, *http.Request)
}

type GetClientByIDAndProjectIDHandler interface {
	Execute(http.ResponseWriter, *http.Request)
}

type CreateClientHandler interface {
	Execute(http.ResponseWriter, *http.Request)
}

type UpdateClientHandler interface {
	Execute(http.ResponseWriter, *http.Request)
}

type DeleteClientHandler interface {
	Execute(http.ResponseWriter, *http.Request)
}
