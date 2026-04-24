package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/locksmithhq/locksmith/api/internal/adapter/database"
	"github.com/locksmithhq/locksmith/api/internal/core/crypto"
	"github.com/spf13/cobra"
)

var migrateSecretsCmd = &cobra.Command{
	Use:   "migrate-secrets",
	Short: "Re-encrypt plaintext secrets with AES-256-GCM at-rest encryption",
	Run: func(cmd *cobra.Command, args []string) {
		if os.Getenv("ENCRYPTION_KEY") == "" {
			log.Fatal("ENCRYPTION_KEY is not set")
		}

		ctx := cmd.Context()
		database.Initialize(ctx)
		db := database.GetConnection()

		type clientRow struct {
			ID           string `db:"id"`
			ClientSecret string `db:"client_secret"`
		}
		var clients []clientRow
		if err := db.SelectContext(ctx, &clients, `SELECT id, client_secret FROM oauth_clients`); err != nil {
			log.Fatalf("fetch oauth_clients: %v", err)
		}

		clientsUpdated := 0
		for _, c := range clients {
			if strings.HasPrefix(c.ClientSecret, "aes256gcm:") {
				continue
			}
			enc, err := crypto.Encrypt(c.ClientSecret)
			if err != nil {
				log.Fatalf("encrypt client %s: %v", c.ID, err)
			}
			if _, err := db.ExecContext(ctx, `UPDATE oauth_clients SET client_secret = $1 WHERE id = $2`, enc, c.ID); err != nil {
				log.Fatalf("update client %s: %v", c.ID, err)
			}
			clientsUpdated++
		}
		fmt.Printf("oauth_clients:                  %d/%d rows re-encrypted\n", clientsUpdated, len(clients))

		type providerRow struct {
			ID           string `db:"id"`
			ClientKey    string `db:"client_key"`
			ClientSecret string `db:"client_secret"`
		}
		var providers []providerRow
		if err := db.SelectContext(ctx, &providers, `SELECT id, client_key, client_secret FROM oauth_client_social_providers`); err != nil {
			log.Fatalf("fetch oauth_client_social_providers: %v", err)
		}

		providersUpdated := 0
		for _, p := range providers {
			needsUpdate := false
			encKey, encSecret := p.ClientKey, p.ClientSecret

			if !strings.HasPrefix(p.ClientKey, "aes256gcm:") {
				enc, err := crypto.Encrypt(p.ClientKey)
				if err != nil {
					log.Fatalf("encrypt provider key %s: %v", p.ID, err)
				}
				encKey = enc
				needsUpdate = true
			}
			if !strings.HasPrefix(p.ClientSecret, "aes256gcm:") {
				enc, err := crypto.Encrypt(p.ClientSecret)
				if err != nil {
					log.Fatalf("encrypt provider secret %s: %v", p.ID, err)
				}
				encSecret = enc
				needsUpdate = true
			}

			if needsUpdate {
				if _, err := db.ExecContext(ctx,
					`UPDATE oauth_client_social_providers SET client_key = $1, client_secret = $2 WHERE id = $3`,
					encKey, encSecret, p.ID,
				); err != nil {
					log.Fatalf("update provider %s: %v", p.ID, err)
				}
				providersUpdated++
			}
		}
		fmt.Printf("oauth_client_social_providers:  %d/%d rows re-encrypted\n", providersUpdated, len(providers))
		fmt.Println("Done.")
	},
}

func init() {
	rootCmd.AddCommand(migrateSecretsCmd)
}
