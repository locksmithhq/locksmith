CREATE TABLE accounts (
    id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    project_id UUID NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    
    -- Constraints de unicidade por projeto
    CONSTRAINT unique_email_per_project UNIQUE (project_id, email),
    CONSTRAINT unique_username_per_project UNIQUE (project_id, username)
);

-- Comentários
COMMENT ON TABLE accounts IS 'Contas de usuários isoladas por projeto';
COMMENT ON COLUMN accounts.project_id IS 'Referência ao projeto - cada conta pertence a um único projeto';