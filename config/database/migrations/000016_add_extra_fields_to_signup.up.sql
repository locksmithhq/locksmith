ALTER TABLE oauth_clients_signup ADD COLUMN background_color VARCHAR(255);
ALTER TABLE oauth_clients_signup ADD COLUMN background_image TEXT;
ALTER TABLE oauth_clients_signup ADD COLUMN background_type VARCHAR(50) DEFAULT 'color';
ALTER TABLE oauth_clients_signup ADD COLUMN primary_color VARCHAR(255);
ALTER TABLE oauth_clients_signup ADD COLUMN logo_url TEXT;
