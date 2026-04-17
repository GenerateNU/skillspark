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
