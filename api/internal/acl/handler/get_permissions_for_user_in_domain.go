package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/booscaaa/locksmith/api/internal/acl/contract"
	"github.com/booscaaa/locksmith/api/internal/core/types/stackerror"
	"github.com/go-chi/chi/v5"
	"github.com/locksmithhq/locksmith-go"
)

type getPermissionsForUserInDomain struct{}

// Execute implements contract.GetPermissionsForUserInDomainHandler.
func (h *getPermissionsForUserInDomain) Execute(w http.ResponseWriter, r *http.Request) {
	user := chi.URLParam(r, "user")
	if user == "" {
		stackerror.HttpResponse(w, "GetPermissionsForUserInDomain", errors.New("user is required"))
		return
	}

	domain := chi.URLParam(r, "domain")
	if domain == "" {
		stackerror.HttpResponse(w, "GetPermissionsForUserInDomain", errors.New("domain is required"))
		return
	}

	// authenticatedUser := locksmith.GetSubFromContext(r.Context())
	// authenticatedUserDomain := locksmith.GetDomainFromContext(r.Context())
	// authenticatedUserSub := fmt.Sprintf("user:%s", authenticatedUser)

	// checkAction := "action:read:own"
	// if authenticatedUser != user {
	// 	checkAction = "action:read:all"
	// }

	// if allowed, err := locksmith.Enforce(authenticatedUserSub, authenticatedUserDomain, "module:acl", checkAction); err != nil || !allowed {
	// 	stackerror.HttpResponse(w, "GetPermissionsForUserInDomain", errors.New("forbidden: you do not have permission to read these permissions"))
	// 	return
	// }

	type actionOutput struct {
		Role   string `json:"role"`
		Domain string `json:"domain"`
		Module string `json:"module"`
		Action string `json:"action"`
	}

	type output struct {
		Permissions []actionOutput `json:"permissions"`
	}
	permissions := locksmith.GetPermissionsForUserInDomain(
		fmt.Sprintf("user:%s", user),
		domain,
	)

	permissionsOutput := make([]actionOutput, 0)
	for _, permission := range permissions {
		permissionsOutput = append(permissionsOutput, actionOutput{
			Role:   permission[0],
			Domain: permission[1],
			Module: permission[2],
			Action: permission[3],
		})
	}

	json.NewEncoder(w).Encode(output{Permissions: permissionsOutput})

}

func NewGetPermissionsForUserInDomain() contract.GetPermissionsForUserInDomainHandler {
	return &getPermissionsForUserInDomain{}
}
