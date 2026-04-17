-- ============================================
-- 1. user (User Accounts)
-- ============================================
-- First, enable UUID extension if not already enabled
INSERT INTO "user" (id, name, email, username, profile_picture_s3_key, language_preference, auth_id) VALUES
('a1b2c3d4-e5f6-4a7b-8c9d-0e1f2a3b4c5d', 'Sarah Johnson', 'sarah.johnson@email.com', 'sarahj', 'profiles/sarah_johnson.jpg', 'en', '00000000-0000-0000-0000-00000000000a'),
('b2c3d4e5-f6a7-4b8c-9d0e-1f2a3b4c5d6e', 'Michael Chen', 'michael.chen@email.com', 'mchen', 'profiles/michael_chen.jpg', 'en', '00000000-0000-0000-0000-00000000000b'),
('c3d4e5f6-a7b8-4c9d-0e1f-2a3b4c5d6e7f', 'Priya Patel', 'priya.patel@email.com', 'priyap', 'profiles/priya_patel.jpg', 'en', '00000000-0000-0000-0000-00000000000c'),
('d4e5f6a7-b8c9-4d0e-1f2a-3b4c5d6e7f8a', 'Carlos Rodriguez', 'carlos.rodriguez@email.com', 'carlosr', NULL, 'es', '00000000-0000-0000-0000-00000000000d'),
('e5f6a7b8-c9d0-4e1f-2a3b-4c5d6e7f8a9b', 'Emma Thompson', 'emma.thompson@email.com', 'emmat', 'profiles/emma_thompson.jpg', 'en', '00000000-0000-0000-0000-00000000000e'),
('f6a7b8c9-d0e1-4f2a-3b4c-5d6e7f8a9b0c', 'Yuki Tanaka', 'yuki.tanaka@email.com', 'yukit', NULL, 'ja', '00000000-0000-0000-0000-00000000000f'),
('a7b8c9d0-e1f2-4a3b-4c5d-6e7f8a9b0c1d', 'Olivia Martinez', 'olivia.martinez@email.com', 'oliviam', 'profiles/olivia_martinez.jpg', 'es', '00000000-0000-0000-0000-000000000001'),
('b8c9d0e1-f2a3-4b4c-5d6e-7f8a9b0c1d2e', 'James Wilson', 'james.wilson@email.com', 'jamesw', NULL, 'en', '00000000-0000-0000-0000-000000000002'),
-- Organization managers
('c9d0e1f2-a3b4-4c5d-6e7f-8a9b0c1d2e3f', 'Dr. Amanda Lee', 'amanda.lee@scienceacademy.com', 'alee', 'profiles/amanda_lee.jpg', 'en', '00000000-0000-0000-0000-000000000003'),
('d0e1f2a3-b4c5-4d6e-7f8a-9b0c1d2e3f4a', 'Marcus Thompson', 'marcus.thompson@sportscenter.com', 'mthompson', 'profiles/marcus_thompson.jpg', 'en', '00000000-0000-0000-0000-000000000004'),
('e1f2a3b4-c5d6-4e7f-8a9b-0c1d2e3f4a5b', 'Sofia Rossi', 'sofia.rossi@artsstudio.com', 'srossi', NULL, 'it', '00000000-0000-0000-0000-000000000005'),
('f2a3b4c5-d6e7-4f8a-9b0c-1d2e3f4a5b6c', 'David Kim', 'david.kim@musicschool.com', 'dkim', 'profiles/david_kim.jpg', 'ko', '00000000-0000-0000-0000-000000000006');
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
-- ============================================
-- 2. GUARDIANS
-- ============================================
INSERT INTO guardian (id, user_id) VALUES
('11111111-1111-1111-1111-111111111111', 'a1b2c3d4-e5f6-4a7b-8c9d-0e1f2a3b4c5d'), -- Sarah Johnson
('22222222-2222-2222-2222-222222222222', 'b2c3d4e5-f6a7-4b8c-9d0e-1f2a3b4c5d6e'), -- Michael Chen
('33333333-3333-3333-3333-333333333333', 'c3d4e5f6-a7b8-4c9d-0e1f-2a3b4c5d6e7f'), -- Priya Patel
('44444444-4444-4444-4444-444444444444', 'd4e5f6a7-b8c9-4d0e-1f2a-3b4c5d6e7f8a'), -- Carlos Rodriguez
('55555555-5555-5555-5555-555555555555', 'e5f6a7b8-c9d0-4e1f-2a3b-4c5d6e7f8a9b'), -- Emma Thompson
('66666666-6666-6666-6666-666666666666', 'f6a7b8c9-d0e1-4f2a-3b4c-5d6e7f8a9b0c'), -- Yuki Tanaka
('77777777-7777-7777-7777-777777777777', 'a7b8c9d0-e1f2-4a3b-4c5d-6e7f8a9b0c1d'), -- Olivia Martinez
('88888888-8888-8888-8888-888888888888', 'b8c9d0e1-f2a3-4b4c-5d6e-7f8a9b0c1d2e'); -- James Wilson
-- ============================================
-- 4. SCHOOLS
-- ============================================
INSERT INTO school (id, name, location_id) VALUES
('20000000-0000-0000-0000-000000000001', 'Bangkok International School', '10000000-0000-0000-0000-000000000001'),
('20000000-0000-0000-0000-000000000002', 'Green Valley Academy', '10000000-0000-0000-0000-000000000002'),
('20000000-0000-0000-0000-000000000003', 'Riverside Elementary', '10000000-0000-0000-0000-000000000003');
-- ============================================
-- 15. NEW SCHOOLS — Boston + Extra Bangkok
-- ============================================
INSERT INTO school (id, name, location_id) VALUES
('dd000001-0000-0000-0000-000000000001', 'Mason-Rice Elementary School',          'b0000001-0000-0000-0000-000000000001'),
('dd000001-0000-0000-0000-000000000002', 'Boston Latin School',                   'b0000001-0000-0000-0000-000000000002'),
('dd000001-0000-0000-0000-000000000003', 'Cambridge Rindge and Latin School',     'b0000001-0000-0000-0000-000000000003'),
('dd000001-0000-0000-0000-000000000004', 'Satit Kasetsart International School',  'b0000001-0000-0000-0000-000000000012');
-- ============================================
-- 6. ORGANIZATIONS
-- ============================================
INSERT INTO organization (id, name, active, pfp_s3_key, location_id, about_en, about_th, links) VALUES
('40000000-0000-0000-0000-000000000001', 'Science Academy Bangkok', true, 'orgs/science_academy.jpg', '10000000-0000-0000-0000-000000000004',
 'Science Academy Bangkok offers hands-on STEM programs for children aged 5–15, fostering curiosity and critical thinking through experiments, robotics, and coding workshops.',
 'Science Academy Bangkok เปิดสอนหลักสูตร STEM แบบลงมือปฏิบัติสำหรับเด็กอายุ 5–15 ปี ส่งเสริมความอยากรู้อยากเห็นและการคิดเชิงวิพากษ์ผ่านการทดลอง หุ่นยนต์ และเวิร์กช็อปการเขียนโค้ด',
 '[{"href": "https://www.scienceacademybkk.com", "label": "Website"}, {"href": "https://www.facebook.com/scienceacademybkk", "label": "Facebook"}, {"href": "https://www.instagram.com/scienceacademybkk", "label": "Instagram"}]'::jsonb),
('40000000-0000-0000-0000-000000000002', 'Champions Sports Center', true, 'orgs/champions_sports.jpg', '10000000-0000-0000-0000-000000000005',
 'Champions Sports Center provides youth sports training in soccer, basketball, and swimming. Our certified coaches build skills, teamwork, and confidence in athletes of all levels.',
 'Champions Sports Center ให้การฝึกสอนกีฬาสำหรับเยาวชนในฟุตบอล บาสเกตบอล และว่ายน้ำ โค้ชที่ได้รับการรับรองของเราสร้างทักษะ การทำงานเป็นทีม และความมั่นใจให้กับนักกีฬาทุกระดับ',
 '[{"href": "https://www.championssportsbkk.com", "label": "Website"}, {"href": "https://www.instagram.com/championssportsbkk", "label": "Instagram"}]'::jsonb),
('40000000-0000-0000-0000-000000000003', 'Creative Arts Studio', true, 'orgs/creative_arts.jpg', '10000000-0000-0000-0000-000000000006',
 'Creative Arts Studio is a vibrant space for young artists to explore painting, sculpture, and mixed media. We inspire self-expression and creativity in children from age 4 and up.',
 'Creative Arts Studio เป็นพื้นที่สำหรับศิลปินรุ่นเยาว์ในการสำรวจการวาดภาพ ประติมากรรม และศิลปะผสม เราสร้างแรงบันดาลใจในการแสดงออกและความคิดสร้างสรรค์ให้เด็กตั้งแต่อายุ 4 ปีขึ้นไป',
 '[{"href": "https://www.creativeartstudio.th", "label": "Website"}, {"href": "https://www.facebook.com/creativeartstudiobkk", "label": "Facebook"}, {"href": "https://www.instagram.com/creativeartstudio_bkk", "label": "Instagram"}]'::jsonb),
('40000000-0000-0000-0000-000000000004', 'Harmony Music School', true, 'orgs/harmony_music_bkk.jpg', '10000000-0000-0000-0000-000000000007',
 'Harmony Music School offers individual and group lessons in piano, guitar, violin, and voice. Our experienced instructors guide students from beginner to advanced levels in a nurturing environment.',
 'Harmony Music School เปิดสอนบทเรียนเดี่ยวและกลุ่มสำหรับเปียโน กีตาร์ ไวโอลิน และการร้องเพลง ครูผู้สอนที่มีประสบการณ์ของเราแนะนำนักเรียนตั้งแต่ระดับเริ่มต้นจนถึงระดับสูงในสภาพแวดล้อมที่อบอุ่น',
 '[{"href": "https://www.harmonymusic.th", "label": "Website"}, {"href": "https://www.facebook.com/harmonymusicbkk", "label": "Facebook"}]'::jsonb),
('40000000-0000-0000-0000-000000000005', 'Tech Kids Workshop', true, 'orgs/tech_kids.jpg', '10000000-0000-0000-0000-000000000008',
 'Tech Kids Workshop teaches children programming, game design, and electronics through project-based learning. We prepare the next generation of innovators with practical technology skills.',
 'Tech Kids Workshop สอนการเขียนโปรแกรม การออกแบบเกม และอิเล็กทรอนิกส์ให้เด็กผ่านการเรียนรู้แบบโครงงาน เราเตรียมนวัตกรรุ่นถัดไปด้วยทักษะเทคโนโลยีที่ใช้งานได้จริง',
 '[{"href": "https://www.techkidsworkshop.com", "label": "Website"}, {"href": "https://www.instagram.com/techkidsworkshop", "label": "Instagram"}, {"href": "https://www.youtube.com/@techkidsworkshop", "label": "YouTube"}]'::jsonb),
('40000000-0000-0000-0000-000000000006', 'Language Learning Center', false, 'orgs/language_learning_center.jpg', '10000000-0000-0000-0000-000000000009',
 'Language Learning Center offers immersive language programs in English, Mandarin, and Japanese for children and teens. Our communicative approach builds fluency and cultural understanding.',
 'Language Learning Center เปิดสอนหลักสูตรภาษาแบบอิมเมอร์สิฟในภาษาอังกฤษ จีนกลาง และญี่ปุ่น สำหรับเด็กและวัยรุ่น แนวทางการสื่อสารของเราสร้างความคล่องแคล่วและความเข้าใจทางวัฒนธรรม',
 '[{"href": "https://www.languagelearningcenter.th", "label": "Website"}, {"href": "https://www.facebook.com/languagelearningcenterbkk", "label": "Facebook"}]'::jsonb);
-- ============================================
-- 17. NEW ORGANIZATIONS — 5 Boston + 3 Bangkok
-- ============================================
INSERT INTO organization (id, name, active, pfp_s3_key, location_id, about_en, about_th, links) VALUES

-- ── Boston Organizations ─────────────────────────────────────────────────────
('ee000001-0000-0000-0000-000000000001', 'MIT Kids Lab', true, 'orgs/mit_kids_lab.jpg', 'b0000001-0000-0000-0000-000000000004',
 'MIT Kids Lab brings university-level STEM education to young learners aged 8–14. Located in Kendall Square, our programs are designed by MIT researchers and educators to spark curiosity, build engineering instincts, and give kids hands-on experience with real technology.',
 'MIT Kids Lab นำการศึกษา STEM ระดับมหาวิทยาลัยมาสู่เยาวชนอายุ 8–14 ปี ตั้งอยู่ใน Kendall Square โปรแกรมของเราออกแบบโดยนักวิจัยและนักการศึกษาจาก MIT เพื่อจุดประกายความอยากรู้อยากเห็น สร้างสัญชาตญาณด้านวิศวกรรม และให้เด็กๆ มีประสบการณ์ตรงกับเทคโนโลยีจริง',
 '[{"href": "https://www.mitkidslab.org", "label": "Website"}, {"href": "https://www.instagram.com/mitkidslab", "label": "Instagram"}, {"href": "https://www.youtube.com/@mitkidslab", "label": "YouTube"}]'::jsonb),

('ee000001-0000-0000-0000-000000000002', 'Boston Athletic Academy', true, 'orgs/boston_athletic_academy.jpg', 'b0000001-0000-0000-0000-000000000005',
 'Boston Athletic Academy offers premier youth sports training in the heart of Fenway. Our certified coaches develop fundamentals in soccer, basketball, and swimming while emphasizing teamwork, sportsmanship, and a love of movement that lasts a lifetime.',
 'Boston Athletic Academy มอบการฝึกกีฬาสำหรับเยาวชนระดับเยี่ยมในย่าน Fenway โค้ชที่ได้รับการรับรองของเราพัฒนาพื้นฐานในฟุตบอล บาสเกตบอล และว่ายน้ำ ขณะที่เน้นการทำงานเป็นทีม น้ำใจนักกีฬา และความรักในการเคลื่อนไหวที่ยั่งยืนตลอดชีวิต',
 '[{"href": "https://www.bostonathleticacademy.com", "label": "Website"}, {"href": "https://www.instagram.com/bostonathleticacademy", "label": "Instagram"}, {"href": "https://www.facebook.com/bostonathleticacademy", "label": "Facebook"}]'::jsonb),

('ee000001-0000-0000-0000-000000000003', 'NEC Youth Programs', true, 'orgs/nec_youth.jpg', 'b0000001-0000-0000-0000-000000000006',
 'New England Conservatory Youth Programs offers world-class music education for children and teens on historic Huntington Avenue. From beginner piano to youth orchestra, our faculty are professional musicians who nurture each student''s unique musical voice.',
 'New England Conservatory Youth Programs มอบการศึกษาดนตรีระดับโลกสำหรับเด็กและวัยรุ่นบนถนน Huntington Ave อันเป็นประวัติศาสตร์ ตั้งแต่เปียโนสำหรับผู้เริ่มต้นจนถึงวงออเคสตราเยาวชน คณาจารย์ของเราเป็นนักดนตรีมืออาชีพที่บ่มเพาะเสียงดนตรีเฉพาะตัวของนักเรียนแต่ละคน',
 '[{"href": "https://www.necmusic.edu/youth", "label": "Website"}, {"href": "https://www.instagram.com/necmusic", "label": "Instagram"}]'::jsonb),

('ee000001-0000-0000-0000-000000000004', 'Boston Art Center Kids', true, 'orgs/boston_art_center.jpg', 'b0000001-0000-0000-0000-000000000008',
 'Boston Art Center Kids is Brookline''s beloved studio arts program for children ages 5–16. We offer watercolor, printmaking, illustration, and mixed media workshops taught by working artists who believe that every child is a creative thinker.',
 'Boston Art Center Kids คือโปรแกรมศิลปะสตูดิโอสำหรับเด็กอายุ 5–16 ปี อันเป็นที่รักของ Brookline เรามีเวิร์กช็อปสีน้ำ การพิมพ์ ภาพประกอบ และสื่อผสม สอนโดยศิลปินที่ทำงานจริงซึ่งเชื่อว่าเด็กทุกคนคือนักคิดเชิงสร้างสรรค์',
 '[{"href": "https://www.bostonartcenter.org", "label": "Website"}, {"href": "https://www.instagram.com/bostonartcenterkids", "label": "Instagram"}, {"href": "https://www.facebook.com/bostonartcenterkids", "label": "Facebook"}]'::jsonb),

('ee000001-0000-0000-0000-000000000005', 'Code & Create Boston', true, 'orgs/code_create_boston.jpg', 'b0000001-0000-0000-0000-000000000007',
 'Code & Create Boston runs immersive coding and game design programs for kids aged 9–16 in the Seaport District. We teach web design, Unity game development, and electronics prototyping in a collaborative studio environment that mirrors real software teams.',
 'Code & Create Boston จัดโปรแกรมการเขียนโค้ดและการออกแบบเกมแบบอิมเมอร์สิฟสำหรับเด็กอายุ 9–16 ปี ในย่าน Seaport District เราสอนการออกแบบเว็บไซต์ การพัฒนาเกม Unity และการสร้างต้นแบบอิเล็กทรอนิกส์ในสภาพแวดล้อมสตูดิโอแบบร่วมมือที่สะท้อนทีมซอฟต์แวร์จริง',
 '[{"href": "https://www.codecreate.boston", "label": "Website"}, {"href": "https://www.instagram.com/codecreateboston", "label": "Instagram"}, {"href": "https://www.youtube.com/@codecreateboston", "label": "YouTube"}]'::jsonb),

-- ── Bangkok Organizations ─────────────────────────────────────────────────────
('ee000001-0000-0000-0000-000000000006', 'Siam Muay Thai Academy', true, 'orgs/siam_muay_thai.jpg', 'b0000001-0000-0000-0000-000000000009',
 'Siam Muay Thai Academy introduces children to Thailand''s national martial art in a safe, disciplined, and joyful environment. Our instructors emphasize respect, fitness, and self-confidence. Classes are structured for all experience levels, from beginners to junior competition.',
 'Siam Muay Thai Academy แนะนำเด็กๆ ให้รู้จักกับศิลปะการต่อสู้ประจำชาติของไทยในสภาพแวดล้อมที่ปลอดภัย มีระเบียบวินัย และสนุกสนาน ครูผู้สอนของเราเน้นความเคารพ ความฟิต และความมั่นใจในตนเอง คลาสถูกจัดโครงสร้างสำหรับทุกระดับประสบการณ์ ตั้งแต่ผู้เริ่มต้นจนถึงการแข่งขันระดับจูเนียร์',
 '[{"href": "https://www.siammuaythaiacademy.th", "label": "Website"}, {"href": "https://www.facebook.com/siammuaythaiakademy", "label": "Facebook"}, {"href": "https://www.instagram.com/siammuaythaiacademy", "label": "Instagram"}]'::jsonb),

