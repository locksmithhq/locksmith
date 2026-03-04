ALTER TABLE user_sessions ADD COLUMN device_id VARCHAR(255);
CREATE INDEX idx_user_sessions_device_lookup ON user_sessions(account_id, client_id, device_id);
