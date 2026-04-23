-- Replace SERIAL PK with UUID and document code hashing.
-- Existing rows (all short-lived, ~5 min TTL) are dropped — they cannot be
-- re-hashed retroactively and will have expired before any operator runs this.

DELETE FROM oauth_authorization_codes;

ALTER TABLE oauth_authorization_codes DROP CONSTRAINT oauth_authorization_codes_pkey;
ALTER TABLE oauth_authorization_codes DROP COLUMN id;
ALTER TABLE oauth_authorization_codes ADD COLUMN id UUID PRIMARY KEY DEFAULT gen_random_uuid();

COMMENT ON COLUMN oauth_authorization_codes.code IS 'SHA-256 hex digest of the raw authorization code. The raw value is sent to the client and never stored.';
