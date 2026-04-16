-- ============================================
-- 5. CHILDREN
-- ============================================
INSERT INTO child (id, name, school_id, birth_month, birth_year, interests, guardian_id) VALUES
-- Sarah Johnson's children
('30000000-0000-0000-0000-000000000001', 'Emily Johnson',      '20000000-0000-0000-0000-000000000001', 3,  2016, ARRAY['science','technology','math']::category[],  '11111111-1111-1111-1111-111111111111'),
('30000000-0000-0000-0000-000000000002', 'Alex Johnson',       '20000000-0000-0000-0000-000000000001', 7,  2018, ARRAY['sports','music']::category[],               '11111111-1111-1111-1111-111111111111'),
-- Michael Chen's children
('30000000-0000-0000-0000-000000000003', 'Sophie Chen',        '20000000-0000-0000-0000-000000000002', 11, 2015, ARRAY['art','language','music']::category[],       '22222222-2222-2222-2222-222222222222'),
-- Priya Patel's children
('30000000-0000-0000-0000-000000000004', 'Aiden Patel',        '20000000-0000-0000-0000-000000000001', 5,  2017, ARRAY['science','sports','technology']::category[], '33333333-3333-3333-3333-333333333333'),
('30000000-0000-0000-0000-000000000005', 'Maya Patel',         '20000000-0000-0000-0000-000000000001', 9,  2019, ARRAY['art','music']::category[],                  '33333333-3333-3333-3333-333333333333'),
-- Carlos Rodriguez's children
('30000000-0000-0000-0000-000000000006', 'Lucas Rodriguez',    '20000000-0000-0000-0000-000000000002', 1,  2016, ARRAY['sports','technology']::category[],          '44444444-4444-4444-4444-444444444444'),
-- Emma Thompson's children
('30000000-0000-0000-0000-000000000007', 'Isabella Thompson',  '20000000-0000-0000-0000-000000000003', 8,  2017, ARRAY['language','art','other']::category[],       '55555555-5555-5555-5555-555555555555'),
('30000000-0000-0000-0000-000000000008', 'Ethan Thompson',     '20000000-0000-0000-0000-000000000003', 12, 2018, ARRAY['math','science']::category[],               '55555555-5555-5555-5555-555555555555'),
-- Yuki Tanaka's children
('30000000-0000-0000-0000-000000000009', 'Hana Tanaka',        '20000000-0000-0000-0000-000000000002', 4,  2016, ARRAY['music','language','art']::category[],       '66666666-6666-6666-6666-666666666666'),
-- Olivia Martinez's children
('30000000-0000-0000-0000-00000000000a', 'Noah Martinez',      '20000000-0000-0000-0000-000000000001', 6,  2017, ARRAY['sports','science']::category[],             '77777777-7777-7777-7777-777777777777'),
-- James Wilson's children
('30000000-0000-0000-0000-00000000000b', 'Liam Wilson',        '20000000-0000-0000-0000-000000000003', 10, 2015, ARRAY['technology','math','science']::category[],  '88888888-8888-8888-8888-888888888888'),
('30000000-0000-0000-0000-00000000000c', 'Ava Wilson',         '20000000-0000-0000-0000-000000000003', 2,  2019, ARRAY['music','art']::category[],                  '88888888-8888-8888-8888-888888888888'),
-- Nattaporn Chaiyasit's children (Bangkok)
('30000000-0000-0000-0000-00000000000d', 'Pear Chaiyasit',     '20000000-0000-0000-0000-000000000001', 9,  2017, ARRAY['science','math']::category[],               '99999999-9999-9999-9999-999999999999'),
('30000000-0000-0000-0000-00000000000e', 'Fah Chaiyasit',      '20000000-0000-0000-0000-000000000002', 3,  2015, ARRAY['art','music']::category[],                  '99999999-9999-9999-9999-999999999999'),
-- Ananya Krishnamurthy's children (Bangkok)
('30000000-0000-0000-0000-00000000000f', 'Aryan Krishnamurthy','20000000-0000-0000-0000-000000000001', 4,  2014, ARRAY['science','technology']::category[],         'bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb'),
('30000000-0000-0000-0000-000000000010', 'Priya Krishnamurthy','20000000-0000-0000-0000-000000000001', 8,  2017, ARRAY['math','art']::category[],                   'bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb'),
-- Siriporn Wattanabe's children (Bangkok)
('30000000-0000-0000-0000-000000000011', 'Krit Wattanabe',     '20000000-0000-0000-0000-000000000002', 2,  2016, ARRAY['technology','science']::category[],         'dddddddd-dddd-dddd-dddd-dddddddddddd'),
('30000000-0000-0000-0000-000000000012', 'Ploy Wattanabe',     '20000000-0000-0000-0000-000000000003', 10, 2013, ARRAY['music','art']::category[],                  'dddddddd-dddd-dddd-dddd-dddddddddddd'),
-- Mei-Ling Huang's children (Bangkok)
('30000000-0000-0000-0000-000000000013', 'Wei Huang',          '20000000-0000-0000-0000-000000000001', 5,  2015, ARRAY['math','science']::category[],               'eeeeeeee-eeee-eeee-eeee-eeeeeeeeeeee'),
('30000000-0000-0000-0000-000000000014', 'Lily Huang',         '20000000-0000-0000-0000-000000000002', 12, 2016, ARRAY['art','music']::category[],                  'eeeeeeee-eeee-eeee-eeee-eeeeeeeeeeee'),
-- Arjun Sharma's children (Bangkok)
('30000000-0000-0000-0000-000000000015', 'Vikram Sharma',      '20000000-0000-0000-0000-000000000001', 11, 2013, ARRAY['math','technology']::category[],            '12121212-1212-1212-1212-121212121212'),
('30000000-0000-0000-0000-000000000016', 'Divya Sharma',       '20000000-0000-0000-0000-000000000002', 6,  2016, ARRAY['science','art']::category[],                '12121212-1212-1212-1212-121212121212'),
-- Somchai Thongprasert's children (Bangkok)
('30000000-0000-0000-0000-000000000017', 'Kamon Thongprasert', '20000000-0000-0000-0000-000000000003', 4,  2017, ARRAY['sports','science']::category[],             '14141414-1414-1414-1414-141414141414'),
('30000000-0000-0000-0000-000000000018', 'Nong Thongprasert',  '20000000-0000-0000-0000-000000000001', 7,  2020, ARRAY['art','other']::category[],                  '14141414-1414-1414-1414-141414141414'),
-- James O'Connor's children (Boston)
('30000000-0000-0000-0000-000000000019', 'Connor O''Connor',   '20000000-0000-0000-0000-000000000004', 6,  2016, ARRAY['sports','math']::category[],                'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa'),
('30000000-0000-0000-0000-00000000001a', 'Fiona O''Connor',    '20000000-0000-0000-0000-000000000004', 11, 2018, ARRAY['art','music']::category[],                  'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa'),
-- Thomas Brennan's children (Boston)
('30000000-0000-0000-0000-00000000001b', 'Sean Brennan',       '20000000-0000-0000-0000-000000000005', 1,  2015, ARRAY['sports','science']::category[],             'cccccccc-cccc-cccc-cccc-cccccccccccc'),
('30000000-0000-0000-0000-00000000001c', 'Aoife Brennan',      '20000000-0000-0000-0000-000000000005', 7,  2017, ARRAY['art','music']::category[],                  'cccccccc-cccc-cccc-cccc-cccccccccccc'),
-- Rachel Kim's children (Boston)
('30000000-0000-0000-0000-00000000001d', 'Hannah Kim',         '20000000-0000-0000-0000-000000000004', 3,  2016, ARRAY['math','science']::category[],               'ffffffff-ffff-ffff-ffff-ffffffffffff'),
('30000000-0000-0000-0000-00000000001e', 'Jake Kim',           '20000000-0000-0000-0000-000000000004', 9,  2014, ARRAY['sports','technology']::category[],          'ffffffff-ffff-ffff-ffff-ffffffffffff'),
-- Patricia Walsh's children (Boston)
('30000000-0000-0000-0000-00000000001f', 'Lena Walsh',         '20000000-0000-0000-0000-000000000005', 8,  2017, ARRAY['art','music']::category[],                  '13131313-1313-1313-1313-131313131313'),
('30000000-0000-0000-0000-000000000020', 'Declan Walsh',       '20000000-0000-0000-0000-000000000005', 2,  2015, ARRAY['sports','math']::category[],                '13131313-1313-1313-1313-131313131313');
