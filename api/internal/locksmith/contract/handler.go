package contract

import "net/http"

type CallbackHandler interface {
	Execute(http.ResponseWriter, *http.Request)
}

type LogoutHandler interface {
	Execute(http.ResponseWriter, *http.Request)
}

type StatusHandler interface {
	Execute(http.ResponseWriter, *http.Request)
}

type RefreshTokenHandler interface {
	Execute(http.ResponseWriter, *http.Request)
}
