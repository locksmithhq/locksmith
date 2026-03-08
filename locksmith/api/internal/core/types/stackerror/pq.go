package stackerror

import (
	"net/http"

	"github.com/lib/pq"
)

func PQError(local string, err error) RepositoryError {
	stack := getStack()
	if pgErr, ok := err.(*pq.Error); ok {
		switch pgErr.Code {
		case "23505": // unique_violation
			ErrAlreadyExists.local = local
			ErrAlreadyExists.err = err
			ErrAlreadyExists.trace = stack
			return ErrAlreadyExists
		default:
			return RepositoryError{
				statusCode: http.StatusInternalServerError,
				message:    "Internal server error",
				local:      local,
				err:        err,
				trace:      stack,
			}
		}
	}

	return RepositoryError{
		statusCode: http.StatusInternalServerError,
		message:    "Internal server error",
		local:      local,
		err:        err,
		trace:      stack,
	}
}
