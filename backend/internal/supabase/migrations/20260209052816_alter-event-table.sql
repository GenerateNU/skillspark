-- Add new columns (_en required, _th nullable for future translations)
ALTER TABLE event
ADD COLUMN title_en TEXT,
ADD COLUMN title_th TEXT,
ADD COLUMN description_en TEXT,
ADD COLUMN description_th TEXT;

-- Make only _en columns NOT NULL (Thai will be populated later via API)
ALTER TABLE event
ALTER COLUMN title_en SET NOT NULL,
ALTER COLUMN description_en SET NOT NULL;

-- Drop old columns
ALTER TABLE event
DROP COLUMN title,
DROP COLUMN description;