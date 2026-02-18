CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

DROP TABLE IF EXISTS school;
DROP TABLE IF EXISTS location;

-- initial migration to set up users table
CREATE TABLE IF NOT EXISTS profile (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    username TEXT NOT NULL UNIQUE,
    profile_picture_s3_key TEXT,
    language_preference TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- initial migration to set up customers table
CREATE TABLE IF NOT EXISTS guardian (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES profile(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- initial migration to set up locations for an organization
create table if not exists location (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    latitude DECIMAL(10, 8) not null,
    longitude DECIMAL(11, 8) not null,
    -- building_name TEXT,
    -- building_number TEXT,
    -- street_name TEXT,
    -- secondary_address TEXT,
    address_line1 TEXT not null,
    address_line2 TEXT,
    subdistrict TEXT not null,
    district TEXT not null,
    province TEXT not null,
    postal_code TEXT not null,
    country TEXT not null,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS school (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    location_id UUID REFERENCES location(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

create type category as enum ('science', 'math', 'music', 'art', 'sports', 'technology', 'language', 'other');

-- initial migration to set up users table
CREATE TABLE IF NOT EXISTS child (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    school_id UUID REFERENCES school(id),
    birth_month INT NOT NULL,
    birth_year INT NOT NULL,
    interests category[],
    guardian_id UUID NOT NULL REFERENCES guardian(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

create type org_type as enum ('educational', 'musical', 'artistic', 'physical', 'other');

-- initial migration to set up organizations table which represents a business that offers skills sessions
create table if not exists organization (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    active BOOLEAN NOT NULL DEFAULT FALSE,
    pfp_s3_key TEXT,
    location_id UUID REFERENCES location(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- initial migration to set up providers table which represents a user who can offer skills sessions within an organization
CREATE TABLE IF NOT EXISTS manager (
    -- on delete id deletion should cascade
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES profile(id),
    organization_id UUID REFERENCES organization(id),
    role TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- initial migration to set up event type that an event can be

-- initial migration to set up events table which represents a specific event for an organization
create table if not exists event (  
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title TEXT NOT NULL,
    description TEXT NOT NULL, 
    organization_id UUID NOT NULL REFERENCES organization(id),
    age_range_min INT,
    age_range_max INT,
    category category[] NOT NULL,
    header_image_s3_key TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT check_age_range CHECK (age_range_min IS NULL OR age_range_max IS NULL OR age_range_min <= age_range_max)
);

-- initial migration to set up event occurrences table which represents a specific instance of an event with a provider and a start and end time
create table if not exists event_occurrence (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),    
    -- cascading null here
    manager_id UUID REFERENCES manager(id),
    event_id UUID NOT NULL REFERENCES event(id),
    location_id UUID NOT NULL REFERENCES location(id),
    start_time TIMESTAMPTZ NOT NULL,
    end_time TIMESTAMPTZ NOT NULL,
    max_attendees INT NOT NULL,
    language TEXT NOT NULL,
    curr_enrolled INT NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TYPE registration_status AS ENUM ('registered', 'cancelled');

CREATE TABLE IF NOT EXISTS registration (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    child_id UUID NOT NULL REFERENCES child(id),
    guardian_id UUID NOT NULL REFERENCES guardian(id),
    event_occurrence_id UUID NOT NULL REFERENCES event_occurrence(id),
    status registration_status NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Create function to automatically update updated_at timestamp
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ language 'plpgsql';

-- Create triggers for each table to update updated_at on row modification
CREATE TRIGGER update_profile_updated_at
BEFORE UPDATE ON profile
FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_column();

CREATE TRIGGER update_guardian_updated_at
BEFORE UPDATE ON guardian
FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_column();

CREATE TRIGGER update_school_updated_at
BEFORE UPDATE ON school
FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_column();

CREATE TRIGGER update_child_updated_at
BEFORE UPDATE ON child
FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_column();

CREATE TRIGGER update_organization_updated_at
BEFORE UPDATE ON organization
FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_column();

CREATE TRIGGER update_manager_updated_at
BEFORE UPDATE ON manager
FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_column();

CREATE TRIGGER update_location_updated_at
BEFORE UPDATE ON location
FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_column();

CREATE TRIGGER update_event_updated_at
BEFORE UPDATE ON event
FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_column();

CREATE TRIGGER update_event_occurrence_updated_at
BEFORE UPDATE ON event_occurrence
FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_column();

CREATE TRIGGER update_registration_updated_at
BEFORE UPDATE ON registration
FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_column();