-- Migration to remove must_change_password column from accounts table
ALTER TABLE accounts DROP COLUMN must_change_password;
