package contract

import "net/http"

type CreateAccountHandler interface {
	Execute(http.ResponseWriter, *http.Request)
}

type FetchAccountsByProjectIDHandler interface {
	Execute(http.ResponseWriter, *http.Request)
}

type UpdateAccountHandler interface {
	Execute(http.ResponseWriter, *http.Request)
}

type ChangePasswordHandler interface {
	Execute(http.ResponseWriter, *http.Request)
}

type GetAccountByProjectIDAndIDHandler interface {
	Execute(http.ResponseWriter, *http.Request)
}

type CountAccountsByProjectIDHandler interface {
	Execute(http.ResponseWriter, *http.Request)
}
