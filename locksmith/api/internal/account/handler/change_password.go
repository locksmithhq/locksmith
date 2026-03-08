package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/locksmithhq/locksmith/api/internal/account/contract"
	"github.com/locksmithhq/locksmith/api/internal/account/types/input"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
)

type changePasswordHandler struct {
	changePasswordUseCase contract.ChangePasswordUseCase
}

func (h *changePasswordHandler) Execute(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	jwtString := strings.TrimPrefix(authHeader, "Bearer ")

	if jwtString == "" {
		stackerror.HttpResponse(w, "ChangePasswordHandler", fmt.Errorf("the jwt is not valid, please check your request"))
		return
	}

	var in input.ChangePassword

	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		stackerror.HttpResponse(w, "ChangePasswordHandler", err)
		return
	}

	in.Jwt = jwtString

	if err := in.ValidateHttp(w); err != nil {
		return
	}

	if err := h.changePasswordUseCase.Execute(r.Context(), in); err != nil {
		stackerror.HttpResponse(w, "ChangePasswordHandler", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func NewChangePasswordHandler(changePasswordUseCase contract.ChangePasswordUseCase) contract.ChangePasswordHandler {
	return &changePasswordHandler{
		changePasswordUseCase: changePasswordUseCase,
	}
}
