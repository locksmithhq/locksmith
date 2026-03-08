package stackerror

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"runtime"
	"strings"
)

type RepositoryError struct {
	statusCode int
	message    string
	local      string
	err        error
	trace      string
}

func (e RepositoryError) Error() string {
	return e.message
}

var (
	ErrNotFound = RepositoryError{
		statusCode: http.StatusNotFound,
		message:    "Not found",
	}
	ErrAlreadyExists = RepositoryError{
		statusCode: http.StatusConflict,
		message:    "Already exists",
	}
)

func NewRepositoryError(local string, err error) RepositoryError {
	local = "REPOSITORY: " + local
	stack := getStack()

	if errors.Is(err, sql.ErrNoRows) {
		ErrNotFound.local = local
		ErrNotFound.err = err
		ErrNotFound.trace = stack
		return ErrNotFound
	}

	return PQError(local, err)
}

func getStack() string {
	var stackBuilder strings.Builder
	for i := 1; ; i++ {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		if strings.Contains(file, "net/http/server") {
			break
		}
		if i > 1 {
			stackBuilder.WriteString("\n")
		}
		if strings.Contains(file, "stackerror") {
			continue
		}
		stackBuilder.WriteString(fmt.Sprintf("%s:%d", file, line))
	}
	stack := stackBuilder.String()

	files := strings.Split(stack, "\n")

	stackBuilder.Reset()
	for i := len(files) - 1; i >= 0; i-- {
		if i < len(files)-1 {
			stackBuilder.WriteString("\n")
		}
		stackBuilder.WriteString(files[i])
	}
	stack = stackBuilder.String()

	return stack
}
