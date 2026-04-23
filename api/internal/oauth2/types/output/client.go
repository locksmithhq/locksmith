package output

import "github.com/locksmithhq/locksmith/api/internal/oauth2/domain"

type Client struct {
	ID              string             `json:"id"`
	ClientID        string             `json:"client_id"`
	Name            string             `json:"name"`
	Login           *OAuthClientLogin  `json:"login"`
	Signup          *OAuthClientSignup `json:"signup"`
	SocialProviders []string           `json:"social_providers"`
}

func NewClient(
	id string,
	clientID string,
	name string,
	login domain.OAuthClientLogin,
	signup domain.OAuthClientSignup,
	socialProviders []string,
) Client {
	client := Client{
		ID:              id,
		ClientID:        clientID,
		Name:            name,
		SocialProviders: socialProviders,
	}
	if client.SocialProviders == nil {
		client.SocialProviders = []string{}
	}

	if login.ID != "" {
		client.Login = &OAuthClientLogin{
			ID:                 login.ID,
			ClientID:           login.ClientID,
			CustomCSS:          login.CustomCSS,
			CustomHTML:         login.CustomHTML,
			InputVariant:       login.InputVariant,
			Layout:             login.Layout,
			ShowForgotPassword: login.ShowForgotPassword,
			ShowRememberMe:     login.ShowRememberMe,
			ShowSignUp:         login.ShowSignUp,
			ShowSocial:         login.ShowSocial,
			UseCustomHTML:      login.UseCustomHTML,
			Enabled:            login.Enabled,
			CreatedAt:          login.CreatedAt,
			UpdatedAt:          login.UpdatedAt,
			BackgroundColor:    login.BackgroundColor,
			BackgroundImage:    login.BackgroundImage,
			BackgroundType:     login.BackgroundType,
			PrimaryColor:       login.PrimaryColor,
			LogoURL:            login.LogoURL,
			FaviconURL:         login.FaviconURL,
		}
	}

	if signup.ID != "" {
		client.Signup = &OAuthClientSignup{
			ID:              signup.ID,
			ClientID:        signup.ClientID,
			CustomCSS:       signup.CustomCSS,
			CustomHTML:      signup.CustomHTML,
			InputVariant:    signup.InputVariant,
			Layout:          signup.Layout,
			ShowSocial:      signup.ShowSocial,
			UseCustomHTML:   signup.UseCustomHTML,
			Enabled:         signup.Enabled,
			BackgroundColor: signup.BackgroundColor,
			BackgroundImage: signup.BackgroundImage,
			BackgroundType:  signup.BackgroundType,
			PrimaryColor:    signup.PrimaryColor,
			LogoURL:         signup.LogoURL,
			FaviconURL:      signup.FaviconURL,
			CreatedAt:       signup.CreatedAt,
			UpdatedAt:       signup.UpdatedAt,
		}
	}

	return client
}
