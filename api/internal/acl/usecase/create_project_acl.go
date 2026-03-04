package usecase

import (
	"context"

	"github.com/booscaaa/initializers/postgres/uow"
	"github.com/locksmithhq/locksmith/api/internal/acl/contract"
	"github.com/locksmithhq/locksmith/api/internal/acl/domain"
	"github.com/locksmithhq/locksmith/api/internal/acl/types/input"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith-go"
)

type createProjectAclUseCase struct {
	deleteProjectAclByProjectIdRepository contract.DeleteProjectAclByProjectIdRepository
	createProjectAclRepository            contract.CreateProjectAclRepository
	getProjectDomainByProjectIDRepository contract.GetProjectDomainByProjectIDRepository
}

func (u *createProjectAclUseCase) Execute(ctx context.Context, projectId string, in input.ProjectAcl) error {
	err := uow.WithTransaction(ctx, func(ctx context.Context) error {
		projectDomain, err := u.getProjectDomainByProjectIDRepository.Execute(ctx, projectId)
		if err != nil {
			return err
		}

		if err := u.deleteProjectAclByProjectIdRepository.Execute(ctx, projectId); err != nil {
			return err
		}

		if _, err := locksmith.RemoveFilteredPolicy(1, projectDomain); err != nil {
			return err
		}

		for _, role := range in.Roles {
			for _, module := range role.Modules {
				for _, action := range module.Actions {
					if err := u.createProjectAclRepository.Execute(ctx, domain.ProjectAcl{
						RoleId:    role.Id,
						ModuleId:  module.Id,
						ActionId:  action.Id,
						ProjectId: projectId,
					}); err != nil {
						return err
					}

					if _, err := locksmith.AddPolicy(role.Title, projectDomain, module.Title, action.Title); err != nil {
						return err
					}
				}
			}
		}
		return nil
	})

	if err != nil {
		return stackerror.NewUseCaseError("CreateProjectAclUseCase", err)
	}

	if _, err := locksmith.AddPolicy("role:admin", "domain:locksmith", "*", "*"); err != nil {
		return err
	}

	return nil
}

func NewCreateProjectAclUseCase(
	deleteProjectAclByProjectIdRepository contract.DeleteProjectAclByProjectIdRepository,
	createProjectAclRepository contract.CreateProjectAclRepository,
	getProjectDomainByProjectIDRepository contract.GetProjectDomainByProjectIDRepository,
) contract.CreateProjectAclUseCase {
	return &createProjectAclUseCase{
		deleteProjectAclByProjectIdRepository: deleteProjectAclByProjectIdRepository,
		createProjectAclRepository:            createProjectAclRepository,
		getProjectDomainByProjectIDRepository: getProjectDomainByProjectIDRepository,
	}
}
