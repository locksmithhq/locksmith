package config

import (
	"context"
	"log"
	"os"

	"github.com/booscaaa/initializers/postgres/types"
	aclRepository "github.com/locksmithhq/locksmith/api/internal/acl/repository"
	aclInput "github.com/locksmithhq/locksmith/api/internal/acl/types/input"
	aclUsecase "github.com/locksmithhq/locksmith/api/internal/acl/usecase"
	accountRepository "github.com/locksmithhq/locksmith/api/internal/account/repository"
	accountInput "github.com/locksmithhq/locksmith/api/internal/account/types/input"
	accountUsecase "github.com/locksmithhq/locksmith/api/internal/account/usecase"
	"github.com/locksmithhq/locksmith/api/internal/adapter/database"
	clientRepository "github.com/locksmithhq/locksmith/api/internal/oauth_clients/repository"
	clientInput "github.com/locksmithhq/locksmith/api/internal/oauth_clients/types/input"
	clientUsecase "github.com/locksmithhq/locksmith/api/internal/oauth_clients/usecase"
	loginRepository "github.com/locksmithhq/locksmith/api/internal/oauth_clients_login/repository"
	loginInput "github.com/locksmithhq/locksmith/api/internal/oauth_clients_login/types/input"
	loginUsecase "github.com/locksmithhq/locksmith/api/internal/oauth_clients_login/usecase"
	signupRepository "github.com/locksmithhq/locksmith/api/internal/oauth_clients_signup/repository"
	signupInput "github.com/locksmithhq/locksmith/api/internal/oauth_clients_signup/types/input"
	signupUsecase "github.com/locksmithhq/locksmith/api/internal/oauth_clients_signup/usecase"
	projectRepository "github.com/locksmithhq/locksmith/api/internal/project/repository"
	projectInput "github.com/locksmithhq/locksmith/api/internal/project/types/input"
	projectUsecase "github.com/locksmithhq/locksmith/api/internal/project/usecase"
	locksmith "github.com/locksmithhq/locksmith-go"
	"gopkg.in/yaml.v2"
)

// ---------------------------------------------------------------------------
// YAML config types
// ---------------------------------------------------------------------------

type seederProjectConfig struct {
	Name        string               `yaml:"name"`
	Description string               `yaml:"description"`
	Domain      string               `yaml:"domain"`
	Roles       []seederRoleConfig   `yaml:"roles"`
	Clients     []seederClientConfig `yaml:"clients"`
	Users       []seederUserConfig   `yaml:"users"`
}

// seederRoleConfig describes a role and its optional per-project ACL policies.
type seederRoleConfig struct {
	Title    string               `yaml:"title"`
	Policies []seederPolicyConfig `yaml:"policies"`
}

// seederPolicyConfig maps a module to the list of actions allowed for a role.
type seederPolicyConfig struct {
	Module  string   `yaml:"module"`
	Actions []string `yaml:"actions"`
}

type seederClientConfig struct {
	Name         string              `yaml:"name"`
	ClientID     string              `yaml:"client_id"`
	ClientSecret string              `yaml:"client_secret"`
	RedirectURIs string              `yaml:"redirect_uris"`
	GrantTypes   string              `yaml:"grant_types"`
	Login        *seederLoginConfig  `yaml:"login"`
	Signup       *seederSignupConfig `yaml:"signup"`
}

type seederLoginConfig struct {
	Layout             string `yaml:"layout"`
	InputVariant       string `yaml:"input_variant"`
	ShowSocial         bool   `yaml:"show_social"`
	ShowRememberMe     bool   `yaml:"show_remember_me"`
	ShowForgotPassword bool   `yaml:"show_forgot_password"`
	ShowSignUp         bool   `yaml:"show_sign_up"`
	UseCustomHTML      bool   `yaml:"use_custom_html"`
	Enabled            bool   `yaml:"enabled"`
	CustomHTML         string `yaml:"custom_html"`
	CustomCSS          string `yaml:"custom_css"`
	BackgroundColor    string `yaml:"background_color"`
	BackgroundImage    string `yaml:"background_image"`
	BackgroundType     string `yaml:"background_type"`
	PrimaryColor       string `yaml:"primary_color"`
	LogoURL            string `yaml:"logo_url"`
}

type seederSignupConfig struct {
	Layout          string `yaml:"layout"`
	InputVariant    string `yaml:"input_variant"`
	ShowSocial      bool   `yaml:"show_social"`
	UseCustomHTML   bool   `yaml:"use_custom_html"`
	Enabled         bool   `yaml:"enabled"`
	CustomHTML      string `yaml:"custom_html"`
	CustomCSS       string `yaml:"custom_css"`
	BackgroundColor string `yaml:"background_color"`
	BackgroundImage string `yaml:"background_image"`
	BackgroundType  string `yaml:"background_type"`
	PrimaryColor    string `yaml:"primary_color"`
	LogoURL         string `yaml:"logo_url"`
	DefaultRoleName string `yaml:"default_role_name"`
}