('ee000001-0000-0000-0000-000000000007', 'Bangkok Ballet & Dance Academy', true, 'orgs/bangkok_ballet.jpg', 'b0000001-0000-0000-0000-000000000010',
 'Bangkok Ballet & Dance Academy provides classical ballet and contemporary dance training for children aged 4–17. Our faculty trained with leading companies across Europe and Asia, and we stage two full productions per year to give students real performance experience.',
 'Bangkok Ballet & Dance Academy ให้การฝึกบัลเล่ต์คลาสสิกและการเต้นรำร่วมสมัยสำหรับเด็กอายุ 4–17 ปี คณาจารย์ของเราผ่านการฝึกฝนกับบริษัทชั้นนำทั่วยุโรปและเอเชีย และเราจัดการแสดงเต็มรูปแบบสองครั้งต่อปีเพื่อให้นักเรียนมีประสบการณ์การแสดงจริง',
 '[{"href": "https://www.bangkokballet.th", "label": "Website"}, {"href": "https://www.instagram.com/bangkokballetacademy", "label": "Instagram"}, {"href": "https://www.facebook.com/bangkokballetdanceacademy", "label": "Facebook"}]'::jsonb),

('ee000001-0000-0000-0000-000000000008', 'Geniuses STEM Thailand', true, 'orgs/geniuses_stem.jpg', 'b0000001-0000-0000-0000-000000000011',
 'Geniuses STEM Thailand is Bangkok''s most dynamic STEM enrichment center for ages 7–15. We run innovation camps, drone programming courses, and engineering challenges that connect classroom learning to real-world problem solving. All programs are available in Thai and English.',
 'Geniuses STEM Thailand คือศูนย์เสริมสร้างความรู้ STEM ที่ไดนามิกที่สุดของกรุงเทพฯ สำหรับอายุ 7–15 ปี เราจัดค่ายนวัตกรรม หลักสูตรการเขียนโปรแกรมโดรน และความท้าทายด้านวิศวกรรมที่เชื่อมโยงการเรียนรู้ในห้องเรียนกับการแก้ปัญหาในโลกจริง โปรแกรมทั้งหมดมีให้ในภาษาไทยและอังกฤษ',
 '[{"href": "https://www.geniusstem.th", "label": "Website"}, {"href": "https://www.instagram.com/geniusstemthailand", "label": "Instagram"}, {"href": "https://www.youtube.com/@geniusstemth", "label": "YouTube"}]'::jsonb);
-- ============================================
-- 5. CHILDREN
-- ============================================
INSERT INTO child (id, name, school_id, birth_month, birth_year, interests, guardian_id) VALUES
('30000000-0000-0000-0000-000000000001', 'Emily Johnson', '20000000-0000-0000-0000-000000000001', 3, 2016, ARRAY['science','technology','math']::category[], '11111111-1111-1111-1111-111111111111'),
('30000000-0000-0000-0000-000000000002', 'Alex Johnson', '20000000-0000-0000-0000-000000000001', 7, 2018, ARRAY['sports','music']::category[], '11111111-1111-1111-1111-111111111111'),
('30000000-0000-0000-0000-000000000003', 'Sophie Chen', '20000000-0000-0000-0000-000000000002', 11, 2015, ARRAY['art','language','music']::category[], '22222222-2222-2222-2222-222222222222'),
('30000000-0000-0000-0000-000000000004', 'Aiden Patel', '20000000-0000-0000-0000-000000000001', 5, 2017, ARRAY['science','sports','technology']::category[], '33333333-3333-3333-3333-333333333333'),
('30000000-0000-0000-0000-000000000005', 'Maya Patel', '20000000-0000-0000-0000-000000000001', 9, 2019, ARRAY['art','music']::category[], '33333333-3333-3333-3333-333333333333'),
('30000000-0000-0000-0000-000000000006', 'Lucas Rodriguez', '20000000-0000-0000-0000-000000000002', 1, 2016, ARRAY['sports','technology']::category[], '44444444-4444-4444-4444-444444444444'),
('30000000-0000-0000-0000-000000000007', 'Isabella Thompson', '20000000-0000-0000-0000-000000000003', 8, 2017, ARRAY['language','art','other']::category[], '55555555-5555-5555-5555-555555555555'),
('30000000-0000-0000-0000-000000000008', 'Ethan Thompson', '20000000-0000-0000-0000-000000000003', 12, 2018, ARRAY['math','science']::category[], '55555555-5555-5555-5555-555555555555'),
('30000000-0000-0000-0000-000000000009', 'Hana Tanaka', '20000000-0000-0000-0000-000000000002', 4, 2016, ARRAY['music','language','art']::category[], '66666666-6666-6666-6666-666666666666'),
('30000000-0000-0000-0000-00000000000a', 'Noah Martinez', '20000000-0000-0000-0000-000000000001', 6, 2017, ARRAY['sports','science']::category[], '77777777-7777-7777-7777-777777777777'),
('30000000-0000-0000-0000-00000000000b', 'Liam Wilson', '20000000-0000-0000-0000-000000000003', 10, 2015, ARRAY['technology','math','science']::category[], '88888888-8888-8888-8888-888888888888'),
('30000000-0000-0000-0000-00000000000c', 'Ava Wilson', '20000000-0000-0000-0000-000000000003', 2, 2019, ARRAY['music','art']::category[], '88888888-8888-8888-8888-888888888888');
-- ============================================
-- 16. NEW GUARDIANS + EMERGENCY CONTACTS — 3 Demo Accounts
-- ============================================
INSERT INTO guardian (id, user_id) VALUES
('cc000001-0000-0000-0000-000000000001', 'aa000001-0000-0000-0000-000000000001'), -- Jennifer Park
('cc000001-0000-0000-0000-000000000002', 'aa000001-0000-0000-0000-000000000002'), -- Nattaya Srisuk
('cc000001-0000-0000-0000-000000000003', 'aa000001-0000-0000-0000-000000000003'); -- Marcus Webb

INSERT INTO emergency_contacts (id, name, guardian_id, phone_number) VALUES
('ec100001-0000-0000-0000-000000000001', 'Daniel Park',        'cc000001-0000-0000-0000-000000000001', '+16175550201'), -- Jennifer's husband
('ec100001-0000-0000-0000-000000000002', 'Somkid Srisuk',      'cc000001-0000-0000-0000-000000000002', '+66812345678'), -- Nattaya's husband
('ec100001-0000-0000-0000-000000000003', 'Angela Webb',        'cc000001-0000-0000-0000-000000000003', '+16175550203'); -- Marcus's wife
-- ============================================
-- 18. NEW CHILDREN — Demo Guardian Kids
-- ============================================
INSERT INTO child (id, name, school_id, birth_month, birth_year, interests, guardian_id) VALUES

-- ── Jennifer Park's children ─────────────────────────────────────────────────
-- Lily Park: 8 years old, science/tech/math kid in Newton
('ad000001-0000-0000-0000-000000000001', 'Lily Park',  'dd000001-0000-0000-0000-000000000001', 6,  2017, ARRAY['science','technology','math','art']::category[],      'cc000001-0000-0000-0000-000000000001'),
-- Max Park: 5 years old, active kid who loves sports and making things
('ad000001-0000-0000-0000-000000000002', 'Max Park',   'dd000001-0000-0000-0000-000000000001', 3,  2020, ARRAY['sports','art']::category[],                           'cc000001-0000-0000-0000-000000000001'),

-- ── Nattaya Srisuk's children ────────────────────────────────────────────────
-- Ploy Srisuk: 10 years old, passionate about dance and languages
('ad000001-0000-0000-0000-000000000003', 'Ploy Srisuk', 'dd000001-0000-0000-0000-000000000004', 8,  2015, ARRAY['music','art','language','science']::category[],       'cc000001-0000-0000-0000-000000000002'),
-- Ton Srisuk: 7 years old, muay thai fan, loves technology
('ad000001-0000-0000-0000-000000000004', 'Ton Srisuk',  '20000000-0000-0000-0000-000000000001', 1,  2018, ARRAY['sports','technology']::category[],                    'cc000001-0000-0000-0000-000000000002'),

-- ── Marcus Webb's children ───────────────────────────────────────────────────
-- Zoe Webb: 11 years old, multi-talented tech + art kid at Boston Latin
('ad000001-0000-0000-0000-000000000005', 'Zoe Webb',   'dd000001-0000-0000-0000-000000000002', 11, 2014, ARRAY['technology','math','science','art']::category[],       'cc000001-0000-0000-0000-000000000003');
-- ============================================
-- 7. MANAGERS
-- ============================================
INSERT INTO manager (id, user_id, organization_id, role) VALUES
('50000000-0000-0000-0000-000000000001', 'c9d0e1f2-a3b4-4c5d-6e7f-8a9b0c1d2e3f', '40000000-0000-0000-0000-000000000001', 'Director'),
('50000000-0000-0000-0000-000000000002', 'd0e1f2a3-b4c5-4d6e-7f8a-9b0c1d2e3f4a', '40000000-0000-0000-0000-000000000002', 'Head Coach'),
('50000000-0000-0000-0000-000000000003', 'e1f2a3b4-c5d6-4e7f-8a9b-0c1d2e3f4a5b', '40000000-0000-0000-0000-000000000003', 'Art Director'),
('50000000-0000-0000-0000-000000000004', 'f2a3b4c5-d6e7-4f8a-9b0c-1d2e3f4a5b6c', '40000000-0000-0000-0000-000000000004', 'Music Director'),
('50000000-0000-0000-0000-000000000005', 'c9d0e1f2-a3b4-4c5d-6e7f-8a9b0c1d2e3f', '40000000-0000-0000-0000-000000000005', 'Tech Instructor'),
('50000000-0000-0000-0000-000000000006', 'd0e1f2a3-b4c5-4d6e-7f8a-9b0c1d2e3f4a', '40000000-0000-0000-0000-000000000002', 'Assistant Coach');
-- ============================================
-- 19. NEW MANAGERS — One per new organization
-- ============================================
INSERT INTO manager (id, user_id, organization_id, role) VALUES
('ff000001-0000-0000-0000-000000000001', 'aa000001-0000-0000-0000-000000000004', 'ee000001-0000-0000-0000-000000000001', 'Director'),         -- Dr. Kevin Walsh @ MIT Kids Lab
('ff000001-0000-0000-0000-000000000002', 'aa000001-0000-0000-0000-000000000005', 'ee000001-0000-0000-0000-000000000002', 'Head Coach'),        -- Tanya Rivers @ Boston Athletic
('ff000001-0000-0000-0000-000000000003', 'aa000001-0000-0000-0000-000000000006', 'ee000001-0000-0000-0000-000000000003', 'Music Director'),    -- Claire Ashford @ NEC Youth
('ff000001-0000-0000-0000-000000000004', 'aa000001-0000-0000-0000-000000000007', 'ee000001-0000-0000-0000-000000000004', 'Art Director'),      -- Maria Santos @ Boston Art Center
('ff000001-0000-0000-0000-000000000005', 'aa000001-0000-0000-0000-000000000008', 'ee000001-0000-0000-0000-000000000005', 'Tech Instructor'),   -- Jake Horton @ Code & Create
('ff000001-0000-0000-0000-000000000006', 'aa000001-0000-0000-0000-000000000009', 'ee000001-0000-0000-0000-000000000006', 'Director'),          -- Somchai Wattana @ Siam Muay Thai
('ff000001-0000-0000-0000-000000000007', 'aa000001-0000-0000-0000-000000000010', 'ee000001-0000-0000-0000-000000000007', 'Art Director'),      -- Plernpit Saengthong @ Bangkok Ballet
('ff000001-0000-0000-0000-000000000008', 'aa000001-0000-0000-0000-000000000011', 'ee000001-0000-0000-0000-000000000008', 'Director');          -- Thitiporn Kasem @ Geniuses STEM
-- ============================================
-- 8. EVENTS
-- ============================================
INSERT INTO event (id, title_en, title_th, description_en, description_th, organization_id, age_range_min, age_range_max, category, header_image_s3_key) VALUES
-- Science Academy Events
('60000000-0000-0000-0000-000000000001', 'Junior Robotics Workshop', 'เวิร์คช็อปหุ่นยนต์สำหรับเด็ก', 'Learn the basics of robotics with hands-on LEGO Mindstorms projects. Build and program your own robots!', 'เรียนรู้พื้นฐานหุ่นยนต์ด้วยโครงการ LEGO Mindstorms สร้างและเขียนโปรแกรมหุ่นยนต์ของคุณเอง!', '40000000-0000-0000-0000-000000000001', 8, 12, ARRAY['science','technology']::category[], 'events/robotics_workshop.jpg'),
('60000000-0000-0000-0000-000000000002', 'Chemistry for Kids', 'เคมีสำหรับเด็ก', 'Exciting chemistry experiments that are safe and fun. Discover reactions, make slime, and learn about molecules!', 'การทดลองเคมีที่น่าตื่นเต้น ปลอดภัย และสนุก ค้นพบปฏิกิริยา ทำสไลม์ และเรียนรู้เกี่ยวกับโมเลกุล!', '40000000-0000-0000-0000-000000000001', 7, 10, ARRAY['science']::category[], 'events/chemistry_kids.jpg'),
('60000000-0000-0000-0000-000000000003', 'Astronomy Club', 'ชมรมดาราศาสตร์', 'Explore the wonders of space! Learn about planets, stars, and galaxies. Includes telescope observation sessions.', 'สำรวจความมหัศจรรย์ของอวกาศ! เรียนรู้เกี่ยวกับดาวเคราะห์ ดาวฤกษ์ และกาแล็กซี รวมถึงการสังเกตการณ์ด้วยกล้องโทรทรรศน์', '40000000-0000-0000-0000-000000000001', 9, 14, ARRAY['science']::category[], 'events/astronomy.jpg'),
-- Sports Center Events
('60000000-0000-0000-0000-000000000004', 'Soccer Skills Training', 'ฝึกทักษะฟุตบอล', 'Develop fundamental soccer skills including dribbling, passing, and teamwork in a fun environment.', 'พัฒนาทักษะฟุตบอลพื้นฐาน รวมถึงการเลี้ยงบอล การส่งบอล และการทำงานเป็นทีมในบรรยากาศที่สนุกสนาน', '40000000-0000-0000-0000-000000000002', 6, 12, ARRAY['sports']::category[], 'events/soccer_training.jpg'),
('60000000-0000-0000-0000-000000000005', 'Basketball Basics', 'พื้นฐานบาสเกตบอล', 'Learn basketball fundamentals: shooting, dribbling, defense, and game strategy.', 'เรียนรู้พื้นฐานบาสเกตบอล: การยิง การเลี้ยงบอล การป้องกัน และกลยุทธ์การเล่น', '40000000-0000-0000-0000-000000000002', 7, 13, ARRAY['sports']::category[], NULL),
('60000000-0000-0000-0000-000000000006', 'Swimming Lessons', 'บทเรียนว่ายน้ำ', 'Professional swimming instruction for beginners to intermediate levels. Focus on technique and water safety.', 'การสอนว่ายน้ำอย่างมืออาชีพสำหรับผู้เริ่มต้นถึงระดับกลาง เน้นเทคนิคและความปลอดภัยในน้ำ', '40000000-0000-0000-0000-000000000002', 5, 15, ARRAY['sports']::category[], 'events/swimming.jpg'),
-- Arts Studio Events
('60000000-0000-0000-0000-000000000007', 'Painting & Drawing Workshop', 'เวิร์คช็อปวาดภาพและระบายสี', 'Explore various art techniques including watercolor, acrylic, and sketching. All materials provided!', 'สำรวจเทคนิคศิลปะต่างๆ รวมถึงสีน้ำ อะคริลิค และการสเก็ตช์ อุปกรณ์ทั้งหมดมีให้!', '40000000-0000-0000-0000-000000000003', 6, 14, ARRAY['art']::category[], 'events/painting_workshop.jpg'),
('60000000-0000-0000-0000-000000000008', 'Pottery for Beginners', 'เครื่องปั้นดินเผาสำหรับผู้เริ่มต้น', 'Learn to work with clay! Create bowls, cups, and sculptures using hand-building and wheel techniques.', 'เรียนรู้การทำงานกับดินเหนียว! สร้างชาม ถ้วย และประติมากรรมโดยใช้เทคนิคการปั้นด้วยมือและแป้นหมุน', '40000000-0000-0000-0000-000000000003', 8, 15, ARRAY['art']::category[], NULL),
('60000000-0000-0000-0000-000000000009', 'Digital Art & Design', 'ศิลปะดิจิทัลและการออกแบบ', 'Introduction to digital illustration using tablets. Learn basic design principles and digital tools.', 'แนะนำการวาดภาพดิจิทัลโดยใช้แท็บเล็ต เรียนรู้หลักการออกแบบพื้นฐานและเครื่องมือดิจิทัล', '40000000-0000-0000-0000-000000000003', 10, 16, ARRAY['art','technology']::category[], 'events/digital_art.jpg'),
-- Music School Events
('60000000-0000-0000-0000-00000000000a', 'Piano for Beginners', 'เปียโนสำหรับผู้เริ่มต้น', 'Start your musical journey with piano! Learn to read music and play simple songs.', 'เริ่มต้นการเดินทางทางดนตรีของคุณด้วยเปียโน! เรียนอ่านโน้ตและเล่นเพลงง่ายๆ', '40000000-0000-0000-0000-000000000004', 6, 12, ARRAY['music']::category[], 'events/piano_lessons.jpg'),
('60000000-0000-0000-0000-00000000000b', 'Guitar Fundamentals', 'พื้นฐานกีตาร์', 'Learn basic chords, strumming patterns, and your first songs on acoustic guitar.', 'เรียนรู้คอร์ดพื้นฐาน รูปแบบการตีคอร์ด และเพลงแรกของคุณบนกีตาร์อะคูสติก', '40000000-0000-0000-0000-000000000004', 8, 15, ARRAY['music']::category[], NULL),
('60000000-0000-0000-0000-00000000000c', 'Kids Choir', 'คณะนักร้องประสานเสียงเด็ก', 'Join our fun choir! Learn harmony, vocal techniques, and perform in recitals.', 'ร่วมคณะนักร้องประสานเสียงสนุกๆ ของเรา! เรียนรู้การประสานเสียง เทคนิคการร้อง และแสดงในคอนเสิร์ต', '40000000-0000-0000-0000-000000000004', 7, 13, ARRAY['music']::category[], 'events/kids_choir.jpg'),
-- Tech Workshop Events
('60000000-0000-0000-0000-00000000000d', 'Coding with Scratch', 'เขียนโค้ดด้วย Scratch', 'Learn programming basics through Scratch! Create games and animations with visual coding blocks.', 'เรียนรู้พื้นฐานการเขียนโปรแกรมผ่าน Scratch! สร้างเกมและแอนิเมชันด้วยบล็อกโค้ดแบบภาพ', '40000000-0000-0000-0000-000000000005', 7, 11, ARRAY['technology']::category[], 'events/scratch_coding.jpg'),
('60000000-0000-0000-0000-00000000000e', 'Python for Kids', 'Python สำหรับเด็ก', 'Introduction to Python programming. Build simple programs and games while learning core concepts.', 'แนะนำการเขียนโปรแกรม Python สร้างโปรแกรมและเกมง่ายๆ ขณะเรียนรู้แนวคิดหลัก', '40000000-0000-0000-0000-000000000005', 10, 15, ARRAY['technology','math']::category[], NULL),
('60000000-0000-0000-0000-00000000000f', '3D Modeling Workshop', 'เวิร์คช็อปโมเดล 3 มิติ', 'Design 3D objects using Tinkercad! Learn basics of 3D design and prepare models for 3D printing.', 'ออกแบบวัตถุ 3 มิติโดยใช้ Tinkercad! เรียนรู้พื้นฐานการออกแบบ 3 มิติและเตรียมโมเดลสำหรับการพิมพ์ 3 มิติ', '40000000-0000-0000-0000-000000000005', 9, 14, ARRAY['technology','art']::category[], 'events/3d_modeling.jpg');
-- ============================================
-- 20. NEW EVENTS — 18 Boston + 6 Bangkok
-- ============================================
INSERT INTO event (id, title_en, title_th, description_en, description_th, organization_id, age_range_min, age_range_max, category, header_image_s3_key) VALUES

