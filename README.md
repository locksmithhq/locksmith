# Locksmith

**Locksmith** is an open-source OAuth2 Identity and Access Management (IAM) platform. It provides multi-tenant user authentication, fine-grained role-based access control, OAuth2 Authorization Code flow with PKCE, session tracking with device fingerprinting, and a fully featured management dashboard — all deployable as a single Docker container.

---

## Table of Contents

- [Features](#features)
- [Architecture](#architecture)
- [Quick Start](#quick-start)
  - [Docker Compose (recommended)](#docker-compose-recommended)
  - [Docker Run](#docker-run)
- [Configuration](#configuration)
- [Default Credentials](#default-credentials)
- [OAuth2 Flow](#oauth2-flow)
- [API Reference](#api-reference)
- [Access Control (ACL)](#access-control-acl)
- [Multi-Tenancy](#multi-tenancy)
- [Session & Device Tracking](#session--device-tracking)
- [Management Dashboard](#management-dashboard)
- [Development](#development)

---

## Features

- **OAuth2 Authorization Code + PKCE** — industry-standard token issuance with code challenge verification
- **Multi-tenant Projects** — fully isolated users, OAuth clients, roles, and sessions per project
- **Role-Based Access Control** — Casbin-powered RBAC with domain, module, and action granularity
- **Session Management** — per-session tracking with browser, OS, device type, IP address, and GeoIP location
- **Refresh Token Rotation** — secure refresh tokens stored as SHA-256 hashes with full rotation audit trail
- **Customizable Login/Register pages** — per-client UI theming, custom CSS/HTML, and field visibility controls
- **User Account Management** — CRUD, password hashing, forced password change on next login
- **Management Dashboard** — Vue 3 + Vuetify web UI for managing all resources
- **Single Binary + SPA** — API and dashboard bundled into one Docker image (no separate containers in production)
- **Database Migrations** — automatic schema creation and seeding on first boot

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
| Layer | Technology |
|---|---|
| Backend | Go 1.25+, Chi v5, Casbin v2, JWT |
| Frontend | Vue 3, Vuetify 3, Vite, Bun |
| Database | PostgreSQL 16 |
| Auth | OAuth2, PKCE, SHA-256 hashed refresh tokens |
| ACL | Casbin RBAC with domain support |

---

## Quick Start

### Docker Compose (recommended)

This is the simplest way to run Locksmith for development or self-hosted production.

**1. Create a `compose.yaml`:**

```yaml
services:
  locksmith:
    image: booscaaa/locksmith:latest
    container_name: locksmith
    ports:
      - "4000:4000"
    environment:
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

### Docker Compose with pgweb (database inspector)

Useful for development — adds a web-based PostgreSQL UI on port 8081.

```yaml
services:
  locksmith:
    image: booscaaa/locksmith:latest
    container_name: locksmith
    ports:
      - "4000:4000"
    environment:
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
    depends_on:
      database:
        condition: service_healthy
    networks:
      - locksmith-network
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
    networks:
      - locksmith-network
    restart: unless-stopped

  pgweb:
    image: sosedoff/pgweb
    container_name: locksmith-pgweb
    ports:
      - "8081:8081"
    environment:
      - DATABASE_URL=postgres://locksmith:locksmith123@database:5432/locksmith?sslmode=disable
    depends_on:
      - database
    networks:
      - locksmith-network
    restart: unless-stopped

volumes:
  locksmith-data:

networks:
  locksmith-network:
    driver: bridge
```

Access pgweb at `http://localhost:8081`.

---

### Docker Run

Run Locksmith with an external PostgreSQL instance using `docker run`.

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

**4. Open the dashboard:**

```
http://localhost:4000
```

---

### Production with custom domain

Replace `localhost:4000` with your actual domain. Locksmith uses `LOCKSMITH_BASE_URL` to build OAuth redirect URIs and set cookies, so it must match the public URL.

```bash
docker run -d \
  --name locksmith \
  --network locksmith-network \
  -p 4000:4000 \
  -e LOCKSMITH_APP_PORT=4000 \
  -e LOCKSMITH_BASE_URL=https://auth.example.com \
  -e LOCKSMITH_APP_CLIENT_ID=my-client-id \
  -e LOCKSMITH_APP_CLIENT_SECRET=my-client-secret \
  -e POSTGRES_HOST=your-postgres-host \
  -e POSTGRES_USER=locksmith \
  -e POSTGRES_PASSWORD=strong-password-here \
  -e POSTGRES_DB=locksmith \
  -e POSTGRES_PORT=5432 \
  -e SCHEMA=locksmith \
  -e SSL_MODE=require \
  --restart unless-stopped \
  booscaaa/locksmith:latest
```

> **Note:** Place a reverse proxy (nginx, Caddy, Traefik) in front of Locksmith to handle TLS termination. Locksmith itself runs on plain HTTP.

---

## Configuration

All configuration is done through environment variables.

### Required Variables

| Variable                      | Description                                                  | Example                 |
| ----------------------------- | ------------------------------------------------------------ | ----------------------- |
| `LOCKSMITH_APP_PORT`          | Port the API listens on                                      | `4000`                  |
| `LOCKSMITH_BASE_URL`          | Public base URL (used for OAuth callbacks and cookie domain) | `http://localhost:4000` |
| `LOCKSMITH_APP_CLIENT_ID`     | Client ID of the built-in Locksmith management client        | `my-client-id`          |
| `LOCKSMITH_APP_CLIENT_SECRET` | Client secret of the built-in management client              | `my-client-secret`      |
| `POSTGRES_HOST`               | PostgreSQL hostname                                          | `database`              |
| `POSTGRES_USER`               | PostgreSQL username                                          | `locksmith`             |
| `POSTGRES_PASSWORD`           | PostgreSQL password                                          | `locksmith123`          |
| `POSTGRES_DB`                 | PostgreSQL database name                                     | `locksmith`             |
| `POSTGRES_PORT`               | PostgreSQL port                                              | `5432`                  |
| `SCHEMA`                      | PostgreSQL schema name                                       | `locksmith`             |
| `SSL_MODE`                    | PostgreSQL SSL mode (`disable`, `require`, `verify-full`)    | `disable`               |

### Optional Variables

| Variable                      | Description                                                                                                                  | Default                  |
| ----------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ------------------------ |
| `VITE_LOCKSMITH_API_BASE_URL` | Override the API base URL used by the frontend SPA. Useful when the frontend is served from a different origin than the API. | `window.location.origin` |

### Notes

- `LOCKSMITH_APP_CLIENT_ID` and `LOCKSMITH_APP_CLIENT_SECRET` are used to bootstrap the management dashboard's own OAuth client on first boot. Change them from the defaults before deploying to production.
- `LOCKSMITH_BASE_URL` must be reachable from the browser — it is used as the OAuth redirect target and for setting the cookie domain.
- `VITE_LOCKSMITH_API_BASE_URL` is a build-time Vite variable. If not set, the SPA defaults to the current page's origin, which works correctly when the API and frontend are served from the same container.
- On first boot, Locksmith automatically runs all database migrations and seeds the default project, admin account, and management OAuth client.

---

## Default Credentials

On first boot, Locksmith creates the following defaults:

| Resource                 | Value                                          |
| ------------------------ | ---------------------------------------------- |
| **Default project**      | `Default Project` (domain: `domain:locksmith`) |
| **Admin email**          | `admin@locksmith.rs`                           |
| **Admin username**       | `admin`                                        |
| **Admin password**       | `admin`                                        |
| **Default OAuth client** | `Default Client`                               |

> **Change the admin password immediately after your first login.**

---

## Seeder Configuration

On first boot, Locksmith reads `/etc/locksmith/config/seeder.yaml` to create the default project, admin account, and management OAuth client. You can override this file to customize the initial seed without rebuilding the image.

### seeder.yaml reference

```yaml
default_project:
  name: Default Project
  description: Default project for Locksmith
  domain: domain:locksmith

default_account:
  name: Default Account
  email: admin@locksmith.rs
  password: admin
  username: admin

default_client:
  name: Default Client
  # Supports $ENV_VAR expansion — LOCKSMITH_BASE_URL is substituted at runtime
  redirect_uris: ${LOCKSMITH_BASE_URL}/api/locksmith/callback ${LOCKSMITH_BASE_URL}
  grant_types: authorization_code
```

> The `redirect_uris` field supports environment variable interpolation (`${VAR}`). The `client_id` and `client_secret` of the default client are always taken from `LOCKSMITH_APP_CLIENT_ID` and `LOCKSMITH_APP_CLIENT_SECRET` — they are not set in this file.

### Mounting a custom seeder with Docker Compose

```yaml
services:
  locksmith:
    image: locksmithhq/locksmith:latest
    volumes:
      - ./my-seeder.yaml:/etc/locksmith/config/seeder.yaml
    environment:
      LOCKSMITH_BASE_URL: https://auth.example.com
      LOCKSMITH_APP_CLIENT_ID: my-app-client
      LOCKSMITH_APP_CLIENT_SECRET: supersecret
      # ... other vars
```

### Mounting a custom seeder with Docker Run

```bash
docker run -d \
  -v $(pwd)/my-seeder.yaml:/etc/locksmith/config/seeder.yaml \
  -e LOCKSMITH_BASE_URL=https://auth.example.com \
  -e LOCKSMITH_APP_CLIENT_ID=my-app-client \
  -e LOCKSMITH_APP_CLIENT_SECRET=supersecret \
  locksmithhq/locksmith:latest
```

### Notes

- The seeder runs **once on startup**. If the project, account, or client already exists (matched by name / email / client_id), the record is reused and not duplicated.
- The file must be present at `/etc/locksmith/config/seeder.yaml` — the process will exit with an error if it is missing.
- After the first boot you can manage all resources through the dashboard or the API; the seeder file is no longer needed.

---

## OAuth2 Flow

Locksmith implements the **Authorization Code flow with PKCE** (RFC 7636).

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

The management dashboard uses Locksmith's own callback handler which sets HTTP-only cookies instead of returning tokens in the response body:

- `LOCKSMITHACCESSTOKEN` — JWT access token (5-minute expiry)
- `LOCKSMITHREFRESHTOKEN` — UUID refresh token (30-day expiry)

The cookie domain is derived automatically from the request's `Origin` header, supporting multi-domain deployments.

---

## API Reference

All API routes are prefixed with `/api`.

### Authentication

Routes marked with 🔒 require a valid `LOCKSMITHACCESSTOKEN` cookie.
Routes marked with 🔑 require HTTP Basic Auth (`client_id:client_secret`).
Routes with no marker are public.

---

### OAuth2

| Method | Path                        | Description                                    |
| ------ | --------------------------- | ---------------------------------------------- |
| `POST` | `/api/oauth2/authorize`     | Start authorization — create auth code         |
| `POST` | `/api/oauth2/login`         | Authenticate user credentials                  |
| `POST` | `/api/oauth2/access-token`  | Exchange auth code for access + refresh tokens |
| `POST` | `/api/oauth2/refresh-token` | Rotate refresh token and get new access token  |

---

### Locksmith Callback

| Method | Path                      | Description                                   |
| ------ | ------------------------- | --------------------------------------------- |
| `GET`  | `/api/locksmith/callback` | OAuth callback — exchanges code, sets cookies |
| `GET`  | `/api/locksmith/status`   | Check if current session is authenticated     |
| `POST` | `/api/locksmith/r`        | Refresh access token using cookie             |

---

### Projects 🔒

| Method   | Path                | Description                   |
| -------- | ------------------- | ----------------------------- |
| `GET`    | `/api/projects`     | List all projects (paginated) |
| `GET`    | `/api/projects/:id` | Get a single project          |
| `POST`   | `/api/projects`     | Create a new project          |
| `PUT`    | `/api/projects/:id` | Update a project              |
| `DELETE` | `/api/projects/:id` | Delete a project              |

---

### Accounts

| Method | Path                                             | Auth   | Description                            |
| ------ | ------------------------------------------------ | ------ | -------------------------------------- |
| `POST` | `/api/projects/:project_id/accounts`             | 🔒     | Create account (dashboard)             |
| `PUT`  | `/api/projects/:project_id/accounts/:account_id` | 🔒     | Update account (dashboard)             |
| `GET`  | `/api/projects/:project_id/accounts`             | 🔒     | List accounts with pagination          |
| `GET`  | `/api/projects/:project_id/accounts/count`       | 🔒     | Count accounts matching filters        |
| `POST` | `/api/accounts`                                  | 🔑     | Create account (client credentials)    |
| `GET`  | `/api/accounts/:id`                              | 🔑     | Get account by ID (client credentials) |
| `PUT`  | `/api/accounts/:account_id`                      | 🔑     | Update account (client credentials)    |
| `POST` | `/api/accounts/change-password`                  | public | Change password                        |

**Pagination query parameters** (for list endpoints):

| Parameter | Description                         | Example        |
| --------- | ----------------------------------- | -------------- |
| `page`    | Page number (1-based)               | `?page=2`      |
| `limit`   | Items per page                      | `?limit=20`    |
| `search`  | Search across name, email, username | `?search=john` |

---

### OAuth Clients 🔒

| Method | Path                                    | Description               |
| ------ | --------------------------------------- | ------------------------- |
| `GET`  | `/api/projects/:project_id/clients`     | List OAuth clients        |
| `GET`  | `/api/projects/:project_id/clients/:id` | Get a single OAuth client |
| `POST` | `/api/projects/:project_id/clients`     | Create OAuth client       |
| `PUT`  | `/api/projects/:project_id/clients/:id` | Update OAuth client       |

**Create/Update OAuth client body:**

```json
{
  "name": "My Application",
  "redirect_uris": "https://myapp.com/callback",
  "grant_types": "authorization_code refresh_token"
}
```

---

### Login Page Configuration 🔒

| Method | Path                                          | Description              |
| ------ | --------------------------------------------- | ------------------------ |
| `GET`  | `/api/projects/:project_id/clients/:id/login` | Get login page config    |
| `POST` | `/api/projects/:project_id/clients/:id/login` | Create login page config |
| `PUT`  | `/api/projects/:project_id/clients/:id/login` | Update login page config |

---

### Sessions 🔒

| Method | Path                                       | Description               |
| ------ | ------------------------------------------ | ------------------------- |
| `GET`  | `/api/projects/:project_id/sessions`       | List sessions (paginated) |
| `GET`  | `/api/projects/:project_id/sessions/count` | Count sessions            |

**Session query parameters:**

| Parameter | Description                             | Example          |
| --------- | --------------------------------------- | ---------------- |
| `page`    | Page number                             | `?page=1`        |
| `limit`   | Items per page (max 100)                | `?limit=20`      |
| `search`  | Search by user name, email, IP, browser | `?search=chrome` |

---

### ACL 🔒

| Method | Path                                             | Description                                             |
| ------ | ------------------------------------------------ | ------------------------------------------------------- |
| `GET`  | `/api/acl`                                       | Fetch full ACL data (roles, modules, actions, policies) |
| `POST` | `/api/acl/role`                                  | Create a role                                           |
| `POST` | `/api/acl/module`                                | Create a module                                         |
| `POST` | `/api/acl/action`                                | Create an action                                        |
| `GET`  | `/api/acl/roles`                                 | List all roles                                          |
| `GET`  | `/api/acl/modules`                               | List all modules                                        |
| `GET`  | `/api/acl/actions`                               | List all actions                                        |
| `GET`  | `/api/acl/projects/:projectId`                   | Get ACL policies for a project                          |
| `POST` | `/api/acl/projects/:projectId`                   | Assign role+module+action to a project                  |
| `POST` | `/api/acl/enforce`                               | Check if a subject has permission                       |
| `GET`  | `/api/acl/permissions/user/:user/domain/:domain` | 🔑 Get user permissions                                 |

**Enforce request body:**

```json
{
  "subject": "role:admin",
  "domain": "domain:locksmith",
  "object": "module:accounts",
  "action": "action:read:all"
}
```

---

### Config (public)

| Method | Path          | Description                                                     |
| ------ | ------------- | --------------------------------------------------------------- |
| `GET`  | `/api/config` | Returns `baseUrl` and `clientId` for the SPA to bootstrap OAuth |

---

## Access Control (ACL)

Locksmith uses **Casbin** with a domain-aware RBAC model. Permissions are structured as:

```
subject → role (e.g. role:admin)
domain  → tenant scope (e.g. domain:locksmith)
object  → module being accessed (e.g. module:accounts)
action  → operation being performed (e.g. action:read:all)
```

### Built-in permission structure

Every route in the dashboard is protected by middleware that checks the authenticated user's role against the project ACL:

```go
// Example: only users with read:all on module:accounts can list accounts
r.With(locksmith.AclMiddleware("domain:locksmith", "module:accounts", "action:read:all"))
```

### Assigning permissions

In the dashboard, go to **Project Details → Roles** to:

1. Create roles (e.g. `admin`, `developer`, `viewer`)
2. Create modules (e.g. `accounts`, `sessions`, `clients`)
3. Create actions (e.g. `read:all`, `create:one`, `delete:one`)
4. Assign role + module + action combinations to a project

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

Response `200 OK` means the permission is granted. Response `403 Forbidden` means it is denied.

---

## Multi-Tenancy

Every resource in Locksmith is scoped to a **Project**. Projects are fully isolated:

| Resource                   | Isolated by project |
| -------------------------- | ------------------- |
| User accounts              | ✅                  |
| OAuth clients              | ✅                  |
| Sessions                   | ✅                  |
| ACL policies               | ✅                  |
| Login/Register page config | ✅                  |

This means you can use a single Locksmith instance to manage authentication for multiple independent applications, each with its own users, clients, and permissions.

**Creating a new project:**

1. Log into the dashboard
2. Go to **Projects → New Project**
3. Fill in name, description, and domain (e.g. `domain:myapp`)
4. Create OAuth clients and invite users within the project

---

## Session & Device Tracking

When a user authenticates via `/api/oauth2/access-token`, Locksmith captures:

| Field              | Source                                                   |
| ------------------ | -------------------------------------------------------- |
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

- **Access tokens**: JWT, signed, 5-minute expiry
- **Refresh tokens**: UUID v4, stored as SHA-256 hash in the database (never plaintext), 30-day expiry
- **Rotation**: each refresh generates a new token and invalidates the previous one; the parent chain is maintained for audit
- **Revocation**: sessions can be revoked per-session via the `revoked` flag; the reason is recorded

---

## Management Dashboard

The web dashboard is served at the root path (`/`) and provides a full UI for managing all Locksmith resources.

### Main sections

| Section              | Path                        | Description                                        |
| -------------------- | --------------------------- | -------------------------------------------------- |
| Dashboard            | `/`                         | Overview                                           |
| Projects             | `/projects`                 | Create and manage projects                         |
| Project Details      | `/projects/:id`             | Tabs for Config, Roles, OAuth Clients, Users, Logs |
| OAuth Client Details | `/projects/:id/clients/:id` | Config, Login page, Register page                  |
| ACL                  | `/acl`                      | Manage global roles, modules, and actions          |

### Project Details tabs

| Tab               | Description                                                 |
| ----------------- | ----------------------------------------------------------- |
| **Config**        | Edit project name, description, and domain                  |
| **Roles**         | Create roles, modules, actions, and assign ACL policies     |
| **OAuth Clients** | Manage OAuth clients and copy credentials                   |
| **Users**         | Create and manage user accounts with server-side pagination |
| **Logs**          | View sessions with device info, IP, location, and status    |

---

## Development

To run Locksmith locally for development with hot reload:

**Prerequisites:** Docker, Docker Compose

**1. Clone the repository:**

```bash
git clone https://github.com/locksmithhq/locksmith.git
cd locksmith
```

**2. Copy the environment file:**

```bash
cp .env.example .env
```

**3. Start all services:**

```bash
make up
# or
docker compose up --build -d
```

**4. Open the dashboard:**

```
http://localhost:4000
```

**Available make commands:**

| Command          | Description                             |
| ---------------- | --------------------------------------- |
| `make up`        | Start all services (builds if needed)   |
| `make down`      | Stop all services                       |
| `make restart`   | Stop and restart all services           |
| `make rebuild`   | Stop, rebuild images, and restart       |
| `make logs`      | Tail logs for all services              |
| `make logs-api`  | Tail API logs only                      |
| `make logs-web`  | Tail frontend logs only                 |
| `make logs-db`   | Tail database logs only                 |
| `make shell-api` | Open shell inside the API container     |
| `make shell-db`  | Open psql inside the database container |
| `make clean`     | Stop services and remove all volumes    |
| `make status`    | Show container status                   |

**Development services:**

| Service         | URL                     | Description           |
| --------------- | ----------------------- | --------------------- |
| Dashboard + API | `http://localhost:4000` | Proxied through nginx |
| pgweb           | `http://localhost:8081` | PostgreSQL web UI     |

The development setup uses **Air** for Go hot reload and **Bun** with Vite HMR for the frontend. Changes to Go or Vue files are reflected without restarting containers.
