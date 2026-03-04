package acl

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/locksmithhq/locksmith-go"
)

// import (
// 	"encoding/json"
// 	"net/http"
// 	"strings"

// 	"github.com/go-chi/chi/v5"
// )

// func GetPermissions(w http.ResponseWriter, r *http.Request) {
// 	type output struct {
// 		Actions []string `json:"actions"`
// 	} //
// 	permissions, err := Control.GetAllActions()
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	json.NewEncoder(w).Encode(output{Actions: permissions})
// }

// func GetDomains(w http.ResponseWriter, r *http.Request) {
// 	type output struct {
// 		Domains []string `json:"domains"`
// 	}
// 	domains, err := Control.GetAllDomains()
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	json.NewEncoder(w).Encode(output{Domains: domains})
// }

// func GetRoles(w http.ResponseWriter, r *http.Request) {
// 	domain := chi.URLParam(r, "domain")
// 	if domain == "" {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	type output struct {
// 		Roles []string `json:"roles"`
// 	}
// 	allRoles, err := Control.GetAllRolesByDomain(domain)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	roles := make([]string, 0)
// 	for _, role := range allRoles {
// 		// if strings.HasPrefix(role, "role:") {
// 		roles = append(roles, role)
// 		// }
// 	}

// 	json.NewEncoder(w).Encode(output{Roles: roles})
// }

// func GetUsers(w http.ResponseWriter, r *http.Request) {
// 	domain := chi.URLParam(r, "domain")
// 	if domain == "" {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	type actionOutput struct {
// 		Role   string `json:"role"`
// 		Domain string `json:"domain"`
// 		Module string `json:"module"`
// 		Action string `json:"action"`
// 	}

// 	type userOutput struct {
// 		ID      string         `json:"id"`
// 		Roles   []string       `json:"roles"`
// 		Actions []actionOutput `json:"actions"`
// 	}

// 	type output struct {
// 		Users []userOutput `json:"users"`
// 	}
// 	users, err := Control.GetAllUsersByDomain(domain)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	usersWithRoles := make([]userOutput, 0)
// 	for _, user := range users {
// 		if strings.HasPrefix(user, "user:") {
// 			roles, _ := Control.GetRolesForUser(user, domain)
// 			actions := Control.GetPermissionsForUserInDomain(user, domain)

// 			usersWithRoles = append(usersWithRoles, userOutput{
// 				ID:      user,
// 				Roles:   roles,
// 				Actions: make([]actionOutput, 0),
// 			})

// 			for _, action := range actions {
// 				usersWithRoles[len(usersWithRoles)-1].Actions = append(usersWithRoles[len(usersWithRoles)-1].Actions, actionOutput{
// 					Role:   action[0],
// 					Domain: action[1],
// 					Module: action[2],
// 					Action: action[3],
// 				})
// 			}
// 		}
// 	}

// 	json.NewEncoder(w).Encode(output{Users: usersWithRoles})
// }

// func GetObjects(w http.ResponseWriter, r *http.Request) {
// 	type output struct {
// 		Objects []string `json:"objects"`
// 	}
// 	objects, err := Control.GetAllObjects()
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	json.NewEncoder(w).Encode(output{Objects: objects})
// }

// func UpdateUserRole(w http.ResponseWriter, r *http.Request) {
// 	user := chi.URLParam(r, "user")
// 	if user == "" {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	domain := chi.URLParam(r, "domain")
// 	if domain == "" {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	newRole := chi.URLParam(r, "role")
// 	if newRole == "" {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	oldRole := r.URL.Query().Get("old_role")

// 	// Se o papel antigo não for fornecido, tentamos encontrar um papel existente para o usuário
// 	if oldRole == "" {
// 		roles, err := Control.GetRolesForUser(user, domain)
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			return
// 		}

// 		if len(roles) > 0 {
// 			oldRole = roles[0]
// 		}
// 	}

// 	if oldRole != "" {
// 		oldRule := []string{user, oldRole, domain}
// 		newRule := []string{user, newRole, domain}

// 		_, err := Control.UpdateGroupingPolicy(oldRule, newRule)
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			return
// 		}
// 	} else {
// 		_, err := Control.AddRoleForUser(user, newRole, domain)
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			return
// 		}
// 	}

// 	w.WriteHeader(http.StatusOK)
// }

