INSERT INTO location (id, latitude, longitude, address, city, state, zip_code, country, created_at, updated_at) VALUES
-- New York location
('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 40.7128, -74.0060, '123 Broadway', 'New York', 'NY', '10001', 'USA', NOW(), NOW()),
('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a12', 40.7589, -73.9851, '456 Times Square', 'New York', 'NY', '10036', 'USA', NOW(), NOW()),

-- Los Angeles location
('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a13', 34.0522, -118.2437, '789 Hollywood Blvd', 'Los Angeles', 'CA', '90028', 'USA', NOW(), NOW()),
('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a14', 34.0195, -118.4912, '101 Santa Monica Pier', 'Santa Monica', 'CA', '90401', 'USA', NOW(), NOW()),
-- Chicago location
('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a15', 41.8781, -87.6298, '200 Michigan Avenue', 'Chicago', 'IL', '60601', 'USA', NOW(), NOW()),

-- Miami location
('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a16', 25.7617, -80.1918, '300 Ocean Drive', 'Miami', 'FL', '33139', 'USA', NOW(), NOW()),

-- Seattle location
('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a17', 47.6062, -122.3321, '400 Pike Street', 'Seattle', 'WA', '98101', 'USA', NOW(), NOW()),

-- Austin location
('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a18', 30.2672, -97.7431, '500 Congress Avenue', 'Austin', 'TX', '78701', 'USA', NOW(), NOW()),

-- Boston location
('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a19', 42.3601, -71.0589, '600 Boylston Street', 'Boston', 'MA', '02116', 'USA', NOW(), NOW()),

-- San Francisco location
('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a20', 37.7749, -122.4194, '700 Market Street', 'San Francisco', 'CA', '94102', 'USA', NOW(), NOW());

-- Verify the inserted data
SELECT COUNT(*) as total_locations FROM location;