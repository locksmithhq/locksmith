# Locksmith

**Locksmith** is an open-source OAuth2 Identity and Access Management (IAM) platform. It provides multi-tenant user authentication, fine-grained role-based access control, OAuth2 Authorization Code flow with PKCE, social login, session tracking with device fingerprinting, and a fully featured management dashboard — all deployable as a single Docker container.

---

## Table of Contents

- [Features](#features)
- [Architecture](#architecture)
- [Quick Start](#quick-start)
  - [Docker Compose (recommended)](#docker-compose-recommended)
  - [Docker Run](#docker-run)
  - [Production with custom domain](#production-with-custom-domain)
- [Configuration](#configuration)
- [Default Credentials](#default-credentials)
- [Seeder Configuration](#seeder-configuration)
- [OAuth2 Flow](#oauth2-flow)
- [Social Login](#social-login)
- [API Reference](#api-reference)
- [Access Control (ACL)](#access-control-acl)
- [Multi-Tenancy](#multi-tenancy)
- [Session & Device Tracking](#session--device-tracking)
- [Security](#security)
- [Management Dashboard](#management-dashboard)
- [Development](#development)

---

## Features

- **OAuth2 Authorization Code + PKCE** — S256-only code challenge enforcement; `require_pkce` flag per client makes PKCE mandatory
- **Social Login** — Google, GitHub, Facebook, Discord, LinkedIn via per-client configurable providers
- **Multi-tenant Projects** — fully isolated users, OAuth clients, roles, and sessions per project
- **Role-Based Access Control** — Casbin-powered RBAC with domain, module, and action granularity
- **Session Management** — per-session tracking with browser, OS, device type, IP address, and GeoIP location
- **Refresh Token Rotation** — SHA-256 hashed refresh tokens with chain revocation on reuse detection
- **Logout Session Revocation** — logout marks the refresh token and its session as revoked server-side
- **Customizable Login/Register pages** — per-client UI theming, custom CSS/HTML, and field visibility controls
- **User Account Management** — CRUD, Argon2id password hashing, forced password change on next login
- **Rate Limiting** — per-IP limits on all authentication endpoints
- **Production Error Mode** — stack traces hidden in production; opaque error IDs logged server-side
- **Management Dashboard** — Vue 3 + Vuetify web UI for managing all resources
- **Single Binary + SPA** — API and dashboard bundled into one Docker image (no separate containers in production)
- **Multi-arch image** — `linux/amd64` and `linux/arm64` published to Docker Hub
- **Runs as non-root** — production container uses a dedicated unprivileged user

---

## Architecture

```
┌─────────────────────────────────────────────┐
│               Locksmith Container            │
│                                             │
│  ┌─────────────────┐  ┌───────────────────┐ │
│  │   Go REST API   │  │  Vue 3 Dashboard  │ │
│  │  (Chi router)   │  │  (served as SPA)  │ │
│  └────────┬────────┘  └───────────────────┘ │
│           │                                 │
└───────────┼─────────────────────────────────┘
            │
     ┌──────▼──────┐
     │  PostgreSQL  │
     └─────────────┘
```

**Tech Stack:**

| Layer    | Technology                                    |
|----------|-----------------------------------------------|
| Backend  | Go 1.25+, Chi v5, Casbin v2, JWT              |
| Frontend | Vue 3, Vuetify 3, Vite, Bun                   |
| Database | PostgreSQL 16                                 |
| Auth     | OAuth2 PKCE (S256), Argon2id, SHA-256 tokens  |
| Social   | goth (Google, GitHub, Facebook, Discord, LinkedIn) |
| ACL      | Casbin RBAC with domain support               |

---

## Quick Start

### Docker Compose (recommended)

**1. Create a `compose.yaml`:**

```yaml
services:
  locksmith:
    image: booscaaa/locksmith:latest
    container_name: locksmith
    ports:
      - "4000:4000"
    environment:
      - APP_ENV=production
      - LOCKSMITH_APP_PORT=4000
      - LOCKSMITH_BASE_URL=http://localhost:4000
      - LOCKSMITH_APP_CLIENT_ID=my-client-id
      - LOCKSMITH_APP_CLIENT_SECRET=my-client-secret
      - POSTGRES_HOST=database
      - POSTGRES_USER=locksmith
      - POSTGRES_PASSWORD=locksmith123
      - POSTGRES_DB=locksmith
      - POSTGRES_PORT=5432
      - SCHEMA=locksmith
      - SSL_MODE=disable
      - SEED_ADMIN_EMAIL=admin@example.com
      - SEED_ADMIN_PASSWORD=changeme123
    depends_on:
      database:
        condition: service_healthy
    restart: unless-stopped

  database:
    image: postgres:16-alpine
    container_name: locksmith-db
    environment:
      - POSTGRES_USER=locksmith
      - POSTGRES_PASSWORD=locksmith123
      - POSTGRES_DB=locksmith
    volumes:
      - locksmith-data:/var/lib/postgresql/data
    healthcheck:
      test: ["CMD", "pg_isready", "-U", "locksmith"]
      interval: 10s
      timeout: 5s
      retries: 5
    restart: unless-stopped

volumes:
  locksmith-data:
```

**2. Start it:**

```bash
docker compose up -d
```

**3. Open the dashboard:**

```
http://localhost:4000
```

---

### Docker Run

**1. Create a Docker network:**

```bash
docker network create locksmith-network
```

**2. Start PostgreSQL:**

```bash
docker run -d \
  --name locksmith-db \
  --network locksmith-network \
  -e POSTGRES_USER=locksmith \
  -e POSTGRES_PASSWORD=locksmith123 \
  -e POSTGRES_DB=locksmith \
  -v locksmith-data:/var/lib/postgresql/data \
  --health-cmd="pg_isready -U locksmith" \
  --health-interval=10s \
  --health-timeout=5s \
  --health-retries=5 \
  --restart unless-stopped \
  postgres:16-alpine
```

**3. Start Locksmith:**

```bash
docker run -d \
  --name locksmith \
  --network locksmith-network \
  -p 4000:4000 \
  -e APP_ENV=production \
  -e LOCKSMITH_APP_PORT=4000 \
  -e LOCKSMITH_BASE_URL=http://localhost:4000 \
  -e LOCKSMITH_APP_CLIENT_ID=my-client-id \
  -e LOCKSMITH_APP_CLIENT_SECRET=my-client-secret \
  -e POSTGRES_HOST=locksmith-db \
  -e POSTGRES_USER=locksmith \
  -e POSTGRES_PASSWORD=locksmith123 \
  -e POSTGRES_DB=locksmith \
  -e POSTGRES_PORT=5432 \
  -e SCHEMA=locksmith \
  -e SSL_MODE=disable \
  --restart unless-stopped \
  booscaaa/locksmith:latest
```

---

### Production with custom domain

```bash
docker run -d \
  --name locksmith \
  --network locksmith-network \
  -p 4000:4000 \
  -e APP_ENV=production \
  -e LOCKSMITH_APP_PORT=4000 \
  -e LOCKSMITH_BASE_URL=https://auth.example.com \
  -e LOCKSMITH_APP_CLIENT_ID=my-client-id \
  -e LOCKSMITH_APP_CLIENT_SECRET=strong-secret-here \
  -e POSTGRES_HOST=your-postgres-host \
  -e POSTGRES_USER=locksmith \
  -e POSTGRES_PASSWORD=strong-password-here \
  -e POSTGRES_DB=locksmith \
  -e POSTGRES_PORT=5432 \
  -e SCHEMA=locksmith \
  -e SSL_MODE=require \
  -e SEED_ADMIN_EMAIL=admin@yourdomain.com \
  -e SEED_ADMIN_PASSWORD=strong-admin-password \
  --restart unless-stopped \
  booscaaa/locksmith:latest
```

> **Note:** Place a reverse proxy (nginx, Caddy, Traefik) in front of Locksmith to handle TLS termination. Locksmith itself runs on plain HTTP internally.

---

## Configuration

All configuration is done through environment variables.

### Required Variables

| Variable                      | Description                                                  | Example                    |
|-------------------------------|--------------------------------------------------------------|----------------------------|
| `LOCKSMITH_APP_PORT`          | Port the API listens on                                      | `4000`                     |
| `LOCKSMITH_BASE_URL`          | Public base URL (used for OAuth callbacks and cookie domain) | `https://auth.example.com` |
| `LOCKSMITH_APP_CLIENT_ID`     | Client ID of the built-in Locksmith management client        | `my-client-id`             |
| `LOCKSMITH_APP_CLIENT_SECRET` | Client secret of the built-in management client              | `my-client-secret`         |
| `POSTGRES_HOST`               | PostgreSQL hostname                                          | `database`                 |
| `POSTGRES_USER`               | PostgreSQL username                                          | `locksmith`                |
| `POSTGRES_PASSWORD`           | PostgreSQL password                                          | `locksmith123`             |
| `POSTGRES_DB`                 | PostgreSQL database name                                     | `locksmith`                |
| `POSTGRES_PORT`               | PostgreSQL port                                              | `5432`                     |
| `SCHEMA`                      | PostgreSQL schema name                                       | `locksmith`                |
| `SSL_MODE`                    | PostgreSQL SSL mode (`disable`, `require`, `verify-full`)    | `require`                  |

### Optional Variables

| Variable                      | Description                                                                                                                  | Default        |
|-------------------------------|------------------------------------------------------------------------------------------------------------------------------|----------------|
| `APP_ENV`                     | Runtime environment. Set to `production` to hide internal error details from HTTP responses.                                 | `development`  |
| `SEED_ADMIN_EMAIL`            | Email of the default admin account created on first boot                                                                     | `admin@locksmith.rs` |
| `SEED_ADMIN_PASSWORD`         | Password of the default admin account                                                                                        | `admin`        |
| `SEED_APP_CLIENT_ID`          | Overrides `LOCKSMITH_APP_CLIENT_ID` in the seeder (useful when they differ)                                                  | —              |
| `SEED_APP_CLIENT_SECRET`      | Overrides `LOCKSMITH_APP_CLIENT_SECRET` in the seeder                                                                        | —              |
| `SEED_APP_USER_PASSWORD`      | Password assigned to seeded user accounts                                                                                    | —              |
| `VITE_LOCKSMITH_API_BASE_URL` | Override the API base URL used by the frontend SPA                                                                           | `window.location.origin` |

### Notes

- **Always set `APP_ENV=production`** in production deployments. In development mode, error responses include stack traces and internal details. In production mode, errors return only a message and an opaque `error_id` (logged server-side for debugging).
- `LOCKSMITH_BASE_URL` must be reachable from the browser — it is used as the OAuth redirect target and for setting the cookie domain.
- `SSL_MODE=disable` is only acceptable in local development with a co-located database. Use `require` or `verify-full` in any networked environment.
- On first boot, Locksmith automatically runs all database migrations and seeds the default project, admin account, and management OAuth client. All seed operations are idempotent.

---

## Default Credentials

On first boot, Locksmith creates the following defaults (overridable via env vars):

| Resource            | Default value                                  | Override via env                |
|---------------------|------------------------------------------------|---------------------------------|
| Default project     | `Default Project` (domain: `domain:locksmith`) | —                               |
| Admin email         | `admin@locksmith.rs`                           | `SEED_ADMIN_EMAIL`              |
| Admin username      | `admin`                                        | —                               |
| Admin password      | `admin`                                        | `SEED_ADMIN_PASSWORD`           |
| Management client   | Uses `LOCKSMITH_APP_CLIENT_ID/SECRET`          | `SEED_APP_CLIENT_ID/SECRET`     |

> **Change the admin password immediately after your first login**, or set `SEED_ADMIN_PASSWORD` before the first boot.

---

## Seeder Configuration

On first boot, Locksmith reads `/etc/locksmith/config/seeder.yaml` to create the default project, admin account, and management OAuth client. All values support `${ENV_VAR}` interpolation at startup.

### seeder.yaml reference

```yaml
default_project:
  name: Default Project
  description: Default project for Locksmith
  domain: domain:locksmith

default_account:
  name: Default Account
  email: ${SEED_ADMIN_EMAIL}
  password: ${SEED_ADMIN_PASSWORD}
  username: admin

default_client:
  name: Default Client
  redirect_uris: ${LOCKSMITH_BASE_URL}/api/locksmith/callback ${LOCKSMITH_BASE_URL}
  grant_types: authorization_code
```

### Seeding additional projects

The `projects` key lets you declare additional projects with roles, OAuth clients, and users that are created automatically on startup. All entries are idempotent.

```yaml
projects:
  - name: My Application
    description: Main application for ACME Corp
    domain: domain:myapp

    roles:
      - title: role:admin
        policies:
          - module: module:accounts
            actions:
              - action:create:one
              - action:read:all
              - action:update:one
              - action:delete:one
          - module: module:sessions
            actions:
              - action:read:all
              - action:revoke

      - title: role:user
        policies:
          - module: module:accounts
            actions:
              - action:read:own
              - action:update:own

    clients:
      - name: My App Web Client
        client_id: ${MY_APP_CLIENT_ID}
        client_secret: ${MY_APP_CLIENT_SECRET}
        redirect_uris: ${MY_APP_BASE_URL}/callback ${MY_APP_BASE_URL}
        grant_types: authorization_code
        login:
          enabled: true
          show_sign_up: true
          show_social: true
          primary_color: "#1976D2"
        signup:
          enabled: true
          default_role_name: role:user

    users:
      - name: App Admin
        email: admin@myapp.com
        username: admin
        password: ${SEED_APP_USER_PASSWORD}
        role: role:admin
```

### Mounting a custom seeder

```yaml
# compose.yaml
services:
  locksmith:
    image: booscaaa/locksmith:latest
    volumes:
      - ./my-seeder.yaml:/etc/locksmith/config/seeder.yaml
    environment:
      LOCKSMITH_BASE_URL: https://auth.example.com
      # ...
```

```bash
# docker run
docker run -d \
  -v $(pwd)/my-seeder.yaml:/etc/locksmith/config/seeder.yaml \
  -e LOCKSMITH_BASE_URL=https://auth.example.com \
  booscaaa/locksmith:latest
```

---

## OAuth2 Flow

Locksmith implements the **Authorization Code flow with PKCE** (RFC 7636). Only `S256` is accepted as `code_challenge_method` — `plain` is rejected with HTTP 400.

### Flow Diagram

<img src="docs/oauth2-flow.svg" width="100%" alt="OAuth2 Authorization Code + PKCE Flow" />

### Step-by-step

**1. Start authorization:**

```
GET /api/oauth2/authorize
  ?client_id=<client_id>
  &redirect_uri=<redirect_uri>
  &response_type=code
  &state=<random_state>
  &code_challenge=<sha256_of_verifier_base64url>
  &code_challenge_method=S256
```

**2. User login:**

```http
POST /api/oauth2/login
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "password",
  "client_id": "<client_id>",
  "code_challenge": "<sha256_of_verifier_base64url>",
  "code_challenge_method": "S256"
}
```

> If the OAuth client has `require_pkce: true`, the `code_challenge` field is mandatory. Requests without it are rejected with HTTP 400.

**2b. (Alternative) User registration:**

```http
POST /api/oauth2/register
Content-Type: application/json

{
  "name": "User Name",
  "email": "user@example.com",
  "username": "username",
  "password": "password",
  "client_id": "<client_id>"
}
```

**3. Exchange code for tokens:**

```http
POST /api/oauth2/access-token
Content-Type: application/json

{
  "code": "<authorization_code>",
  "client_id": "<client_id>",
  "client_secret": "<client_secret>",
  "grant_type": "authorization_code",
  "code_verifier": "<original_verifier>"
}
```

Response:

```json
{
  "access_token": "<jwt>",
  "refresh_token": "<uuid>",
  "token_type": "Bearer",
  "expires_in": 300
}
```

**4. Refresh tokens:**

```http
POST /api/oauth2/refresh-token
Content-Type: application/json

{
  "refresh_token": "<refresh_token>",
  "client_id": "<client_id>",
  "client_secret": "<client_secret>"
}
```

### Cookie-based flow (dashboard)

The management dashboard uses Locksmith's own callback handler which sets HTTP-only cookies:

- `LOCKSMITHACCESSTOKEN` — JWT access token (5-minute expiry, `HttpOnly`, `Secure`, `SameSite=Strict`)
- `LOCKSMITHREFRESHTOKEN` — UUID refresh token (15-day expiry, `HttpOnly`, `Secure`, `SameSite=Strict`)

The cookie domain is derived automatically from the request's `Origin` header.

---

## Social Login

Locksmith supports social login via OAuth2 providers configured per OAuth client. Each provider requires a `client_key` (OAuth App Client ID) and `client_secret` from the provider's developer console.

### Supported providers

| Provider  | Provider name (API) | Required scopes            |
|-----------|--------------------|-----------------------------|
| Google    | `google`           | `openid email profile`      |
| GitHub    | `github`           | `user:email`                |
| Facebook  | `facebook`         | `email`                     |
| Discord   | `discord`          | `identify email`            |
| LinkedIn  | `linkedin`         | `r_liteprofile r_emailaddress` |

### Configuring a provider

1. Go to **Project Details → OAuth Client → Social Providers** in the dashboard
2. Enable the desired provider and paste your `client_key` and `client_secret`
3. Set the allowed scopes (space-separated)
4. Configure the OAuth callback URI in your provider's developer console:

```
{LOCKSMITH_BASE_URL}/api/oauth2/social/{provider}/callback
```

For example: `https://auth.example.com/api/oauth2/social/google/callback`

### Social login flow

**1. Begin social authorization:**

```
GET /api/oauth2/social/{provider}/begin
  ?client_id=<client_id>
  &redirect_uri=<redirect_uri>
  &state=<random_state>
  &code_challenge=<sha256_of_verifier_base64url>
  &code_challenge_method=S256
```

Returns:

```json
{
  "auth_url": "https://accounts.google.com/o/oauth2/auth?..."
}
```

**2.** Redirect the user to `auth_url`. After the provider authenticates the user, it redirects back to Locksmith's callback endpoint, which issues an auth code and redirects to your `redirect_uri`.

**3.** Exchange the auth code for tokens using the same `/api/oauth2/access-token` endpoint.

### Account linking

On social login:
- If an account with the provider's user ID already exists → authenticated directly
- If an account with the same email exists → the social provider is linked to that account
- If no account exists and signup is enabled for the client → a new account is created automatically
- If no account exists and signup is disabled → returns HTTP 403

---

## API Reference

All API routes are prefixed with `/api`.

**Auth notation:**
- 🔒 Requires a valid `LOCKSMITHACCESSTOKEN` cookie (dashboard session)
- 🔑 Requires HTTP Basic Auth (`client_id:client_secret`)
- *(no marker)* Public

---

### OAuth2

| Method | Path                                      | Description                                    |
|--------|-------------------------------------------|------------------------------------------------|
| `POST` | `/api/oauth2/authorize`                   | Start authorization — create auth code         |
| `POST` | `/api/oauth2/login`                       | Authenticate user credentials                  |
| `POST` | `/api/oauth2/register`                    | Register a new user account                    |
| `POST` | `/api/oauth2/access-token`                | Exchange auth code for access + refresh tokens |
| `POST` | `/api/oauth2/refresh-token`               | Rotate refresh token and get new access token  |
| `GET`  | `/api/oauth2/resolve-domain`              | Resolve OAuth client by custom domain          |
| `GET`  | `/api/oauth2/social/{provider}/begin`     | Start social OAuth flow                        |
| `GET`  | `/api/oauth2/social/{provider}/callback`  | Social OAuth callback (provider → Locksmith)   |

---

### Locksmith Callback

| Method | Path                      | Description                                    |
|--------|---------------------------|------------------------------------------------|
| `GET`  | `/api/locksmith/callback` | OAuth callback — exchanges code, sets cookies  |
| `GET`  | `/api/locksmith/status`   | Check if current session is authenticated      |
| `POST` | `/api/locksmith/r`        | Refresh access token using cookie              |
| `POST` | `/api/locksmith/logout`   | Revoke session + tokens and clear cookies      |

---

### Projects 🔒

| Method   | Path                | Description                   |
|----------|---------------------|-------------------------------|
| `GET`    | `/api/projects`     | List all projects (paginated) |
| `GET`    | `/api/projects/:id` | Get a single project          |
| `POST`   | `/api/projects`     | Create a new project          |
| `PUT`    | `/api/projects/:id` | Update a project              |
| `DELETE` | `/api/projects/:id` | Delete a project              |

---

### Accounts

| Method | Path                                                 | Auth | Description                            |
|--------|------------------------------------------------------|------|----------------------------------------|
| `POST` | `/api/projects/:project_id/accounts`                 | 🔒   | Create account (dashboard)             |
| `PUT`  | `/api/projects/:project_id/accounts/:account_id`     | 🔒   | Update account (dashboard)             |
| `GET`  | `/api/projects/:project_id/accounts`                 | 🔒   | List accounts (paginated)              |
| `GET`  | `/api/projects/:project_id/accounts/count`           | 🔒   | Count accounts                         |
| `GET`  | `/api/projects/:project_id/accounts/:id`             | 🔒   | Get a single account                   |
| `POST` | `/api/accounts`                                      | 🔑   | Create account (client credentials)    |
| `GET`  | `/api/accounts/:id`                                  | 🔑   | Get account by ID                      |
| `PUT`  | `/api/accounts/:account_id`                          | 🔑   | Update account                         |
| `POST` | `/api/accounts/change-password`                      | —    | Change password (JWT verified inline)  |

---

### OAuth Clients 🔒

| Method | Path                                              | Description                        |
|--------|---------------------------------------------------|------------------------------------|
| `GET`  | `/api/projects/:project_id/clients`               | List OAuth clients                 |
| `GET`  | `/api/projects/:project_id/clients/:id`           | Get a single OAuth client          |
| `POST` | `/api/projects/:project_id/clients`               | Create OAuth client                |
| `PUT`  | `/api/projects/:project_id/clients/:id`           | Update OAuth client                |

**OAuth client fields:**

| Field           | Type   | Description                                                             |
|-----------------|--------|-------------------------------------------------------------------------|
| `name`          | string | Display name                                                            |
| `redirect_uris` | string | Space-separated allowed redirect URIs                                   |
| `grant_types`   | string | Space-separated grant types (e.g. `authorization_code`)                 |
| `require_pkce`  | bool   | When `true`, `code_challenge` is mandatory on login and social begin    |

> The `client_secret` is returned in full only on creation. Subsequent reads show only the last 4 characters (`****xxxx`) for security. The field is never overwritten if you submit a masked value.

---

### Social Providers 🔒

| Method | Path                                                        | Description                           |
|--------|-------------------------------------------------------------|---------------------------------------|
| `GET`  | `/api/projects/:project_id/clients/:id/social-providers`    | List social providers for a client    |
| `POST` | `/api/projects/:project_id/clients/:id/social-providers`    | Create or update a social provider    |

**Social provider fields:**

| Field           | Type   | Description                                              |
|-----------------|--------|----------------------------------------------------------|
| `provider`      | string | Provider name: `google`, `github`, `facebook`, `discord`, `linkedin` |
| `client_key`    | string | OAuth App Client ID from the provider's developer console |
| `client_secret` | string | OAuth App Client Secret (write-only; submit `****` to keep existing) |
| `scopes`        | string | Space-separated OAuth scopes                             |
| `enabled`       | bool   | Whether this provider is active for login/signup         |

---

### Login / Signup Page Configuration 🔒

| Method | Path                                                         | Description              |
|--------|--------------------------------------------------------------|--------------------------|
| `GET`  | `/api/projects/:project_id/clients/:id/login`                | Get login page config    |
| `POST` | `/api/projects/:project_id/clients/:id/login`                | Create login page config |
| `PUT`  | `/api/projects/:project_id/clients/:id/login`                | Update login page config |
| `GET`  | `/api/projects/:project_id/clients/:id/signup`               | Get signup page config   |
| `POST` | `/api/projects/:project_id/clients/:id/signup`               | Create signup page config|
| `PUT`  | `/api/projects/:project_id/clients/:id/signup`               | Update signup page config|

---

### Sessions

| Method   | Path                                                                | Auth | Description                             |
|----------|---------------------------------------------------------------------|------|-----------------------------------------|
| `GET`    | `/api/projects/:project_id/sessions`                                | 🔒   | List sessions for a project (paginated) |
| `GET`    | `/api/projects/:project_id/sessions/count`                          | 🔒   | Count sessions                          |
| `GET`    | `/api/projects/:project_id/accounts/:account_id/sessions`           | 🔒   | List sessions for a specific account    |
| `DELETE` | `/api/projects/:project_id/sessions/:session_id`                    | 🔒   | Revoke a session                        |
| `GET`    | `/api/projects/:project_id/accounts/:account_id/refresh-tokens`     | 🔒   | List refresh tokens for an account      |
| `GET`    | `/api/accounts/:account_id/sessions`                                | 🔑   | List sessions (client auth)             |
| `DELETE` | `/api/accounts/:account_id/sessions/:session_id`                    | 🔑   | Revoke a session (client auth)          |

---

### Dashboard 🔒

| Method | Path                   | Description                                          |
|--------|------------------------|------------------------------------------------------|
| `GET`  | `/api/dashboard/stats` | Aggregated statistics (projects, accounts, sessions) |

---

### ACL 🔒

| Method | Path                                             | Description                                             |
|--------|--------------------------------------------------|---------------------------------------------------------|
| `GET`  | `/api/acl`                                       | Fetch full ACL data (roles, modules, actions, policies) |
| `POST` | `/api/acl/role`                                  | Create a role                                           |
| `POST` | `/api/acl/module`                                | Create a module                                         |
| `POST` | `/api/acl/action`                                | Create an action                                        |
| `GET`  | `/api/acl/roles`                                 | List all roles                                          |
| `GET`  | `/api/acl/modules`                               | List all modules                                        |
| `GET`  | `/api/acl/actions`                               | List all actions                                        |
| `POST` | `/api/acl/enforce`                               | Check if a subject has permission                       |
| `GET`  | `/api/acl/permissions/user/:user/domain/:domain` | 🔑 Get user permissions in a domain                    |

**Enforce request body:**

```json
{
  "subject": "role:admin",
  "domain": "domain:locksmith",
  "object": "module:accounts",
  "action": "action:read:all"
}
```

Response `200 OK` = permission granted. `403 Forbidden` = denied.

---

### Config (public)

| Method | Path          | Description                                                     |
|--------|---------------|-----------------------------------------------------------------|
| `GET`  | `/api/config` | Returns `baseUrl` and `clientId` for the SPA to bootstrap OAuth |

---

## Access Control (ACL)

Locksmith uses **Casbin** with a domain-aware RBAC model:

```
subject → role   (e.g. role:admin)
domain  → tenant (e.g. domain:locksmith)
object  → module (e.g. module:accounts)
action  → op     (e.g. action:read:all)
```

### Enforcing permissions from your application

```http
POST /api/acl/enforce
Content-Type: application/json
Cookie: LOCKSMITHACCESSTOKEN=<token>

{
  "subject": "role:admin",
  "domain": "domain:myproject",
  "object": "module:orders",
  "action": "action:delete:one"
}
```

---

## Multi-Tenancy

Every resource is scoped to a **Project**. Projects are fully isolated:

| Resource                   | Isolated by project |
|----------------------------|---------------------|
| User accounts              | ✅                  |
| OAuth clients              | ✅                  |
| Social provider configs    | ✅                  |
| Sessions                   | ✅                  |
| ACL policies               | ✅                  |
| Login/Register page config | ✅                  |

---

## Session & Device Tracking

When a user authenticates, Locksmith captures:

| Field              | Source                                                   |
|--------------------|----------------------------------------------------------|
| `ip_address`       | `X-Forwarded-For` / `X-Real-IP` header                   |
| `device_type`      | User-Agent parsing (`mobile`, `desktop`, `tablet`)       |
| `browser`          | User-Agent parsing (Chrome, Firefox, Safari, etc.)       |
| `browser_version`  | User-Agent parsing                                       |
| `os`               | User-Agent parsing (Windows, macOS, Linux, iOS, Android) |
| `os_version`       | User-Agent parsing                                       |
| `location_country` | GeoIP via ip-api.com                                     |
| `location_city`    | GeoIP via ip-api.com                                     |

Sessions are visible in the dashboard under **Project Details → Logs**.

### Token security

| Token         | Format  | Storage          | Expiry   |
|---------------|---------|------------------|----------|
| Access token  | JWT     | HTTP-only cookie | 5 min    |
| Refresh token | UUID v4 | SHA-256 hash in DB | 15 days |

- **Rotation** — each refresh issues a new token and revokes the previous one
- **Reuse detection** — if a revoked refresh token is presented, the entire session is revoked immediately and all tokens in the chain are invalidated (`revoked_reason = 'token_reuse_detected'`)
- **Logout revocation** — logout marks the refresh token and its associated session as revoked server-side (`revoked_reason = 'user_logout'`); the tokens cannot be reused even if stolen

---

## Security

### Authentication hardening

| Mechanism                  | Details                                                              |
|----------------------------|----------------------------------------------------------------------|
| Password hashing           | Argon2id (64 MB memory, 3 iterations, parallelism 2)                 |
| Timing attack protection   | Dummy Argon2id hash always run on "email not found" path (CWE-204)  |
| PKCE                       | S256 only — `plain` rejected; `require_pkce` flag per client         |
| Rate limiting              | Per-IP limits on all auth endpoints (login: 5 req/min)               |
| Request size limit         | 1 MB maximum body size on all endpoints                              |
| Soft-deleted accounts      | `deleted_at IS NULL` enforced on every authentication query          |

### Cookie security

| Cookie                  | HttpOnly | Secure | SameSite |
|-------------------------|----------|--------|----------|
| `LOCKSMITHACCESSTOKEN`  | ✅       | ✅     | Strict   |
| `LOCKSMITHREFRESHTOKEN` | ✅       | ✅     | Strict   |
| `pkce_cv`               | —        | ✅     | Strict   |

### Secrets handling

- `client_secret` is masked in all read responses (shows only last 4 chars: `****xxxx`)
- Social provider `client_secret` is write-only (always returned as `****`)
- Submitting a masked value (`****`) in an update request preserves the existing credential without overwriting it
- Seeder credentials are injected via environment variables — no hardcoded secrets in config files

### Production mode

Set `APP_ENV=production` to enable opaque error responses. In production:
- HTTP error responses contain only `{ "message": "...", "error_id": "<uuid>" }`
- Full error details (stack trace, internal error) are logged server-side with the same `error_id` for correlation
- In development mode, full details are returned in the response for easier debugging

### CI/CD security scanning

The repository includes GitHub Actions workflows:

- **`security.yml`** — runs `govulncheck`, `gosec`, `golangci-lint` (Go) and `npm audit` (frontend) on every push and pull request
- **`codeql.yml`** — GitHub CodeQL SAST for Go and JavaScript (taint analysis, weekly schedule + push/PR)
- **`dependabot.yml`** — automated weekly dependency update PRs for Go, npm, and GitHub Actions

### Container security

The production Docker image:
- Based on `scratch` (no OS, no shell, no package manager)
- Runs as a dedicated non-root user (`locksmith`, UID assigned by Alpine)
- Multi-arch: `linux/amd64` and `linux/arm64`
- Binary compiled with `-ldflags="-s -w" -trimpath` (debug symbols stripped, build paths removed)

---

## Management Dashboard

The web dashboard is served at the root path (`/`) and provides a full UI for managing all Locksmith resources.

| Section              | Description                                                  |
|----------------------|--------------------------------------------------------------|
| Dashboard            | Overview with aggregated statistics                          |
| Projects             | Create and manage projects                                   |
| Project Details      | Config, Roles, OAuth Clients, Users, Logs tabs               |
| OAuth Client Details | Config, Login page, Signup page, Social Providers            |
| ACL                  | Manage global roles, modules, and actions                    |

---

## Development

### Prerequisites

Docker and Docker Compose.

### Setup

```bash
git clone https://github.com/locksmithhq/locksmith.git
cd locksmith
cp .env.example .env
make up
```

Open the dashboard at `http://localhost:${LOCKSMITH_APP_PORT}` (default: `http://localhost:4000`).

The development setup uses **Air** for Go hot reload and **Vite HMR** for the frontend. Changes to Go or Vue files are reflected without restarting containers.

### Available make commands

| Command          | Description                                               |
|------------------|-----------------------------------------------------------|
| `make up`        | Start all services (builds if needed)                     |
| `make down`      | Stop all services                                         |
| `make restart`   | Stop and restart all services                             |
| `make rebuild`   | Stop, rebuild images, and restart                         |
| `make logs`      | Tail logs for all services                                |
| `make logs-api`  | Tail API logs only                                        |
| `make logs-web`  | Tail frontend logs only                                   |
| `make logs-db`   | Tail database logs only                                   |
| `make shell-api` | Open shell inside the API container                       |
| `make shell-db`  | Open psql inside the database container                   |
| `make clean`     | Stop services and remove all volumes                      |
| `make status`    | Show container status                                     |
| `make open-web`  | Open dashboard in browser                                 |
| `make open-api`  | Open API directly in browser                              |
| `make build-prod`| Build multi-arch production image and push to Docker Hub  |

### Building and publishing the production image

```bash
make build-prod
```

This runs `docker buildx build --platform linux/amd64,linux/arm64` using the optimized multi-stage `locksmith/Dockerfile` with the project root as the build context. Requires `docker buildx` and an authenticated Docker Hub session (`docker login`).
