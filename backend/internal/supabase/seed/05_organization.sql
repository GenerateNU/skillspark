-- ============================================
-- 6. ORGANIZATIONS
-- ============================================
INSERT INTO organization (id, name, active, pfp_s3_key, location_id, about, links) VALUES
('40000000-0000-0000-0000-000000000001', 'Science Academy Bangkok', true, 'orgs/science_academy.jpg', '10000000-0000-0000-0000-000000000004',
 'Science Academy Bangkok offers hands-on STEM programs for children aged 5–15, fostering curiosity and critical thinking through experiments, robotics, and coding workshops.',
 '[{"href": "https://www.scienceacademybkk.com", "label": "Website"}, {"href": "https://www.facebook.com/scienceacademybkk", "label": "Facebook"}, {"href": "https://www.instagram.com/scienceacademybkk", "label": "Instagram"}]'::jsonb),
('40000000-0000-0000-0000-000000000002', 'Champions Sports Center', true, 'orgs/champions_sports.jpg', '10000000-0000-0000-0000-000000000005',
 'Champions Sports Center provides youth sports training in soccer, basketball, and swimming. Our certified coaches build skills, teamwork, and confidence in athletes of all levels.',
 '[{"href": "https://www.championssportsbkk.com", "label": "Website"}, {"href": "https://www.instagram.com/championssportsbkk", "label": "Instagram"}]'::jsonb),
('40000000-0000-0000-0000-000000000003', 'Creative Arts Studio', true, 'orgs/creative_arts.jpg', '10000000-0000-0000-0000-000000000006',
 'Creative Arts Studio is a vibrant space for young artists to explore painting, sculpture, and mixed media. We inspire self-expression and creativity in children from age 4 and up.',
 '[{"href": "https://www.creativeartstudio.th", "label": "Website"}, {"href": "https://www.facebook.com/creativeartstudiobkk", "label": "Facebook"}, {"href": "https://www.instagram.com/creativeartstudio_bkk", "label": "Instagram"}]'::jsonb),
('40000000-0000-0000-0000-000000000004', 'Harmony Music School', true, NULL, '10000000-0000-0000-0000-000000000007',
 'Harmony Music School offers individual and group lessons in piano, guitar, violin, and voice. Our experienced instructors guide students from beginner to advanced levels in a nurturing environment.',
 '[{"href": "https://www.harmonymusic.th", "label": "Website"}, {"href": "https://www.facebook.com/harmonymusicbkk", "label": "Facebook"}]'::jsonb),
('40000000-0000-0000-0000-000000000005', 'Tech Kids Workshop', true, 'orgs/tech_kids.jpg', '10000000-0000-0000-0000-000000000008',
 'Tech Kids Workshop teaches children programming, game design, and electronics through project-based learning. We prepare the next generation of innovators with practical technology skills.',
 '[{"href": "https://www.techkidsworkshop.com", "label": "Website"}, {"href": "https://www.instagram.com/techkidsworkshop", "label": "Instagram"}, {"href": "https://www.youtube.com/@techkidsworkshop", "label": "YouTube"}]'::jsonb),
('40000000-0000-0000-0000-000000000006', 'Language Learning Center', false, NULL, '10000000-0000-0000-0000-000000000009',
 'Language Learning Center offers immersive language programs in English, Mandarin, and Japanese for children and teens. Our communicative approach builds fluency and cultural understanding.',
 '[{"href": "https://www.languagelearningcenter.th", "label": "Website"}, {"href": "https://www.facebook.com/languagelearningcenterbkk", "label": "Facebook"}]'::jsonb);
