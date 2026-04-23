DELETE FROM oauth_authorization_codes;

ALTER TABLE oauth_authorization_codes DROP CONSTRAINT oauth_authorization_codes_pkey;
ALTER TABLE oauth_authorization_codes DROP COLUMN id;
ALTER TABLE oauth_authorization_codes ADD COLUMN id SERIAL PRIMARY KEY;

COMMENT ON COLUMN oauth_authorization_codes.code IS NULL;