type seederUserConfig struct {
	Name     string `yaml:"name"`
	Email    string `yaml:"email"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Role     string `yaml:"role"`
}

// ---------------------------------------------------------------------------
// Idempotent DB helpers — upsert by title and return the row ID
// ---------------------------------------------------------------------------

func upsertRole(ctx context.Context, conn types.Database, title string) (string, error) {
	var id string
	err := conn.QueryRowxContext(ctx,
		`INSERT INTO roles (title) VALUES ($1)
		 ON CONFLICT (title) DO UPDATE SET title = EXCLUDED.title
		 RETURNING id`,
		title,
	).Scan(&id)
	return id, err
}

func upsertModule(ctx context.Context, conn types.Database, title string) (string, error) {
	var id string
	err := conn.QueryRowxContext(ctx,
		`INSERT INTO modules (title) VALUES ($1)
		 ON CONFLICT (title) DO UPDATE SET title = EXCLUDED.title
		 RETURNING id`,
		title,
	).Scan(&id)
	return id, err
}

func upsertAction(ctx context.Context, conn types.Database, title string) (string, error) {
	var id string
	err := conn.QueryRowxContext(ctx,
		`INSERT INTO actions (title) VALUES ($1)
		 ON CONFLICT (title) DO UPDATE SET title = EXCLUDED.title
		 RETURNING id`,
		title,
	).Scan(&id)
	return id, err
}

func upsertProjectAcl(ctx context.Context, conn types.Database, roleID, moduleID, actionID, projectID string) error {
	_, err := conn.ExecContext(ctx,
		`INSERT INTO project_acl (role_id, module_id, action_id, project_id)
		 VALUES ($1, $2, $3, $4)
		 ON CONFLICT DO NOTHING`,
		roleID, moduleID, actionID, projectID,
	)
	return err
}

// ---------------------------------------------------------------------------
// seedProjects — called from InitializeSeeder on every startup
// ---------------------------------------------------------------------------

func seedProjects(ctx context.Context) {
	filePath := "/etc/locksmith/config/seeder.yaml"
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("SEED: failed to read seeder config: %v", err)
		return
	}

	expandedData := os.ExpandEnv(string(data))

	var seederConfig struct {
		Projects []seederProjectConfig `yaml:"projects"`
	}

	if err := yaml.Unmarshal([]byte(expandedData), &seederConfig); err != nil {
		log.Printf("SEED: failed to parse seeder config: %v", err)
		return
	}

	if len(seederConfig.Projects) == 0 {
		return
	}

	conn := database.GetConnection()

	createProjectUseCase := projectUsecase.NewCreateProjectUseCase(
		projectRepository.NewCreateProjectRepository(conn),
		projectRepository.NewGetProjectByNameRepository(conn),
	)

	createRoleUseCase := aclUsecase.NewCreateRoleUseCase(
		aclRepository.NewCreateRoleRepository(conn),
	)

	createClientUseCase := clientUsecase.NewCreateClientUseCase(
		clientRepository.NewCreateClientRepository(conn),
		clientRepository.NewGetClientByProjectIDAndClientIDRepository(conn),
	)

	createLoginUseCase := loginUsecase.NewCreateLoginByClientIDUseCase(
		loginRepository.NewCreateLoginByClientIDRepository(conn),
	)

	createSignupUseCase := signupUsecase.NewCreateSignupByClientIDUseCase(
		signupRepository.NewCreateSignupByClientIDRepository(conn),
	)

	createAccountUseCase := accountUsecase.NewCreateAccountUseCase(
		accountRepository.NewCreateAccountRepository(conn),
		accountRepository.NewGetAccountByEmailAndProjectIDRepository(conn),
		aclRepository.NewGetProjectDomainByProjectIDRepository(conn),
	)

	for _, p := range seederConfig.Projects {
		project, err := createProjectUseCase.Execute(ctx, projectInput.Project{
			Name:        p.Name,
			Description: p.Description,
			Domain:      p.Domain,
		})
		if err != nil {
			log.Printf("SEED: failed to create project %q: %v", p.Name, err)
			continue
		}

		log.Printf("SEED: project %q ready (id=%s)", p.Name, project.ID)

		// --- Roles + ACL policies ---
		for _, r := range p.Roles {
			// Ensure the role exists globally.
			if err := createRoleUseCase.Execute(ctx, aclInput.Role{Title: r.Title}); err != nil {
				log.Printf("SEED: role %q in project %q: %v", r.Title, p.Name, err)
			}

			if len(r.Policies) == 0 {
				continue
			}

			// Get (or create) the role ID for the project_acl table.
			roleID, err := upsertRole(ctx, conn, r.Title)
			if err != nil {
				log.Printf("SEED: failed to get id for role %q: %v", r.Title, err)
				continue
			}

			for _, policy := range r.Policies {
				moduleID, err := upsertModule(ctx, conn, policy.Module)
				if err != nil {
					log.Printf("SEED: failed to get id for module %q: %v", policy.Module, err)
					continue
				}

				for _, actionTitle := range policy.Actions {
					actionID, err := upsertAction(ctx, conn, actionTitle)
					if err != nil {
						log.Printf("SEED: failed to get id for action %q: %v", actionTitle, err)
						continue
					}

					// Persist to the project_acl table (visible in the dashboard).
					if err := upsertProjectAcl(ctx, conn, roleID, moduleID, actionID, project.ID); err != nil {
						log.Printf("SEED: project_acl %s → %s → %s in project %q: %v",
							r.Title, policy.Module, actionTitle, p.Name, err)
					}

					// Register in Casbin (idempotent — no-op if policy already exists).
					if _, err := locksmith.AddPolicy(r.Title, project.Domain, policy.Module, actionTitle); err != nil {
						log.Printf("SEED: casbin policy %s/%s/%s/%s: %v",
							r.Title, project.Domain, policy.Module, actionTitle, err)
					}
				}
			}
		}

		// --- OAuth clients ---
		for _, c := range p.Clients {
			client, err := createClientUseCase.Execute(ctx, clientInput.Client{
				ProjectID:    project.ID,
				Name:         c.Name,
				ClientID:     c.ClientID,
				ClientSecret: c.ClientSecret,
				RedirectURIs: c.RedirectURIs,
				GrantTypes:   c.GrantTypes,
			})
			if err != nil {
				log.Printf("SEED: failed to create client %q in project %q: %v", c.Name, p.Name, err)
				continue
			}

			// Use the primary key ID (oauth_clients.id) as FK for login/signup tables.
			clientID := client.ID

			if c.Login != nil {
				if err := createLoginUseCase.Execute(ctx, clientID, loginInput.Login{
					Layout:             c.Login.Layout,
					InputVariant:       c.Login.InputVariant,
					ShowSocial:         c.Login.ShowSocial,
					ShowRememberMe:     c.Login.ShowRememberMe,
					ShowForgotPassword: c.Login.ShowForgotPassword,
					ShowSignUp:         c.Login.ShowSignUp,
					UseCustomHTML:      c.Login.UseCustomHTML,
					Enabled:            c.Login.Enabled,
					CustomHTML:         c.Login.CustomHTML,
					CustomCSS:          c.Login.CustomCSS,
					BackgroundColor:    c.Login.BackgroundColor,
					BackgroundImage:    c.Login.BackgroundImage,
					BackgroundType:     c.Login.BackgroundType,
					PrimaryColor:       c.Login.PrimaryColor,
					LogoURL:            c.Login.LogoURL,
				}); err != nil {
					log.Printf("SEED: login config for client %q: %v", c.Name, err)
				}
			}

			if c.Signup != nil {
				if err := createSignupUseCase.Execute(ctx, clientID, signupInput.Signup{
					Layout:          c.Signup.Layout,
					InputVariant:    c.Signup.InputVariant,
					ShowSocial:      c.Signup.ShowSocial,
					UseCustomHTML:   c.Signup.UseCustomHTML,
					Enabled:         c.Signup.Enabled,
					CustomHTML:      c.Signup.CustomHTML,
					CustomCSS:       c.Signup.CustomCSS,
					BackgroundColor: c.Signup.BackgroundColor,
					BackgroundImage: c.Signup.BackgroundImage,
					BackgroundType:  c.Signup.BackgroundType,
					PrimaryColor:    c.Signup.PrimaryColor,
					LogoURL:         c.Signup.LogoURL,
					DefaultRoleName: c.Signup.DefaultRoleName,
				}); err != nil {
					log.Printf("SEED: signup config for client %q: %v", c.Name, err)
				}
			}
		}

		// --- Users ---
		for _, u := range p.Users {
			if _, err := createAccountUseCase.Execute(ctx, accountInput.Account{
				Name:      u.Name,
				Email:     u.Email,
				Username:  u.Username,
				Password:  u.Password,
				ProjectID: project.ID,
				RoleName:  u.Role,
			}); err != nil {
				log.Printf("SEED: user %q in project %q: %v", u.Email, p.Name, err)
			}
		}
	}
}
