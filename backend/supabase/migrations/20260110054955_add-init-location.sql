-- initial migration to set up locations for an organization

-- Enable UUID generation extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table if not exists locations (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    latitude DECIMAL(10, 8) not null,
    longitude DECIMAL(11, 8) not null,
    organization_id UUID NOT NULL,-- REFERENCES organizations(id) not null,
    address TEXT not null,
    city TEXT not null,
    state TEXT not null,
    zip_code TEXT not null,
    country TEXT not null,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);