package domain

import "time"

type Module struct {
	Id        string    `json:"id" db:"id"`
	Title     string    `json:"title" db:"title"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type ModuleOption func(*Module)

func WithModuleId(id string) ModuleOption {
	return func(m *Module) {
		m.Id = id
	}
}

func WithModuleTitle(title string) ModuleOption {
	return func(m *Module) {
		m.Title = title
	}
}

func WithModuleCreatedAt(createdAt time.Time) ModuleOption {
	return func(m *Module) {
		m.CreatedAt = createdAt
	}
}

func WithModuleUpdatedAt(updatedAt time.Time) ModuleOption {
	return func(m *Module) {
		m.UpdatedAt = updatedAt
	}
}

func NewModule(opts ...ModuleOption) Module {
	module := Module{}
	for _, opt := range opts {
		opt(&module)
	}
	return module
}
