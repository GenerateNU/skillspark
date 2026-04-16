-- ============================================
-- 1. user (User Accounts)
-- ============================================
INSERT INTO "user" (id, name, email, username, profile_picture_s3_key, language_preference, auth_id) VALUES
-- Guardians
('a1b2c3d4-e5f6-4a7b-8c9d-0e1f2a3b4c5d', 'Sarah Johnson',     'sarah.johnson@email.com',    'sarahj',    'profiles/sarah_johnson.jpg',    'en', '00000000-0000-0000-0000-00000000000a'),
('b2c3d4e5-f6a7-4b8c-9d0e-1f2a3b4c5d6e', 'Michael Chen',      'michael.chen@email.com',     'mchen',     'profiles/michael_chen.jpg',     'en', '00000000-0000-0000-0000-00000000000b'),
('c3d4e5f6-a7b8-4c9d-0e1f-2a3b4c5d6e7f', 'Priya Patel',       'priya.patel@email.com',      'priyap',    'profiles/priya_patel.jpg',      'en', '00000000-0000-0000-0000-00000000000c'),
('d4e5f6a7-b8c9-4d0e-1f2a-3b4c5d6e7f8a', 'Carlos Rodriguez',  'carlos.rodriguez@email.com', 'carlosr',   'profiles/carlos_rodriguez.jpg', 'es', '00000000-0000-0000-0000-00000000000d'),
('e5f6a7b8-c9d0-4e1f-2a3b-4c5d6e7f8a9b', 'Emma Thompson',     'emma.thompson@email.com',    'emmat',     'profiles/emma_thompson.jpg',    'en', '00000000-0000-0000-0000-00000000000e'),
('f6a7b8c9-d0e1-4f2a-3b4c-5d6e7f8a9b0c', 'Yuki Tanaka',       'yuki.tanaka@email.com',      'yukit',     'profiles/yuki_tanaka.jpg',      'ja', '00000000-0000-0000-0000-00000000000f'),
('a7b8c9d0-e1f2-4a3b-4c5d-6e7f8a9b0c1d', 'Olivia Martinez',   'olivia.martinez@email.com',  'oliviam',   'profiles/olivia_martinez.jpg',  'es', '00000000-0000-0000-0000-000000000001'),
('b8c9d0e1-f2a3-4b4c-5d6e-7f8a9b0c1d2e', 'James Wilson',      'james.wilson@email.com',     'jamesw',    'profiles/james_wilson.jpg',     'en', '00000000-0000-0000-0000-000000000002'),
-- Bangkok org managers
('c9d0e1f2-a3b4-4c5d-6e7f-8a9b0c1d2e3f', 'Dr. Amanda Lee',    'amanda.lee@scienceacademy.com',    'alee',      'profiles/amanda_lee.jpg',      'en', '00000000-0000-0000-0000-000000000003'),
('d0e1f2a3-b4c5-4d6e-7f8a-9b0c1d2e3f4a', 'Marcus Thompson',   'marcus.thompson@sportscenter.com', 'mthompson', 'profiles/marcus_thompson.jpg', 'en', '00000000-0000-0000-0000-000000000004'),
('e1f2a3b4-c5d6-4e7f-8a9b-0c1d2e3f4a5b', 'Sofia Rossi',       'sofia.rossi@artsstudio.com',       'srossi',    'profiles/sofia_rossi.jpg',     'it', '00000000-0000-0000-0000-000000000005'),
('f2a3b4c5-d6e7-4f8a-9b0c-1d2e3f4a5b6c', 'David Kim',         'david.kim@musicschool.com',        'dkim',      'profiles/david_kim.jpg',       'ko', '00000000-0000-0000-0000-000000000006'),
-- Boston org managers
('a8b9c0d1-e2f3-4a4b-5c6d-7e8f9a0b1c2d', 'Jennifer Walsh',    'jennifer.walsh@bostonstemlab.com', 'jwalsh',    'profiles/jennifer_walsh.jpg',  'en', '00000000-0000-0000-0000-000000000007'),
('b9c0d1e2-f3a4-4b5c-6d7e-8f9a0b1c2d3e', 'Maria Fontaine',    'maria.fontaine@nedance.com',       'mfontaine', 'profiles/maria_fontaine.jpg',  'en', '00000000-0000-0000-0000-000000000008'),
('c0d1e2f3-a4b5-4c6d-7e8f-9a0b1c2d3e4f', 'Patrick O''Brien',  'patrick.obrien@fenwaychess.com',   'pobrien',   'profiles/patrick_obrien.jpg',  'en', '00000000-0000-0000-0000-000000000009'),
-- Additional Bangkok guardians
('d1e2f3a4-b5c6-4d7e-8f9a-0b1c2d3e4f5a', 'Nattaporn Chaiyasit',   'nattaporn.chaiyasit@email.com',  'nattapornc', 'profiles/nattaporn_chaiyasit.jpg', 'th', '00000000-0000-0000-0000-000000000010'),
('f3a4b5c6-d7e8-4f9a-0b1c-2d3e4f5a6b7c', 'Ananya Krishnamurthy',  'ananya.krishnamurthy@email.com', 'ananyak',    'profiles/ananya_krishnamurthy.jpg','en', '00000000-0000-0000-0000-000000000012'),
('b5c6d7e8-f9a0-4b1c-2d3e-4f5a6b7c8d9e', 'Siriporn Wattanabe',    'siriporn.wattanabe@email.com',   'siripornw',  'profiles/siriporn_wattanabe.jpg',  'th', '00000000-0000-0000-0000-000000000014'),
('c6d7e8f9-a0b1-4c2d-3e4f-5a6b7c8d9e0f', 'Mei-Ling Huang',        'meiling.huang@email.com',        'meilingh',   'profiles/meiling_huang.jpg',       'zh', '00000000-0000-0000-0000-000000000015'),
('e8f9a0b1-c2d3-4e4f-5a6b-7c8d9e0f1a2b', 'Arjun Sharma',          'arjun.sharma@email.com',         'arjuns',     'profiles/arjun_sharma.jpg',        'en', '00000000-0000-0000-0000-000000000017'),
('a0b1c2d3-e4f5-4a6b-7c8d-9e0f1a2b3c4d', 'Somchai Thongprasert',  'somchai.thongprasert@email.com', 'somchait',   'profiles/somchai_thongprasert.jpg','th', '00000000-0000-0000-0000-000000000019'),
-- Additional Boston guardians
('e2f3a4b5-c6d7-4e8f-9a0b-1c2d3e4f5a6b', 'James O''Connor',       'james.oconnor@email.com',        'joconnor',   'profiles/james_oconnor.jpg',       'en', '00000000-0000-0000-0000-000000000011'),
('a4b5c6d7-e8f9-4a0b-1c2d-3e4f5a6b7c8d', 'Thomas Brennan',        'thomas.brennan@email.com',       'tbrennan',   'profiles/thomas_brennan.jpg',      'en', '00000000-0000-0000-0000-000000000013'),
('d7e8f9a0-b1c2-4d3e-4f5a-6b7c8d9e0f1a', 'Rachel Kim',            'rachel.kim@email.com',           'rachelk',    'profiles/rachel_kim.jpg',          'en', '00000000-0000-0000-0000-000000000016'),
('f9a0b1c2-d3e4-4f5a-6b7c-8d9e0f1a2b3c', 'Patricia Walsh',        'patricia.walsh@email.com',       'pwalsh',     'profiles/patricia_walsh.jpg',      'en', '00000000-0000-0000-0000-000000000018');
