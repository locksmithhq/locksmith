-- Migration to add must_change_password column to accounts table
ALTER TABLE accounts ADD COLUMN must_change_password BOOLEAN NOT NULL DEFAULT FALSE;
