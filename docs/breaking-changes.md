# Breaking Changes

This document aggregates all breaking changes across Locksmith versions in one place.
It is intended for operators who need to quickly assess the impact of an upgrade,
especially when skipping multiple versions.

Each section links to the full upgrade guide in the corresponding version file.

---

## Severity legend

| Symbol | Meaning |
|---|---|
| 🔴 **Required** | The application will not start or will malfunction without this action |
| 🟡 **Action needed** | Data loss or incorrect behaviour without this action |
| 🟢 **Optional** | Backwards-compatible; recommended but not mandatory |

---

## v0.1.0

Full details and step-by-step instructions: [docs/changelog/v0.1.0.md — Upgrade Guide](changelog/v0.1.0.md#upgrade-guide)

| Severity | Area | Change | Action required |
|---|---|---|---|
| 🔴 | Secrets — `ENCRYPTION_KEY` | `oauth_clients.client_secret` and social provider credentials are now encrypted at rest (AES-256-GCM). The app will return errors when reading any encrypted field without the key. | Set `ENCRYPTION_KEY=<base64-32-bytes>` in your environment. Generate with `openssl rand -base64 32`. |
| 🔴 | Seeder | Admin and app credentials are no longer hardcoded in `seeder.yaml`. | Add `SEED_ADMIN_EMAIL`, `SEED_ADMIN_PASSWORD`, `SEED_APP_CLIENT_ID`, `SEED_APP_CLIENT_SECRET`, `SEED_APP_USER_PASSWORD` to your `.env`. |
| 🟡 | Secrets — existing rows | Existing `client_secret` and social provider credentials remain in plaintext until next write. They still work, but are unencrypted at rest. | Run the one-time migration tool: `ENCRYPTION_KEY=<key> go run ./cmd/migrate-secrets/main.go` |
| 🟡 | PKCE — `plain` rejected | `code_challenge_method=plain` is now explicitly rejected. Only `S256` is accepted. | Update any application sending `plain` to use `S256`. |
| 🟡 | API response — `client_secret` masked | `GET /clients/:id` and list endpoints now return `****<last-4>` instead of the full secret. | Clients that read the secret from GET responses must regenerate it via the admin UI or API. |
| 🟡 | Auth codes — existing rows deleted | Migration 000027 deletes all rows from `oauth_authorization_codes` before changing the PK to UUID. Auth codes have a 5-minute TTL, so in practice no valid code should exist at migration time. | Schedule the migration during low-traffic or after confirming no authorization flows are in progress. |
| 🟢 | Docker Compose — `pgweb` removed | The `pgweb` container is no longer part of `compose.yaml`. | Add it back in `compose.override.yaml` if needed. |

---

<!--
## v0.2.0

Full details: [docs/changelog/v0.2.0.md — Upgrade Guide](changelog/v0.2.0.md#upgrade-guide)

| Severity | Area | Change | Action required |
|---|---|---|---|
| ... | ... | ... | ... |
-->
