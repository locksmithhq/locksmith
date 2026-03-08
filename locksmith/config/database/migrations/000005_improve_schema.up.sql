-- Melhoria de Schema: Índices e Consistência

-- 1. Padronizar timestamps na tabela accounts para TIMESTAMPTZ
ALTER TABLE accounts 
    ALTER COLUMN created_at TYPE TIMESTAMP WITH TIME ZONE,
    ALTER COLUMN updated_at TYPE TIMESTAMP WITH TIME ZONE,
    ALTER COLUMN deleted_at TYPE TIMESTAMP WITH TIME ZONE;

-- 2. Adicionar índices em Foreign Keys para performance
-- accounts
CREATE INDEX IF NOT EXISTS idx_accounts_project_id ON accounts(project_id);

-- oauth_clients
CREATE INDEX IF NOT EXISTS idx_oauth_clients_project_id ON oauth_clients(project_id);

-- user_sessions
CREATE INDEX IF NOT EXISTS idx_user_sessions_account_id ON user_sessions(account_id);
CREATE INDEX IF NOT EXISTS idx_user_sessions_client_id ON user_sessions(client_id);

-- refresh_tokens
CREATE INDEX IF NOT EXISTS idx_refresh_tokens_session_id ON refresh_tokens(session_id);
CREATE INDEX IF NOT EXISTS idx_refresh_tokens_account_id ON refresh_tokens(account_id);
CREATE INDEX IF NOT EXISTS idx_refresh_tokens_client_id ON refresh_tokens(client_id);
CREATE INDEX IF NOT EXISTS idx_refresh_tokens_parent_token_id ON refresh_tokens(parent_token_id);

-- oauth_authorization_codes
CREATE INDEX IF NOT EXISTS idx_oauth_codes_client_id ON oauth_authorization_codes(client_id);
CREATE INDEX IF NOT EXISTS idx_oauth_codes_account_id ON oauth_authorization_codes(account_id);

-- 3. Adicionar índices para limpeza de dados expirados (TTL)
CREATE INDEX IF NOT EXISTS idx_user_sessions_expires_at ON user_sessions(expires_at);
CREATE INDEX IF NOT EXISTS idx_refresh_tokens_expires_at ON refresh_tokens(expires_at);
CREATE INDEX IF NOT EXISTS idx_oauth_codes_expires_at ON oauth_authorization_codes(expires_at);

-- 4. Adicionar constraints de integridade de dados
-- Garantir valores válidos para device_type
ALTER TABLE user_sessions 
    ADD CONSTRAINT check_device_type 
    CHECK (device_type IN ('mobile', 'desktop', 'tablet', 'tv', 'watch', 'other'));

-- Comentários
COMMENT ON INDEX idx_accounts_project_id IS 'Performance para buscar contas por projeto';
COMMENT ON INDEX idx_user_sessions_expires_at IS 'Performance para jobs de limpeza de sessões expiradas';
