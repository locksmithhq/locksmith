package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/locksmithhq/locksmith/api/internal/oauth_clients/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients/types/input"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients/types/output"
	"gopkg.in/yaml.v2"
)

type configDefaultClientCMD struct {
	createClientUseCase contract.CreateClientUseCase
}

func (c *configDefaultClientCMD) Execute(ctx context.Context, projectID string) (output.Client, error) {
	secretKey := os.Getenv("LOCKSMITH_APP_CLIENT_SECRET")
	if secretKey == "" {
		return output.Client{}, fmt.Errorf("LOCKSMITH_APP_CLIENT_SECRET is required")
	}

	clientID := os.Getenv("LOCKSMITH_APP_CLIENT_ID")
	if clientID == "" {
		return output.Client{}, fmt.Errorf("LOCKSMITH_APP_CLIENT_ID is required")
	}

	filePath := "/etc/locksmith/config/seeder.yaml"
	data, err := os.ReadFile(filePath)
	if err != nil {
		return output.Client{}, fmt.Errorf("failed to read seeder config file %s: %w", filePath, err)
	}

	// Expand environment variables in the YAML content
	expandedData := os.ExpandEnv(string(data))

	var seederConfig struct {
		DefaultClient struct {
			Name         string `yaml:"name"`
			RedirectURIs string `yaml:"redirect_uris"`
			GrantTypes   string `yaml:"grant_types"`
		} `yaml:"default_client"`
	}

	if err := yaml.Unmarshal([]byte(expandedData), &seederConfig); err != nil {
		return output.Client{}, fmt.Errorf("failed to unmarshal seeder config from %s: %w", filePath, err)
	}

	in := input.Client{
		ProjectID:    projectID,
		Name:         seederConfig.DefaultClient.Name,
		ClientID:     clientID,
		ClientSecret: secretKey,
		RedirectURIs: seederConfig.DefaultClient.RedirectURIs,
		GrantTypes:   seederConfig.DefaultClient.GrantTypes,
	}

	client, err := c.createClientUseCase.Execute(ctx, in)
	if err != nil {
		return output.Client{}, fmt.Errorf("failed to create default client: %w", err)
	}

	return client, nil
}

func NewConfigDefaultClientCMD(createClientUseCase contract.CreateClientUseCase) contract.ConfigDefaultClientCMD {
	return &configDefaultClientCMD{createClientUseCase: createClientUseCase}
}
