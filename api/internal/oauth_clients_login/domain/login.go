package domain

import "github.com/locksmithhq/locksmith/api/internal/core/types/database"

type Login struct {
	ID                 string        `json:"id" db:"id" paginate:"id"`
	ClientID           string        `json:"client_id" db:"client_id" paginate:"client_id"`
	CustomCSS          string        `json:"custom_css" db:"custom_css" paginate:"custom_css"`
	CustomHTML         string        `json:"custom_html" db:"custom_html" paginate:"custom_html"`
	InputVariant       string        `json:"input_variant" db:"input_variant" paginate:"input_variant"`
	Layout             string        `json:"layout" db:"layout" paginate:"layout"`
	ShowForgotPassword bool          `json:"show_forgot_password" db:"show_forgot_password" paginate:"show_forgot_password"`
	ShowRememberMe     bool          `json:"show_remember_me" db:"show_remember_me" paginate:"show_remember_me"`
	ShowSignUp         bool          `json:"show_sign_up" db:"show_sign_up" paginate:"show_sign_up"`
	ShowSocial         bool          `json:"show_social" db:"show_social" paginate:"show_social"`
	UseCustomHTML      bool          `json:"use_custom_html" db:"use_custom_html" paginate:"use_custom_html"`
	Enabled            bool          `json:"enabled" db:"enabled" paginate:"enabled"`
	BackgroundColor    database.Null `json:"background_color" db:"background_color" paginate:"background_color"`
	BackgroundImage    database.Null `json:"background_image" db:"background_image" paginate:"background_image"`
	BackgroundType     database.Null `json:"background_type" db:"background_type" paginate:"background_type"`
	PrimaryColor       database.Null `json:"primary_color" db:"primary_color" paginate:"primary_color"`
	LogoURL            database.Null `json:"logo_url" db:"logo_url" paginate:"logo_url"`
	CreatedAt          string        `json:"created_at" db:"created_at" paginate:"created_at"`
	UpdatedAt          string        `json:"updated_at" db:"updated_at" paginate:"updated_at"`
}

func WithID(id string) func(*Login) error {
	return func(login *Login) error {
		login.ID = id
		return nil
	}
}

func WithClientID(clientID string) func(*Login) error {
	return func(login *Login) error {
		login.ClientID = clientID
		return nil
	}
}

func WithCustomCSS(customCSS string) func(*Login) error {
	return func(login *Login) error {
		login.CustomCSS = customCSS
		return nil
	}
}

func WithCustomHTML(customHTML string) func(*Login) error {
	return func(login *Login) error {
		login.CustomHTML = customHTML
		return nil
	}
}

func WithInputVariant(inputVariant string) func(*Login) error {
	return func(login *Login) error {
		login.InputVariant = inputVariant
		return nil
	}
}

func WithLayout(layout string) func(*Login) error {
	return func(login *Login) error {
		login.Layout = layout
		return nil
	}
}

func WithShowForgotPassword(showForgotPassword bool) func(*Login) error {
	return func(login *Login) error {
		login.ShowForgotPassword = showForgotPassword
		return nil
	}
}

func WithShowRememberMe(showRememberMe bool) func(*Login) error {
	return func(login *Login) error {
		login.ShowRememberMe = showRememberMe
		return nil
	}
}

func WithShowSignUp(showSignUp bool) func(*Login) error {
	return func(login *Login) error {
		login.ShowSignUp = showSignUp
		return nil
	}
}

func WithShowSocial(showSocial bool) func(*Login) error {
	return func(login *Login) error {
		login.ShowSocial = showSocial
		return nil
	}
}

func WithUseCustomHTML(useCustomHTML bool) func(*Login) error {
	return func(login *Login) error {
		login.UseCustomHTML = useCustomHTML
		return nil
	}
}

func WithEnabled(enabled bool) func(*Login) error {
	return func(login *Login) error {
		login.Enabled = enabled
		return nil
	}
}

func WithBackgroundColor(backgroundColor string) func(*Login) error {
	return func(login *Login) error {
		login.BackgroundColor = database.ParseNull(backgroundColor)
		return nil
	}
}

func WithBackgroundImage(backgroundImage string) func(*Login) error {
	return func(login *Login) error {
		login.BackgroundImage = database.ParseNull(backgroundImage)
		return nil
	}
}

func WithBackgroundType(backgroundType string) func(*Login) error {
	return func(login *Login) error {
		login.BackgroundType = database.ParseNull(backgroundType)
		return nil
	}
}

func WithPrimaryColor(primaryColor string) func(*Login) error {
	return func(login *Login) error {
		login.PrimaryColor = database.ParseNull(primaryColor)
		return nil
	}
}

func WithLogoURL(logoURL string) func(*Login) error {
	return func(login *Login) error {
		login.LogoURL = database.ParseNull(logoURL)
		return nil
	}
}

func WithCreatedAt(createdAt string) func(*Login) error {
	return func(login *Login) error {
		login.CreatedAt = createdAt
		return nil
	}
}

func WithUpdatedAt(updatedAt string) func(*Login) error {
	return func(login *Login) error {
		login.UpdatedAt = updatedAt
		return nil
	}
}

func NewLogin(
	options ...func(*Login) error,
) Login {
	login := &Login{}
	for _, option := range options {
		option(login)
	}
	return *login
}
