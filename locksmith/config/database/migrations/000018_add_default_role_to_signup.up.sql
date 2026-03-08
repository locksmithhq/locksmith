ALTER TABLE oauth_clients_signup ADD COLUMN default_role_name VARCHAR(100) NOT NULL DEFAULT 'user';
