-- Passwords previously hashed with pgcrypto (bcrypt) are incompatible with Argon2id.
-- Force all affected users to reset their password on next login.
UPDATE accounts SET must_change_password = true WHERE password NOT LIKE '$argon2id$%';