-- ── MIT Kids Lab (Boston) ────────────────────────────────────────────────────
('6b000000-0000-0000-0000-000000000001',
 'Young Engineers Workshop',
 'เวิร์กช็อปวิศวกรรมเยาวชน',
 'Teams of kids design and build functional machines using MIT-supplied components. Projects range from simple circuits to motorized contraptions. Participants leave with their creation and a solid grasp of engineering design thinking.',
 'ทีมเด็กๆ ออกแบบและสร้างเครื่องจักรที่ใช้งานได้โดยใช้ชิ้นส่วนที่ MIT จัดเตรียม โครงการมีตั้งแต่วงจรอย่างง่ายจนถึงอุปกรณ์ที่ขับเคลื่อนด้วยมอเตอร์ ผู้เข้าร่วมจะได้กลับบ้านพร้อมผลงานของตนเองและความเข้าใจที่มั่นคงในการคิดเชิงออกแบบทางวิศวกรรม',
 'ee000001-0000-0000-0000-000000000001', 8, 14, ARRAY['technology','science','robotics']::category[], 'events/young_engineers.jpg'),

('6b000000-0000-0000-0000-000000000002',
 'Science Explorers Lab',
 'ห้องปฏิบัติการนักสำรวจวิทยาศาสตร์',
 'Hands-on experiments in chemistry, biology, and physics. Each session tackles one big question — why do volcanoes erupt? how does DNA work? — through guided experimentation. Safety equipment provided; all materials are age-appropriate.',
 'การทดลองภาคปฏิบัติด้านเคมี ชีววิทยา และฟิสิกส์ แต่ละเซสชั่นจะตอบคำถามใหญ่หนึ่งข้อ — ทำไมภูเขาไฟระเบิด? DNA ทำงานอย่างไร? — ผ่านการทดลองที่มีการแนะนำ อุปกรณ์ความปลอดภัยมีให้ วัสดุทั้งหมดเหมาะสมกับช่วงอายุ',
 'ee000001-0000-0000-0000-000000000001', 7, 12, ARRAY['science','chemistry','biology']::category[], 'events/science_explorers.jpg'),

('6b000000-0000-0000-0000-000000000003',
 'App Inventor Academy',
 'สถาบัน App Inventor',
 'Build your first mobile app — no prior coding required! Students use MIT App Inventor to design interfaces, add logic, and test on real Android devices. By the end of the day each participant has a working app they can show friends and family.',
 'สร้างแอปมือถือแรกของคุณ — ไม่จำเป็นต้องมีประสบการณ์การเขียนโค้ดมาก่อน! นักเรียนใช้ MIT App Inventor เพื่อออกแบบอินเตอร์เฟซ เพิ่มตรรกะ และทดสอบบนอุปกรณ์ Android จริง ในตอนท้ายของวันผู้เข้าร่วมแต่ละคนจะมีแอปที่ใช้งานได้ซึ่งสามารถแสดงให้เพื่อนและครอบครัวได้',
 'ee000001-0000-0000-0000-000000000001', 10, 15, ARRAY['technology','coding']::category[], 'events/app_inventor.jpg'),

-- ── Boston Athletic Academy ──────────────────────────────────────────────────
('6b000000-0000-0000-0000-000000000004',
 'Youth Soccer Academy',
 'สถาบันฟุตบอลเยาวชน',
 'Professional-grade coaching in a welcoming environment. Sessions cover dribbling, passing, shooting, and small-sided games. Coaches are USSF-licensed and dedicated to building confident athletes regardless of starting skill level.',
 'การโค้ชระดับมืออาชีพในสภาพแวดล้อมที่เป็นมิตร เซสชั่นครอบคลุมการเลี้ยงบอล การส่งบอล การยิง และเกมขนาดเล็ก โค้ชมีใบอนุญาต USSF และมุ่งมั่นสร้างนักกีฬาที่มั่นใจโดยไม่คำนึงถึงระดับทักษะเริ่มต้น',
 'ee000001-0000-0000-0000-000000000002', 6, 13, ARRAY['sports','soccer']::category[], 'events/youth_soccer.jpg'),

('6b000000-0000-0000-0000-000000000005',
 'Basketball Training Camp',
 'ค่ายฝึกบาสเกตบอล',
 'A fast-paced, high-energy session covering ball handling, layups, jump shots, and defensive footwork. Ends with a 3-on-3 tournament. Suitable for beginners through intermediate players.',
 'เซสชั่นที่รวดเร็วและมีพลังงานสูงครอบคลุมการจับลูกบอล เลย์อัป จัมพ์ช็อต และการเคลื่อนเท้าในการป้องกัน จบด้วยการแข่งขัน 3 ต่อ 3 เหมาะสำหรับผู้เริ่มต้นจนถึงผู้เล่นระดับกลาง',
 'ee000001-0000-0000-0000-000000000002', 7, 14, ARRAY['sports','basketball']::category[], 'events/basketball_camp.jpg'),

('6b000000-0000-0000-0000-000000000006',
 'Swim Team Prep',
 'เตรียมทีมว่ายน้ำ',
 'Coach-led swim technique sessions in a 25-meter pool. Focuses on freestyle and backstroke fundamentals, flip turns, and race starts. A great feeder program for competitive swim teams at local schools.',
 'เซสชั่นเทคนิคว่ายน้ำที่นำโดยโค้ชในสระ 25 เมตร เน้นพื้นฐานฟรีสไตล์และท่าว่ายหงาย การพลิกตัว และการเริ่มต้นการแข่งขัน เป็นโปรแกรมป้อนที่ดีสำหรับทีมว่ายน้ำแข่งขันในโรงเรียนท้องถิ่น',
 'ee000001-0000-0000-0000-000000000002', 7, 14, ARRAY['sports','swimming']::category[], 'events/swim_team_prep.jpg'),

-- ── NEC Youth Programs ───────────────────────────────────────────────────────
('6b000000-0000-0000-0000-000000000007',
 'Piano Foundations',
 'รากฐานเปียโน',
 'A nurturing introduction to the piano for absolute beginners. Students learn posture, hand position, basic theory, and play simple melodies from day one. Upright practice pianos available for all students during the session.',
 'การแนะนำเปียโนที่อบอุ่นสำหรับผู้เริ่มต้นอย่างแท้จริง นักเรียนเรียนท่าทาง ตำแหน่งมือ ทฤษฎีพื้นฐาน และเล่นทำนองง่ายๆ ตั้งแต่วันแรก มีเปียโนแบบตั้งตรงสำหรับนักเรียนทุกคนในระหว่างเซสชั่น',
 'ee000001-0000-0000-0000-000000000003', 6, 13, ARRAY['music','instrumental music']::category[], 'events/piano_foundations.jpg'),

('6b000000-0000-0000-0000-000000000008',
 'Strings for Beginners',
 'เครื่องสายสำหรับผู้เริ่มต้น',
 'Introduction to violin and viola in a small-group format. Students learn to hold the bow, produce their first notes, and play simple folk tunes. Instruments available to borrow. A parent or guardian is welcome to observe.',
 'แนะนำไวโอลินและวิโอลาในรูปแบบกลุ่มเล็ก นักเรียนเรียนวิธีจับคันชัก ออกเสียงโน้ตแรก และเล่นเพลงพื้นบ้านง่ายๆ มีเครื่องดนตรีให้ยืม ผู้ปกครองยินดีต้อนรับให้สังเกตการณ์',
 'ee000001-0000-0000-0000-000000000003', 7, 14, ARRAY['music','instrumental music']::category[], 'events/strings_beginners.jpg'),

('6b000000-0000-0000-0000-000000000009',
 'Youth Orchestra Ensemble',
 'วงออเคสตราเยาวชน',
 'A full ensemble rehearsal experience for students with at least one year of their instrument. Students sight-read new pieces, work on blend and intonation, and perform a short concert for family at session end.',
 'ประสบการณ์การซ้อมวงออเคสตราเต็มรูปแบบสำหรับนักเรียนที่เล่นเครื่องดนตรีมาอย่างน้อยหนึ่งปี นักเรียนอ่านโน้ตครั้งแรก ฝึกการผสมผสานและการปรับเสียง และแสดงคอนเสิร์ตสั้นๆ ให้ครอบครัวชมเมื่อจบเซสชั่น',
 'ee000001-0000-0000-0000-000000000003', 10, 17, ARRAY['music','instrumental music']::category[], 'events/youth_orchestra.jpg'),

-- ── Boston Art Center Kids ───────────────────────────────────────────────────
('6b000000-0000-0000-0000-00000000000a',
 'Watercolor & Mixed Media',
 'สีน้ำและสื่อผสม',
 'Explore the luminous world of watercolor combined with collage, ink, and pastel. Students work on themed compositions — cityscapes, botanicals, abstract — and leave with a finished piece suitable for framing. All materials included.',
 'สำรวจโลกที่สว่างไสวของสีน้ำร่วมกับคอลลาจ หมึก และพาสเทล นักเรียนทำงานในองค์ประกอบตามธีม — ภาพเมือง พฤกษศาสตร์ นามธรรม — และออกไปพร้อมชิ้นงานสำเร็จที่เหมาะสำหรับใส่กรอบ รวมวัสดุทั้งหมด',
 'ee000001-0000-0000-0000-000000000004', 7, 15, ARRAY['art','painting']::category[], 'events/watercolor_workshop.jpg'),

('6b000000-0000-0000-0000-00000000000b',
 'Printmaking Lab',
 'ห้องปฏิบัติการพิมพ์',
 'Students learn linoleum block carving and monoprinting. Each participant designs their own relief image, carves the block, and pulls multiple prints. Learn about repetition, texture, and the satisfying surprise of lifting a fresh print.',
 'นักเรียนเรียนการแกะสลักบล็อกลิโนเลียมและโมโนพริ้นต์ ผู้เข้าร่วมแต่ละคนออกแบบภาพนูนของตนเอง แกะสลักบล็อก และดึงพิมพ์หลายแผ่น เรียนรู้เกี่ยวกับการซ้ำ พื้นผิว และความประหลาดใจที่น่าพึงพอใจของการยกพิมพ์ใหม่',
 'ee000001-0000-0000-0000-000000000004', 9, 15, ARRAY['art','crafts']::category[], 'events/printmaking_lab.jpg'),

('6b000000-0000-0000-0000-00000000000c',
 'Comic Art & Illustration',
 'ศิลปะการ์ตูนและภาพประกอบ',
 'Create your own comic strip from scratch: character design, paneling, inking, and lettering. Students study real comics for technique before building their own 4–8 panel story. Perfect for kids who love to draw and tell stories.',
 'สร้างการ์ตูนสตริปของคุณเองตั้งแต่ต้น: การออกแบบตัวละคร การจัดเฟรม การลงหมึก และการใส่ตัวอักษร นักเรียนศึกษาการ์ตูนจริงสำหรับเทคนิคก่อนสร้างเรื่องราว 4–8 เฟรมของตนเอง เหมาะสำหรับเด็กที่ชอบวาดและเล่าเรื่อง',
 'ee000001-0000-0000-0000-000000000004', 8, 15, ARRAY['art','drawing','creative writing']::category[], 'events/comic_illustration.jpg'),

-- ── Code & Create Boston ─────────────────────────────────────────────────────
('6b000000-0000-0000-0000-00000000000d',
 'Web Design Basics',
 'พื้นฐานการออกแบบเว็บ',
 'Build a personal website from zero to deployed in one session. Students write real HTML and CSS, learn about layout and typography, and publish their site using GitHub Pages. No experience needed — just curiosity.',
 'สร้างเว็บไซต์ส่วนตัวจากศูนย์จนเผยแพร่ได้ในเซสชั่นเดียว นักเรียนเขียน HTML และ CSS จริง เรียนรู้เกี่ยวกับเลย์เอาต์และการพิมพ์ และเผยแพร่ไซต์โดยใช้ GitHub Pages ไม่จำเป็นต้องมีประสบการณ์ เพียงแค่ความอยากรู้อยากเห็น',
 'ee000001-0000-0000-0000-000000000005', 11, 16, ARRAY['technology','coding']::category[], 'events/web_design.jpg'),

('6b000000-0000-0000-0000-00000000000e',
 'Game Design with Unity',
 'การออกแบบเกมด้วย Unity',
 'Design a playable 2D game using Unity and C#. Students learn the game loop, sprites, collision detection, and scoring. By session end each student exports a working game they can send to friends. Intermediate programmers preferred.',
 'ออกแบบเกม 2D ที่เล่นได้โดยใช้ Unity และ C# นักเรียนเรียนรู้ game loop สไปรท์ การตรวจจับการชน และการนับคะแนน เมื่อสิ้นสุดเซสชั่นนักเรียนแต่ละคนจะส่งออกเกมที่ใช้งานได้ซึ่งสามารถส่งให้เพื่อนได้ ต้องการผู้เขียนโปรแกรมระดับกลาง',
 'ee000001-0000-0000-0000-000000000005', 11, 16, ARRAY['technology','coding']::category[], 'events/game_design_unity.jpg'),

('6b000000-0000-0000-0000-00000000000f',
 'Robotics & Electronics Lab',
 'ห้องปฏิบัติการหุ่นยนต์และอิเล็กทรอนิกส์',
 'Wire up real circuits, program Arduino microcontrollers, and assemble a robot that responds to light and distance sensors. Students work in pairs on structured challenges and take home a kit with components to continue experimenting.',
 'ต่อวงจรจริง เขียนโปรแกรมไมโครคอนโทรลเลอร์ Arduino และประกอบหุ่นยนต์ที่ตอบสนองต่อเซ็นเซอร์แสงและระยะทาง นักเรียนทำงานเป็นคู่บนความท้าทายที่มีโครงสร้างและกลับบ้านพร้อมชุดอุปกรณ์เพื่อทดลองต่อ',
 'ee000001-0000-0000-0000-000000000005', 9, 15, ARRAY['technology','robotics','engineering']::category[], 'events/robotics_electronics.jpg'),

-- ── Siam Muay Thai Academy (Bangkok) ────────────────────────────────────────
('6b000000-0000-0000-0000-000000000010',
 'Kids Muay Thai — Beginner',
 'มวยไทยเด็ก — ระดับเริ่มต้น',
 'A fun, safe introduction to Muay Thai for children with no prior martial arts experience. Sessions cover the Wai Kru ceremony, stance, basic strikes, and pad work. Heavy emphasis on respect, discipline, and having fun. Gloves and pads provided.',
 'การแนะนำมวยไทยที่สนุกและปลอดภัยสำหรับเด็กที่ไม่มีประสบการณ์ศิลปะการต่อสู้มาก่อน เซสชั่นครอบคลุมพิธีไหว้ครู ท่าทาง การตีพื้นฐาน และการตีแป้น เน้นความเคารพ วินัย และความสนุกสนาน มีถุงมือและแป้นให้',
 'ee000001-0000-0000-0000-000000000006', 6, 12, ARRAY['sports','martial arts','fitness']::category[], 'events/muay_thai_kids.jpg'),

('6b000000-0000-0000-0000-000000000011',
 'Junior Fighters Training',
 'ฝึกนักสู้จูเนียร์',
 'For kids with 3+ months of Muay Thai experience. Structured sparring drills, clinch work, combo sequences, and conditioning circuits. Coach Somchai has trained multiple national junior champions and brings that expertise to every session.',
 'สำหรับเด็กที่มีประสบการณ์มวยไทย 3 เดือนขึ้นไป การฝึกซ้อมการชกแบบมีโครงสร้าง งานคลิ้นช์ ลำดับคอมโบ และวงจรการปรับสภาพ โค้ชสมชายฝึกแชมป์จูเนียร์แห่งชาติหลายคนและนำความเชี่ยวชาญนั้นมาสู่ทุกเซสชั่น',
 'ee000001-0000-0000-0000-000000000006', 8, 15, ARRAY['sports','martial arts','fitness']::category[], 'events/muay_thai_junior.jpg'),

-- ── Bangkok Ballet & Dance Academy ───────────────────────────────────────────
('6b000000-0000-0000-0000-000000000012',
 'Classical Ballet for Kids',
 'บัลเล่ต์คลาสสิกสำหรับเด็ก',
 'A structured introduction to classical ballet technique: positions, barre work, center exercises, and a short combination across the floor. Faculty trained at the Royal Ballet School and Paris Opéra Ballet. Students should wear ballet shoes and comfortable attire.',
 'การแนะนำที่มีโครงสร้างสำหรับเทคนิคบัลเล่ต์คลาสสิก: ตำแหน่ง งานบาร์ แบบฝึกหัดกลาง และการรวมสั้นๆ ข้ามพื้น คณาจารย์ได้รับการฝึกฝนที่ Royal Ballet School และ Paris Opéra Ballet นักเรียนควรสวมรองเท้าบัลเล่ต์และเสื้อผ้าสบาย',
 'ee000001-0000-0000-0000-000000000007', 5, 14, ARRAY['dance','art']::category[], 'events/ballet_kids.jpg'),

('6b000000-0000-0000-0000-000000000013',
 'Contemporary Dance Workshop',
 'เวิร์กช็อปการเต้นร่วมสมัย',
 'Explore movement vocabulary beyond classical technique. Students work on floor work, improvisation, and structured group choreography inspired by Pina Bausch and Thai contemporary artists. No prior dance experience needed.',
 'สำรวจคำศัพท์การเคลื่อนไหวที่ไปไกลกว่าเทคนิคคลาสสิก นักเรียนทำงานด้วยงานพื้น การด้นสด และการออกแบบท่าเต้นกลุ่มที่มีโครงสร้างโดยได้แรงบันดาลใจจาก Pina Bausch และศิลปินร่วมสมัยชาวไทย ไม่จำเป็นต้องมีประสบการณ์การเต้นมาก่อน',
 'ee000001-0000-0000-0000-000000000007', 9, 17, ARRAY['dance','art']::category[], 'events/contemporary_dance.jpg'),

