CREATE TABLE IF NOT EXISTS saved (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    guardian_id UUID NOT NULL REFERENCES guardian(id) ON DELETE CASCADE,
    event_occurrence_id UUID NOT NULL REFERENCES event_occurrence(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TRIGGER update_saved_updated_at
BEFORE UPDATE ON saved
FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_column();