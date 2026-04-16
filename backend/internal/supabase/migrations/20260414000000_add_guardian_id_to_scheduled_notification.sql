ALTER TABLE scheduled_notification
ADD COLUMN IF NOT EXISTS guardian_id UUID REFERENCES guardian(id);
