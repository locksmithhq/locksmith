package domain

import "github.com/locksmithhq/locksmith/api/internal/core/types/database"

type Account struct {
	ID                 string        `json:"id" db:"id"`
	ProjectID          string        `json:"project_id" db:"project_id"`
	Name               string        `json:"name" db:"name"`
	Email              string        `json:"email" db:"email"`
	Username           string        `json:"username" db:"username"`
	Password           string        `json:"password" db:"password"`
	RoleName           string        `json:"role_name" db:"role_name"`
	CreatedAt          string        `json:"created_at" db:"created_at"`
	UpdatedAt          string        `json:"updated_at" db:"updated_at"`
	DeletedAt          database.Null `json:"deleted_at,omitempty" db:"deleted_at"`
	MustChangePassword bool          `json:"must_change_password" db:"must_change_password"`
}