-- ── Geniuses STEM Thailand ───────────────────────────────────────────────────
('6b000000-0000-0000-0000-000000000014',
 'STEM Innovation Camp',
 'ค่ายนวัตกรรม STEM',
 'A full-day project-based camp where teams tackle a real engineering challenge: design a water filtration system, build a bridge from limited materials, or prototype a solar-powered device. Judges evaluate creativity, function, and presentation.',
 'ค่ายโครงงานตลอดวันที่ทีมเผชิญกับความท้าทายทางวิศวกรรมจริง: ออกแบบระบบกรองน้ำ สร้างสะพานจากวัสดุจำกัด หรือสร้างต้นแบบอุปกรณ์พลังงานแสงอาทิตย์ ผู้ตัดสินประเมินความคิดสร้างสรรค์ การทำงาน และการนำเสนอ',
 'ee000001-0000-0000-0000-000000000008', 8, 15, ARRAY['science','technology','engineering','robotics']::category[], 'events/stem_camp.jpg'),

('6b000000-0000-0000-0000-000000000015',
 'Drone Programming Workshop',
 'เวิร์กช็อปการเขียนโปรแกรมโดรน',
 'Students program DJI Tello drones using Python and Scratch. Learn about altitude sensors, coordinate systems, and flight path algorithms. Fly your own programmed routine in the indoor arena. Safety briefing and mandatory supervision throughout.',
 'นักเรียนเขียนโปรแกรมโดรน DJI Tello โดยใช้ Python และ Scratch เรียนรู้เกี่ยวกับเซ็นเซอร์ความสูง ระบบพิกัด และอัลกอริทึมเส้นทางการบิน บินตามรูทีนที่เขียนโปรแกรมเองในสนามในร่ม มีการบรรยายความปลอดภัยและการดูแลบังคับตลอดเวลา',
 'ee000001-0000-0000-0000-000000000008', 10, 16, ARRAY['technology','coding','robotics']::category[], 'events/drone_programming.jpg');
-- ============================================
-- 9. EVENT OCCURRENCES
-- ============================================
INSERT INTO event_occurrence (id, manager_id, event_id, start_time, end_time, max_attendees, language, curr_enrolled, price, currency) VALUES
-- Robotics Workshop occurrences
('70000000-0000-0000-0000-000000000001', '50000000-0000-0000-0000-000000000001', '60000000-0000-0000-0000-000000000001', '2028-05-15 09:00:00+07', '2028-05-15 11:00:00+07', 15, 'en', 8,  50000, 'thb'),
('70000000-0000-0000-0000-000000000002', '50000000-0000-0000-0000-000000000001', '60000000-0000-0000-0000-000000000001', '2026-05-22 09:00:00+07', '2026-05-22 11:00:00+07', 15, 'en', 5,  50000, 'thb'),
-- Chemistry for Kids
('70000000-0000-0000-0000-000000000003', '50000000-0000-0000-0000-000000000001', '60000000-0000-0000-0000-000000000002',  '2026-05-16 14:00:00+07', '2026-05-16 15:30:00+07', 12, 'en', 10, 40000, 'thb'),
('70000000-0000-0000-0000-000000000004', '50000000-0000-0000-0000-000000000001', '60000000-0000-0000-0000-000000000002', '2026-05-23 14:00:00+07', '2026-05-23 15:30:00+07', 12, 'en', 7,  40000, 'thb'),
-- Astronomy Club
('70000000-0000-0000-0000-000000000005', '50000000-0000-0000-0000-000000000001', '60000000-0000-0000-0000-000000000003', '2026-05-20 18:00:00+07', '2026-05-20 20:00:00+07', 20, 'en', 12, 35000, 'thb'),
-- Soccer Training
('70000000-0000-0000-0000-000000000006', '50000000-0000-0000-0000-000000000002', '60000000-0000-0000-0000-000000000004', '2026-05-17 16:00:00+07', '2026-05-17 17:30:00+07', 20, 'en', 15, 30000, 'thb'),
('70000000-0000-0000-0000-000000000007', '50000000-0000-0000-0000-000000000006', '60000000-0000-0000-0000-000000000004', '2026-05-24 16:00:00+07', '2026-05-24 17:30:00+07', 20, 'en', 11, 30000, 'thb'),
-- Basketball
('70000000-0000-0000-0000-000000000008', '50000000-0000-0000-0000-000000000002', '60000000-0000-0000-0000-000000000005', '2026-05-18 15:00:00+07', '2026-05-18 16:30:00+07', 16, 'en', 9,  30000, 'thb'),
-- Swimming
('70000000-0000-0000-0000-000000000009', '50000000-0000-0000-0000-000000000006', '60000000-0000-0000-0000-000000000006', '2026-05-19 10:00:00+07', '2026-05-19 11:00:00+07', 10, 'en', 8,  45000, 'thb'),
('70000000-0000-0000-0000-00000000000a', '50000000-0000-0000-0000-000000000006', '60000000-0000-0000-0000-000000000006', '2026-05-26 10:00:00+07', '2026-05-26 11:00:00+07', 10, 'en', 6,  45000, 'thb'),
-- Painting Workshop
('70000000-0000-0000-0000-00000000000b', '50000000-0000-0000-0000-000000000003', '60000000-0000-0000-0000-000000000007', '2026-05-21 13:00:00+07', '2026-05-21 15:00:00+07', 12, 'en', 10, 25000, 'thb'),
-- Pottery
('70000000-0000-0000-0000-00000000000c', '50000000-0000-0000-0000-000000000003', '60000000-0000-0000-0000-000000000008', '2026-05-22 14:00:00+07', '2026-05-22 16:00:00+07', 10, 'en', 6,  25000, 'thb'),
-- Digital Art
('70000000-0000-0000-0000-00000000000d', '50000000-0000-0000-0000-000000000003', '60000000-0000-0000-0000-000000000009', '2026-05-25 15:00:00+07', '2026-05-25 17:00:00+07', 14, 'en', 7,  25000, 'thb'),
-- Piano
('70000000-0000-0000-0000-00000000000e', '50000000-0000-0000-0000-000000000004', '60000000-0000-0000-0000-00000000000a', '2026-05-16 10:00:00+07', '2026-05-16 11:00:00+07', 8,  'en', 6,  60000, 'thb'),
('70000000-0000-0000-0000-00000000000f', '50000000-0000-0000-0000-000000000004', '60000000-0000-0000-0000-00000000000a', '2026-05-23 10:00:00+07', '2026-05-23 11:00:00+07', 8,  'en', 4,  60000, 'thb'),
-- Guitar
('70000000-0000-0000-0000-000000000010', '50000000-0000-0000-0000-000000000004', '60000000-0000-0000-0000-00000000000b', '2026-05-17 14:00:00+07', '2026-05-17 15:00:00+07', 10, 'en', 7,  55000, 'thb'),
-- Choir
('70000000-0000-0000-0000-000000000011', '50000000-0000-0000-0000-000000000004', '60000000-0000-0000-0000-00000000000c', '2026-05-20 16:00:00+07', '2026-05-20 17:30:00+07', 25, 'en', 18, 20000, 'thb'),
-- Scratch Coding
('70000000-0000-0000-0000-000000000012', '50000000-0000-0000-0000-000000000005', '60000000-0000-0000-0000-00000000000d', '2026-06-18 13:00:00+07', '2026-06-18 15:00:00+07', 15, 'en', 11, 45000, 'thb'),
('70000000-0000-0000-0000-000000000013', '50000000-0000-0000-0000-000000000005', '60000000-0000-0000-0000-00000000000d', '2026-06-25 13:00:00+07', '2026-06-25 15:00:00+07', 15, 'en', 8,  45000, 'thb'),
-- Python
('70000000-0000-0000-0000-000000000014', '50000000-0000-0000-0000-000000000005', '60000000-0000-0000-0000-00000000000e', '2026-08-21 15:00:00+07', '2026-08-21 17:00:00+07', 12, 'en', 9,  45000, 'thb'),
-- 3D Modeling
('70000000-0000-0000-0000-000000000015', '50000000-0000-0000-0000-000000000005', '60000000-0000-0000-0000-00000000000f', '2027-05-24 14:00:00+07', '2027-05-24 16:00:00+07', 12, 'en', 5,  45000, 'thb');-- ============================================
-- 21. NEW EVENT OCCURRENCES
-- Past occurrences (Jan–Apr 2026): demo guardians can have completed registrations + reviews
-- Upcoming occurrences (May–Aug 2026): demo guardians have active registrations
--
-- Boston events use Eastern time: -05 (EST, winter) / -04 (EDT, spring+)
-- Bangkok events use +07
--
-- Columns: id, manager_id, event_id, start_time, end_time, max_attendees, language, curr_enrolled, price, currency
-- ============================================

INSERT INTO event_occurrence (id, manager_id, event_id, start_time, end_time, max_attendees, language, curr_enrolled, price, currency) VALUES

-- ════════════════════════════════════════════════════════════════════════════
-- PAST OCCURRENCES (completed — basis for demo reviews)
-- ════════════════════════════════════════════════════════════════════════════

-- MIT Kids Lab
('7b000000-0000-0000-0000-000000000001', 'ff000001-0000-0000-0000-000000000001', '6b000000-0000-0000-0000-000000000001', '2026-02-14 10:00:00-05', '2026-02-14 13:00:00-05', 16, 'en', 14, 7500,  'usd'), -- Young Engineers (past) ← Lily, Zoe
('7b000000-0000-0000-0000-000000000002', 'ff000001-0000-0000-0000-000000000001', '6b000000-0000-0000-0000-000000000002', '2026-01-24 10:00:00-05', '2026-01-24 12:30:00-05', 14, 'en', 13, 7500,  'usd'), -- Science Explorers (past) ← Lily, Zoe
('7b000000-0000-0000-0000-000000000003', 'ff000001-0000-0000-0000-000000000001', '6b000000-0000-0000-0000-000000000003', '2026-03-07 10:00:00-05', '2026-03-07 14:00:00-05', 12, 'en', 10, 7500,  'usd'), -- App Inventor (past) ← Lily

-- Boston Athletic Academy
('7b000000-0000-0000-0000-000000000004', 'ff000001-0000-0000-0000-000000000002', '6b000000-0000-0000-0000-000000000004', '2026-01-10 09:00:00-05', '2026-01-10 10:30:00-05', 20, 'en', 19, 5500,  'usd'), -- Youth Soccer (past) ← Max
('7b000000-0000-0000-0000-000000000005', 'ff000001-0000-0000-0000-000000000002', '6b000000-0000-0000-0000-000000000005', '2026-02-07 09:00:00-05', '2026-02-07 10:30:00-05', 18, 'en', 15, 5500,  'usd'), -- Basketball (past) ← Max
('7b000000-0000-0000-0000-000000000006', 'ff000001-0000-0000-0000-000000000002', '6b000000-0000-0000-0000-000000000006', '2026-03-14 09:00:00-05', '2026-03-14 10:30:00-05', 10, 'en', 9,  5500,  'usd'), -- Swim Team (past)

-- NEC Youth
('7b000000-0000-0000-0000-000000000007', 'ff000001-0000-0000-0000-000000000003', '6b000000-0000-0000-0000-000000000007', '2026-02-07 10:00:00-05', '2026-02-07 11:30:00-05', 8,  'en', 7,  8000,  'usd'), -- Piano Foundations (past)
('7b000000-0000-0000-0000-000000000008', 'ff000001-0000-0000-0000-000000000003', '6b000000-0000-0000-0000-000000000008', '2026-03-07 10:00:00-05', '2026-03-07 11:30:00-05', 8,  'en', 6,  8000,  'usd'), -- Strings (past)
('7b000000-0000-0000-0000-000000000009', 'ff000001-0000-0000-0000-000000000003', '6b000000-0000-0000-0000-000000000009', '2026-03-21 13:00:00-05', '2026-03-21 15:00:00-05', 30, 'en', 22, 6000,  'usd'), -- Youth Orchestra (past)

-- Boston Art Center
('7b000000-0000-0000-0000-00000000000a', 'ff000001-0000-0000-0000-000000000004', '6b000000-0000-0000-0000-00000000000a', '2026-01-17 13:00:00-05', '2026-01-17 15:30:00-05', 14, 'en', 12, 4500,  'usd'), -- Watercolor (past) ← Max, Zoe
('7b000000-0000-0000-0000-00000000000b', 'ff000001-0000-0000-0000-000000000004', '6b000000-0000-0000-0000-00000000000b', '2026-02-28 13:00:00-05', '2026-02-28 15:30:00-05', 12, 'en', 10, 4500,  'usd'), -- Printmaking (past)
('7b000000-0000-0000-0000-00000000000c', 'ff000001-0000-0000-0000-000000000004', '6b000000-0000-0000-0000-00000000000c', '2026-02-14 13:00:00-05', '2026-02-14 15:30:00-05', 14, 'en', 11, 4500,  'usd'), -- Comic Art (past) ← Zoe

-- Code & Create Boston
('7b000000-0000-0000-0000-00000000000d', 'ff000001-0000-0000-0000-000000000005', '6b000000-0000-0000-0000-00000000000d', '2026-01-31 10:00:00-05', '2026-01-31 14:00:00-05', 16, 'en', 14, 7500,  'usd'), -- Web Design (past) ← Lily
('7b000000-0000-0000-0000-00000000000e', 'ff000001-0000-0000-0000-000000000005', '6b000000-0000-0000-0000-00000000000e', '2026-03-07 10:00:00-05', '2026-03-07 14:00:00-05', 14, 'en', 12, 7500,  'usd'), -- Game Design (past) ← Lily, Zoe
('7b000000-0000-0000-0000-00000000000f', 'ff000001-0000-0000-0000-000000000005', '6b000000-0000-0000-0000-00000000000f', '2026-02-21 10:00:00-05', '2026-02-21 14:00:00-05', 14, 'en', 11, 7500,  'usd'), -- Robotics Lab (past) ← Zoe

-- Siam Muay Thai (Bangkok, past)
('7b000000-0000-0000-0000-000000000010', 'ff000001-0000-0000-0000-000000000006', '6b000000-0000-0000-0000-000000000010', '2026-01-17 09:00:00+07', '2026-01-17 10:30:00+07', 20, 'th', 18, 35000, 'thb'), -- Muay Thai Beginner (past) ← Ton

-- Bangkok Ballet (past)
('7b000000-0000-0000-0000-000000000011', 'ff000001-0000-0000-0000-000000000007', '6b000000-0000-0000-0000-000000000012', '2026-02-07 10:00:00+07', '2026-02-07 12:00:00+07', 18, 'th', 16, 40000, 'thb'), -- Ballet (past) ← Ploy

-- Geniuses STEM (Bangkok, past)
('7b000000-0000-0000-0000-000000000012', 'ff000001-0000-0000-0000-000000000008', '6b000000-0000-0000-0000-000000000014', '2026-02-21 09:00:00+07', '2026-02-21 17:00:00+07', 16, 'en', 14, 45000, 'thb'), -- STEM Camp (past) ← Ploy

-- Extra past occurrences for existing Bangkok events (for Nattaya + Marcus cross-city registrations)
('7b000000-0000-0000-0000-000000000028', '50000000-0000-0000-0000-000000000001', '60000000-0000-0000-0000-000000000001', '2026-03-14 09:00:00+07', '2026-03-14 11:00:00+07', 15, 'en', 13, 50000, 'thb'), -- Robotics Workshop past ← Ploy, Zoe
('7b000000-0000-0000-0000-000000000029', '50000000-0000-0000-0000-000000000001', '60000000-0000-0000-0000-000000000002', '2026-02-21 14:00:00+07', '2026-02-21 15:30:00+07', 12, 'en', 10, 40000, 'thb'), -- Chemistry (past) ← Ploy
('7b000000-0000-0000-0000-00000000002a', '50000000-0000-0000-0000-000000000002', '60000000-0000-0000-0000-000000000004', '2026-01-24 16:00:00+07', '2026-01-24 17:30:00+07', 20, 'en', 17, 30000, 'thb'), -- Soccer Skills past ← Ton

-- ════════════════════════════════════════════════════════════════════════════
-- UPCOMING OCCURRENCES (active registrations, no reviews yet)
-- ════════════════════════════════════════════════════════════════════════════

-- MIT Kids Lab
('7b000000-0000-0000-0000-000000000013', 'ff000001-0000-0000-0000-000000000001', '6b000000-0000-0000-0000-000000000001', '2026-05-09 10:00:00-04', '2026-05-09 13:00:00-04', 16, 'en', 7,  7500,  'usd'), -- Young Engineers ← Lily, Zoe
('7b000000-0000-0000-0000-000000000014', 'ff000001-0000-0000-0000-000000000001', '6b000000-0000-0000-0000-000000000002', '2026-05-30 10:00:00-04', '2026-05-30 12:30:00-04', 14, 'en', 5,  7500,  'usd'), -- Science Explorers
('7b000000-0000-0000-0000-000000000015', 'ff000001-0000-0000-0000-000000000001', '6b000000-0000-0000-0000-000000000003', '2026-06-06 10:00:00-04', '2026-06-06 14:00:00-04', 12, 'en', 8,  7500,  'usd'), -- App Inventor ← Lily
('7b000000-0000-0000-0000-000000000016', 'ff000001-0000-0000-0000-000000000001', '6b000000-0000-0000-0000-000000000001', '2026-07-18 10:00:00-04', '2026-07-18 13:00:00-04', 16, 'en', 3,  7500,  'usd'), -- Young Engineers (summer)
('7b000000-0000-0000-0000-000000000017', 'ff000001-0000-0000-0000-000000000001', '6b000000-0000-0000-0000-000000000002', '2026-07-25 10:00:00-04', '2026-07-25 12:30:00-04', 14, 'en', 2,  7500,  'usd'), -- Science Explorers (summer)

