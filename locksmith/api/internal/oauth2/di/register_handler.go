package di

import (
	accountRepository "github.com/locksmithhq/locksmith/api/internal/account/repository"
	accountUseCase "github.com/locksmithhq/locksmith/api/internal/account/usecase"
	aclRepository "github.com/locksmithhq/locksmith/api/internal/acl/repository"
	"github.com/locksmithhq/locksmith/api/internal/adapter/database"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/handler"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/repository"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/usecase"
)

func NewRegisterHandler() contract.RegisterHandler {
	db := database.GetConnection()
	return handler.NewRegisterHandler(
		usecase.NewRegisterUseCase(
			repository.NewGetClientByClientIDRepository(db),
			repository.NewGetSignupByClientIDRepository(db),
			accountUseCase.NewCreateAccountUseCase(
				accountRepository.NewCreateAccountRepository(db),
				accountRepository.NewGetAccountByEmailAndProjectIDRepository(db),
				aclRepository.NewGetProjectDomainByProjectIDRepository(db),
			),
			repository.NewCreateAuthCodeRepository(db),
		),
	)
}
