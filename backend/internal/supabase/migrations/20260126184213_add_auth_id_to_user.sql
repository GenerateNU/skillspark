-- Add auth_id column to user table
ALTER TABLE "user"
ADD COLUMN auth_id TEXT UNIQUE;

