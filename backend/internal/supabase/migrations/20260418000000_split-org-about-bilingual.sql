ALTER TABLE organization
ADD COLUMN about_en TEXT,
ADD COLUMN about_th TEXT;

UPDATE organization SET about_en = about WHERE about IS NOT NULL;

ALTER TABLE organization
DROP COLUMN about;
