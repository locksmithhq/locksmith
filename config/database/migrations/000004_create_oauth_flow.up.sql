

-- Tabela de clientes OAuth
CREATE TABLE IF NOT EXISTS oauth_clients (
    id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    project_id UUID NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
    client_id VARCHAR(255) NOT NULL,
    client_secret VARCHAR(255) NOT NULL,
    redirect_uris VARCHAR(255) NOT NULL,
    grant_types VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    
    -- Constraint de unicidade por projeto
    CONSTRAINT unique_client_id_per_project UNIQUE (project_id, client_id)
);

-- Comentários
COMMENT ON TABLE oauth_clients IS 'Clientes OAuth isolados por projeto';
COMMENT ON COLUMN oauth_clients.project_id IS 'Referência ao projeto - cada cliente OAuth pertence a um único projeto';


-- Tabela de sessões de usuário (rastreamento de dispositivos e logins)
CREATE TABLE IF NOT EXISTS user_sessions (
    id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    account_id UUID NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
    client_id UUID NOT NULL REFERENCES oauth_clients(id) ON DELETE CASCADE,
    jti VARCHAR(255) UNIQUE NOT NULL, -- JWT ID (claim 'jti' do token)
    
    -- Informações de rede
    ip_address INET,
    
    -- Informações do dispositivo
    user_agent TEXT,
    device_name VARCHAR(255), -- Nome personalizado do dispositivo (ex: "iPhone de João")
    device_type VARCHAR(50), -- "mobile", "desktop", "tablet", "tv", "watch", "other"
    browser VARCHAR(100), -- "Chrome", "Safari", "Firefox", etc
    browser_version VARCHAR(50),
    os VARCHAR(100), -- "iOS", "Android", "Windows", "macOS", "Linux"
    os_version VARCHAR(50),
    
    -- Informações de localização
    location_country VARCHAR(2), -- Código ISO do país (ex: "BR")
    location_region VARCHAR(100), -- Estado/Região (ex: "SP")
    location_city VARCHAR(100), -- Cidade (ex: "São Paulo")
    
    -- Controle de acesso
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    revoked BOOLEAN DEFAULT FALSE,
    revoked_at TIMESTAMP WITH TIME ZONE,
    revoked_reason VARCHAR(255), -- "user_logout", "admin_revoke", "suspicious_activity", etc
    
    -- Auditoria
    last_activity TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);


-- Tabela de refresh tokens (para rotação segura)
CREATE TABLE IF NOT EXISTS refresh_tokens (
    id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    token_hash VARCHAR(255) UNIQUE NOT NULL, -- Hash SHA256 do refresh token (nunca salvar em texto plano)
    session_id UUID REFERENCES user_sessions(id) ON DELETE CASCADE,
    account_id UUID NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
    client_id UUID REFERENCES oauth_clients(id) ON DELETE CASCADE,
    
    -- Controle de rotação
    rotation_count INTEGER DEFAULT 0, -- Quantas vezes este token foi rotacionado
    parent_token_id UUID REFERENCES refresh_tokens(id) ON DELETE SET NULL, -- Token anterior na cadeia de rotação
    
    -- Controle de validade
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    revoked BOOLEAN DEFAULT FALSE,
    revoked_at TIMESTAMP WITH TIME ZONE,
    revoked_reason VARCHAR(255),
    
    -- Auditoria
    last_used_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS oauth_authorization_codes (
    id SERIAL PRIMARY KEY,
    code VARCHAR(255) UNIQUE NOT NULL,
    client_id UUID NOT NULL REFERENCES oauth_clients(id) ON DELETE CASCADE,
    account_id UUID NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
    redirect_uri VARCHAR(255) NOT NULL,
    code_challenge VARCHAR(255), -- Para PKCE
    code_challenge_method VARCHAR(10), -- 'S256' ou 'plain'
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    used BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Comentários para documentação
COMMENT ON TABLE user_sessions IS 'Rastreia todas as sessões/logins de usuários com informações detalhadas de dispositivo e localização';
COMMENT ON COLUMN user_sessions.jti IS 'JWT ID único para cada access token, permite revogação individual';
COMMENT ON COLUMN user_sessions.device_name IS 'Nome personalizado do dispositivo definido pelo usuário';
COMMENT ON COLUMN user_sessions.last_activity IS 'Atualizado a cada requisição autenticada para rastrear atividade';

COMMENT ON TABLE refresh_tokens IS 'Tokens de longa duração para renovar access tokens expirados, com suporte a rotação segura';
COMMENT ON COLUMN refresh_tokens.token_hash IS 'Hash SHA256 do refresh token - NUNCA armazenar o token em texto plano';
COMMENT ON COLUMN refresh_tokens.rotation_count IS 'Contador de rotações, útil para detectar reutilização suspeita';
COMMENT ON COLUMN refresh_tokens.parent_token_id IS 'Mantém cadeia de rotação para auditoria e detecção de ataques';
