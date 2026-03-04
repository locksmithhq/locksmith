package di

import (
	"github.com/locksmithhq/locksmith/api/internal/acl/contract"
	"github.com/locksmithhq/locksmith/api/internal/acl/handler"
)

func NewGetPermissionsForUserInDomainHandler() contract.GetPermissionsForUserInDomainHandler {
	return handler.NewGetPermissionsForUserInDomain()
}
