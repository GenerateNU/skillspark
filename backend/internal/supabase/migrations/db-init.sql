-- initial migration to set up users table
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    username TEXT NOT NULL UNIQUE,
    profile_picture_s3_key TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- initial migration to set up customers table
CREATE TABLE IF NOT EXISTS customers (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);


-- initial migration to set up users table
CREATE TABLE IF NOT EXISTS children (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL,
    age_years INT NOT NULL,
    customer_id UUID NOT NULL REFERENCES customers(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- initial migration to set up organizations table which represents a business that offers skills sessions
create table if not exists organizations (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL,
    active BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- initial migration to set up locations for an organization
create table if not exists locations (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    latitude DECIMAL(10, 8) not null,
    longitude DECIMAL(11, 8) not null,
    organization_id UUID NOT NULL REFERENCES organizations(id) not null,
    address TEXT not null,
    city TEXT not null,
    state TEXT not null,
    zip_code TEXT not null,
    country TEXT not null,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- initial migration to set up providers table which represents a user who can offer skills sessions within an organization
create table if not exists providers (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id),
    organization_id UUID NOT NULL REFERENCES organizations(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- initial migration to set up event type that an event can be
create type enum event_type as enum ('tutor', 'activity');

-- initial migration to set up events table which represents a specific event for an organization
create table if not exists events (  
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    location_id UUID NOT NULL REFERENCES locations(id),
    organization_id UUID NOT NULL REFERENCES organizations(id),
    max_attendees INT NOT NULL,
    curr_enrolled INT NOT NULL DEFAULT 0,
    type event_type NOT NULL DEFAULT,
    tags TEXT[] NOT NULL DEFAULT '{}',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
)

-- initial migration to set up event occurrences table which represents a specific instance of an event with a provider and a start and end time
create table if not exists event_occurrences (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),    
    provider_id UUID REFERENCES providers(id),
    event_id UUID NOT NULL REFERENCES events(id),
    start_time TIME NOT NULL,
    end_time TIME NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

