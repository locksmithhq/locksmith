package contract

import "net/http"

type GetSocialProvidersByClientIDHandler interface {
	Execute(http.ResponseWriter, *http.Request)
}

type UpsertSocialProviderHandler interface {
	Execute(http.ResponseWriter, *http.Request)
}
