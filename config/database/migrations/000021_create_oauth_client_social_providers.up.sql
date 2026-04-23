CREATE TABLE oauth_client_social_providers (
  id            UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  client_id     UUID NOT NULL REFERENCES oauth_clients(id) ON DELETE CASCADE,
  provider      TEXT NOT NULL,
  client_key    TEXT NOT NULL,
  client_secret TEXT NOT NULL,
  enabled       BOOLEAN NOT NULL DEFAULT false,
  scopes        TEXT NOT NULL DEFAULT 'email profile',
  created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  UNIQUE (client_id, provider)
);
