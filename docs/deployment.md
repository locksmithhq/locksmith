# Deployment

This guide covers running Locksmith in production using either **Docker Compose** or plain **`docker run`** commands.

The production image is available at:

```
booscaaa/locksmith:latest
```

It is a multi-arch `scratch`-based image (`linux/amd64` + `linux/arm64`) that bundles the Go API, the Vue frontend, database migrations, and the seeder config. It runs as a non-root user (`locksmith:locksmith`).

---

## Environment variables reference

| Variable | Required | Description |
|---|---|---|
| `LOCKSMITH_APP_PORT` | Yes | Port the application listens on |
| `LOCKSMITH_BASE_URL` | Yes | Public base URL (e.g. `https://auth.example.com`) |
| `LOCKSMITH_APP_CLIENT_ID` | Yes | Internal Locksmith client ID |
| `LOCKSMITH_APP_CLIENT_SECRET` | Yes | Internal Locksmith client secret |
| `ENCRYPTION_KEY` | Yes | Base64-encoded 32-byte AES-256 key. Generate: `openssl rand -base64 32` |
| `APP_ENV` | Yes | Set to `production` to hide stack traces from API responses |
| `POSTGRES_HOST` | Yes | PostgreSQL host |
| `POSTGRES_PORT` | Yes | PostgreSQL port |
| `POSTGRES_USER` | Yes | PostgreSQL user |
| `POSTGRES_PASSWORD` | Yes | PostgreSQL password |
| `POSTGRES_DB` | Yes | PostgreSQL database name |
| `SCHEMA` | No | PostgreSQL schema name (default: `locksmith`) |
| `SSL_MODE` | No | PostgreSQL SSL mode (default: `disable`) |
| `SEED_ADMIN_EMAIL` | First boot | Admin account e-mail for seeder |
| `SEED_ADMIN_PASSWORD` | First boot | Admin account password for seeder |
| `SEED_APP_CLIENT_ID` | First boot | App client ID for seeder |
| `SEED_APP_CLIENT_SECRET` | First boot | App client secret for seeder |
| `SEED_APP_USER_PASSWORD` | First boot | Default user password for seeder |

---

## Option 1 — Docker Compose (recommended)

The easiest way to run the full stack (Postgres + Locksmith + reverse proxy).

**1. Copy the example env file:**

```bash
cp .env.example .env
```

**2. Fill in your values in `.env`:**

```env
POSTGRES_HOST=database
POSTGRES_USER=locksmith
POSTGRES_PASSWORD=changeme
POSTGRES_DB=locksmith
POSTGRES_PORT=5432
SCHEMA=locksmith
SSL_MODE=disable

LOCKSMITH_APP_PORT=4000
LOCKSMITH_BASE_URL=http://localhost:4000
LOCKSMITH_APP_CLIENT_ID=your-client-id
LOCKSMITH_APP_CLIENT_SECRET=your-client-secret

APP_ENV=production
ENCRYPTION_KEY=<output of: openssl rand -base64 32>

SEED_ADMIN_EMAIL=admin@example.com
SEED_ADMIN_PASSWORD=changeme
SEED_APP_CLIENT_ID=your-seed-client-id
SEED_APP_CLIENT_SECRET=your-seed-client-secret
SEED_APP_USER_PASSWORD=changeme
```

**3. Start:**

```bash
docker compose up -d
```

---

## Option 2 — `docker run` (no Compose)

Use this when you already have a PostgreSQL instance running and want to manage containers individually.

### Step 1 — Create a shared network

```bash
docker network create locksmith-network
```

### Step 2 — Run PostgreSQL

Skip this step if you already have a PostgreSQL instance.

```bash
docker run -d \
  --name postgres-locksmith \
  --network locksmith-network \
  --restart unless-stopped \
  -e POSTGRES_USER=locksmith \
  -e POSTGRES_PASSWORD=changeme \
  -e POSTGRES_DB=locksmith \
  -v locksmith-data:/var/lib/postgresql/data \
  postgres:16-alpine
```

