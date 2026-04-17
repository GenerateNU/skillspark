-- ============================================
-- 14. NEW LOCATIONS — Boston-area + Extra Bangkok
-- ============================================
INSERT INTO location (id, latitude, longitude, address_line1, address_line2, subdistrict, district, province, postal_code, country) VALUES

-- ── Boston-area Schools ──────────────────────────────────────────────────────
('b0000001-0000-0000-0000-000000000001', 42.3336, -71.2092, '89 Kingsberry St',          NULL,        'Newton Centre',  'Middlesex County', 'MA', '02459', 'USA'),
('b0000001-0000-0000-0000-000000000002', 42.3450, -71.0820, '78 Avenue Louis Pasteur',   NULL,        'Fenway',         'Suffolk County',   'MA', '02115', 'USA'),
('b0000001-0000-0000-0000-000000000003', 42.3726, -71.1097, '459 Broadway',              NULL,        'Cambridge',      'Middlesex County', 'MA', '02139', 'USA'),

-- ── Boston-area Organizations ────────────────────────────────────────────────
-- MIT Kids Lab — Kendall Square, Cambridge
('b0000001-0000-0000-0000-000000000004', 42.3627, -71.0876, '100 Binney St',             'Suite 200', 'Kendall Square', 'Middlesex County', 'MA', '02142', 'USA'),
-- Boston Athletic Academy — Fenway
('b0000001-0000-0000-0000-000000000005', 42.3467, -71.0972, '4 Jersey St',               NULL,        'Fenway',         'Suffolk County',   'MA', '02215', 'USA'),
-- NEC Youth Programs — Huntington Ave
('b0000001-0000-0000-0000-000000000006', 42.3401, -71.0868, '290 Huntington Ave',        NULL,        'Back Bay',       'Suffolk County',   'MA', '02115', 'USA'),
-- Code & Create Boston — Seaport
('b0000001-0000-0000-0000-000000000007', 42.3530, -71.0474, '249 A St',                  'Floor 2',   'Seaport',        'Suffolk County',   'MA', '02210', 'USA'),
-- Boston Art Center Kids — Brookline
('b0000001-0000-0000-0000-000000000008', 42.3397, -71.1213, '8 Eliot St',                NULL,        'Brookline',      'Norfolk County',   'MA', '02146', 'USA'),

-- ── Extra Bangkok Organizations ──────────────────────────────────────────────
-- Siam Muay Thai Academy — Rama III
('b0000001-0000-0000-0000-000000000009', 13.7050, 100.5300, '512 Rama III Road',         'Unit A',    'Bang Kho Laem',  'Bang Kho Laem',   'Bangkok', '10120', 'Thailand'),
-- Bangkok Ballet & Dance Academy — Sukhumvit Soi 12
('b0000001-0000-0000-0000-000000000010', 13.7400, 100.5630, '14/1 Sukhumvit Soi 12',    '2nd Floor', 'Khlong Toei Nuea','Watthana',        'Bangkok', '10110', 'Thailand'),
-- Geniuses STEM Thailand — Si Ayutthaya Rd, Phaya Thai
('b0000001-0000-0000-0000-000000000011', 13.7620, 100.5390, '45 Si Ayutthaya Road',      'Building B','Thung Phaya Thai','Ratchathewi',     'Bangkok', '10400', 'Thailand'),

-- ── Extra Bangkok School ─────────────────────────────────────────────────────
('b0000001-0000-0000-0000-000000000012', 13.8480, 100.5710, '50 Ngam Wong Wan Road',     NULL,        'Lat Yao',        'Chatuchak',       'Bangkok', '10900', 'Thailand');