// func UpdatePolicy(w http.ResponseWriter, r *http.Request) {
// 	type input struct {
// 		Old []string `json:"old"`
// 		New []string `json:"new"`
// 	}

// 	var data input
// 	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	if len(data.Old) != 4 || len(data.New) != 4 {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	_, err := Control.UpdatePolicy(data.Old, data.New)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// }

// func GetPolicies(w http.ResponseWriter, r *http.Request) {
// 	domain := chi.URLParam(r, "domain")
// 	if domain == "" {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	type actionOutput struct {
// 		Role   string `json:"role"`
// 		Domain string `json:"domain"`
// 		Module string `json:"module"`
// 		Action string `json:"action"`
// 	}

// 	type output struct {
// 		Policies []actionOutput `json:"policies"`
// 	}

// 	policies, err := Control.GetFilteredPolicy(1, domain)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	policiesOutput := make([]actionOutput, 0)
// 	for _, policy := range policies {
// 		policiesOutput = append(policiesOutput, actionOutput{
// 			Role:   policy[0],
// 			Domain: policy[1],
// 			Module: policy[2],
// 			Action: policy[3],
// 		})
// 	}

// 	json.NewEncoder(w).Encode(output{Policies: policiesOutput})
// }

// func CreatePolicy(w http.ResponseWriter, r *http.Request) {
// 	type input struct {
// 		Sub string `json:"sub"`
// 		Dom string `json:"dom"`
// 		Obj string `json:"obj"`
// 		Act string `json:"act"`
// 	}

// 	var data input
// 	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	_, err := Control.AddPolicy(data.Sub, data.Dom, data.Obj, data.Act)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusCreated)
// }

// func DeletePolicy(w http.ResponseWriter, r *http.Request) {
// 	type input struct {
// 		Sub string `json:"sub"`
// 		Dom string `json:"dom"`
// 		Obj string `json:"obj"`
// 		Act string `json:"act"`
// 	}

// 	var data input
// 	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	_, err := Control.RemovePolicy(data.Sub, data.Dom, data.Obj, data.Act)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusNoContent)
// }

// func AddRoleForUser(w http.ResponseWriter, r *http.Request) {
// 	type input struct {
// 		User string `json:"user"`
// 		Role string `json:"role"`
// 		Dom  string `json:"dom"`
// 	}

// 	var data input
// 	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	_, err := Control.AddRoleForUser(data.User, data.Role, data.Dom)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusCreated)
// }

// func RemoveRoleForUser(w http.ResponseWriter, r *http.Request) {
// 	type input struct {
// 		User string `json:"user"`
// 		Role string `json:"role"`
// 		Dom  string `json:"dom"`
// 	}

// 	var data input
// 	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	_, err := Control.RemoveRoleForUser(data.User, data.Role, data.Dom)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusNoContent)
// }

// func RemoveUserFromDomain(w http.ResponseWriter, r *http.Request) {
// 	user := chi.URLParam(r, "user")
// 	domain := chi.URLParam(r, "domain")

// 	if user == "" || domain == "" {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	_, err := Control.RemoveFilteredGroupingPolicy(0, user, "", domain)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusNoContent)
// }

// func RemoveRoleFromDomain(w http.ResponseWriter, r *http.Request) {
// 	role := chi.URLParam(r, "role")
// 	domain := chi.URLParam(r, "domain")

// 	if role == "" || domain == "" {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	// Remove all users from this role in this domain
// 	_, err := Control.RemoveFilteredGroupingPolicy(1, "", role, domain)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	// Remove all permissions for this role in this domain
// 	_, err = Control.RemoveFilteredPolicy(0, role, domain)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusNoContent)
// }

func Enforcer(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Sub string `json:"sub"`
		Dom string `json:"dom"`
		Obj string `json:"obj"`
		Act string `json:"act"`
		Key string `json:"key"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	bearerToken := r.Header.Get("Authorization")
	if bearerToken == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	bearerToken = strings.TrimPrefix(bearerToken, "Bearer ")

	if _, valid := locksmith.VerifyTokenWithClientSecret(bearerToken, input.Key); !valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	allowed, err := locksmith.Enforce(input.Sub, input.Dom, input.Obj, input.Act)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !allowed {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	w.WriteHeader(http.StatusOK)
}
