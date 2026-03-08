-- No rollback possible: original bcrypt hashes were overwritten by new logins.
-- This migration cannot be reversed automatically.
SELECT 1;
