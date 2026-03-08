package usecase

import (
	"context"

	"github.com/locksmithhq/locksmith/api/internal/acl/contract"
	"github.com/locksmithhq/locksmith/api/internal/acl/domain"
	"github.com/locksmithhq/locksmith/api/internal/acl/types/output"
)

type fetchProjectAclUseCase struct {
	fetchProjectAclRepository contract.FetchProjectAclRepository
}

func (u *fetchProjectAclUseCase) Execute(ctx context.Context, projectId string) (output.ProjectAcl, error) {
	projectAcls, err := u.fetchProjectAclRepository.Execute(ctx, projectId)
	if err != nil {
		return output.ProjectAcl{}, err
	}

	return GroupPermissions(projectAcls), nil
}

func NewFetchProjectAclUseCase(fetchProjectAclRepository contract.FetchProjectAclRepository) contract.FetchProjectAclUseCase {
	return &fetchProjectAclUseCase{
		fetchProjectAclRepository: fetchProjectAclRepository,
	}
}

func GroupPermissions(permissions []domain.ProjectAcl) output.ProjectAcl {
	roleMap := make(map[string]*output.Role)

	for _, p := range permissions {
		// Verifica se a role já existe
		role, ok := roleMap[p.RoleId]
		if !ok {
			role = &output.Role{
				Id:      p.RoleId,
				Title:   p.RoleName,
				Modules: []output.Module{},
			}
			roleMap[p.RoleId] = role
		}

		// Verifica se o módulo já existe dentro da role
		var module *output.Module
		for i := range role.Modules {
			if role.Modules[i].Id == p.ModuleId {
				module = &role.Modules[i]
				break
			}
		}
		if module == nil {
			newModule := output.Module{
				Id:      p.ModuleId,
				Title:   p.ModuleName,
				Actions: []output.Action{},
			}
			role.Modules = append(role.Modules, newModule)
			module = &role.Modules[len(role.Modules)-1]
		}

		// Adiciona a action se ainda não existir
		actionExists := false
		for _, a := range module.Actions {
			if a.Id == p.ActionId {
				actionExists = true
				break
			}
		}
		if !actionExists {
			module.Actions = append(module.Actions, output.Action{
				Id:    p.ActionId,
				Title: p.ActionName,
			})
		}
	}

	// Converte o map em slice
	roles := make([]output.Role, 0, len(roleMap))
	for _, r := range roleMap {
		roles = append(roles, *r)
	}

	return output.ProjectAcl{Roles: roles}
}
