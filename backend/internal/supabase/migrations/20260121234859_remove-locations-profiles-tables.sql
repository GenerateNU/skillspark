drop table if exists "locations";

ALTER TABLE manager DROP CONSTRAINT IF EXISTS manager_user_id_fkey;
ALTER TABLE guardian DROP CONSTRAINT IF EXISTS guardian_user_id_fkey;

drop table if exists profile;

-- remove the profile table and create a new user table
CREATE TABLE IF NOT EXISTS "user" (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    username TEXT NOT NULL UNIQUE,
    profile_picture_s3_key TEXT,
    language_preference TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

DROP TRIGGER IF EXISTS update_user_updated_at ON "user";
CREATE TRIGGER update_user_updated_at

BEFORE UPDATE ON "user"
FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_column();

ALTER TABLE IF EXISTS manager
ADD CONSTRAINT manager_user_id_fkey
FOREIGN KEY (user_id)
REFERENCES "user" (id);

ALTER TABLE IF EXISTS guardian
ADD CONSTRAINT guardian_user_id_fkey
FOREIGN KEY (user_id)
REFERENCES "user" (id);