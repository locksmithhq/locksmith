CREATE TABLE oauth_clients_login (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	client_id UUID UNIQUE REFERENCES oauth_clients(id) ON DELETE CASCADE,
	custom_css TEXT,
	custom_html TEXT,
	input_variant TEXT,
	layout TEXT,
	show_forgot_password BOOLEAN DEFAULT FALSE,
	show_remember_me BOOLEAN DEFAULT FALSE,
	show_sign_up BOOLEAN DEFAULT FALSE,
	show_social BOOLEAN DEFAULT FALSE,
	use_custom_html BOOLEAN DEFAULT FALSE,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE oauth_clients_signup (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	client_id UUID UNIQUE REFERENCES oauth_clients(id) ON DELETE CASCADE,
	custom_css TEXT,
	custom_html TEXT,
	input_variant TEXT,
	layout TEXT,
	show_social BOOLEAN DEFAULT FALSE,
	use_custom_html BOOLEAN DEFAULT FALSE,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);