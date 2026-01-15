INSERT INTO school (id, name, location_id, created_at, updated_at) VALUES
-- New York schools
('b1eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'Manhattan Academy of Arts & Sciences', 'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', NOW(), NOW()),
('b1eebc99-9c0b-4ef8-bb6d-6bb9bd380a12', 'Times Square Preparatory School', 'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', NOW(), NOW()),
('b1eebc99-9c0b-4ef8-bb6d-6bb9bd380a13', 'Broadway High School', 'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', NOW(), NOW()),
-- Los Angeles schools
('b1eebc99-9c0b-4ef8-bb6d-6bb9bd380a14', 'Hollywood Arts Academy', 'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', NOW(), NOW()),
('b1eebc99-9c0b-4ef8-bb6d-6bb9bd380a15', 'Santa Monica Bay School', 'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', NOW(), NOW()),
('b1eebc99-9c0b-4ef8-bb6d-6bb9bd380a16', 'Pacific Coast Preparatory', 'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', NOW(), NOW()),
-- Chicago schools
('b1eebc99-9c0b-4ef8-bb6d-6bb9bd380a17', 'Lakeshore Academy', 'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', NOW(), NOW()),
('b1eebc99-9c0b-4ef8-bb6d-6bb9bd380a18', 'Michigan Avenue School of Excellence', 'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', NOW(), NOW()),
-- Miami schools
('b1eebc99-9c0b-4ef8-bb6d-6bb9bd380a19', 'Ocean Drive Academy', 'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', NOW(), NOW()),
('b1eebc99-9c0b-4ef8-bb6d-6bb9bd380a20', 'South Beach International School', 'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', NOW(), NOW()),
-- Seattle schools
('b1eebc99-9c0b-4ef8-bb6d-6bb9bd380a21', 'Pike Place School of Innovation', 'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', NOW(), NOW()),
('b1eebc99-9c0b-4ef8-bb6d-6bb9bd380a22', 'Emerald City Preparatory', 'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', NOW(), NOW()),
-- Austin schools
('b1eebc99-9c0b-4ef8-bb6d-6bb9bd380a23', 'Congress Avenue Academy', 'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', NOW(), NOW()),
('b1eebc99-9c0b-4ef8-bb6d-6bb9bd380a24', 'Austin STEM School', 'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', NOW(), NOW()),
-- Boston schools
('b1eebc99-9c0b-4ef8-bb6d-6bb9bd380a25', 'Boylston Academy', 'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a19', NOW(), NOW()),
('b1eebc99-9c0b-4ef8-bb6d-6bb9bd380a26', 'Back Bay Preparatory School', 'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a19', NOW(), NOW()),
-- San Francisco schools
('b1eebc99-9c0b-4ef8-bb6d-6bb9bd380a27', 'Market Street Academy', 'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a20', NOW(), NOW()),
('b1eebc99-9c0b-4ef8-bb6d-6bb9bd380a28', 'Golden Gate School of Technology', 'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a20', NOW(), NOW());

-- Verify the inserted data
SELECT COUNT(*) as total_schools FROM school;