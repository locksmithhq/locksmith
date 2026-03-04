package stackerror

import (
	"errors"
	"fmt"
	"net/http"
	"os"
)

type UseCaseOption func(c *UseCaseError)

type UseCaseError struct {
	statusCode int
	message    string
	local      string
	err        error
	trace      string
}

func (e UseCaseError) Error() string {
	return e.message
}

func (e UseCaseError) StdoutResponse(caller string) {
	fmt.Fprintf(os.Stderr, "Error: %s\n", e.message)
	fmt.Fprintf(os.Stderr, "  at %s\n", caller+" -> "+e.local)
	fmt.Fprintf(os.Stderr, "  error: %s\n", e.err.Error())
	fmt.Fprintf(os.Stderr, "  trace: %s\n", e.trace)
}

func WithMessage(message string) UseCaseOption {
	return func(c *UseCaseError) {
		c.message = message
	}
}

func WithStatusCode(statusCode int) UseCaseOption {
	return func(c *UseCaseError) {
		c.statusCode = statusCode
	}
}

func NewUseCaseError(local string, err error, options ...UseCaseOption) UseCaseError {
	var repositoryError RepositoryError
	if errors.As(err, &repositoryError) {
		usecaseError := UseCaseError{
			statusCode: repositoryError.statusCode,
			message:    repositoryError.message,
			local:      "USECASE: " + local + " -> " + repositoryError.local,
			err:        repositoryError.err,
			trace:      repositoryError.trace,
		}

		return parseOptions(usecaseError, options...)
	}

	usecaseError := UseCaseError{
		statusCode: http.StatusInternalServerError,
		message:    err.Error(),
		local:      "USECASE: " + local,
		err:        err,
		trace:      getStack(),
	}

	return parseOptions(usecaseError, options...)
}

func parseOptions(usecaseError UseCaseError, options ...UseCaseOption) UseCaseError {
	for _, option := range options {
		option(&usecaseError)
	}

	return usecaseError
}
