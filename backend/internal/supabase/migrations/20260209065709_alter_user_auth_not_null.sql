-- drop column and re-add with auth_id as UUID
ALTER TABLE "user"
DROP COLUMN IF EXISTS auth_id;

ALTER TABLE "user"
ADD COLUMN auth_id UUID UNIQUE;

-- make auth_id a required field
ALTER TABLE "user"
ALTER COLUMN auth_id SET NOT NULL;