package di

import (
	"github.com/locksmithhq/locksmith/api/internal/locksmith/contract"
	"github.com/locksmithhq/locksmith/api/internal/locksmith/handler"
)

func NewStatusHandler() contract.StatusHandler {
	return handler.NewStatusHandler()
}
