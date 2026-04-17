-- ============================================
-- 13. NEW USER ACCOUNTS — Demo Guardians + New Org Staff
-- ============================================
-- 3 hero demo guardian accounts + 8 new organization staff members
INSERT INTO "user" (id, name, email, username, profile_picture_s3_key, language_preference, auth_id) VALUES

-- ── Demo Guardian Accounts ──────────────────────────────────────────────────
-- Jennifer Park: Boston mom, 2 kids (STEM + sports/art), very active
('aa000001-0000-0000-0000-000000000001', 'Jennifer Park',       'jennifer.park@gmail.com',                  'jenniferpark',  'profiles/jennifer_park.jpg',       'en', '00000000-0000-0000-0000-0000000000a1'),
-- Nattaya Srisuk: Bangkok mom, 2 kids (dance + martial arts), Thai-speaking
('aa000001-0000-0000-0000-000000000002', 'Nattaya Srisuk',      'nattaya.srisuk@gmail.com',                 'nattayath',     'profiles/nattaya_srisuk.jpg',      'th', '00000000-0000-0000-0000-0000000000a2'),
-- Marcus Webb: Boston dad, 1 kid (tech/art), power user with many reviews
('aa000001-0000-0000-0000-000000000003', 'Marcus Webb',         'marcus.webb@gmail.com',                    'marcuswebb',    'profiles/marcus_webb.jpg',         'en', '00000000-0000-0000-0000-0000000000a3'),

-- ── Boston Organization Staff ────────────────────────────────────────────────
('aa000001-0000-0000-0000-000000000004', 'Dr. Kevin Walsh',     'kevin.walsh@mitkidslab.org',               'kevinwalsh',    'profiles/kevin_walsh.jpg',         'en', '00000000-0000-0000-0000-0000000000a4'),
('aa000001-0000-0000-0000-000000000005', 'Tanya Rivers',        'tanya.rivers@bostonathleticacademy.com',   'tanyarivers',   'profiles/tanya_rivers.jpg',        'en', '00000000-0000-0000-0000-0000000000a5'),
('aa000001-0000-0000-0000-000000000006', 'Claire Ashford',      'claire.ashford@necyouth.org',              'claireasford',  'profiles/claire_ashford.jpg',      'en', '00000000-0000-0000-0000-0000000000a6'),
('aa000001-0000-0000-0000-000000000007', 'Maria Santos',        'maria.santos@bostonartcenter.org',         'mariasantos',   'profiles/maria_santos.jpg',        'en', '00000000-0000-0000-0000-0000000000a7'),
('aa000001-0000-0000-0000-000000000008', 'Jake Horton',         'jake.horton@codecreate.boston',            'jakehorton',    'profiles/jake_horton.jpg',         'en', '00000000-0000-0000-0000-0000000000a8'),

-- ── Bangkok Organization Staff ───────────────────────────────────────────────
('aa000001-0000-0000-0000-000000000009', 'Somchai Wattana',     'somchai@siammuaythai.th',                  'krusom',        'profiles/somchai_wattana.jpg',     'th', '00000000-0000-0000-0000-0000000000a9'),
('aa000001-0000-0000-0000-000000000010', 'Plernpit Saengthong', 'plernpit@bangkokballet.th',                'ajarnplern',    'profiles/plernpit_saengthong.jpg', 'th', '00000000-0000-0000-0000-0000000000aa'),
('aa000001-0000-0000-0000-000000000011', 'Thitiporn Kasem',     'thitiporn@geniusstem.th',                  'drkasem',       'profiles/thitiporn_kasem.jpg',     'th', '00000000-0000-0000-0000-0000000000ab');
