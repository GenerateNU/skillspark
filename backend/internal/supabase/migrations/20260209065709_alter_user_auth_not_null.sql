-- make auth_id a required field
ALTER TABLE "user"
ALTER COLUMN auth_id SET NOT NULL;

-- make auth_id type UUID
ALTER TABLE "user"
ALTER COLUMN auth_id SET DATA TYPE UUID;