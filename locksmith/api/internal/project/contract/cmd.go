package contract

import (
	"context"

	"github.com/locksmithhq/locksmith/api/internal/project/types/output"
)

type ConfigDefaultProjectCMD interface {
	Execute(context.Context) (output.Project, error)
}
