-- ============================================
-- 6. ORGANIZATIONS
-- ============================================
INSERT INTO organization (id, name, active, pfp_s3_key, location_id, about, links) VALUES
-- Bangkok organizations
('40000000-0000-0000-0000-000000000001', 'Science Academy Bangkok', true, 'orgs/science_academy_bkk.jpg', '10000000-0000-0000-0000-000000000004',
 'Science Academy Bangkok offers hands-on STEM programs for children aged 5–15, fostering curiosity and critical thinking through experiments, robotics, and coding workshops. Our internationally certified instructors make complex concepts accessible and exciting.',
 '[{"href": "https://www.scienceacademybkk.com", "label": "Website"}, {"href": "https://www.facebook.com/scienceacademybkk", "label": "Facebook"}, {"href": "https://www.instagram.com/scienceacademybkk", "label": "Instagram"}]'::jsonb),

('40000000-0000-0000-0000-000000000002', 'Champions Sports Center', true, 'orgs/champions_sports_bkk.jpg', '10000000-0000-0000-0000-000000000005',
 'Champions Sports Center provides elite youth sports training in soccer, basketball, and swimming. Our UEFA and FIBA certified coaches build technical skills, teamwork, and mental resilience in athletes aged 5–17.',
 '[{"href": "https://www.championssportsbkk.com", "label": "Website"}, {"href": "https://www.instagram.com/championssportsbkk", "label": "Instagram"}, {"href": "https://www.facebook.com/championssportsbkk", "label": "Facebook"}]'::jsonb),

('40000000-0000-0000-0000-000000000003', 'Creative Arts Studio', true, 'orgs/creative_arts_bkk.jpg', '10000000-0000-0000-0000-000000000006',
 'Creative Arts Studio is a vibrant space for young artists to explore painting, sculpture, pottery, and digital media. Founded in 2012, we have nurtured thousands of young creators from age 4 and up with a philosophy that every child is an artist.',
 '[{"href": "https://www.creativeartstudio.th", "label": "Website"}, {"href": "https://www.facebook.com/creativeartstudiobkk", "label": "Facebook"}, {"href": "https://www.instagram.com/creativeartstudio_bkk", "label": "Instagram"}]'::jsonb),

('40000000-0000-0000-0000-000000000004', 'Harmony Music School', true, 'orgs/harmony_music_bkk.jpg', '10000000-0000-0000-0000-000000000007',
 'Harmony Music School offers individual and group lessons in piano, guitar, violin, and voice. Our experienced instructors — many with conservatory backgrounds — guide students from beginner to advanced levels in a nurturing, performance-focused environment.',
 '[{"href": "https://www.harmonymusic.th", "label": "Website"}, {"href": "https://www.facebook.com/harmonymusicbkk", "label": "Facebook"}, {"href": "https://www.youtube.com/@harmonymusicbkk", "label": "YouTube"}]'::jsonb),

('40000000-0000-0000-0000-000000000005', 'Tech Kids Workshop', true, 'orgs/tech_kids_bkk.jpg', '10000000-0000-0000-0000-000000000008',
 'Tech Kids Workshop teaches children programming, game design, and electronics through project-based learning. We partner with leading tech companies to ensure our curriculum stays current, preparing the next generation of innovators with real-world skills.',
 '[{"href": "https://www.techkidsworkshop.com", "label": "Website"}, {"href": "https://www.instagram.com/techkidsworkshop", "label": "Instagram"}, {"href": "https://www.youtube.com/@techkidsworkshop", "label": "YouTube"}]'::jsonb),

('40000000-0000-0000-0000-000000000006', 'Language Learning Center', false, 'orgs/language_center_bkk.jpg', '10000000-0000-0000-0000-000000000009',
 'Language Learning Center offers immersive language programs in English, Mandarin, and Japanese for children and teens. Our communicative approach builds fluency and cultural understanding through storytelling, games, and conversation.',
 '[{"href": "https://www.languagelearningcenter.th", "label": "Website"}, {"href": "https://www.facebook.com/languagelearningcenterbkk", "label": "Facebook"}]'::jsonb),

-- Boston organizations
('40000000-0000-0000-0000-000000000007', 'Boston STEM Lab', true, 'orgs/boston_stem_lab.jpg', '10000000-0000-0000-0000-00000000000d',
 'Boston STEM Lab is a hands-on science and engineering enrichment center in Cambridge for children ages 6–14. Our programs are developed in collaboration with MIT researchers and cover circuits, marine biology, chemistry, and engineering design. We believe every child deserves access to world-class science education.',
 '[{"href": "https://www.bostonstemlab.com", "label": "Website"}, {"href": "https://www.instagram.com/bostonstemlab", "label": "Instagram"}, {"href": "https://www.facebook.com/bostonstemlab", "label": "Facebook"}]'::jsonb),

('40000000-0000-0000-0000-000000000008', 'New England Dance Academy', true, 'orgs/ne_dance_academy.jpg', '10000000-0000-0000-0000-00000000000e',
 'New England Dance Academy has trained young dancers in Boston''s South End since 2008. We offer ballet, hip hop, jazz, and creative movement for ages 4–18. Our faculty includes former members of the Boston Ballet and Alvin Ailey American Dance Theater.',
 '[{"href": "https://www.nedanceacademy.com", "label": "Website"}, {"href": "https://www.instagram.com/nedanceacademy", "label": "Instagram"}, {"href": "https://www.facebook.com/nedanceacademy", "label": "Facebook"}]'::jsonb),

('40000000-0000-0000-0000-000000000009', 'Fenway Chess & Math Club', true, 'orgs/fenway_chess_math.jpg', '10000000-0000-0000-0000-00000000000f',
 'Fenway Chess & Math Club develops critical thinking, pattern recognition, and problem-solving skills through chess and math enrichment. Our head instructor is a FIDE-rated master with 15 years of youth coaching experience. We have produced multiple state chess champions and Math Olympiad qualifiers.',
 '[{"href": "https://www.fenwaychessmath.com", "label": "Website"}, {"href": "https://www.instagram.com/fenwaychessmath", "label": "Instagram"}]'::jsonb);
