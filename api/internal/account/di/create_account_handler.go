package di

import (
	"github.com/booscaaa/locksmith/api/internal/account/contract"
	"github.com/booscaaa/locksmith/api/internal/account/handler"
	"github.com/booscaaa/locksmith/api/internal/account/repository"
	"github.com/booscaaa/locksmith/api/internal/account/usecase"
	aclRepository "github.com/booscaaa/locksmith/api/internal/acl/repository"
	"github.com/booscaaa/locksmith/api/internal/adapter/database"
)

func NewCreateAccountHandler() contract.CreateAccountHandler {
	return handler.NewCreateAccountHandler(
		usecase.NewCreateAccountUseCase(
			repository.NewCreateAccountRepository(database.GetConnection()),
			repository.NewGetAccountByEmailAndProjectIDRepository(database.GetConnection()),
			aclRepository.NewGetProjectDomainByProjectIDRepository(database.GetConnection()),
		),
	)
}
