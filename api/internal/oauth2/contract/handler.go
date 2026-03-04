package contract

import "net/http"

type AuthorizeClientHandler interface {
	Execute(w http.ResponseWriter, r *http.Request)
}

type LoginHandler interface {
	Execute(w http.ResponseWriter, r *http.Request)
}

type GenerateAccessTokenHandler interface {
	Execute(w http.ResponseWriter, r *http.Request)
}

type GenerateRefreshTokenHandler interface {
	Execute(w http.ResponseWriter, r *http.Request)
}

type GetClientByClientIDHandler interface {
	Execute(w http.ResponseWriter, r *http.Request)
}

type RegisterHandler interface {
	Execute(w http.ResponseWriter, r *http.Request)
}