### Step 3 — Run Locksmith

```bash
docker run -d \
  --name locksmith \
  --network locksmith-network \
  --restart unless-stopped \
  -p 4000:4000 \
  -e LOCKSMITH_APP_PORT=4000 \
  -e LOCKSMITH_BASE_URL=http://localhost:4000 \
  -e LOCKSMITH_APP_CLIENT_ID=your-client-id \
  -e LOCKSMITH_APP_CLIENT_SECRET=your-client-secret \
  -e ENCRYPTION_KEY="$(openssl rand -base64 32)" \
  -e APP_ENV=production \
  -e POSTGRES_HOST=postgres-locksmith \
  -e POSTGRES_PORT=5432 \
  -e POSTGRES_USER=locksmith \
  -e POSTGRES_PASSWORD=changeme \
  -e POSTGRES_DB=locksmith \
  -e SCHEMA=locksmith \
  -e SSL_MODE=disable \
  -e SEED_ADMIN_EMAIL=admin@example.com \
  -e SEED_ADMIN_PASSWORD=changeme \
  -e SEED_APP_CLIENT_ID=your-seed-client-id \
  -e SEED_APP_CLIENT_SECRET=your-seed-client-secret \
  -e SEED_APP_USER_PASSWORD=changeme \
  booscaaa/locksmith:latest
```

> **Note:** The `ENCRYPTION_KEY` above is generated inline for illustration. In production, generate it once, store it safely, and reuse it across restarts. Changing the key after data has been written will make all encrypted secrets unreadable.

### Step 3 (alternative) — Using an `.env` file

Instead of passing each variable inline, load them from a file:

```bash
docker run -d \
  --name locksmith \
  --network locksmith-network \
  --restart unless-stopped \
  -p 4000:4000 \
  --env-file .env \
  booscaaa/locksmith:latest
```

---

## Option 3 — Connecting to an external PostgreSQL (no container)

Same as Option 2 but skip Step 2 and point `POSTGRES_HOST` at your external instance:

```bash
docker run -d \
  --name locksmith \
  --restart unless-stopped \
  -p 4000:4000 \
  -e LOCKSMITH_APP_PORT=4000 \
  -e LOCKSMITH_BASE_URL=https://auth.example.com \
  -e LOCKSMITH_APP_CLIENT_ID=your-client-id \
  -e LOCKSMITH_APP_CLIENT_SECRET=your-client-secret \
  -e ENCRYPTION_KEY=<your-stable-key> \
  -e APP_ENV=production \
  -e POSTGRES_HOST=db.example.com \
  -e POSTGRES_PORT=5432 \
  -e POSTGRES_USER=locksmith \
  -e POSTGRES_PASSWORD=changeme \
  -e POSTGRES_DB=locksmith \
  -e SCHEMA=locksmith \
  -e SSL_MODE=require \
  -e SEED_ADMIN_EMAIL=admin@example.com \
  -e SEED_ADMIN_PASSWORD=changeme \
  -e SEED_APP_CLIENT_ID=your-seed-client-id \
  -e SEED_APP_CLIENT_SECRET=your-seed-client-secret \
  -e SEED_APP_USER_PASSWORD=changeme \
  booscaaa/locksmith:latest
```

---

## Verifying the deployment

```bash
# Check the application is healthy
curl http://localhost:4000/health

# Follow logs
docker logs -f locksmith
```

---

## Re-encrypting existing secrets after deployment

If you are upgrading from a version prior to v0.1.0, run the one-time migration tool to encrypt any plaintext secrets already in the database:

```bash
docker run --rm \
  --network locksmith-network \
  --env-file .env \
  --entrypoint /locksmith \
  booscaaa/locksmith:latest \
  migrate-secrets
```

> See [docs/changelog/v0.1.0.md — Upgrade Guide](changelog/v0.1.0.md#upgrade-guide) for full details.
