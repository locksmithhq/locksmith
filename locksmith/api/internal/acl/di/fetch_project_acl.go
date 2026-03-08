package di

import (
	"github.com/locksmithhq/locksmith/api/internal/acl/contract"
	"github.com/locksmithhq/locksmith/api/internal/acl/handler"
	"github.com/locksmithhq/locksmith/api/internal/acl/repository"
	"github.com/locksmithhq/locksmith/api/internal/acl/usecase"
	"github.com/locksmithhq/locksmith/api/internal/adapter/database"
)

func NewFetchProjectAclHandler() contract.FetchProjectAclHandler {
	return handler.NewFetchProjectAclHandler(
		usecase.NewFetchProjectAclUseCase(
			repository.NewFetchProjectAclRepository(
				database.GetConnection(),
			),
		),
	)
}
