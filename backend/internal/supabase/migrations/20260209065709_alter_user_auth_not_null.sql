-- make auth_id a required field
ALTER TABLE "user"
ALTER COLUMN auth_id SET NOT NULL;