DROP INDEX IF EXISTS idx_user_sessions_device_lookup;
ALTER TABLE user_sessions DROP COLUMN device_id;
