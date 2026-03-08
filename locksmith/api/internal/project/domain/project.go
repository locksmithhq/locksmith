package domain

import "github.com/locksmithhq/locksmith/api/internal/core/types/database"

type Project struct {
	ID          string        `json:"id" db:"id" paginate:"id"`
	Name        string        `json:"name" db:"name" paginate:"name"`
	Description string        `json:"description" db:"description" paginate:"description"`
	CreatedAt   string        `json:"created_at" db:"created_at" paginate:"created_at"`
	UpdatedAt   string        `json:"updated_at" db:"updated_at" paginate:"updated_at"`
	DeletedAt   database.Null `json:"deleted_at" db:"deleted_at" paginate:"deleted_at"`
	Domain      string        `json:"domain" db:"domain" paginate:"domain"`
}

type ProjectOption func(*Project)

func WithID(id string) ProjectOption {
	return func(p *Project) {
		p.ID = id
	}
}

func WithCreatedAt(createdAt string) ProjectOption {
	return func(p *Project) {
		p.CreatedAt = createdAt
	}
}

func WithUpdatedAt(updatedAt string) ProjectOption {
	return func(p *Project) {
		p.UpdatedAt = updatedAt
	}
}

func WithDeletedAt(deletedAt database.Null) ProjectOption {
	return func(p *Project) {
		p.DeletedAt = deletedAt
	}
}

func WithName(name string) ProjectOption {
	return func(p *Project) {
		p.Name = name
	}
}

func WithDescription(description string) ProjectOption {
	return func(p *Project) {
		p.Description = description
	}
}

func WithDomain(domain string) ProjectOption {
	return func(p *Project) {
		p.Domain = domain
	}
}

func NewProject(options ...ProjectOption) Project {
	project := &Project{}
	for _, option := range options {
		option(project)
	}
	return *project
}
