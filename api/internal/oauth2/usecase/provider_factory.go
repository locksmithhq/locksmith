package usecase

import (
	"fmt"
	"net/http"

	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/discord"
	"github.com/markbates/goth/providers/facebook"
	"github.com/markbates/goth/providers/github"
	"github.com/markbates/goth/providers/google"
	"github.com/markbates/goth/providers/linkedin"
)

// newGothProvider instantiates the correct goth.Provider for the given provider
// name. Add a new case here — and a matching one in newGothSession — to support
// additional OAuth2 providers without touching any other code.
func newGothProvider(provider, clientKey, clientSecret, callbackURL string, scopes ...string) (goth.Provider, error) {
	switch provider {
	case "google":
		return google.New(clientKey, clientSecret, callbackURL, scopes...), nil
	case "github":
		return github.New(clientKey, clientSecret, callbackURL, scopes...), nil
	case "facebook":
		return facebook.New(clientKey, clientSecret, callbackURL, scopes...), nil
	case "discord":
		return discord.New(clientKey, clientSecret, callbackURL, scopes...), nil
	case "linkedin":
		return linkedin.New(clientKey, clientSecret, callbackURL, scopes...), nil
	default:
		return nil, stackerror.NewUseCaseError(
			"ProviderFactory",
			fmt.Errorf("unsupported social provider: %q", provider),
			stackerror.WithMessage(fmt.Sprintf("social provider %q is not supported", provider)),
			stackerror.WithStatusCode(http.StatusBadRequest),
		)
	}
}

// newGothSession returns an empty session of the correct type for the provider.
// Used during the OAuth2 callback to exchange the authorization code for tokens.
func newGothSession(provider string) (goth.Session, error) {
	switch provider {
	case "google":
		return &google.Session{}, nil
	case "github":
		return &github.Session{}, nil
	case "facebook":
		return &facebook.Session{}, nil
	case "discord":
		return &discord.Session{}, nil
	case "linkedin":
		return &linkedin.Session{}, nil
	default:
		return nil, stackerror.NewUseCaseError(
			"ProviderFactory",
			fmt.Errorf("unsupported social provider: %q", provider),
			stackerror.WithMessage(fmt.Sprintf("social provider %q is not supported", provider)),
			stackerror.WithStatusCode(http.StatusBadRequest),
		)
	}
}
