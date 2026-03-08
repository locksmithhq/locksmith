package di

import (
	"github.com/locksmithhq/locksmith/api/internal/account/contract"
	"github.com/locksmithhq/locksmith/api/internal/account/handler"
	"github.com/locksmithhq/locksmith/api/internal/account/repository"
	"github.com/locksmithhq/locksmith/api/internal/account/usecase"
	"github.com/locksmithhq/locksmith/api/internal/adapter/database"
	oauth2Repository "github.com/locksmithhq/locksmith/api/internal/oauth2/repository"
)

func NewChangePasswordHandler() contract.ChangePasswordHandler {
	database := database.GetConnection()
	updateAccountPasswordRepository := repository.NewUpdateAccountPasswordRepository(database)
	getClientByClientIDRepository := oauth2Repository.NewGetClientByClientIDRepository(database)
	changePasswordUseCase := usecase.NewChangePasswordUseCase(
		updateAccountPasswordRepository,
		getClientByClientIDRepository,
	)
	return handler.NewChangePasswordHandler(changePasswordUseCase)
}