-- Boston Athletic Academy
('7b000000-0000-0000-0000-000000000018', 'ff000001-0000-0000-0000-000000000002', '6b000000-0000-0000-0000-000000000004', '2026-05-02 09:00:00-04', '2026-05-02 10:30:00-04', 20, 'en', 8,  5500,  'usd'), -- Soccer ← Max
('7b000000-0000-0000-0000-000000000019', 'ff000001-0000-0000-0000-000000000002', '6b000000-0000-0000-0000-000000000005', '2026-05-09 09:00:00-04', '2026-05-09 10:30:00-04', 18, 'en', 6,  5500,  'usd'), -- Basketball
('7b000000-0000-0000-0000-00000000001a', 'ff000001-0000-0000-0000-000000000002', '6b000000-0000-0000-0000-000000000006', '2026-05-16 09:00:00-04', '2026-05-16 10:30:00-04', 10, 'en', 4,  5500,  'usd'), -- Swim Team ← (Max cancelled)
('7b000000-0000-0000-0000-00000000001b', 'ff000001-0000-0000-0000-000000000002', '6b000000-0000-0000-0000-000000000004', '2026-06-06 09:00:00-04', '2026-06-06 10:30:00-04', 20, 'en', 11, 5500,  'usd'), -- Soccer (summer)
('7b000000-0000-0000-0000-00000000001c', 'ff000001-0000-0000-0000-000000000002', '6b000000-0000-0000-0000-000000000005', '2026-06-13 09:00:00-04', '2026-06-13 10:30:00-04', 18, 'en', 9,  5500,  'usd'), -- Basketball (summer)

-- NEC Youth
('7b000000-0000-0000-0000-00000000001d', 'ff000001-0000-0000-0000-000000000003', '6b000000-0000-0000-0000-000000000007', '2026-05-02 10:00:00-04', '2026-05-02 11:30:00-04', 8,  'en', 5,  8000,  'usd'), -- Piano Foundations
('7b000000-0000-0000-0000-00000000001e', 'ff000001-0000-0000-0000-000000000003', '6b000000-0000-0000-0000-000000000008', '2026-05-16 10:00:00-04', '2026-05-16 11:30:00-04', 8,  'en', 3,  8000,  'usd'), -- Strings
('7b000000-0000-0000-0000-00000000001f', 'ff000001-0000-0000-0000-000000000003', '6b000000-0000-0000-0000-000000000009', '2026-06-06 13:00:00-04', '2026-06-06 15:00:00-04', 30, 'en', 12, 6000,  'usd'), -- Youth Orchestra

-- Boston Art Center
('7b000000-0000-0000-0000-000000000020', 'ff000001-0000-0000-0000-000000000004', '6b000000-0000-0000-0000-00000000000a', '2026-05-09 13:00:00-04', '2026-05-09 15:30:00-04', 14, 'en', 6,  4500,  'usd'), -- Watercolor
('7b000000-0000-0000-0000-000000000021', 'ff000001-0000-0000-0000-000000000004', '6b000000-0000-0000-0000-00000000000b', '2026-05-23 13:00:00-04', '2026-05-23 15:30:00-04', 12, 'en', 4,  4500,  'usd'), -- Printmaking
('7b000000-0000-0000-0000-000000000022', 'ff000001-0000-0000-0000-000000000004', '6b000000-0000-0000-0000-00000000000c', '2026-05-02 13:00:00-04', '2026-05-02 15:30:00-04', 14, 'en', 6,  4500,  'usd'), -- Comic Art ← Zoe, Max
('7b000000-0000-0000-0000-000000000023', 'ff000001-0000-0000-0000-000000000004', '6b000000-0000-0000-0000-00000000000a', '2026-07-11 13:00:00-04', '2026-07-11 15:30:00-04', 14, 'en', 2,  4500,  'usd'), -- Watercolor (summer)
('7b000000-0000-0000-0000-000000000024', 'ff000001-0000-0000-0000-000000000004', '6b000000-0000-0000-0000-00000000000c', '2026-07-18 13:00:00-04', '2026-07-18 15:30:00-04', 14, 'en', 3,  4500,  'usd'), -- Comic Art (summer)

-- Code & Create Boston
('7b000000-0000-0000-0000-000000000025', 'ff000001-0000-0000-0000-000000000005', '6b000000-0000-0000-0000-00000000000d', '2026-05-16 10:00:00-04', '2026-05-16 14:00:00-04', 16, 'en', 7,  7500,  'usd'), -- Web Design ← (Zoe cancelled)
('7b000000-0000-0000-0000-000000000026', 'ff000001-0000-0000-0000-000000000005', '6b000000-0000-0000-0000-00000000000e', '2026-06-06 10:00:00-04', '2026-06-06 14:00:00-04', 14, 'en', 7,  7500,  'usd'), -- Game Design ← Zoe
('7b000000-0000-0000-0000-000000000027', 'ff000001-0000-0000-0000-000000000005', '6b000000-0000-0000-0000-00000000000f', '2026-05-23 10:00:00-04', '2026-05-23 14:00:00-04', 14, 'en', 8,  7500,  'usd'), -- Robotics Lab ← Zoe
('7b000000-0000-0000-0000-000000000030', 'ff000001-0000-0000-0000-000000000005', '6b000000-0000-0000-0000-00000000000f', '2026-07-11 10:00:00-04', '2026-07-11 14:00:00-04', 14, 'en', 4,  7500,  'usd'), -- Robotics Lab (summer)

-- Siam Muay Thai (Bangkok, upcoming)
('7b000000-0000-0000-0000-000000000031', 'ff000001-0000-0000-0000-000000000006', '6b000000-0000-0000-0000-000000000010', '2026-05-02 09:00:00+07', '2026-05-02 10:30:00+07', 20, 'th', 15, 35000, 'thb'), -- Muay Thai Beginner ← Ton
('7b000000-0000-0000-0000-000000000032', 'ff000001-0000-0000-0000-000000000006', '6b000000-0000-0000-0000-000000000011', '2026-05-09 09:00:00+07', '2026-05-09 10:30:00+07', 14, 'th', 11, 40000, 'thb'), -- Junior Fighters
('7b000000-0000-0000-0000-000000000033', 'ff000001-0000-0000-0000-000000000006', '6b000000-0000-0000-0000-000000000010', '2026-06-06 09:00:00+07', '2026-06-06 10:30:00+07', 20, 'th', 8,  35000, 'thb'), -- Muay Thai (summer)

-- Bangkok Ballet (upcoming)
('7b000000-0000-0000-0000-000000000034', 'ff000001-0000-0000-0000-000000000007', '6b000000-0000-0000-0000-000000000012', '2026-05-09 10:00:00+07', '2026-05-09 12:00:00+07', 18, 'en', 9,  40000, 'thb'), -- Ballet ← Ploy
('7b000000-0000-0000-0000-000000000035', 'ff000001-0000-0000-0000-000000000007', '6b000000-0000-0000-0000-000000000013', '2026-05-23 10:00:00+07', '2026-05-23 12:00:00+07', 16, 'en', 8,  40000, 'thb'), -- Contemporary Dance
('7b000000-0000-0000-0000-000000000036', 'ff000001-0000-0000-0000-000000000007', '6b000000-0000-0000-0000-000000000012', '2026-06-13 10:00:00+07', '2026-06-13 12:00:00+07', 18, 'en', 5,  40000, 'thb'), -- Ballet (summer)

-- Geniuses STEM Thailand (upcoming)
('7b000000-0000-0000-0000-000000000037', 'ff000001-0000-0000-0000-000000000008', '6b000000-0000-0000-0000-000000000014', '2026-05-16 09:00:00+07', '2026-05-16 17:00:00+07', 16, 'en', 11, 45000, 'thb'), -- STEM Camp ← Ploy
('7b000000-0000-0000-0000-000000000038', 'ff000001-0000-0000-0000-000000000008', '6b000000-0000-0000-0000-000000000015', '2026-06-13 09:00:00+07', '2026-06-13 13:00:00+07', 12, 'en', 6,  45000, 'thb'), -- Drone Programming
('7b000000-0000-0000-0000-000000000039', 'ff000001-0000-0000-0000-000000000008', '6b000000-0000-0000-0000-000000000014', '2026-07-04 09:00:00+07', '2026-07-04 17:00:00+07', 16, 'en', 4,  45000, 'thb'); -- STEM Camp (summer)
-- ============================================
-- 10. REGISTRATIONS
-- ============================================
INSERT INTO registration (
    id, child_id, guardian_id, event_occurrence_id, status
) VALUES
-- Emily Johnson (science, technology, math interests)
('80000000-0000-0000-0000-000000000001', '30000000-0000-0000-0000-000000000001', '11111111-1111-1111-1111-111111111111', '70000000-0000-0000-0000-000000000001', 'registered'),
('80000000-0000-0000-0000-000000000002', '30000000-0000-0000-0000-000000000001', '11111111-1111-1111-1111-111111111111', '70000000-0000-0000-0000-000000000003', 'registered'),
('80000000-0000-0000-0000-000000000003', '30000000-0000-0000-0000-000000000001', '11111111-1111-1111-1111-111111111111', '70000000-0000-0000-0000-000000000012', 'registered'),
-- Alex Johnson (sports, music interests)
('80000000-0000-0000-0000-000000000004', '30000000-0000-0000-0000-000000000002', '11111111-1111-1111-1111-111111111111', '70000000-0000-0000-0000-000000000006', 'registered'),
('80000000-0000-0000-0000-000000000005', '30000000-0000-0000-0000-000000000002', '11111111-1111-1111-1111-111111111111', '70000000-0000-0000-0000-00000000000e', 'registered'),
-- Sophie Chen (art, language, music interests)
('80000000-0000-0000-0000-000000000006', '30000000-0000-0000-0000-000000000003', '22222222-2222-2222-2222-222222222222', '70000000-0000-0000-0000-00000000000b', 'registered'),
('80000000-0000-0000-0000-000000000007', '30000000-0000-0000-0000-000000000003', '22222222-2222-2222-2222-222222222222', '70000000-0000-0000-0000-000000000010', 'registered'),
('80000000-0000-0000-0000-000000000008', '30000000-0000-0000-0000-000000000003', '22222222-2222-2222-2222-222222222222', '70000000-0000-0000-0000-000000000011', 'registered'),
-- Aiden Patel (science, sports, technology interests)
('80000000-0000-0000-0000-000000000009', '30000000-0000-0000-0000-000000000004', '33333333-3333-3333-3333-333333333333', '70000000-0000-0000-0000-000000000001', 'registered'),
('80000000-0000-0000-0000-00000000000a', '30000000-0000-0000-0000-000000000004', '33333333-3333-3333-3333-333333333333', '70000000-0000-0000-0000-000000000008', 'registered'),
('80000000-0000-0000-0000-00000000000b', '30000000-0000-0000-0000-000000000004', '33333333-3333-3333-3333-333333333333', '70000000-0000-0000-0000-000000000014', 'registered'),
-- Maya Patel (art, music interests)
('80000000-0000-0000-0000-00000000000c', '30000000-0000-0000-0000-000000000005', '33333333-3333-3333-3333-333333333333', '70000000-0000-0000-0000-00000000000b', 'registered'),
('80000000-0000-0000-0000-00000000000d', '30000000-0000-0000-0000-000000000005', '33333333-3333-3333-3333-333333333333', '70000000-0000-0000-0000-000000000011', 'registered'),
-- Lucas Rodriguez (sports, technology interests)
('80000000-0000-0000-0000-00000000000e', '30000000-0000-0000-0000-000000000006', '44444444-4444-4444-4444-444444444444', '70000000-0000-0000-0000-000000000006', 'registered'),
('80000000-0000-0000-0000-00000000000f', '30000000-0000-0000-0000-000000000006', '44444444-4444-4444-4444-444444444444', '70000000-0000-0000-0000-000000000013', 'registered'),
-- Isabella Thompson (language, art interests)
('80000000-0000-0000-0000-000000000010', '30000000-0000-0000-0000-000000000007', '55555555-5555-5555-5555-555555555555', '70000000-0000-0000-0000-00000000000b', 'registered'),
('80000000-0000-0000-0000-000000000011', '30000000-0000-0000-0000-000000000007', '55555555-5555-5555-5555-555555555555', '70000000-0000-0000-0000-00000000000c', 'registered'),
-- Ethan Thompson (math, science interests)
('80000000-0000-0000-0000-000000000012', '30000000-0000-0000-0000-000000000008', '55555555-5555-5555-5555-555555555555', '70000000-0000-0000-0000-000000000001', 'registered'),
('80000000-0000-0000-0000-000000000013', '30000000-0000-0000-0000-000000000008', '55555555-5555-5555-5555-555555555555', '70000000-0000-0000-0000-000000000003', 'registered'),
-- Hana Tanaka (music, language, art interests)
('80000000-0000-0000-0000-000000000014', '30000000-0000-0000-0000-000000000009', '66666666-6666-6666-6666-666666666666', '70000000-0000-0000-0000-00000000000e', 'registered'),
('80000000-0000-0000-0000-000000000015', '30000000-0000-0000-0000-000000000009', '66666666-6666-6666-6666-666666666666', '70000000-0000-0000-0000-00000000000b', 'registered'),
('80000000-0000-0000-0000-000000000016', '30000000-0000-0000-0000-000000000009', '66666666-6666-6666-6666-666666666666', '70000000-0000-0000-0000-000000000011', 'registered'),
-- Noah Martinez (sports, science interests)
('80000000-0000-0000-0000-000000000017', '30000000-0000-0000-0000-00000000000a', '77777777-7777-7777-7777-777777777777', '70000000-0000-0000-0000-000000000007', 'registered'),
('80000000-0000-0000-0000-000000000018', '30000000-0000-0000-0000-00000000000a', '77777777-7777-7777-7777-777777777777', '70000000-0000-0000-0000-000000000005', 'registered'),
-- Liam Wilson (technology, math, science interests)
('80000000-0000-0000-0000-000000000019', '30000000-0000-0000-0000-00000000000b', '88888888-8888-8888-8888-888888888888', '70000000-0000-0000-0000-000000000002', 'registered'),
('80000000-0000-0000-0000-00000000001a', '30000000-0000-0000-0000-00000000000b', '88888888-8888-8888-8888-888888888888', '70000000-0000-0000-0000-000000000014', 'registered'),
('80000000-0000-0000-0000-00000000001b', '30000000-0000-0000-0000-00000000000b', '88888888-8888-8888-8888-888888888888', '70000000-0000-0000-0000-000000000015', 'registered'),
-- Ava Wilson (music, art interests)
('80000000-0000-0000-0000-00000000001c', '30000000-0000-0000-0000-00000000000c', '88888888-8888-8888-8888-888888888888', '70000000-0000-0000-0000-00000000000f', 'registered'),
('80000000-0000-0000-0000-00000000001d', '30000000-0000-0000-0000-00000000000c', '88888888-8888-8888-8888-888888888888', '70000000-0000-0000-0000-000000000011', 'registered'),
-- Additional registrations showing variety
('80000000-0000-0000-0000-00000000001e', '30000000-0000-0000-0000-000000000001', '11111111-1111-1111-1111-111111111111', '70000000-0000-0000-0000-000000000004', 'registered'),
('80000000-0000-0000-0000-00000000001f', '30000000-0000-0000-0000-000000000004', '33333333-3333-3333-3333-333333333333', '70000000-0000-0000-0000-000000000009', 'registered'),
('80000000-0000-0000-0000-000000000020', '30000000-0000-0000-0000-000000000006', '44444444-4444-4444-4444-444444444444', '70000000-0000-0000-0000-00000000000a', 'registered'),
('80000000-0000-0000-0000-000000000021', '30000000-0000-0000-0000-000000000003', '22222222-2222-2222-2222-222222222222', '70000000-0000-0000-0000-00000000000d', 'registered'),
-- Some cancelled registrations
('80000000-0000-0000-0000-000000000022', '30000000-0000-0000-0000-000000000002', '11111111-1111-1111-1111-111111111111', '70000000-0000-0000-0000-000000000009', 'cancelled'),
('80000000-0000-0000-0000-000000000023', '30000000-0000-0000-0000-000000000005', '33333333-3333-3333-3333-333333333333', '70000000-0000-0000-0000-00000000000c', 'cancelled'),
('80000000-0000-0000-0000-000000000024', '30000000-0000-0000-0000-00000000000a', '77777777-7777-7777-7777-777777777777', '70000000-0000-0000-0000-000000000003', 'cancelled');

