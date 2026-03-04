package di

import (
	"github.com/locksmithhq/locksmith/api/internal/acl/contract"
	"github.com/locksmithhq/locksmith/api/internal/acl/handler"
	"github.com/locksmithhq/locksmith/api/internal/acl/repository"
	"github.com/locksmithhq/locksmith/api/internal/acl/usecase"
	"github.com/locksmithhq/locksmith/api/internal/adapter/database"
)

func NewCreateProjectAclHandler() contract.CreateProjectAclHandler {
	return handler.NewCreateProjectAclHandler(
		usecase.NewCreateProjectAclUseCase(
			repository.NewDeleteProjectAclByProjectIdRepository(database.GetConnection()),
			repository.NewCreateProjectAclRepository(database.GetConnection()),
			repository.NewGetProjectDomainByProjectIDRepository(database.GetConnection()),
		),
	)
}
