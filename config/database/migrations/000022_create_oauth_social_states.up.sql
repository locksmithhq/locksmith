CREATE TABLE oauth_social_states (
  nonce                 TEXT PRIMARY KEY,
  client_id             TEXT NOT NULL,
  redirect_uri          TEXT NOT NULL,
  state                 TEXT NOT NULL,
  code_challenge        TEXT,
  code_challenge_method TEXT,
  expires_at            TIMESTAMPTZ NOT NULL
);
