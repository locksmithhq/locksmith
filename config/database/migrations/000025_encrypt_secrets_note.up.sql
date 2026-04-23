-- Secrets encryption notice (migration 000025)
--
-- oauth_clients.client_secret and oauth_client_social_providers.(client_key, client_secret)
-- are now encrypted at rest with AES-256-GCM by the application layer.
--
-- Existing plaintext values are handled transparently: on read, values without the
-- "aes256gcm:" prefix are returned as-is; on the next write they are re-encrypted.
--
-- REQUIRED: set ENCRYPTION_KEY in your environment before starting the application.
-- Generate a key:  openssl rand -base64 32
--
-- To re-encrypt all existing rows at once, run:
--   go run ./cmd/migrate-secrets/main.go
-- (see api/cmd/migrate-secrets for the one-time migration tool)

COMMENT ON COLUMN oauth_clients.client_secret IS 'AES-256-GCM encrypted (aes256gcm: prefix). See migration 000025.';
COMMENT ON COLUMN oauth_client_social_providers.client_key IS 'AES-256-GCM encrypted (aes256gcm: prefix). See migration 000025.';
COMMENT ON COLUMN oauth_client_social_providers.client_secret IS 'AES-256-GCM encrypted (aes256gcm: prefix). See migration 000025.';
