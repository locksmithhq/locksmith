package contract

import "net/http"

type FetchSessionsByProjectIDHandler interface {
	Execute(http.ResponseWriter, *http.Request)
}

type CountSessionsByProjectIDHandler interface {
	Execute(http.ResponseWriter, *http.Request)
}

type FetchSessionsByAccountIDHandler interface {
	Execute(http.ResponseWriter, *http.Request)
}

type CountSessionsByAccountIDHandler interface {
	Execute(http.ResponseWriter, *http.Request)
}

type RevokeSessionHandler interface {
	Execute(http.ResponseWriter, *http.Request)
}

type FetchRefreshTokensByAccountIDHandler interface {
	Execute(http.ResponseWriter, *http.Request)
}

type CountRefreshTokensByAccountIDHandler interface {
	Execute(http.ResponseWriter, *http.Request)
}
