-- ============================================
-- 3. LOCATIONS
-- ============================================
INSERT INTO location (id, latitude, longitude, address_line1, address_line2, subdistrict, district, province, postal_code, country) VALUES
-- Test locations for unit tests
('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 40.7128, -74.0060, '123 Broadway', NULL, 'Manhattan', 'New York County', 'NY', '10001', 'USA'),
('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a19', 42.3501, -71.0786, '456 Boylston St', NULL, 'Back Bay', 'Suffolk County', 'MA', '02116', 'USA'),
('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a20', 37.7749, -122.4194, '789 Market St', NULL, 'Financial District', 'San Francisco County', 'CA', '94102', 'USA'),
-- Schools
('10000000-0000-0000-0000-000000000001', 13.7563000, 100.5018000, '123 Sukhumvit Road', 'Building A', 'Khlong Toei', 'Khlong Toei', 'Bangkok', '10110', 'Thailand'),
('10000000-0000-0000-0000-000000000002', 13.7467000, 100.5350000, '456 Rama IV Road', NULL, 'Pathum Wan', 'Pathum Wan', 'Bangkok', '10330', 'Thailand'),
('10000000-0000-0000-0000-000000000003', 13.8200000, 100.5600000, '789 Vibhavadi Rangsit Road', 'Floor 2', 'Chatuchak', 'Chatuchak', 'Bangkok', '10900', 'Thailand'),
-- Organizations
('10000000-0000-0000-0000-000000000004', 13.7650000, 100.5380000, '321 Phetchaburi Road', 'Suite 15', 'Ratchathewi', 'Ratchathewi', 'Bangkok', '10400', 'Thailand'),
('10000000-0000-0000-0000-000000000005', 13.7300000, 100.5240000, '654 Sathorn Road', NULL, 'Yan Nawa', 'Sathorn', 'Bangkok', '10120', 'Thailand'),
('10000000-0000-0000-0000-000000000006', 13.7890000, 100.5600000, '147 Lat Phrao Road', 'Building C', 'Lat Phrao', 'Lat Phrao', 'Bangkok', '10230', 'Thailand'),
('10000000-0000-0000-0000-000000000007', 13.7250000, 100.4950000, '258 Silom Road', 'Floor 3', 'Suriyawong', 'Bang Rak', 'Bangkok', '10500', 'Thailand'),
-- Event occurrence locations
('10000000-0000-0000-0000-000000000008', 13.7400000, 100.5450000, '369 Wireless Road', NULL, 'Lumphini', 'Pathum Wan', 'Bangkok', '10330', 'Thailand'),
('10000000-0000-0000-0000-000000000009', 13.7100000, 100.5300000, '741 Narathiwat Road', 'Hall B', 'Chong Nonsi', 'Yan Nawa', 'Bangkok', '10120', 'Thailand'),
('10000000-0000-0000-0000-00000000000a', 13.8000000, 100.5500000, '852 Pradipat Road', NULL, 'Sam Sen Nai', 'Phaya Thai', 'Bangkok', '10400', 'Thailand');
