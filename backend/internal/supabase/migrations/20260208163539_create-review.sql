create type review_category as enum ('fun', 'engaging','interesting', 'informative');

CREATE TABLE IF NOT EXISTS review (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    registration_id UUID NOT NULL REFERENCES registration(id),
    guardian_id UUID NOT NULL REFERENCES guardian(id),
    description TEXT, 
    categories review_category[] NOT NULL DEFAULT '{}',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TRIGGER update_review_updated_at
BEFORE UPDATE ON review
FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_column();
