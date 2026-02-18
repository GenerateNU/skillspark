-- Add new columns (_en required, _th nullable for future translations)
ALTER TABLE review
ADD COLUMN description_en TEXT,
ADD COLUMN description_th TEXT;

-- Make only _en columns NOT NULL (Thai will be populated later via API)
ALTER TABLE review
ALTER COLUMN description_en SET NOT NULL;

-- Drop old columns
ALTER TABLE review
DROP COLUMN description;
