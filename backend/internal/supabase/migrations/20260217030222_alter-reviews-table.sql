-- Add new columns (_en required, _th nullable for future translations)
ALTER TABLE review
ADD COLUMN description_en TEXT,
ADD COLUMN description_th TEXT;

-- Drop old columns
ALTER TABLE review
DROP COLUMN description;
