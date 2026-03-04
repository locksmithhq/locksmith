package rest

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"

	accountsRest "github.com/booscaaa/locksmith/api/internal/account/rest"
	"github.com/booscaaa/locksmith/api/internal/acl"
	aclRest "github.com/booscaaa/locksmith/api/internal/acl/rest"
	locksmithRest "github.com/booscaaa/locksmith/api/internal/locksmith/rest"
	oauth2Rest "github.com/booscaaa/locksmith/api/internal/oauth2/rest"
	oauthClientsRest "github.com/booscaaa/locksmith/api/internal/oauth_clients/rest"
	oauthClientsLoginRest "github.com/booscaaa/locksmith/api/internal/oauth_clients_login/rest"
	oauthClientsSignupRest "github.com/booscaaa/locksmith/api/internal/oauth_clients_signup/rest"
	projectRest "github.com/booscaaa/locksmith/api/internal/project/rest"
	sessionRest "github.com/booscaaa/locksmith/api/internal/session/rest"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func Initialize() {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/api", func(r chi.Router) {
		oauth2Rest.InitializeOAuth2Router(r)
		locksmithRest.InitializeLocksmithRouter(r)
		projectRest.InitializeProjectRouter(r)
		oauthClientsRest.InitializeOauthClientsRouter(r)
		oauthClientsLoginRest.InitializeOauthClientsLoginRouter(r)
		oauthClientsSignupRest.InitializeOauthClientsSignupRouter(r)
		accountsRest.InitializeAccountsRouter(r)
		sessionRest.InitializeSessionRouter(r)
		aclRest.InitializeAclRouter(r)
		r.Post("/acl/enforce", acl.Enforcer)

		// GET /config
		r.Get("/config", func(w http.ResponseWriter, r *http.Request) {
			config := map[string]string{
				"baseUrl":  os.Getenv("LOCKSMITH_BASE_URL"),
				"clientId": os.Getenv("LOCKSMITH_APP_CLIENT_ID"),
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(config)
		})

		// r.Get("/acl/roles/{domain}", acl.GetRoles)
		// r.Get("/acl/users/{domain}", acl.GetUsers)

		// r.Get("/acl/policies/{domain}", acl.GetPolicies)
		// r.Post("/acl/policies", acl.CreatePolicy)
		// r.Put("/acl/policies", acl.UpdatePolicy)
		// r.Delete("/acl/policies", acl.DeletePolicy)

		// r.Post("/acl/roles", acl.AddRoleForUser)
		// r.Delete("/acl/roles", acl.RemoveRoleForUser)
		// r.Put("/acl/roles/{user}/domain/{domain}/role/{role}", acl.UpdateUserRole)

		// r.Delete("/acl/users/{domain}/user/{user}", acl.RemoveUserFromDomain)
		// r.Delete("/acl/roles/{domain}/role/{role}", acl.RemoveRoleFromDomain)
	})

	spaPath := "/web/dist"
	fs := http.FileServer(http.Dir(spaPath))

	r.Get("/*", func(w http.ResponseWriter, req *http.Request) {
		fullPath := filepath.Join(spaPath, req.URL.Path)
		if _, err := os.Stat(fullPath); os.IsNotExist(err) {
			http.ServeFile(w, req, filepath.Join(spaPath, "index.html"))
			return
		}

		fs.ServeHTTP(w, req)
	})

	apiPort := os.Getenv("LOCKSMITH_APP_PORT")
	if apiPort == "" {
		apiPort = "10001"
	}

	log.Printf("REST API serve in port: %s\n", apiPort)
	http.ListenAndServe(":"+apiPort, r) //
}
