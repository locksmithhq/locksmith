package config

import (
	"context"
	"errors"
	"fmt"
	"os"

	diAccount "github.com/locksmithhq/locksmith/api/internal/account/di"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	diOAuthClients "github.com/locksmithhq/locksmith/api/internal/oauth_clients/di"
	diProject "github.com/locksmithhq/locksmith/api/internal/project/di"
)

func exitOnSeedError(label string, err error) {
	if err == nil {
		return
	}
	var useCaseError stackerror.UseCaseError
	if errors.As(err, &useCaseError) {
		useCaseError.StdoutResponse(label)
	} else {
		fmt.Fprintln(os.Stderr, label+":", err)
	}
	os.Exit(1)
}

func InitializeSeeder(ctx context.Context) {
	project, err := diProject.NewConfigDefaultProjectCMD().Execute(ctx)
	exitOnSeedError("CONFIG: InitializeSeeder project", err)

	_, err = diOAuthClients.NewConfigDefaultClientCMD().Execute(ctx, project.ID)
	exitOnSeedError("CONFIG: InitializeSeeder client", err)

	err = diAccount.NewConfigDefaultAccountCMD().Execute(ctx, project.ID)
	exitOnSeedError("CONFIG: InitializeSeeder account", err)

	seedProjects(ctx)
}
