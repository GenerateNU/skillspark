ALTER TABLE event_occurrence
DROP COLUMN location_id;

ALTER TABLE organization
ALTER COLUMN location_id SET NOT NULL;