DROP TABLE IF EXISTS locations;

-- initial migration to set up users table
CREATE TABLE IF NOT EXISTS profiles (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    username TEXT NOT NULL UNIQUE,
    profile_picture_s3_key TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- initial migration to set up customers table
CREATE TABLE IF NOT EXISTS guardian (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES profiles(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS school (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL,
    location_id UUID REFERENCES location(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

create type category as enum ('science', 'math', 'music', 'art', 'sports', 'technology', 'language', 'other');

-- initial migration to set up users table
CREATE TABLE IF NOT EXISTS child (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
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
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL,
    active BOOLEAN NOT NULL DEFAULT FALSE,
    pfp_s3_key TEXT,
    location_id UUID REFERENCES location(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- initial migration to set up providers table which represents a user who can offer skills sessions within an organization
CREATE TABLE IF NOT EXISTS manager (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES profiles(id),
    organization_id UUID REFERENCES organization(id),
    role TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- initial migration to set up locations for an organization
create table if not exists location (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    latitude DECIMAL(10, 8) not null,
    longitude DECIMAL(11, 8) not null,
    street_number TEXT not null,
    street_name TEXT not null,
    secondary_address TEXT,
    city TEXT not null,
    state TEXT not null,
    postal_code TEXT not null,
    country TEXT not null,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- initial migration to set up event type that an event can be

-- initial migration to set up events table which represents a specific event for an organization
create table if not exists event (  
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    location_id UUID NOT NULL REFERENCES location(id),
    organization_id UUID NOT NULL REFERENCES organization(id),
    language TEXT NOT NULL,
    max_attendees INT NOT NULL,
    curr_enrolled INT NOT NULL DEFAULT 0,
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
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),    
    manager_id UUID REFERENCES manager(id),
    event_id UUID NOT NULL REFERENCES event(id),
    start_time TIMESTAMPTZ NOT NULL,
    end_time TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TYPE registration_status AS ENUM ('registered', 'cancelled');

CREATE TABLE IF NOT EXISTS registration (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    child_id UUID NOT NULL REFERENCES child(id),
    guardian_id UUID NOT NULL REFERENCES guardian(id),
    event_occurrence_id UUID NOT NULL REFERENCES event_occurrence(id),
    status enum registration_status NOT NULL,
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
CREATE TRIGGER update_profiles_updated_at
BEFORE UPDATE ON profiles
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