-- ============================================
-- 10b. PAYMENTS (seed data for registrations above)
-- ============================================
INSERT INTO payment (
    registration_id, stripe_payment_intent_id, stripe_customer_id, org_stripe_account_id,
    stripe_payment_method_id, total_amount, provider_amount, platform_fee_amount,
    currency, payment_intent_status
) VALUES
('80000000-0000-0000-0000-000000000001', 'pi_seed_001', 'cus_11111111', 'acct_seed_123', 'pm_seed_123', 10000, 8500, 1500, 'usd', 'requires_capture'),
('80000000-0000-0000-0000-000000000002', 'pi_seed_002', 'cus_11111111', 'acct_seed_123', 'pm_seed_123', 10000, 8500, 1500, 'usd', 'requires_capture'),
('80000000-0000-0000-0000-000000000003', 'pi_seed_003', 'cus_11111111', 'acct_seed_123', 'pm_seed_123', 10000, 8500, 1500, 'usd', 'requires_capture'),
('80000000-0000-0000-0000-000000000004', 'pi_seed_004', 'cus_11111111', 'acct_seed_123', 'pm_seed_123', 10000, 8500, 1500, 'usd', 'requires_capture'),
('80000000-0000-0000-0000-000000000005', 'pi_seed_005', 'cus_11111111', 'acct_seed_123', 'pm_seed_123', 10000, 8500, 1500, 'usd', 'requires_capture'),
('80000000-0000-0000-0000-000000000006', 'pi_seed_006', 'cus_22222222', 'acct_seed_123', 'pm_seed_123', 10000, 8500, 1500, 'usd', 'requires_capture'),
('80000000-0000-0000-0000-000000000007', 'pi_seed_007', 'cus_22222222', 'acct_seed_123', 'pm_seed_123', 10000, 8500, 1500, 'usd', 'requires_capture'),
('80000000-0000-0000-0000-000000000008', 'pi_seed_008', 'cus_22222222', 'acct_seed_123', 'pm_seed_123', 10000, 8500, 1500, 'usd', 'requires_capture'),
('80000000-0000-0000-0000-000000000009', 'pi_seed_009', 'cus_33333333', 'acct_seed_123', 'pm_seed_123', 10000, 8500, 1500, 'usd', 'requires_capture'),
('80000000-0000-0000-0000-00000000000a', 'pi_seed_00a', 'cus_33333333', 'acct_seed_123', 'pm_seed_123', 10000, 8500, 1500, 'usd', 'requires_capture'),
('80000000-0000-0000-0000-00000000000b', 'pi_seed_00b', 'cus_33333333', 'acct_seed_123', 'pm_seed_123', 10000, 8500, 1500, 'usd', 'requires_capture'),
('80000000-0000-0000-0000-00000000000c', 'pi_seed_00c', 'cus_33333333', 'acct_seed_123', 'pm_seed_123', 10000, 8500, 1500, 'usd', 'requires_capture'),
('80000000-0000-0000-0000-00000000000d', 'pi_seed_00d', 'cus_33333333', 'acct_seed_123', 'pm_seed_123', 10000, 8500, 1500, 'usd', 'requires_capture'),
('80000000-0000-0000-0000-00000000000e', 'pi_seed_00e', 'cus_44444444', 'acct_seed_123', 'pm_seed_123', 10000, 8500, 1500, 'usd', 'requires_capture'),
('80000000-0000-0000-0000-00000000000f', 'pi_seed_00f', 'cus_44444444', 'acct_seed_123', 'pm_seed_123', 10000, 8500, 1500, 'usd', 'requires_capture'),
('80000000-0000-0000-0000-000000000010', 'pi_seed_010', 'cus_55555555', 'acct_seed_123', 'pm_seed_123', 10000, 8500, 1500, 'usd', 'requires_capture'),
('80000000-0000-0000-0000-000000000011', 'pi_seed_011', 'cus_55555555', 'acct_seed_123', 'pm_seed_123', 10000, 8500, 1500, 'usd', 'requires_capture'),
('80000000-0000-0000-0000-000000000012', 'pi_seed_012', 'cus_55555555', 'acct_seed_123', 'pm_seed_123', 10000, 8500, 1500, 'usd', 'requires_capture'),
('80000000-0000-0000-0000-000000000013', 'pi_seed_013', 'cus_55555555', 'acct_seed_123', 'pm_seed_123', 10000, 8500, 1500, 'usd', 'requires_capture'),
('80000000-0000-0000-0000-000000000014', 'pi_seed_014', 'cus_66666666', 'acct_seed_123', 'pm_seed_123', 10000, 8500, 1500, 'usd', 'requires_capture'),
('80000000-0000-0000-0000-000000000015', 'pi_seed_015', 'cus_66666666', 'acct_seed_123', 'pm_seed_123', 10000, 8500, 1500, 'usd', 'requires_capture'),
('80000000-0000-0000-0000-000000000016', 'pi_seed_016', 'cus_66666666', 'acct_seed_123', 'pm_seed_123', 10000, 8500, 1500, 'usd', 'requires_capture'),
('80000000-0000-0000-0000-000000000017', 'pi_seed_017', 'cus_77777777', 'acct_seed_123', 'pm_seed_123', 10000, 8500, 1500, 'usd', 'requires_capture'),
('80000000-0000-0000-0000-000000000018', 'pi_seed_018', 'cus_77777777', 'acct_seed_123', 'pm_seed_123', 10000, 8500, 1500, 'usd', 'requires_capture'),
('80000000-0000-0000-0000-000000000019', 'pi_seed_019', 'cus_88888888', 'acct_seed_123', 'pm_seed_123', 10000, 8500, 1500, 'usd', 'requires_capture'),
('80000000-0000-0000-0000-00000000001a', 'pi_seed_01a', 'cus_88888888', 'acct_seed_123', 'pm_seed_123', 10000, 8500, 1500, 'usd', 'requires_capture'),
('80000000-0000-0000-0000-00000000001b', 'pi_seed_01b', 'cus_88888888', 'acct_seed_123', 'pm_seed_123', 10000, 8500, 1500, 'usd', 'requires_capture'),
('80000000-0000-0000-0000-00000000001c', 'pi_seed_01c', 'cus_88888888', 'acct_seed_123', 'pm_seed_123', 10000, 8500, 1500, 'usd', 'requires_capture'),
('80000000-0000-0000-0000-00000000001d', 'pi_seed_01d', 'cus_88888888', 'acct_seed_123', 'pm_seed_123', 10000, 8500, 1500, 'usd', 'requires_capture'),
('80000000-0000-0000-0000-00000000001e', 'pi_seed_01e', 'cus_11111111', 'acct_seed_123', 'pm_seed_123', 10000, 8500, 1500, 'usd', 'requires_capture'),
('80000000-0000-0000-0000-00000000001f', 'pi_seed_01f', 'cus_33333333', 'acct_seed_123', 'pm_seed_123', 10000, 8500, 1500, 'usd', 'requires_capture'),
('80000000-0000-0000-0000-000000000020', 'pi_seed_020', 'cus_44444444', 'acct_seed_123', 'pm_seed_123', 10000, 8500, 1500, 'usd', 'requires_capture'),
('80000000-0000-0000-0000-000000000021', 'pi_seed_021', 'cus_22222222', 'acct_seed_123', 'pm_seed_123', 10000, 8500, 1500, 'usd', 'requires_capture'),
-- Cancelled registrations
('80000000-0000-0000-0000-000000000022', 'pi_seed_022', 'cus_11111111', 'acct_seed_123', 'pm_seed_123', 10000, 8500, 1500, 'usd', 'succeeded'),
('80000000-0000-0000-0000-000000000023', 'pi_seed_023', 'cus_33333333', 'acct_seed_123', 'pm_seed_123', 10000, 8500, 1500, 'usd', 'succeeded'),
('80000000-0000-0000-0000-000000000024', 'pi_seed_024', 'cus_77777777', 'acct_seed_123', 'pm_seed_123', 10000, 8500, 1500, 'usd', 'succeeded');
-- ============================================
-- 10. EMERGENCY CONTACTS
-- ============================================
INSERT INTO emergency_contacts (id, name, guardian_id, phone_number) VALUES
('ec000000-0000-0000-0000-000000000001', 'David Johnson',   '11111111-1111-1111-1111-111111111111', '+16175550101'), -- Sarah Johnson's contact
('ec000000-0000-0000-0000-000000000002', 'Linda Chen',      '22222222-2222-2222-2222-222222222222', '+16175550102'), -- Michael Chen's contact
('ec000000-0000-0000-0000-000000000003', 'Raj Patel',       '33333333-3333-3333-3333-333333333333', '+16175550103'), -- Priya Patel's contact
('ec000000-0000-0000-0000-000000000004', 'Maria Rodriguez', '44444444-4444-4444-4444-444444444444', '+16175550104'), -- Carlos Rodriguez's contact
('ec000000-0000-0000-0000-000000000005', 'Robert Thompson', '55555555-5555-5555-5555-555555555555', '+16175550105'), -- Emma Thompson's contact
('ec000000-0000-0000-0000-000000000006', 'Kenji Tanaka',    '66666666-6666-6666-6666-666666666666', '+16175550106'), -- Yuki Tanaka's contact
('ec000000-0000-0000-0000-000000000007', 'Luis Martinez',   '77777777-7777-7777-7777-777777777777', '+16175550107'), -- Olivia Martinez's contact
('ec000000-0000-0000-0000-000000000008', 'Patricia Wilson', '88888888-8888-8888-8888-888888888888', '+16175550108'); -- James Wilson's contact
-- ============================================
-- 22. REGISTRATIONS + PAYMENTS — Demo Guardian Accounts
--
-- Jennifer Park  (cc000001-...001)  cus_jenniferpark
--   children: Lily (ad000001-...001), Max (ad000001-...002)
-- Nattaya Srisuk (cc000001-...002)  cus_nattayath
--   children: Ploy (ad000001-...003), Ton  (ad000001-...004)
-- Marcus Webb    (cc000001-...003)  cus_marcuswebb
--   children: Zoe  (ad000001-...005)
-- ============================================

INSERT INTO registration (id, child_id, guardian_id, event_occurrence_id, status) VALUES

-- ── Jennifer Park — COMPLETED (past) ────────────────────────────────────────
('ae000001-0000-0000-0000-000000000001', 'ad000001-0000-0000-0000-000000000001', 'cc000001-0000-0000-0000-000000000001', '7b000000-0000-0000-0000-000000000001', 'registered'), -- Lily → Young Engineers (past)
('ae000001-0000-0000-0000-000000000002', 'ad000001-0000-0000-0000-000000000001', 'cc000001-0000-0000-0000-000000000001', '7b000000-0000-0000-0000-000000000002', 'registered'), -- Lily → Science Explorers (past)
('ae000001-0000-0000-0000-000000000003', 'ad000001-0000-0000-0000-000000000001', 'cc000001-0000-0000-0000-000000000001', '7b000000-0000-0000-0000-000000000003', 'registered'), -- Lily → App Inventor (past)
('ae000001-0000-0000-0000-000000000004', 'ad000001-0000-0000-0000-000000000001', 'cc000001-0000-0000-0000-000000000001', '7b000000-0000-0000-0000-00000000000d', 'registered'), -- Lily → Web Design (past)
('ae000001-0000-0000-0000-000000000005', 'ad000001-0000-0000-0000-000000000001', 'cc000001-0000-0000-0000-000000000001', '7b000000-0000-0000-0000-00000000000e', 'registered'), -- Lily → Game Design (past)
('ae000001-0000-0000-0000-000000000006', 'ad000001-0000-0000-0000-000000000002', 'cc000001-0000-0000-0000-000000000001', '7b000000-0000-0000-0000-000000000004', 'registered'), -- Max  → Youth Soccer (past)
('ae000001-0000-0000-0000-000000000007', 'ad000001-0000-0000-0000-000000000002', 'cc000001-0000-0000-0000-000000000001', '7b000000-0000-0000-0000-000000000005', 'registered'), -- Max  → Basketball (past)
('ae000001-0000-0000-0000-000000000008', 'ad000001-0000-0000-0000-000000000002', 'cc000001-0000-0000-0000-000000000001', '7b000000-0000-0000-0000-00000000000a', 'registered'), -- Max  → Watercolor (past)

-- ── Jennifer Park — UPCOMING ─────────────────────────────────────────────────
('ae000001-0000-0000-0000-000000000009', 'ad000001-0000-0000-0000-000000000001', 'cc000001-0000-0000-0000-000000000001', '7b000000-0000-0000-0000-000000000015', 'registered'), -- Lily → App Inventor (upcoming)
('ae000001-0000-0000-0000-00000000000a', 'ad000001-0000-0000-0000-000000000001', 'cc000001-0000-0000-0000-000000000001', '7b000000-0000-0000-0000-000000000013', 'registered'), -- Lily → Young Engineers (upcoming)
('ae000001-0000-0000-0000-00000000000b', 'ad000001-0000-0000-0000-000000000002', 'cc000001-0000-0000-0000-000000000001', '7b000000-0000-0000-0000-000000000018', 'registered'), -- Max  → Soccer (upcoming)
('ae000001-0000-0000-0000-00000000000c', 'ad000001-0000-0000-0000-000000000002', 'cc000001-0000-0000-0000-000000000001', '7b000000-0000-0000-0000-000000000022', 'registered'), -- Max  → Comic Art (upcoming)

-- ── Jennifer Park — CANCELLED ────────────────────────────────────────────────
('ae000001-0000-0000-0000-00000000000d', 'ad000001-0000-0000-0000-000000000002', 'cc000001-0000-0000-0000-000000000001', '7b000000-0000-0000-0000-00000000001a', 'cancelled'),  -- Max  → Swim Team (cancelled)

-- ── Nattaya Srisuk — COMPLETED (past) ───────────────────────────────────────
('ae000001-0000-0000-0000-00000000000e', 'ad000001-0000-0000-0000-000000000003', 'cc000001-0000-0000-0000-000000000002', '7b000000-0000-0000-0000-000000000011', 'registered'), -- Ploy → Ballet (past)
('ae000001-0000-0000-0000-00000000000f', 'ad000001-0000-0000-0000-000000000003', 'cc000001-0000-0000-0000-000000000002', '7b000000-0000-0000-0000-000000000012', 'registered'), -- Ploy → STEM Camp (past)
('ae000001-0000-0000-0000-000000000010', 'ad000001-0000-0000-0000-000000000003', 'cc000001-0000-0000-0000-000000000002', '7b000000-0000-0000-0000-000000000028', 'registered'), -- Ploy → Robotics Workshop (past)
('ae000001-0000-0000-0000-000000000011', 'ad000001-0000-0000-0000-000000000003', 'cc000001-0000-0000-0000-000000000002', '7b000000-0000-0000-0000-000000000029', 'registered'), -- Ploy → Chemistry (past)
('ae000001-0000-0000-0000-000000000012', 'ad000001-0000-0000-0000-000000000004', 'cc000001-0000-0000-0000-000000000002', '7b000000-0000-0000-0000-000000000010', 'registered'), -- Ton  → Muay Thai Beginner (past)
('ae000001-0000-0000-0000-000000000013', 'ad000001-0000-0000-0000-000000000004', 'cc000001-0000-0000-0000-000000000002', '7b000000-0000-0000-0000-00000000002a', 'registered'), -- Ton  → Soccer Skills (past)

-- ── Nattaya Srisuk — UPCOMING ────────────────────────────────────────────────
('ae000001-0000-0000-0000-000000000014', 'ad000001-0000-0000-0000-000000000003', 'cc000001-0000-0000-0000-000000000002', '7b000000-0000-0000-0000-000000000034', 'registered'), -- Ploy → Ballet (upcoming)
('ae000001-0000-0000-0000-000000000015', 'ad000001-0000-0000-0000-000000000003', 'cc000001-0000-0000-0000-000000000002', '7b000000-0000-0000-0000-000000000037', 'registered'), -- Ploy → STEM Camp (upcoming)
('ae000001-0000-0000-0000-000000000016', 'ad000001-0000-0000-0000-000000000004', 'cc000001-0000-0000-0000-000000000002', '7b000000-0000-0000-0000-000000000031', 'registered'), -- Ton  → Muay Thai (upcoming)
('ae000001-0000-0000-0000-000000000017', 'ad000001-0000-0000-0000-000000000004', 'cc000001-0000-0000-0000-000000000002', '7b000000-0000-0000-0000-00000000001b', 'registered'), -- Ton  → Soccer (upcoming, Boston!)

-- ── Marcus Webb — COMPLETED (past) ──────────────────────────────────────────
('ae000001-0000-0000-0000-000000000018', 'ad000001-0000-0000-0000-000000000005', 'cc000001-0000-0000-0000-000000000003', '7b000000-0000-0000-0000-000000000001', 'registered'), -- Zoe → Young Engineers (past)
('ae000001-0000-0000-0000-000000000019', 'ad000001-0000-0000-0000-000000000005', 'cc000001-0000-0000-0000-000000000003', '7b000000-0000-0000-0000-00000000000f', 'registered'), -- Zoe → Robotics Lab (past)
('ae000001-0000-0000-0000-00000000001a', 'ad000001-0000-0000-0000-000000000005', 'cc000001-0000-0000-0000-000000000003', '7b000000-0000-0000-0000-00000000000e', 'registered'), -- Zoe → Game Design (past)
('ae000001-0000-0000-0000-00000000001b', 'ad000001-0000-0000-0000-000000000005', 'cc000001-0000-0000-0000-000000000003', '7b000000-0000-0000-0000-00000000000a', 'registered'), -- Zoe → Watercolor (past)
('ae000001-0000-0000-0000-00000000001c', 'ad000001-0000-0000-0000-000000000005', 'cc000001-0000-0000-0000-000000000003', '7b000000-0000-0000-0000-00000000000c', 'registered'), -- Zoe → Comic Art (past)
('ae000001-0000-0000-0000-00000000001d', 'ad000001-0000-0000-0000-000000000005', 'cc000001-0000-0000-0000-000000000003', '7b000000-0000-0000-0000-000000000028', 'registered'), -- Zoe → Robotics Workshop Bangkok (past)
('ae000001-0000-0000-0000-00000000001e', 'ad000001-0000-0000-0000-000000000005', 'cc000001-0000-0000-0000-000000000003', '7b000000-0000-0000-0000-000000000002', 'registered'), -- Zoe → Science Explorers (past)

-- ── Marcus Webb — UPCOMING ───────────────────────────────────────────────────
('ae000001-0000-0000-0000-00000000001f', 'ad000001-0000-0000-0000-000000000005', 'cc000001-0000-0000-0000-000000000003', '7b000000-0000-0000-0000-000000000027', 'registered'), -- Zoe → Robotics Lab (upcoming)
('ae000001-0000-0000-0000-000000000020', 'ad000001-0000-0000-0000-000000000005', 'cc000001-0000-0000-0000-000000000003', '7b000000-0000-0000-0000-000000000013', 'registered'), -- Zoe → Young Engineers (upcoming)
('ae000001-0000-0000-0000-000000000021', 'ad000001-0000-0000-0000-000000000005', 'cc000001-0000-0000-0000-000000000003', '7b000000-0000-0000-0000-000000000026', 'registered'), -- Zoe → Game Design (upcoming)
('ae000001-0000-0000-0000-000000000022', 'ad000001-0000-0000-0000-000000000005', 'cc000001-0000-0000-0000-000000000003', '7b000000-0000-0000-0000-000000000022', 'registered'), -- Zoe → Comic Art (upcoming)

-- ── Marcus Webb — CANCELLED ──────────────────────────────────────────────────
('ae000001-0000-0000-0000-000000000023', 'ad000001-0000-0000-0000-000000000005', 'cc000001-0000-0000-0000-000000000003', '7b000000-0000-0000-0000-000000000025', 'cancelled');  -- Zoe → Web Design (cancelled)


-- ============================================
-- 22b. PAYMENTS for the registrations above
-- ============================================
-- Org Stripe account mapping:
--   MIT Kids Lab        → acct_mitkidslab
--   Boston Athletic     → acct_bostonathletic
--   Boston Art Center   → acct_bostonartcenter
--   Code & Create       → acct_codecreate
--   Siam Muay Thai      → acct_siammuaythai
--   Bangkok Ballet      → acct_bangkokballet
--   Geniuses STEM       → acct_geniusstem
--   Science Academy BKK → acct_seed_123  (existing)
--   Champions Sports BKK→ acct_seed_123  (existing)

