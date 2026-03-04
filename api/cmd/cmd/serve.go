package cmd

import (
	"log"
	"os"

	"github.com/booscaaa/initializers/validator"
	aclDi "github.com/locksmithhq/locksmith/api/internal/acl/di"
	"github.com/locksmithhq/locksmith/api/internal/acl/types/input"
	"github.com/locksmithhq/locksmith/api/internal/adapter/config"
	"github.com/locksmithhq/locksmith/api/internal/adapter/database"
	"github.com/locksmithhq/locksmith/api/internal/adapter/rest"
	"github.com/locksmithhq/locksmith-go"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use: "serve",
	Run: func(cmd *cobra.Command, args []string) {

		ctx := cmd.Context()
		database.Initialize(ctx)

		if err := locksmith.Initialize(
			database.GetConnection(),
			os.Getenv("LOCKSMITH_BASE_URL"),
			os.Getenv("LOCKSMITH_APP_CLIENT_ID"),
			os.Getenv("LOCKSMITH_APP_CLIENT_SECRET"),
		); err != nil {
			log.Fatalf("Failed to initialize ACL: %v", err)
		}

		config.InitializeSeeder(ctx)

		validator.Initialize(
			validator.WithDicionario(map[string]map[string]string{
				"pt": {
					"CpfCnpj": "CPF ou CNPJ",
				},
			}),
			validator.WithTraducoes(map[string]map[string]string{
				"pt": {
					"required":  "{0} é obrigatório.",
					"email":     "{0} inválido.",
					"gte":       "{0} deve ter {1} ou mais caracteres.",
					"lte":       "{0} deve ter {1} ou menos caracteres.",
					"min":       "{0} deve ter {1} ou mais caracteres.",
					"max":       "{0} deve ter {1} ou menos caracteres.",
					"datetime":  "{0} inválido.",
					"ltefield":  "{0} deve ser menor ou igual que {1}.",
					"CPForCNPJ": "{0} inválido.",
				},
			}),
		)

		locksmith.AddPolicy("role:admin", "domain:locksmith", "*", "*")
		locksmith.AddPolicy("*", "*", "module:acl", "action:read:own")

		_ = aclDi.NewCreateActionHandler()
		_ = input.Action{Title: "action:delete:one"}

		rest.Initialize()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
