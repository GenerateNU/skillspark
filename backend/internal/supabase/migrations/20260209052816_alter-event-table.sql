
ALTER TABLE event
ADD COLUMN title_en TEXT,
ADD COLUMN title_th TEXT,
ADD COLUMN description_en TEXT,
ADD COLUMN description_th TEXT;

-- Drop old columns
ALTER TABLE event
DROP COLUMN title,
DROP COLUMN description;