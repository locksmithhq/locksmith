package contract

import (
	"context"

	"github.com/locksmithhq/locksmith/api/internal/oauth_clients/types/output"
)

type ConfigDefaultClientCMD interface {
	Execute(context.Context, string) (output.Client, error)
}
