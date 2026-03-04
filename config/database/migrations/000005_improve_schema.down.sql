-- Reverter Melhoria de Schema

-- 1. Remover constraints
ALTER TABLE user_sessions DROP CONSTRAINT IF EXISTS check_device_type;

-- 2. Remover índices de limpeza
DROP INDEX IF EXISTS idx_oauth_codes_expires_at;
DROP INDEX IF EXISTS idx_refresh_tokens_expires_at;
DROP INDEX IF EXISTS idx_user_sessions_expires_at;

-- 3. Remover índices de Foreign Keys
DROP INDEX IF EXISTS idx_oauth_codes_account_id;
DROP INDEX IF EXISTS idx_oauth_codes_client_id;
DROP INDEX IF EXISTS idx_refresh_tokens_parent_token_id;
DROP INDEX IF EXISTS idx_refresh_tokens_client_id;
DROP INDEX IF EXISTS idx_refresh_tokens_account_id;
DROP INDEX IF EXISTS idx_refresh_tokens_session_id;
DROP INDEX IF EXISTS idx_user_sessions_client_id;
DROP INDEX IF EXISTS idx_user_sessions_account_id;
DROP INDEX IF EXISTS idx_oauth_clients_project_id;
DROP INDEX IF EXISTS idx_accounts_project_id;

-- 4. Reverter timestamps para TIMESTAMP (sem time zone)
ALTER TABLE accounts 
    ALTER COLUMN created_at TYPE TIMESTAMP,
    ALTER COLUMN updated_at TYPE TIMESTAMP,
    ALTER COLUMN deleted_at TYPE TIMESTAMP;
