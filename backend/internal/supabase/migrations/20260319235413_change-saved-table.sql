ALTER TABLE saved
DROP COLUMN event_occurrence_id;

ALTER TABLE saved
ADD COLUMN event_id UUID NOT NULL REFERENCES event(id);