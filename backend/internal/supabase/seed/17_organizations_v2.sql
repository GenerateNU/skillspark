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
