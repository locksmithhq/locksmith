ALTER TABLE account_social_providers
    ADD COLUMN updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW();
