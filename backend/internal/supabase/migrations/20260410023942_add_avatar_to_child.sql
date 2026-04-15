ALTER TABLE child
    ADD COLUMN IF NOT EXISTS avatar_face TEXT,
    ADD COLUMN IF NOT EXISTS avatar_background TEXT;
