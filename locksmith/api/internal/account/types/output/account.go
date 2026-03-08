package output

import (
	"github.com/locksmithhq/locksmith/api/internal/account/domain"

	"github.com/locksmithhq/locksmith/api/internal/core/types/database"
)

type Action struct {
	Role   string `json:"role"`
	Domain string `json:"domain"`
	Module string `json:"module"`
	Action string `json:"action"`
}

type Domain struct {
	Name    string   `json:"name"`
	Roles   []string `json:"roles"`
	Actions []Action `json:"actions"`
}

type Account struct {
	ID                 string        `json:"id"`
	ProjectID          string        `json:"project_id"`
	Name               string        `json:"name"`
	Email              string        `json:"email"`
	Username           string        `json:"username"`
	CreatedAt          string        `json:"created_at"`
	UpdatedAt          string        `json:"updated_at"`
	DeletedAt          database.Null `json:"deleted_at,omitempty"`
	RoleName           string        `json:"role_name"`
	MustChangePassword bool          `json:"must_change_password"`
}

func NewAccountFromDomain(
	account domain.Account,
) Account {
	return Account{
		ID:                 account.ID,
		ProjectID:          account.ProjectID,
		Name:               account.Name,
		Email:              account.Email,
		Username:           account.Username,
		CreatedAt:          account.CreatedAt,
		UpdatedAt:          account.UpdatedAt,
		DeletedAt:          account.DeletedAt,
		RoleName:           account.RoleName,
		MustChangePassword: account.MustChangePassword,
	}
}

func NewAccountsFromDomain(accounts []domain.Account) []Account {
	out := make([]Account, len(accounts))
	for i, account := range accounts {
		out[i] = NewAccountFromDomain(account)
	}
	return out
}
