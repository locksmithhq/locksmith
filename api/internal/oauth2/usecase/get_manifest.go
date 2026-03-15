package usecase

import (
	"context"
	"net/http"
	"net/url"

	"github.com/locksmithhq/locksmith/api/internal/core/types/database"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/types/output"
)

type getPWAManifestUseCase struct {
	getClientByClientIDRepository contract.GetClientByClientIDRepository
	getLoginByClientIDRepository  contract.GetLoginByClientIDRepository
}

func nullString(n database.Null) string {
	v, _ := n.Value()
	if s, ok := v.(string); ok {
		return s
	}
	return ""
}

func (u *getPWAManifestUseCase) Execute(ctx context.Context, clientID, redirectURI, locale string) (output.Manifest, error) {
	client, err := u.getClientByClientIDRepository.Execute(ctx, clientID)
	if err != nil {
		return output.Manifest{}, stackerror.NewUseCaseError(
			"GetPWAManifestUseCase",
			err,
			stackerror.WithMessage("the client_id is not valid"),
			stackerror.WithStatusCode(http.StatusBadRequest),
		)
	}

	login, _ := u.getLoginByClientIDRepository.Execute(ctx, client.ID)

	themeColor := "#1a1a2e"
	if s := nullString(login.PrimaryColor); s != "" {
		themeColor = s
	}

	bgColor := "#1a1a2e"
	if s := nullString(login.BackgroundColor); s != "" {
		bgColor = s
	}

	iconSrc := "/locksmith.png"
	if s := nullString(login.FaviconURL); s != "" {
		iconSrc = "/api/oauth2/favicon?client_id=" + clientID
	}

	startURL := "/" + locale + "/auth?client_id=" + url.QueryEscape(clientID)
	if redirectURI != "" {
		startURL += "&redirect_uri=" + url.QueryEscape(redirectURI)
	}

	return output.Manifest{
		Name:            client.Name,
		ShortName:       client.Name,
		Description:     client.Name,
		ThemeColor:      themeColor,
		BackgroundColor: bgColor,
		Display:         "standalone",
		StartURL:        startURL,
		Icons: []output.ManifestIcon{
			{Src: iconSrc, Sizes: "192x192", Type: "image/png"},
			{Src: iconSrc, Sizes: "512x512", Type: "image/png"},
		},
	}, nil
}

func NewGetPWAManifestUseCase(
	getClientByClientIDRepository contract.GetClientByClientIDRepository,
	getLoginByClientIDRepository contract.GetLoginByClientIDRepository,
) contract.GetPWAManifestUseCase {
	return &getPWAManifestUseCase{
		getClientByClientIDRepository: getClientByClientIDRepository,
		getLoginByClientIDRepository:  getLoginByClientIDRepository,
	}
}
