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

type PWAManifestHandler interface {
	Execute(w http.ResponseWriter, r *http.Request)
}

type FaviconHandler interface {
	Execute(w http.ResponseWriter, r *http.Request)
}

type SocialBeginHandler interface {
	Execute(w http.ResponseWriter, r *http.Request)
}

type SocialCallbackHandler interface {
	Execute(w http.ResponseWriter, r *http.Request)
}
