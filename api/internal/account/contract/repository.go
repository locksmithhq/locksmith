package contract

import (
	"context"

	"github.com/booscaaa/go-paginate/v3/paginate"
	"github.com/locksmithhq/locksmith/api/internal/account/domain"
)

type CreateAccountRepository interface {
	Execute(context.Context, domain.Account) (domain.Account, error)
}

type GetAccountByEmailAndProjectIDRepository interface {
	Execute(context.Context, string, string) (domain.Account, error)
}

type FetchAccountsByProjectIDRepository interface {
	Execute(context.Context, string, paginate.PaginationParams) ([]domain.Account, error)
}

type UpdateAccountRepository interface {
	Execute(context.Context, domain.Account) (domain.Account, error)
}

type UpdateAccountPasswordRepository interface {
	Execute(ctx context.Context, id string, password string) error
}

type GetAccountByProjectIDAndIDRepository interface {
	Execute(context.Context, string, string) (domain.Account, error)
}

type CountAccountsByProjectIDRepository interface {
	Execute(context.Context, string, paginate.PaginationParams) (int64, error)
}

type FetchSocialProvidersByAccountIDsRepository interface {
	Execute(context.Context, []string) (map[string][]string, error)
}
