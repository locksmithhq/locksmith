package domain

import "time"

type Action struct {
	Id        string    `json:"id" db:"id"`
	Title     string    `json:"title" db:"title"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type ActionOption func(*Action)

func WithActionId(id string) ActionOption {
	return func(a *Action) {
		a.Id = id
	}
}

func WithActionTitle(title string) ActionOption {
	return func(a *Action) {
		a.Title = title
	}
}

func WithActionCreatedAt(createdAt time.Time) ActionOption {
	return func(a *Action) {
		a.CreatedAt = createdAt
	}
}

func WithActionUpdatedAt(updatedAt time.Time) ActionOption {
	return func(a *Action) {
		a.UpdatedAt = updatedAt
	}
}

func NewAction(opts ...ActionOption) Action {
	action := Action{}
	for _, opt := range opts {
		opt(&action)
	}
	return action
}
