package domain

import "time"

type ProjectAcl struct {
	Id         string    `json:"id" db:"id"`
	RoleId     string    `json:"role_id" db:"role_id"`
	RoleName   string    `json:"role_name" db:"role_name"`
	ModuleId   string    `json:"module_id" db:"module_id"`
	ModuleName string    `json:"module_name" db:"module_name"`
	ActionId   string    `json:"action_id" db:"action_id"`
	ActionName string    `json:"action_name" db:"action_name"`
	ProjectId  string    `json:"project_id" db:"project_id"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}

type ProjectAclOption func(*ProjectAcl)

func WithProjectAclId(id string) ProjectAclOption {
	return func(p *ProjectAcl) {
		p.Id = id
	}
}

func WithProjectAclRoleId(roleId string) ProjectAclOption {
	return func(p *ProjectAcl) {
		p.RoleId = roleId
	}
}

func WithProjectAclModuleId(moduleId string) ProjectAclOption {
	return func(p *ProjectAcl) {
		p.ModuleId = moduleId
	}
}

func WithProjectAclActionId(actionId string) ProjectAclOption {
	return func(p *ProjectAcl) {
		p.ActionId = actionId
	}
}

func WithProjectAclProjectId(projectId string) ProjectAclOption {
	return func(p *ProjectAcl) {
		p.ProjectId = projectId
	}
}

func WithProjectAclCreatedAt(createdAt time.Time) ProjectAclOption {
	return func(p *ProjectAcl) {
		p.CreatedAt = createdAt
	}
}

func WithProjectAclUpdatedAt(updatedAt time.Time) ProjectAclOption {
	return func(p *ProjectAcl) {
		p.UpdatedAt = updatedAt
	}
}

func NewProjectAcl(opts ...ProjectAclOption) *ProjectAcl {
	projectAcl := &ProjectAcl{}
	for _, opt := range opts {
		opt(projectAcl)
	}
	return projectAcl
}
