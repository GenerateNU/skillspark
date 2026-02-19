-- Create notification type enum
CREATE TYPE notification_type AS ENUM ('email', 'push', 'both');

-- Create notification status enum
CREATE TYPE notification_status AS ENUM ('pending', 'sent', 'failed');

-- Create scheduled_notification table
CREATE TABLE IF NOT EXISTS scheduled_notification (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    notification_type notification_type NOT NULL,
    recipient_email TEXT,
    recipient_push_token TEXT,
    subject TEXT,
    body TEXT NOT NULL,
    metadata JSONB,
    scheduled_for TIMESTAMPTZ NOT NULL,
    sent_at TIMESTAMPTZ,
    status notification_status NOT NULL DEFAULT 'pending',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Create index on scheduled_for and status for efficient querying
CREATE INDEX idx_scheduled_notification_scheduled_for_status 
ON scheduled_notification(scheduled_for, status) 
WHERE status = 'pending';

-- Create trigger to update updated_at timestamp
CREATE TRIGGER update_scheduled_notification_updated_at
    BEFORE UPDATE ON scheduled_notification
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

