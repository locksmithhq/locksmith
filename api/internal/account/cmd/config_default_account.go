package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/booscaaa/locksmith/api/internal/account/contract"
	"github.com/booscaaa/locksmith/api/internal/account/types/input"
	"github.com/booscaaa/locksmith/api/internal/acl"
	"gopkg.in/yaml.v2"
)

type configDefaultAccountCMD struct {
	createAccountUseCase contract.CreateAccountUseCase
}

func (c *configDefaultAccountCMD) Execute(ctx context.Context, projectID string) error {
	filePath := "/etc/locksmith/config/seeder.yaml"
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read seeder config file %s: %w", filePath, err)
	}

	var seederConfig struct {
		DefaultAccount struct {
			Name     string `yaml:"name"`
			Email    string `yaml:"email"`
			Username string `yaml:"username"`
			Password string `yaml:"password"`
		} `yaml:"default_account"`
	}

	if err := yaml.Unmarshal(data, &seederConfig); err != nil {
		return fmt.Errorf("failed to unmarshal seeder config from %s: %w", filePath, err)
	}

	in := input.Account{
		Name:      seederConfig.DefaultAccount.Name,
		Email:     seederConfig.DefaultAccount.Email,
		Username:  seederConfig.DefaultAccount.Username,
		Password:  seederConfig.DefaultAccount.Password,
		ProjectID: projectID,
		RoleName:  acl.RoleAdmin,
	}

	_, err = c.createAccountUseCase.Execute(ctx, in)
	if err != nil {
		return fmt.Errorf("failed to create default account: %w", err)
	}

	return nil
}

func NewConfigDefaultAccount(createAccountUseCase contract.CreateAccountUseCase) contract.ConfigDefaultAccountCMD {
	return &configDefaultAccountCMD{createAccountUseCase: createAccountUseCase}
}