INSERT INTO payment (
    registration_id, stripe_payment_intent_id, stripe_customer_id, org_stripe_account_id,
    stripe_payment_method_id, total_amount, provider_amount, platform_fee_amount,
    currency, payment_intent_status
) VALUES
-- Jennifer — past (requires_capture = completed/held)
('ae000001-0000-0000-0000-000000000001', 'pi_demo_jen_001', 'cus_jenniferpark', 'acct_mitkidslab',     'pm_demo_jen', 7500,  6375,  1125, 'usd', 'requires_capture'),
('ae000001-0000-0000-0000-000000000002', 'pi_demo_jen_002', 'cus_jenniferpark', 'acct_mitkidslab',     'pm_demo_jen', 7500,  6375,  1125, 'usd', 'requires_capture'),
('ae000001-0000-0000-0000-000000000003', 'pi_demo_jen_003', 'cus_jenniferpark', 'acct_mitkidslab',     'pm_demo_jen', 7500,  6375,  1125, 'usd', 'requires_capture'),
('ae000001-0000-0000-0000-000000000004', 'pi_demo_jen_004', 'cus_jenniferpark', 'acct_codecreate',     'pm_demo_jen', 7500,  6375,  1125, 'usd', 'requires_capture'),
('ae000001-0000-0000-0000-000000000005', 'pi_demo_jen_005', 'cus_jenniferpark', 'acct_codecreate',     'pm_demo_jen', 7500,  6375,  1125, 'usd', 'requires_capture'),
('ae000001-0000-0000-0000-000000000006', 'pi_demo_jen_006', 'cus_jenniferpark', 'acct_bostonathletic', 'pm_demo_jen', 5500,  4675,  825,  'usd', 'requires_capture'),
('ae000001-0000-0000-0000-000000000007', 'pi_demo_jen_007', 'cus_jenniferpark', 'acct_bostonathletic', 'pm_demo_jen', 5500,  4675,  825,  'usd', 'requires_capture'),
('ae000001-0000-0000-0000-000000000008', 'pi_demo_jen_008', 'cus_jenniferpark', 'acct_bostonartcenter','pm_demo_jen', 4500,  3825,  675,  'usd', 'requires_capture'),
-- Jennifer — upcoming
('ae000001-0000-0000-0000-000000000009', 'pi_demo_jen_009', 'cus_jenniferpark', 'acct_mitkidslab',     'pm_demo_jen', 7500,  6375,  1125, 'usd', 'requires_capture'),
('ae000001-0000-0000-0000-00000000000a', 'pi_demo_jen_00a', 'cus_jenniferpark', 'acct_mitkidslab',     'pm_demo_jen', 7500,  6375,  1125, 'usd', 'requires_capture'),
('ae000001-0000-0000-0000-00000000000b', 'pi_demo_jen_00b', 'cus_jenniferpark', 'acct_bostonathletic', 'pm_demo_jen', 5500,  4675,  825,  'usd', 'requires_capture'),
('ae000001-0000-0000-0000-00000000000c', 'pi_demo_jen_00c', 'cus_jenniferpark', 'acct_bostonartcenter','pm_demo_jen', 4500,  3825,  675,  'usd', 'requires_capture'),
-- Jennifer — cancelled (succeeded = refunded)
('ae000001-0000-0000-0000-00000000000d', 'pi_demo_jen_00d', 'cus_jenniferpark', 'acct_bostonathletic', 'pm_demo_jen', 5500,  4675,  825,  'usd', 'succeeded'),

-- Nattaya — past
('ae000001-0000-0000-0000-00000000000e', 'pi_demo_nat_001', 'cus_nattayath', 'acct_bangkokballet',  'pm_demo_nat', 40000, 34000, 6000, 'thb', 'requires_capture'),
('ae000001-0000-0000-0000-00000000000f', 'pi_demo_nat_002', 'cus_nattayath', 'acct_geniusstem',     'pm_demo_nat', 45000, 38250, 6750, 'thb', 'requires_capture'),
('ae000001-0000-0000-0000-000000000010', 'pi_demo_nat_003', 'cus_nattayath', 'acct_seed_123',       'pm_demo_nat', 50000, 42500, 7500, 'thb', 'requires_capture'),
('ae000001-0000-0000-0000-000000000011', 'pi_demo_nat_004', 'cus_nattayath', 'acct_seed_123',       'pm_demo_nat', 40000, 34000, 6000, 'thb', 'requires_capture'),
('ae000001-0000-0000-0000-000000000012', 'pi_demo_nat_005', 'cus_nattayath', 'acct_siammuaythai',   'pm_demo_nat', 35000, 29750, 5250, 'thb', 'requires_capture'),
('ae000001-0000-0000-0000-000000000013', 'pi_demo_nat_006', 'cus_nattayath', 'acct_seed_123',       'pm_demo_nat', 30000, 25500, 4500, 'thb', 'requires_capture'),
-- Nattaya — upcoming
('ae000001-0000-0000-0000-000000000014', 'pi_demo_nat_007', 'cus_nattayath', 'acct_bangkokballet',  'pm_demo_nat', 40000, 34000, 6000, 'thb', 'requires_capture'),
('ae000001-0000-0000-0000-000000000015', 'pi_demo_nat_008', 'cus_nattayath', 'acct_geniusstem',     'pm_demo_nat', 45000, 38250, 6750, 'thb', 'requires_capture'),
('ae000001-0000-0000-0000-000000000016', 'pi_demo_nat_009', 'cus_nattayath', 'acct_siammuaythai',   'pm_demo_nat', 35000, 29750, 5250, 'thb', 'requires_capture'),
('ae000001-0000-0000-0000-000000000017', 'pi_demo_nat_00a', 'cus_nattayath', 'acct_bostonathletic', 'pm_demo_nat', 5500,  4675,  825,  'usd', 'requires_capture'),

-- Marcus — past
('ae000001-0000-0000-0000-000000000018', 'pi_demo_mar_001', 'cus_marcuswebb', 'acct_mitkidslab',     'pm_demo_mar', 7500,  6375,  1125, 'usd', 'requires_capture'),
('ae000001-0000-0000-0000-000000000019', 'pi_demo_mar_002', 'cus_marcuswebb', 'acct_codecreate',     'pm_demo_mar', 7500,  6375,  1125, 'usd', 'requires_capture'),
('ae000001-0000-0000-0000-00000000001a', 'pi_demo_mar_003', 'cus_marcuswebb', 'acct_codecreate',     'pm_demo_mar', 7500,  6375,  1125, 'usd', 'requires_capture'),
('ae000001-0000-0000-0000-00000000001b', 'pi_demo_mar_004', 'cus_marcuswebb', 'acct_bostonartcenter','pm_demo_mar', 4500,  3825,  675,  'usd', 'requires_capture'),
('ae000001-0000-0000-0000-00000000001c', 'pi_demo_mar_005', 'cus_marcuswebb', 'acct_bostonartcenter','pm_demo_mar', 4500,  3825,  675,  'usd', 'requires_capture'),
('ae000001-0000-0000-0000-00000000001d', 'pi_demo_mar_006', 'cus_marcuswebb', 'acct_seed_123',       'pm_demo_mar', 50000, 42500, 7500, 'thb', 'requires_capture'),
('ae000001-0000-0000-0000-00000000001e', 'pi_demo_mar_007', 'cus_marcuswebb', 'acct_mitkidslab',     'pm_demo_mar', 7500,  6375,  1125, 'usd', 'requires_capture'),
-- Marcus — upcoming
('ae000001-0000-0000-0000-00000000001f', 'pi_demo_mar_008', 'cus_marcuswebb', 'acct_codecreate',     'pm_demo_mar', 7500,  6375,  1125, 'usd', 'requires_capture'),
('ae000001-0000-0000-0000-000000000020', 'pi_demo_mar_009', 'cus_marcuswebb', 'acct_mitkidslab',     'pm_demo_mar', 7500,  6375,  1125, 'usd', 'requires_capture'),
('ae000001-0000-0000-0000-000000000021', 'pi_demo_mar_00a', 'cus_marcuswebb', 'acct_codecreate',     'pm_demo_mar', 7500,  6375,  1125, 'usd', 'requires_capture'),
('ae000001-0000-0000-0000-000000000022', 'pi_demo_mar_00b', 'cus_marcuswebb', 'acct_bostonartcenter','pm_demo_mar', 4500,  3825,  675,  'usd', 'requires_capture'),
-- Marcus — cancelled
('ae000001-0000-0000-0000-000000000023', 'pi_demo_mar_00c', 'cus_marcuswebb', 'acct_codecreate',     'pm_demo_mar', 7500,  6375,  1125, 'usd', 'succeeded');
INSERT INTO review (
    id,
    registration_id,
    guardian_id,
    rating,
    description_en,
    description_th,
    categories
) VALUES

(
    '90000000-0000-0000-0000-000000000001',
    '80000000-0000-0000-0000-000000000001',
    '11111111-1111-1111-1111-111111111111',
    5,
    'Emily had a great time and came home excited to talk about everything she learned.',
    'เอมิลี่สนุกมากและกลับบ้านด้วยความตื่นเต้นที่จะเล่าทุกสิ่งที่เธอได้เรียนรู้',
    ARRAY['informative', 'interesting', 'engaging']::review_category[]
),
(
    '90000000-0000-0000-0000-000000000002',
    '80000000-0000-0000-0000-000000000002',
    '11111111-1111-1111-1111-111111111111',
    4,
    'The event was excellently organized. The activities were fun and educational.',
    'งานจัดได้อย่างดีเยี่ยม กิจกรรมสนุกและให้ความรู้',
    ARRAY['fun', 'informative']::review_category[]
),
(
    '90000000-0000-0000-0000-000000000003',
    '80000000-0000-0000-0000-000000000003',
    '11111111-1111-1111-1111-111111111111',
    5,
    'Engaging instructors and hands-on activities kept Emily interested the whole time.',
    'ครูผู้สอนที่น่าสนใจและกิจกรรมภาคปฏิบัติช่วยให้เอมิลี่สนใจตลอดเวลา',
    ARRAY['engaging', 'interesting']::review_category[]
),

(
    '90000000-0000-0000-0000-000000000004',
    '80000000-0000-0000-0000-000000000004',
    '11111111-1111-1111-1111-111111111111',
    4,
    'Alex really enjoyed the physical activities and music at the event.',
    'อเล็กซ์สนุกกับกิจกรรมทางกายภาพและดนตรีในงานเป็นอย่างมาก',
    ARRAY['fun', 'engaging']::review_category[]
);
-- ============================================
-- 23. REVIEWS — 21 detailed bilingual reviews
--     Jennifer Park  (cc000001-...001): 8 reviews for registrations ae...001–008
--     Nattaya Srisuk (cc000001-...002): 6 reviews for registrations ae...00e–013
--     Marcus Webb    (cc000001-...003): 7 reviews for registrations ae...018–01e
-- ============================================

INSERT INTO review (id, registration_id, guardian_id, rating, description_en, description_th, categories) VALUES

-- ════════════════════════════════════════════════════════════════════════════
-- JENNIFER PARK
-- ════════════════════════════════════════════════════════════════════════════

-- 1. Young Engineers — 5 stars
('af000001-0000-0000-0000-000000000001',
 'ae000001-0000-0000-0000-000000000001',
 'cc000001-0000-0000-0000-000000000001',
 5,
 'Lily came home absolutely buzzing. She spent the entire dinner explaining how her gear mechanism worked and then asked when she could sign up again. Dr. Walsh''s team has a real talent for meeting kids where they are — Lily is only eight and never felt lost. The ratio of instructors to kids was excellent, materials were high quality, and the pacing kept everyone engaged from start to finish. We''ve done a lot of enrichment programs and this is the best one we''ve found.',
 'ลิลี่กลับบ้านมาตื่นเต้นมาก เธอใช้เวลาอาหารเย็นทั้งหมดอธิบายว่ากลไกเฟืองของเธอทำงานอย่างไร แล้วถามว่าจะสมัครใหม่ได้เมื่อไหร่ ทีมของ Dr. Walsh มีพรสวรรค์จริงๆ ในการพบเด็กๆ ในจุดที่พวกเขาอยู่',
 ARRAY['engaging', 'educational', 'interactive', 'well structured', 'beginner friendly']::review_category[]),

-- 2. Science Explorers — 5 stars
('af000001-0000-0000-0000-000000000002',
 'ae000001-0000-0000-0000-000000000002',
 'cc000001-0000-0000-0000-000000000001',
 5,
 'The instructors were brilliant at breaking down complex concepts without dumbing them down. Lily''s group did a baking soda and vinegar volcano but then immediately moved into the actual chemistry of why it works — acids, bases, pH — and Lily retained all of it. She''s now asking for a chemistry set. The facilities are clean, everything felt safe, and the staff communicated beautifully with parents before and after.',
 'ครูผู้สอนยอดเยี่ยมมากในการอธิบายแนวคิดที่ซับซ้อนโดยไม่ทำให้ง่ายเกินไป กลุ่มของลิลี่ทำภูเขาไฟโซดาเบกกิ้งและน้ำส้มสายชู แล้วทันทีก็เรียนรู้เคมีจริงๆ ว่าทำไมมันถึงทำงาน สิ่งอำนวยความสะดวกสะอาด ทุกอย่างรู้สึกปลอดภัย',
 ARRAY['educational', 'informative', 'insightful', 'clear instruction', 'welcoming']::review_category[]),

-- 3. App Inventor — 4 stars
('af000001-0000-0000-0000-000000000003',
 'ae000001-0000-0000-0000-000000000003',
 'cc000001-0000-0000-0000-000000000001',
 4,
 'Lily built a working quiz app by the end of the day — I was honestly impressed. The instructors are clearly experienced and the curriculum is well thought through. My only minor note is that the second half of the session moved faster than the first and a few kids struggled to keep up. But Lily loved it and immediately wanted to add more features to her app at home. Four stars only because I think an extra 30 minutes would make it perfect.',
 'ลิลี่สร้างแอปทดสอบที่ใช้งานได้ภายในสิ้นวัน ฉันประทับใจจริงๆ ครูผู้สอนมีประสบการณ์อย่างชัดเจนและหลักสูตรได้รับการคิดมาอย่างดี ข้อสังเกตเล็กน้อยของฉันคือครึ่งหลังของเซสชั่นเร็วกว่าครึ่งแรก',
 ARRAY['educational', 'engaging', 'interactive', 'challenging', 'well structured']::review_category[]),

-- 4. Web Design — 5 stars
('af000001-0000-0000-0000-000000000004',
 'ae000001-0000-0000-0000-000000000004',
 'cc000001-0000-0000-0000-000000000001',
 5,
 'Lily built her first real website in a single afternoon and published it live on the internet before we even got home. Jake is an exceptional teacher — he explains every concept clearly, checks for understanding constantly, and makes every kid feel like a real developer. Lily''s been updating her site every evening since. Worth every penny.',
 'ลิลี่สร้างเว็บไซต์จริงแรกของเธอในช่วงบ่ายเดียวและเผยแพร่ live บนอินเทอร์เน็ตก่อนที่เราจะกลับถึงบ้านด้วยซ้ำ Jake เป็นครูที่ยอดเยี่ยม เขาอธิบายทุกแนวคิดอย่างชัดเจน ตรวจสอบความเข้าใจอยู่ตลอดเวลา',
 ARRAY['educational', 'engaging', 'clear instruction', 'well structured', 'interactive']::review_category[]),

-- 5. Game Design — 4 stars
('af000001-0000-0000-0000-000000000005',
 'ae000001-0000-0000-0000-000000000005',
 'cc000001-0000-0000-0000-000000000001',
 4,
 'Lily had a great time and made a simple but actually playable platformer game. The Unity environment is a bit advanced for a 9-year-old but the step-by-step guidance made it work. The highlight was watching all the kids play each other''s games at the end and the genuine excitement on their faces. Would recommend for kids who already have some coding exposure — it''s a big step up from block coding.',
 'ลิลี่สนุกมากและสร้างเกม platformer ที่เล่นได้จริงแม้จะเรียบง่าย สภาพแวดล้อม Unity ค่อนข้างล้ำสำหรับเด็กอายุ 9 ปี แต่คำแนะนำทีละขั้นทำให้มันได้ผล จุดเด่นคือการดูเด็กๆ ทุกคนเล่นเกมของกันและกันในตอนท้าย',
 ARRAY['engaging', 'educational', 'interactive', 'challenging', 'intermediate']::review_category[]),

-- 6. Youth Soccer — 5 stars
('af000001-0000-0000-0000-000000000006',
 'ae000001-0000-0000-0000-000000000006',
 'cc000001-0000-0000-0000-000000000001',
 5,
 'Max absolutely loves this program. He''s only five so I wasn''t sure how he would do, but Coach Rivers and her team are incredibly patient with the younger kids and make sure everyone is included. Max came home talking about the drills like he was a professional. He asks every week if it''s "soccer day." We''ve already signed him up for the next session.',
 'แม็กซ์รักโปรแกรมนี้มากเลย เขาอายุแค่ห้าขวบดังนั้นฉันไม่แน่ใจว่าเขาจะทำได้แค่ไหน แต่โค้ช Rivers และทีมของเธอมีความอดทนอย่างไม่น่าเชื่อกับเด็กเล็กและทำให้แน่ใจว่าทุกคนมีส่วนร่วม',
 ARRAY['welcoming', 'engaging', 'fun', 'beginner friendly', 'inclusive']::review_category[]),

-- 7. Basketball — 4 stars
('af000001-0000-0000-0000-000000000007',
 'ae000001-0000-0000-0000-000000000007',
 'cc000001-0000-0000-0000-000000000001',
 4,
 'Really solid coaching and great energy in the gym. Max loved the 3-on-3 game at the end. My only note is that the session could use slightly shorter water breaks — the momentum dipped a little in the middle. But the fundamentals work was top notch and Max came away understanding how to dribble with both hands, which I wasn''t expecting after a single session.',
 'การโค้ชที่แน่นหนาและพลังงานที่ยอดเยี่ยมในยิม แม็กซ์ชอบเกม 3 ต่อ 3 ในตอนท้าย ข้อสังเกตของฉันคือเซสชั่นอาจใช้เวลาพักดื่มน้ำสั้นลงเล็กน้อย แต่การฝึกพื้นฐานยอดเยี่ยมมาก',
 ARRAY['educational', 'engaging', 'fun', 'clear instruction']::review_category[]),

-- 8. Watercolor — 5 stars
('af000001-0000-0000-0000-000000000008',
 'ae000001-0000-0000-0000-000000000008',
 'cc000001-0000-0000-0000-000000000001',
 5,
 'I signed Max up for watercolor thinking it would be a fun afternoon activity — I had no idea he would discover a genuine talent for it. Maria and her team created a completely non-judgmental space where every child''s work was celebrated. Max''s painting is now framed above our fireplace. He''s asked to do every session going forward. Truly a special program.',
 'ฉันสมัครให้แม็กซ์เรียนสีน้ำโดยคิดว่ามันจะเป็นกิจกรรมบ่ายสนุกๆ ฉันไม่รู้เลยว่าเขาจะค้นพบความสามารถแท้จริงในสิ่งนั้น ภาพวาดของแม็กซ์ตอนนี้อยู่ในกรอบเหนือเตาผิงของเรา',
 ARRAY['welcoming', 'engaging', 'fun', 'inclusive', 'interactive']::review_category[]),


-- ════════════════════════════════════════════════════════════════════════════
-- NATTAYA SRISUK
-- ════════════════════════════════════════════════════════════════════════════

-- 9. Ballet — 5 stars
('af000001-0000-0000-0000-000000000009',
 'ae000001-0000-0000-0000-00000000000e',
 'cc000001-0000-0000-0000-000000000002',
 5,
 'Ploy has been dancing around the house every day since this session. Ajarn Plernpit is phenomenal — so precise with corrections but so encouraging at the same time. Ploy is naturally shy and she came out of this session beaming with confidence. The facility is beautiful, the instruction is world-class, and the love for ballet in that room is contagious.',
 'พลอยเต้นรำอยู่รอบบ้านทุกวันตั้งแต่เซสชั่นนี้ อาจารย์เพลินพิทยอดเยี่ยมมาก — แม่นยำมากในการแก้ไขแต่ก็ให้กำลังใจในเวลาเดียวกัน พลอยขี้อายโดยธรรมชาติและเธอออกมาจากเซสชั่นนี้อย่างมีความมั่นใจ',
 ARRAY['welcoming', 'engaging', 'educational', 'clear instruction', 'inclusive']::review_category[]),

