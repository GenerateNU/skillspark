CREATE TYPE event_occurrence_status AS ENUM ('scheduled', 'cancelled');

ALTER TABLE event_occurrence
ADD COLUMN status event_occurrence_status NOT NULL DEFAULT 'scheduled';