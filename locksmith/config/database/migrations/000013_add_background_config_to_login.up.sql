ALTER TABLE oauth_clients_login ADD COLUMN background_color VARCHAR(255);
ALTER TABLE oauth_clients_login ADD COLUMN background_image TEXT;
ALTER TABLE oauth_clients_login ADD COLUMN background_type VARCHAR(50) DEFAULT 'color';
