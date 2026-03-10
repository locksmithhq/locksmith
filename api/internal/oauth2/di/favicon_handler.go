package di

import (
	"github.com/locksmithhq/locksmith/api/internal/adapter/database"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/handler"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/repository"
)

func NewFaviconHandler() contract.FaviconHandler {
	return handler.NewFaviconHandler(
		repository.NewGetClientByClientIDRepository(database.GetConnection()),
		repository.NewGetLoginByClientIDRepository(database.GetConnection()),
	)
}