-- 10. STEM Camp — 5 stars
('af000001-0000-0000-0000-00000000000a',
 'ae000001-0000-0000-0000-00000000000f',
 'cc000001-0000-0000-0000-000000000002',
 5,
 'Geniuses STEM absolutely lived up to its name. Ploy''s team built a solar-powered water pump and presented it to a panel of "judges" at the end of the day. She was so proud and I honestly could not believe how much she learned in a single session. Dr. Kasem''s curriculum is exceptional — it feels like a mini university for kids. Already booked the next session.',
 'Geniuses STEM มีชีวิตอยู่ตามชื่อของมันอย่างแน่นอน ทีมของพลอยสร้างปั๊มน้ำพลังงานแสงอาทิตย์และนำเสนอต่อคณะกรรมการตัดสินในตอนท้ายของวัน เธอภาคภูมิใจมากและฉันแทบไม่เชื่อว่าเธอเรียนรู้ได้มากขนาดนี้ในเซสชั่นเดียว',
 ARRAY['educational', 'engaging', 'immersive', 'well structured', 'informative']::review_category[]),

-- 11. Robotics Workshop BKK — 4 stars
('af000001-0000-0000-0000-00000000000b',
 'ae000001-0000-0000-0000-000000000010',
 'cc000001-0000-0000-0000-000000000002',
 4,
 'Ploy loved every minute of the hands-on robot building. Dr. Lee''s teaching style is very clear and she made the programming concepts accessible to kids who had never coded before. Ploy''s robot was navigating a simple maze by the end! My only feedback is that the room was quite loud which made it a little hard for some kids to hear. Still a fantastic program overall.',
 'พลอยรักทุกนาทีของการสร้างหุ่นยนต์ภาคปฏิบัติ สไตล์การสอนของ Dr. Lee ชัดเจนมากและเธอทำให้แนวคิดการเขียนโปรแกรมเข้าถึงได้สำหรับเด็กที่ไม่เคยเขียนโค้ดมาก่อน หุ่นยนต์ของพลอยกำลังนำทางในเขาวงกตเรียบง่ายในตอนท้าย',
 ARRAY['educational', 'engaging', 'interactive', 'informative']::review_category[]),

-- 12. Chemistry BKK — 5 stars
('af000001-0000-0000-0000-00000000000c',
 'ae000001-0000-0000-0000-000000000011',
 'cc000001-0000-0000-0000-000000000002',
 5,
 'The experiments were absolutely safe but the results were dramatic enough to capture every child''s imagination. Ploy made slime, grew crystals, and conducted a chromatography experiment to separate ink pigments. She explained the whole process to her grandmother that evening without any prompting. The instructors clearly love what they do and it shows in how curious the kids become.',
 'การทดลองปลอดภัยอย่างแน่นอนแต่ผลลัพธ์น่าตื่นเต้นพอที่จะดึงดูดจินตนาการของเด็กทุกคน พลอยทำสไลม์ ปลูกคริสตัล และทำการทดลองโครมาโตกราฟีเพื่อแยกสีหมึก เธออธิบายกระบวนการทั้งหมดให้คุณยายฟังในตอนเย็นโดยไม่ต้องบอกให้ทำ',
 ARRAY['educational', 'engaging', 'immersive', 'informative', 'interactive']::review_category[]),

-- 13. Muay Thai Beginner — 5 stars
('af000001-0000-0000-0000-00000000000d',
 'ae000001-0000-0000-0000-000000000012',
 'cc000001-0000-0000-0000-000000000002',
 5,
 'Ton found his confidence through this program. He was reluctant to try Muay Thai at first — too shy — but Kru Somchai has this incredible ability to make every child feel brave and capable. The Wai Kru ceremony was beautiful and instilled a real sense of tradition and respect. Ton now greets adults with a Wai every morning. That alone is worth the price of admission.',
 'ต้นพบความมั่นใจของเขาผ่านโปรแกรมนี้ เขาลังเลที่จะลองมวยไทยในตอนแรก — ขี้อายเกินไป — แต่ครูสมชายมีความสามารถที่น่าทึ่งในการทำให้เด็กทุกคนรู้สึกกล้าหาญและมีความสามารถ พิธีไหว้ครูสวยงามและปลูกฝังความรู้สึกของประเพณีและการเคารพอย่างแท้จริง',
 ARRAY['welcoming', 'inclusive', 'engaging', 'educational', 'fun']::review_category[]),

-- 14. Soccer Skills BKK — 4 stars
('af000001-0000-0000-0000-00000000000e',
 'ae000001-0000-0000-0000-000000000013',
 'cc000001-0000-0000-0000-000000000002',
 4,
 'Ton is always excited for Saturdays now. He''s improved so much in just two months. The coaches are attentive and manage the group really well — 20 kids is a lot but it never felt chaotic. I gave four stars instead of five only because the field can get quite hot by midday and I''d love to see the session time moved 30 minutes earlier in the summer months.',
 'ต้นตื่นเต้นสำหรับวันเสาร์ทุกสัปดาห์ตอนนี้ เขาพัฒนาขึ้นมากในเพียงสองเดือน โค้ชใส่ใจและจัดการกลุ่มได้ดีมาก ฉันให้ 4 ดาวแทนที่จะเป็น 5 ดาวเพราะสนามค่อนข้างร้อนในช่วงเที่ยงวัน',
 ARRAY['engaging', 'fun', 'educational', 'welcoming']::review_category[]),


-- ════════════════════════════════════════════════════════════════════════════
-- MARCUS WEBB
-- ════════════════════════════════════════════════════════════════════════════

-- 15. Young Engineers — 5 stars
('af000001-0000-0000-0000-00000000000f',
 'ae000001-0000-0000-0000-000000000018',
 'cc000001-0000-0000-0000-000000000003',
 5,
 'Zoe is already asking when the next session is. As someone who works in tech myself I was curious to see how MIT Kids Lab would translate university-level thinking to an 11-year-old, and I was blown away. They don''t talk down to the kids at all — they give them real constraints, real tools, and trust them to figure it out. Zoe''s team designed a motorized lift mechanism and she can explain every engineering decision they made. Exceptional.',
 'โซอี้ถามแล้วว่าเซสชั่นถัดไปจะเมื่อไหร่ ในฐานะที่ทำงานด้านเทคโนโลยีเองฉันอยากรู้ว่า MIT Kids Lab จะแปลการคิดระดับมหาวิทยาลัยให้กับเด็กอายุ 11 ปีได้อย่างไร และฉันตะลึง พวกเขาไม่พูดโดยไม่นับถือเด็กเลย',
 ARRAY['educational', 'engaging', 'immersive', 'well structured', 'interactive']::review_category[]),

-- 16. Robotics Lab — 5 stars
('af000001-0000-0000-0000-000000000010',
 'ae000001-0000-0000-0000-000000000019',
 'cc000001-0000-0000-0000-000000000003',
 5,
 'Outstanding instructors and perfect challenge level for Zoe. She''s done block coding before but never soldered a circuit or programmed a microcontroller. By the end she had a robot that avoided obstacles using an ultrasonic sensor — she built and programmed it herself. Jake''s ability to teach Arduino C to a kid who''s never seen C before is genuinely impressive. This is the program I wish had existed when I was growing up.',
 'ครูผู้สอนที่โดดเด่นและระดับความท้าทายที่สมบูรณ์แบบสำหรับโซอี้ ในตอนท้ายเธอมีหุ่นยนต์ที่หลีกเลี่ยงสิ่งกีดขวางโดยใช้เซ็นเซอร์อัลตราโซนิก เธอสร้างและเขียนโปรแกรมมันเอง นี่คือโปรแกรมที่ฉันอยากให้มีตั้งแต่ตอนที่ฉันยังเด็ก',
 ARRAY['educational', 'engaging', 'immersive', 'well structured', 'challenging']::review_category[]),

-- 17. Game Design — 5 stars
('af000001-0000-0000-0000-000000000011',
 'ae000001-0000-0000-0000-00000000001a',
 'cc000001-0000-0000-0000-000000000003',
 5,
 'Zoe designed her first game concept from scratch and shipped it by 3pm. She''d never touched Unity before and left with a working 2D platformer with custom sprites, sound effects, and three levels. The instructors balance free creative time with structured guidance beautifully. When I picked her up she immediately pulled out her phone to send the game to her friends. Cannot recommend this highly enough.',
 'โซอี้ออกแบบแนวคิดเกมแรกของเธอตั้งแต่ต้นและส่งออกมาภายในบ่ายสามโมง เธอไม่เคยแตะ Unity มาก่อนและออกมาพร้อมกับเกม 2D platformer ที่ทำงานได้ด้วย sprites ที่กำหนดเอง เอฟเฟกต์เสียง และสามเลเวล',
 ARRAY['educational', 'engaging', 'immersive', 'interactive', 'well structured']::review_category[]),

-- 18. Watercolor — 4 stars
('af000001-0000-0000-0000-000000000012',
 'ae000001-0000-0000-0000-00000000001b',
 'cc000001-0000-0000-0000-000000000003',
 4,
 'Nice balance of free expression and guided technique. Maria is a working artist and it shows — she has a great eye and explains composition and color theory in ways that actually stick. Zoe loved it and wants to go back. I''m giving four stars because the session felt a bit short for what they were trying to accomplish; an extra 30–45 minutes would give the kids more time to develop their pieces. But the quality of instruction is excellent.',
 'สมดุลที่ดีระหว่างการแสดงออกอย่างอิสระและเทคนิคที่มีการแนะนำ Maria เป็นศิลปินที่ทำงานจริงและมันเห็นได้ชัด โซอี้ชอบและต้องการกลับไปอีก ฉันให้ 4 ดาวเพราะเซสชั่นรู้สึกสั้นเล็กน้อย',
 ARRAY['educational', 'engaging', 'fun', 'informative']::review_category[]),

-- 19. Comic Art — 5 stars
('af000001-0000-0000-0000-000000000013',
 'ae000001-0000-0000-0000-00000000001c',
 'cc000001-0000-0000-0000-000000000003',
 5,
 'Zoe filled up a whole notebook with new characters on the drive home. She''d always doodled but this session taught her actual comic structure — how to panel, how to use gutters for time, how to make a face convey emotion with four lines. She''s been writing a graphic novel ever since. The instructor''s knowledge of comic history (Moebius to Persepolis) made the session feel genuinely educational, not just arts and crafts.',
 'โซอี้เติมสมุดโน้ตทั้งเล่มด้วยตัวละครใหม่ระหว่างการขับรถกลับบ้าน เธอวาดรูปเสมอแต่เซสชั่นนี้สอนโครงสร้างการ์ตูนจริงๆ ว่าต้องจัดเฟรมอย่างไร วิธีใช้ gutters สำหรับเวลา วิธีทำหน้าให้แสดงอารมณ์ด้วยสี่เส้น เธอเขียนนิยายภาพมาตั้งแต่นั้น',
 ARRAY['educational', 'engaging', 'immersive', 'insightful', 'interactive']::review_category[]),

-- 20. Robotics Workshop Bangkok — 4 stars (Marcus visiting Bangkok)
('af000001-0000-0000-0000-000000000014',
 'ae000001-0000-0000-0000-00000000001d',
 'cc000001-0000-0000-0000-000000000003',
 4,
 'We were visiting Bangkok for two weeks and I found this program last-minute. I''m so glad I did. The quality is genuinely comparable to what we do in Boston. Zoe slotted right in with the local kids (most sessions are in English) and had a brilliant afternoon. The only reason I''m not giving five stars is that the air conditioning wasn''t working properly on the day we went. Otherwise flawless.',
 'เราเยี่ยมชมกรุงเทพฯ เป็นเวลาสองสัปดาห์และฉันพบโปรแกรมนี้ในนาทีสุดท้าย ฉันดีใจมากที่ทำ คุณภาพเทียบเคียงได้กับสิ่งที่เราทำในบอสตันจริงๆ โซอี้เข้ากันได้ดีกับเด็กในท้องถิ่นและมีบ่ายที่ยอดเยี่ยม',
 ARRAY['educational', 'engaging', 'interactive', 'welcoming']::review_category[]),

-- 21. Science Explorers — 5 stars
('af000001-0000-0000-0000-000000000015',
 'ae000001-0000-0000-0000-00000000001e',
 'cc000001-0000-0000-0000-000000000003',
 5,
 'Zoe has done a lot of science programs but this one stands apart because it never treats science like a performance — it treats it like an actual inquiry process. The kids form hypotheses, run experiments, and revise their thinking when results are unexpected. That''s real science. By the end Zoe was genuinely curious about enzyme reactions and asked me to get her a book on biochemistry. That''s the best possible outcome.',
 'โซอี้ได้ทำโปรแกรมวิทยาศาสตร์มามากมายแต่โปรแกรมนี้โดดเด่นเพราะไม่ถือว่าวิทยาศาสตร์เป็นการแสดง มันถือว่าเป็นกระบวนการสืบค้นจริงๆ เด็กๆ ตั้งสมมติฐาน ทำการทดลอง และปรับความคิดเมื่อผลลัพธ์ไม่คาดคิด นั่นคือวิทยาศาสตร์จริงๆ',
 ARRAY['educational', 'immersive', 'insightful', 'informative', 'well structured']::review_category[]);
INSERT INTO saved (guardian_id, event_id) VALUES
-- Sarah Johnson
('11111111-1111-1111-1111-111111111111', '60000000-0000-0000-0000-000000000001'),
('11111111-1111-1111-1111-111111111111', '60000000-0000-0000-0000-000000000003'),

-- Michael Chen
('22222222-2222-2222-2222-222222222222', '60000000-0000-0000-0000-000000000001'),
('22222222-2222-2222-2222-222222222222', '60000000-0000-0000-0000-000000000002'),

-- Priya Patel
('33333333-3333-3333-3333-333333333333', '60000000-0000-0000-0000-000000000003'),
('33333333-3333-3333-3333-333333333333', '60000000-0000-0000-0000-000000000004'),

-- Carlos Rodriguez
('44444444-4444-4444-4444-444444444444', '60000000-0000-0000-0000-000000000001'),

-- Emma Thompson
('55555555-5555-5555-5555-555555555555', '60000000-0000-0000-0000-000000000002'),
('55555555-5555-5555-5555-555555555555', '60000000-0000-0000-0000-000000000004'),

-- Yuki Tanaka
('66666666-6666-6666-6666-666666666666', '60000000-0000-0000-0000-000000000003'),

-- Olivia Martinez
('77777777-7777-7777-7777-777777777777', '60000000-0000-0000-0000-000000000001'),
('77777777-7777-7777-7777-777777777777', '60000000-0000-0000-0000-000000000004'),

-- James Wilson
('88888888-8888-8888-8888-888888888888', '60000000-0000-0000-0000-000000000002');-- ============================================
-- 24. SAVED EVENTS — Demo Guardian Wishlists
-- ============================================
INSERT INTO saved (guardian_id, event_id) VALUES

-- ── Jennifer Park ────────────────────────────────────────────────────────────
-- Saves tech + art events she's eyeing for both kids
('cc000001-0000-0000-0000-000000000001', '6b000000-0000-0000-0000-000000000001'), -- Young Engineers
('cc000001-0000-0000-0000-000000000001', '6b000000-0000-0000-0000-000000000003'), -- App Inventor
('cc000001-0000-0000-0000-000000000001', '6b000000-0000-0000-0000-000000000005'), -- Basketball
('cc000001-0000-0000-0000-000000000001', '6b000000-0000-0000-0000-00000000000c'), -- Comic Art
('cc000001-0000-0000-0000-000000000001', '6b000000-0000-0000-0000-000000000007'), -- Piano Foundations (thinking about Max)
('cc000001-0000-0000-0000-000000000001', '6b000000-0000-0000-0000-000000000014'), -- STEM Camp Bangkok (upcoming trip?)
('cc000001-0000-0000-0000-000000000001', '6b000000-0000-0000-0000-000000000015'), -- Drone Programming Bangkok
('cc000001-0000-0000-0000-000000000001', '60000000-0000-0000-0000-000000000001'), -- Robotics Workshop BKK (original)

-- ── Nattaya Srisuk ───────────────────────────────────────────────────────────
-- Saves Boston events (planning a visit) + more Bangkok options
('cc000001-0000-0000-0000-000000000002', '6b000000-0000-0000-0000-000000000012'), -- Classical Ballet
('cc000001-0000-0000-0000-000000000002', '6b000000-0000-0000-0000-000000000013'), -- Contemporary Dance
('cc000001-0000-0000-0000-000000000002', '6b000000-0000-0000-0000-000000000014'), -- STEM Camp
('cc000001-0000-0000-0000-000000000002', '6b000000-0000-0000-0000-000000000001'), -- Young Engineers (for Ton)
('cc000001-0000-0000-0000-000000000002', '60000000-0000-0000-0000-00000000000a'), -- Piano for Beginners (existing BKK)
('cc000001-0000-0000-0000-000000000002', '60000000-0000-0000-0000-000000000001'), -- Robotics Workshop (existing BKK)
('cc000001-0000-0000-0000-000000000002', '6b000000-0000-0000-0000-000000000009'), -- Youth Orchestra (Boston visit)

-- ── Marcus Webb ──────────────────────────────────────────────────────────────
-- Heavy saver — researches everything before committing
('cc000001-0000-0000-0000-000000000003', '6b000000-0000-0000-0000-000000000001'), -- Young Engineers
('cc000001-0000-0000-0000-000000000003', '6b000000-0000-0000-0000-000000000002'), -- Science Explorers
('cc000001-0000-0000-0000-000000000003', '6b000000-0000-0000-0000-000000000003'), -- App Inventor
('cc000001-0000-0000-0000-000000000003', '6b000000-0000-0000-0000-00000000000f'), -- Robotics Lab
('cc000001-0000-0000-0000-000000000003', '6b000000-0000-0000-0000-00000000000e'), -- Game Design
('cc000001-0000-0000-0000-000000000003', '6b000000-0000-0000-0000-00000000000c'), -- Comic Art
('cc000001-0000-0000-0000-000000000003', '6b000000-0000-0000-0000-00000000000d'), -- Web Design
('cc000001-0000-0000-0000-000000000003', '6b000000-0000-0000-0000-000000000007'), -- Piano Foundations (diversifying)
('cc000001-0000-0000-0000-000000000003', '6b000000-0000-0000-0000-000000000014'), -- STEM Camp Bangkok
('cc000001-0000-0000-0000-000000000003', '6b000000-0000-0000-0000-000000000015'), -- Drone Programming
('cc000001-0000-0000-0000-000000000003', '60000000-0000-0000-0000-000000000003'), -- Astronomy Club (existing BKK)
('cc000001-0000-0000-0000-000000000003', '60000000-0000-0000-0000-00000000000f'); -- 3D Modeling (existing BKK)
