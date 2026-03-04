package contract

import "net/http"

type GetLoginByClientIDHandler interface {
	Execute(http.ResponseWriter, *http.Request)
}

type CreateLoginByClientIDHandler interface {
	Execute(http.ResponseWriter, *http.Request)
}

type UpdateLoginByClientIDHandler interface {
	Execute(http.ResponseWriter, *http.Request)
}
