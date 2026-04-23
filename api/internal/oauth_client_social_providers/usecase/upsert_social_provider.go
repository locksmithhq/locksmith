package usecase

import (
	"context"
	"errors"
	"net/http"

	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/oauth_client_social_providers/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth_client_social_providers/domain"
	"github.com/locksmithhq/locksmith/api/internal/oauth_client_social_providers/types/input"
	"github.com/locksmithhq/locksmith/api/internal/oauth_client_social_providers/types/output"
)

const maskedSentinel = "****"

type upsertSocialProviderUseCase struct {
	upsertSocialProviderRepository            contract.UpsertSocialProviderRepository
	getSocialProviderByClientAndProviderRepository contract.GetSocialProviderByClientAndProviderRepository
}

func (u *upsertSocialProviderUseCase) Execute(ctx context.Context, clientID string, provider string, in input.SocialProvider) (output.SocialProvider, error) {
	if in.ClientKey == "" {
		return output.SocialProvider{}, stackerror.NewUseCaseError(
			"UpsertSocialProviderUseCase",
			errors.New("client_key is required"),
			stackerror.WithMessage("client_key is required"),
			stackerror.WithStatusCode(http.StatusBadRequest),
		)
	}
	if in.ClientSecret == "" {
		return output.SocialProvider{}, stackerror.NewUseCaseError(
			"UpsertSocialProviderUseCase",
			errors.New("client_secret is required"),
			stackerror.WithMessage("client_secret is required"),
			stackerror.WithStatusCode(http.StatusBadRequest),
		)
	}

	clientKey := in.ClientKey
	clientSecret := in.ClientSecret

	// If either field carries the masked sentinel, fetch the real values from DB.
	if clientKey == maskedSentinel || clientSecret == maskedSentinel {
		existing, err := u.getSocialProviderByClientAndProviderRepository.Execute(ctx, clientID, provider)
		if err != nil {
			return output.SocialProvider{}, stackerror.NewUseCaseError(
				"UpsertSocialProviderUseCase",
				errors.New("cannot resolve masked credentials: provider not found"),
				stackerror.WithMessage("cannot update credentials: provider not configured yet"),
				stackerror.WithStatusCode(http.StatusBadRequest),
			)
		}
		if clientKey == maskedSentinel {
			clientKey = existing.ClientKey
		}
		if clientSecret == maskedSentinel {
			clientSecret = existing.ClientSecret
		}
	}

	scopes := in.Scopes
	if scopes == "" {
		scopes = "email profile"
	}

	result, err := u.upsertSocialProviderRepository.Execute(ctx, domain.SocialProvider{
		ClientID:     clientID,
		Provider:     provider,
		ClientKey:    clientKey,
		ClientSecret: clientSecret,
		Enabled:      in.Enabled,
		Scopes:       scopes,
	})
	if err != nil {
		return output.SocialProvider{}, stackerror.NewUseCaseError("UpsertSocialProviderUseCase", err)
	}
	return output.NewSocialProviderFromDomain(result), nil
}

func NewUpsertSocialProviderUseCase(
	upsertSocialProviderRepository contract.UpsertSocialProviderRepository,
	getSocialProviderByClientAndProviderRepository contract.GetSocialProviderByClientAndProviderRepository,
) contract.UpsertSocialProviderUseCase {
	return &upsertSocialProviderUseCase{
		upsertSocialProviderRepository:            upsertSocialProviderRepository,
		getSocialProviderByClientAndProviderRepository: getSocialProviderByClientAndProviderRepository,
	}
}
