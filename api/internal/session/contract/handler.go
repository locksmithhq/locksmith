package contract

import "net/http"

type FetchSessionsByProjectIDHandler interface {
	Execute(http.ResponseWriter, *http.Request)
}

type CountSessionsByProjectIDHandler interface {
	Execute(http.ResponseWriter, *http.Request)
}
