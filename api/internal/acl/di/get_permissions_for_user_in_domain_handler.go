package di

import (
	"github.com/booscaaa/locksmith/api/internal/acl/contract"
	"github.com/booscaaa/locksmith/api/internal/acl/handler"
)

func NewGetPermissionsForUserInDomainHandler() contract.GetPermissionsForUserInDomainHandler {
	return handler.NewGetPermissionsForUserInDomain()
}
