package contract

import (
	"context"
)

type ConfigDefaultAccountCMD interface {
	Execute(context.Context, string) error
}
