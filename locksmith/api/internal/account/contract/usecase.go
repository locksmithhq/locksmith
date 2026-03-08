package contract

import (
	"context"

	"github.com/booscaaa/go-paginate/v3/paginate"
	"github.com/locksmithhq/locksmith/api/internal/account/types/input"
	"github.com/locksmithhq/locksmith/api/internal/account/types/output"
)

type CreateAccountUseCase interface {
	Execute(context.Context, input.Account) (output.Account, error)
}

type FetchAccountsByProjectIDUseCase interface {
	Execute(context.Context, string, paginate.PaginationParams) ([]output.Account, error)
}

type UpdateAccountUseCase interface {
	Execute(context.Context, input.UpdateAccount) (output.Account, error)
}

type ChangePasswordUseCase interface {
	Execute(context.Context, input.ChangePassword) error
}

type GetAccountByProjectIDAndIDUseCase interface {
	Execute(context.Context, string, string) (output.Account, error)
}

type CountAccountsByProjectIDUseCase interface {
	Execute(context.Context, string, paginate.PaginationParams) (int64, error)
}
