package domain

import "github.com/locksmithhq/locksmith/api/internal/core/types/database"

type Account struct {
	ID                 string        `json:"id" db:"id" paginate:"id"`
	ProjectID          string        `json:"project_id" db:"project_id" paginate:"project_id"`
	Name               string        `json:"name" db:"name" paginate:"name"`
	Email              string        `json:"email" db:"email" paginate:"email"`
	Username           string        `json:"username" db:"username" paginate:"username"`
	Password           string        `json:"-" db:"password" paginate:"-"`
	CreatedAt          string        `json:"created_at" db:"created_at" paginate:"created_at"`
	UpdatedAt          string        `json:"updated_at" db:"updated_at" paginate:"updated_at"`
	DeletedAt          database.Null `json:"deleted_at" db:"deleted_at" paginate:"deleted_at"`
	RoleName           string        `json:"role_name" db:"role_name" paginate:"role_name"`
	MustChangePassword bool          `json:"must_change_password" db:"must_change_password" paginate:"must_change_password"`
}

type AccountOption func(*Account)

func WithID(id string) AccountOption {
	return func(a *Account) {
		a.ID = id
	}
}

func WithProjectID(projectID string) AccountOption {
	return func(a *Account) {
		a.ProjectID = projectID
	}
}

func WithName(name string) AccountOption {
	return func(a *Account) {
		a.Name = name
	}
}

func WithEmail(email string) AccountOption {
	return func(a *Account) {
		a.Email = email
	}
}

func WithUsername(username string) AccountOption {
	return func(a *Account) {
		a.Username = username
	}
}

func WithPassword(password string) AccountOption {
	return func(a *Account) {
		a.Password = password
	}
}

func WithCreatedAt(createdAt string) AccountOption {
	return func(a *Account) {
		a.CreatedAt = createdAt
	}
}

func WithUpdatedAt(updatedAt string) AccountOption {
	return func(a *Account) {
		a.UpdatedAt = updatedAt
	}
}

func WithDeletedAt(deletedAt database.Null) AccountOption {
	return func(a *Account) {
		a.DeletedAt = deletedAt
	}
}

func WithRoleName(roleName string) AccountOption {
	return func(a *Account) {
		a.RoleName = roleName
	}
}

func WithMustChangePassword(mustChangePassword bool) AccountOption {
	return func(a *Account) {
		a.MustChangePassword = mustChangePassword
	}
}

func NewAccount(
	options ...AccountOption,
) Account {
	account := Account{}

	for _, option := range options {
		option(&account)
	}

	return account
}
