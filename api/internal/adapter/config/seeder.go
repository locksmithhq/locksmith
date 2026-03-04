package config

import (
	"context"
	"errors"
	"os"

	diAccount "github.com/booscaaa/locksmith/api/internal/account/di"
	"github.com/booscaaa/locksmith/api/internal/core/types/stackerror"
	diOAuthClients "github.com/booscaaa/locksmith/api/internal/oauth_clients/di"
	diProject "github.com/booscaaa/locksmith/api/internal/project/di"
)

func InitializeSeeder(ctx context.Context) {
	var useCaseError stackerror.UseCaseError
	project, err := diProject.NewConfigDefaultProjectCMD().Execute(ctx)
	if errors.As(err, &useCaseError) {
		useCaseError.StdoutResponse("CONFIG: InitializeSeeder")
		os.Exit(1)
	}

	_, err = diOAuthClients.NewConfigDefaultClientCMD().Execute(ctx, project.ID)
	if errors.As(err, &useCaseError) {
		useCaseError.StdoutResponse("CONFIG: InitializeSeeder")
		os.Exit(1)
	}

	diAccount.NewConfigDefaultAccountCMD().Execute(ctx, project.ID)
}
