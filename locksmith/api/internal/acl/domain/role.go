package domain

import "time"

type Role struct {
	Id        string    `json:"id" db:"id"`
	Title     string    `json:"title" db:"title"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type RoleOption func(*Role)

func WithRoleId(id string) RoleOption {
	return func(r *Role) {
		r.Id = id
	}
}

func WithRoleTitle(title string) RoleOption {
	return func(r *Role) {
		r.Title = title
	}
}

func WithRoleCreatedAt(createdAt time.Time) RoleOption {
	return func(r *Role) {
		r.CreatedAt = createdAt
	}
}

func WithRoleUpdatedAt(updatedAt time.Time) RoleOption {
	return func(r *Role) {
		r.UpdatedAt = updatedAt
	}
}

func NewRole(opts ...RoleOption) Role {
	role := Role{}
	for _, opt := range opts {
		opt(&role)
	}
	return role
}
