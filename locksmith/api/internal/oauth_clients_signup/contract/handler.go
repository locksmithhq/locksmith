package contract

import "net/http"

type GetSignupByClientIDHandler interface {
	Execute(http.ResponseWriter, *http.Request)
}

type CreateSignupByClientIDHandler interface {
	Execute(http.ResponseWriter, *http.Request)
}

type UpdateSignupByClientIDHandler interface {
	Execute(http.ResponseWriter, *http.Request)
}
