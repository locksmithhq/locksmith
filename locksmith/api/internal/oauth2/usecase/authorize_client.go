package usecase

import (
	"context"
	"fmt"
	"net/http"
	"slices"
	"strings"

	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/types/input"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/types/output"
)

type authorizeClientUseCase struct {
	getClientByClientIDRepository contract.GetClientByClientIDRepository
	getLoginByClientIDRepository  contract.GetLoginByClientIDRepository
	getSignupByClientIDRepository contract.GetSignupByClientIDRepository
}

// Execute implements contract.AuthorizeClient.
func (u *authorizeClientUseCase) Execute(ctx context.Context, in input.Authorization) (output.Client, error) {
	client, err := u.getClientByClientIDRepository.Execute(ctx, in.ClientID)
	if err != nil {
		return output.Client{}, stackerror.NewUseCaseError(
			"AuthorizeClientUseCase",
			err,
			stackerror.WithMessage("the client_id is not valid, please check your request"),
			stackerror.WithStatusCode(http.StatusBadRequest),
		)
	}

	clientRedirectURIs := strings.Split(client.RedirectURIs, " ")
	if !slices.Contains(clientRedirectURIs, in.RedirectURI) {
		return output.Client{}, stackerror.NewUseCaseError(
			"AuthorizeClientUseCase",
			fmt.Errorf("invalid redirect_uri"),
			stackerror.WithMessage("the redirect_uri is not valid, please check your request"),
			stackerror.WithStatusCode(http.StatusBadRequest),
		)
	}

	login, _ := u.getLoginByClientIDRepository.Execute(ctx, client.ID)
	signup, _ := u.getSignupByClientIDRepository.Execute(ctx, client.ID)

	return output.NewClient(
		client.ID,
		client.ClientID,
		client.Name,
		login,
		signup,
	), nil
}

func NewAuthorizeClientUseCase(
	getClientByClientIDRepository contract.GetClientByClientIDRepository,
	getLoginByClientIDRepository contract.GetLoginByClientIDRepository,
	getSignupByClientIDRepository contract.GetSignupByClientIDRepository,
) contract.AuthorizeClient {
	return &authorizeClientUseCase{
		getClientByClientIDRepository: getClientByClientIDRepository,
		getLoginByClientIDRepository:  getLoginByClientIDRepository,
		getSignupByClientIDRepository: getSignupByClientIDRepository,
	}
}
