-- ============================================
-- 8. EVENTS
-- ============================================
INSERT INTO event (id, title_en, title_th, description_en, description_th, organization_id, age_range_min, age_range_max, category, header_image_s3_key) VALUES

-- ── Science Academy Bangkok ──────────────────────────────────────────────────
('60000000-0000-0000-0000-000000000001',
 'Junior Robotics Workshop', 'เวิร์คช็อปหุ่นยนต์สำหรับเด็ก',
 'Learn the basics of robotics with hands-on LEGO Mindstorms projects. Children design, build, and program their own robots to complete real-world challenges. No prior experience required — just curiosity!',
 'เรียนรู้พื้นฐานหุ่นยนต์ด้วยโครงการ LEGO Mindstorms แบบลงมือทำ เด็กๆ จะได้ออกแบบ สร้าง และเขียนโปรแกรมหุ่นยนต์ของตัวเองเพื่อรับมือกับความท้าทายในชีวิตจริง ไม่จำเป็นต้องมีประสบการณ์มาก่อน',
 '40000000-0000-0000-0000-000000000001', 8, 12, ARRAY['science','technology']::category[], 'events/robotics_workshop.jpg'),

('60000000-0000-0000-0000-000000000002',
 'Chemistry for Kids', 'เคมีสำหรับเด็ก',
 'Exciting and safe chemistry experiments that spark a love for science. Children discover chemical reactions, make slime and bath bombs, grow crystals, and learn how everyday materials behave at the molecular level.',
 'การทดลองเคมีที่น่าตื่นเต้นและปลอดภัย เด็กๆ จะค้นพบปฏิกิริยาเคมี ทำสไลม์และบาธบอมบ์ เพาะผลึก และเรียนรู้ว่าวัสดุในชีวิตประจำวันทำงานอย่างไรในระดับโมเลกุล',
 '40000000-0000-0000-0000-000000000001', 7, 10, ARRAY['science']::category[], 'events/chemistry_kids.jpg'),

('60000000-0000-0000-0000-000000000003',
 'Astronomy Club', 'ชมรมดาราศาสตร์',
 'Explore the wonders of the universe! Children learn about planets, stars, black holes, and galaxies through interactive lessons and hands-on model building. Evening sessions include live telescope observation and star-chart navigation.',
 'สำรวจความมหัศจรรย์ของจักรวาล! เด็กๆ จะเรียนรู้เกี่ยวกับดาวเคราะห์ ดาวฤกษ์ หลุมดำ และกาแล็กซี ผ่านบทเรียนแบบโต้ตอบและการสร้างโมเดลลงมือทำ กิจกรรมตอนเย็นรวมถึงการสังเกตการณ์ด้วยกล้องโทรทรรศน์สดและการอ่านแผนที่ดาว',
 '40000000-0000-0000-0000-000000000001', 9, 14, ARRAY['science']::category[], 'events/astronomy_club.jpg'),

-- ── Champions Sports Center ──────────────────────────────────────────────────
('60000000-0000-0000-0000-000000000004',
 'Soccer Skills Training', 'ฝึกทักษะฟุตบอล',
 'Develop fundamental soccer skills in a fun, high-energy environment. Certified coaches cover dribbling, first touch, passing, positioning, and small-sided games. All sessions end with a scrimmage to put skills into practice.',
 'พัฒนาทักษะฟุตบอลพื้นฐานในบรรยากาศที่สนุกสนานและเต็มพลังงาน โค้ชที่ผ่านการรับรองจะสอนการเลี้ยงบอล การสัมผัสบอลครั้งแรก การส่งบอล ตำแหน่ง และเกมขนาดเล็ก ทุกเซสชั่นจบด้วยการแข่งขันเพื่อฝึกทักษะ',
 '40000000-0000-0000-0000-000000000002', 6, 12, ARRAY['sports']::category[], 'events/soccer_training.jpg'),

('60000000-0000-0000-0000-000000000005',
 'Basketball Basics', 'พื้นฐานบาสเกตบอล',
 'Learn basketball from the ground up. Sessions cover ball handling, shooting form, defensive footwork, pick-and-roll concepts, and game strategy. Players develop both individual skills and team play in a supportive, competitive environment.',
 'เรียนรู้บาสเกตบอลตั้งแต่พื้นฐาน เซสชั่นครอบคลุมการจับบอล ท่ายิง การเคลื่อนไหวป้องกัน แนวคิดปิ้กแอนด์โรล และกลยุทธ์เกม นักกีฬาพัฒนาทักษะส่วนบุคคลและการเล่นเป็นทีมในสภาพแวดล้อมที่สนับสนุนและแข่งขัน',
 '40000000-0000-0000-0000-000000000002', 7, 13, ARRAY['sports']::category[], 'events/basketball_basics.jpg'),

('60000000-0000-0000-0000-000000000006',
 'Swimming Lessons', 'บทเรียนว่ายน้ำ',
 'Professional swimming instruction from beginner to intermediate levels. Certified coaches teach freestyle, backstroke, and breaststroke technique, flip turns, and open-water awareness. Class sizes are kept small (max 8) for maximum individual attention.',
 'การสอนว่ายน้ำอย่างมืออาชีพตั้งแต่ผู้เริ่มต้นถึงระดับกลาง โค้ชที่ผ่านการรับรองสอนเทคนิคฟรีสไตล์ กรรเชียง และกบ การกลับตัว และความตระหนักในน้ำเปิด ชั้นเรียนมีขนาดเล็ก (สูงสุด 8 คน) เพื่อความสนใจส่วนบุคคลสูงสุด',
 '40000000-0000-0000-0000-000000000002', 5, 15, ARRAY['sports']::category[], 'events/swimming_lessons.jpg'),

-- ── Creative Arts Studio ─────────────────────────────────────────────────────
('60000000-0000-0000-0000-000000000007',
 'Painting & Drawing Workshop', 'เวิร์คช็อปวาดภาพและระบายสี',
 'A multi-media art workshop exploring watercolor, acrylic, gouache, and pencil sketching. Children work on structured projects each session while developing their personal style. All materials are provided. Completed works are framed and sent home.',
 'เวิร์คช็อปศิลปะมัลติมีเดียที่สำรวจสีน้ำ อะคริลิค กัวช และการสเก็ตช์ดินสอ เด็กๆ ทำงานในโครงการที่มีโครงสร้างในแต่ละเซสชั่นขณะพัฒนาสไตล์ส่วนตัว อุปกรณ์ทั้งหมดมีให้ ผลงานที่เสร็จสมบูรณ์จะถูกใส่กรอบและส่งกลับบ้าน',
 '40000000-0000-0000-0000-000000000003', 6, 14, ARRAY['art']::category[], 'events/painting_workshop.jpg'),

('60000000-0000-0000-0000-000000000008',
 'Pottery for Beginners', 'เครื่องปั้นดินเผาสำหรับผู้เริ่มต้น',
 'Discover the joy of working with clay! Children learn hand-building techniques (pinch, coil, slab) to create functional bowls, cups, and decorative sculptures. Finished pieces are kiln-fired and glazed. A take-home creation is guaranteed every session.',
 'ค้นพบความสุขของการทำงานกับดินเหนียว! เด็กๆ เรียนรู้เทคนิคการสร้างด้วยมือ (บีบ ม้วน แผ่น) เพื่อสร้างชาม ถ้วย และประติมากรรมตกแต่ง ชิ้นงานที่เสร็จแล้วจะถูกเผาและเคลือบ รับประกันผลงานที่สามารถนำกลับบ้านได้ทุกเซสชั่น',
 '40000000-0000-0000-0000-000000000003', 8, 15, ARRAY['art']::category[], 'events/pottery_beginners.jpg'),

('60000000-0000-0000-0000-000000000009',
 'Digital Art & Design', 'ศิลปะดิจิทัลและการออกแบบ',
 'Introduction to digital illustration using Wacom tablets and Procreate. Children learn composition, color theory, layering, and digital brush techniques. By the end of the program, each student creates a full illustrated character with background.',
 'แนะนำการวาดภาพดิจิทัลโดยใช้แท็บเล็ต Wacom และ Procreate เด็กๆ เรียนรู้การจัดองค์ประกอบ ทฤษฎีสี การแบ่งชั้น และเทคนิคพู่กันดิจิทัล เมื่อสิ้นสุดโปรแกรม นักเรียนแต่ละคนจะสร้างตัวละครที่มีภาพประกอบเต็มพร้อมพื้นหลัง',
 '40000000-0000-0000-0000-000000000003', 10, 16, ARRAY['art','technology']::category[], 'events/digital_art_design.jpg'),

-- ── Harmony Music School ─────────────────────────────────────────────────────
('60000000-0000-0000-0000-00000000000a',
 'Piano for Beginners', 'เปียโนสำหรับผู้เริ่มต้น',
 'Start your musical journey at the piano! Students learn proper hand position, posture, note reading, and simple songs from the classical and popular repertoire. Lessons use the Alfred Premier Piano Course method. A digital keyboard is available to borrow for home practice.',
 'เริ่มต้นการเดินทางทางดนตรีที่เปียโน! นักเรียนเรียนรู้ตำแหน่งมือที่ถูกต้อง ท่าทาง การอ่านโน้ต และเพลงง่ายๆ จากคลังเพลงคลาสสิกและยอดนิยม บทเรียนใช้วิธีการ Alfred Premier Piano Course มีคีย์บอร์ดดิจิทัลให้ยืมสำหรับฝึกที่บ้าน',
 '40000000-0000-0000-0000-000000000004', 6, 12, ARRAY['music']::category[], 'events/piano_beginners.jpg'),

('60000000-0000-0000-0000-00000000000b',
 'Guitar Fundamentals', 'พื้นฐานกีตาร์',
 'Learn acoustic guitar from the ground up. Students master open chords, basic strumming patterns, fingerpicking, and song structure. By the end of the course, students can play 6–8 songs across multiple genres. Guitars are available to borrow for the session.',
 'เรียนรู้กีตาร์อะคูสติกตั้งแต่พื้นฐาน นักเรียนเรียนรู้คอร์ดเปิด รูปแบบการตีคอร์ดพื้นฐาน การดีดนิ้ว และโครงสร้างเพลง เมื่อสิ้นสุดหลักสูตร นักเรียนสามารถเล่นเพลงได้ 6–8 เพลงในหลายแนว มีกีตาร์ให้ยืมสำหรับเซสชั่น',
 '40000000-0000-0000-0000-000000000004', 8, 15, ARRAY['music']::category[], 'events/guitar_fundamentals.jpg'),

('60000000-0000-0000-0000-00000000000c',
 'Kids Choir', 'คณะนักร้องประสานเสียงเด็ก',
 'Join our award-winning children''s choir! Members learn two-part harmony, vocal warm-ups, breathing technique, and stage presence. The choir performs at school events and an annual public concert. Repertoire spans pop, folk, Broadway, and classical.',
 'ร่วมคณะนักร้องประสานเสียงเด็กที่ได้รับรางวัลของเรา! สมาชิกเรียนรู้การประสานเสียงสองส่วน การอุ่นเสียง เทคนิคการหายใจ และบุคลิกบนเวที คณะนักร้องแสดงในงานโรงเรียนและคอนเสิร์ตสาธารณะประจำปี คลังเพลงครอบคลุมป๊อป โฟล์ค บรอดเวย์ และคลาสสิก',
 '40000000-0000-0000-0000-000000000004', 7, 13, ARRAY['music']::category[], 'events/kids_choir.jpg'),

-- ── Tech Kids Workshop ───────────────────────────────────────────────────────
('60000000-0000-0000-0000-00000000000d',
 'Coding with Scratch', 'เขียนโค้ดด้วย Scratch',
 'Learn the fundamentals of programming through Scratch, MIT''s visual coding platform. Children create interactive stories, animated games, and music projects using drag-and-drop code blocks. No typing required — just logical thinking and creativity.',
 'เรียนรู้พื้นฐานการเขียนโปรแกรมผ่าน Scratch แพลตฟอร์มโค้ดภาพของ MIT เด็กๆ สร้างเรื่องราวแบบโต้ตอบ เกมเคลื่อนไหว และโครงการดนตรีโดยใช้บล็อกโค้ดแบบลากและวาง ไม่จำเป็นต้องพิมพ์ — เพียงแค่การคิดเชิงตรรกะและความคิดสร้างสรรค์',
 '40000000-0000-0000-0000-000000000005', 7, 11, ARRAY['technology']::category[], 'events/scratch_coding.jpg'),

('60000000-0000-0000-0000-00000000000e',
 'Python for Kids', 'Python สำหรับเด็ก',
 'A project-based introduction to Python programming. Students write real code from day one, building text adventures, simple web scrapers, and graphical games using Pygame. By the final session, every student ships a working project they built themselves.',
 'การแนะนำการเขียนโปรแกรม Python แบบโครงการ นักเรียนเขียนโค้ดจริงตั้งแต่วันแรก สร้างการผจญภัยข้อความ เว็บสเครปเปอร์อย่างง่าย และเกมกราฟิกโดยใช้ Pygame ภายในเซสชั่นสุดท้าย นักเรียนทุกคนจะส่งมอบโปรเจ็กต์ที่ใช้งานได้ที่ตัวเองสร้าง',
 '40000000-0000-0000-0000-000000000005', 10, 15, ARRAY['technology','math']::category[], 'events/python_kids.jpg'),

('60000000-0000-0000-0000-00000000000f',
 '3D Modeling Workshop', 'เวิร์คช็อปโมเดล 3 มิติ',
 'Design 3D objects in Tinkercad and prepare them for 3D printing. Students learn spatial reasoning, measurement, and iterative design by creating keychains, phone stands, architectural models, and their own original inventions. Finished prints are taken home.',
 'ออกแบบวัตถุ 3 มิติใน Tinkercad และเตรียมสำหรับการพิมพ์ 3 มิติ นักเรียนเรียนรู้การคิดเชิงพื้นที่ การวัด และการออกแบบซ้ำโดยการสร้างพวงกุญแจ ที่วางโทรศัพท์ โมเดลสถาปัตยกรรม และสิ่งประดิษฐ์ต้นฉบับของตัวเอง ผลงานพิมพ์ที่เสร็จแล้วสามารถนำกลับบ้านได้',
 '40000000-0000-0000-0000-000000000005', 9, 14, ARRAY['technology','art']::category[], 'events/3d_modeling.jpg'),

-- ── Boston STEM Lab ──────────────────────────────────────────────────────────
('60000000-0000-0000-0000-000000000010',
 'Intro to Circuits & Electronics', NULL,
 'Children explore the fundamentals of electricity and circuit design using snap-together components and breadboards. Sessions cover series and parallel circuits, LEDs, resistors, and basic sensors. Each child builds and takes home a working LED light-up badge.',
 NULL,
 '40000000-0000-0000-0000-000000000007', 7, 12, ARRAY['science','technology']::category[], 'events/intro_circuits.jpg'),

('60000000-0000-0000-0000-000000000011',
 'Marine Biology Explorer', NULL,
 'Dive into ocean science! Children study real specimens under microscopes, dissect squid, test water salinity, and learn about marine food webs and conservation threats. Sessions are led by a marine biologist and every class includes a virtual ocean dive.',
 NULL,
 '40000000-0000-0000-0000-000000000007', 8, 13, ARRAY['science']::category[], 'events/marine_biology.jpg'),

('60000000-0000-0000-0000-000000000012',
 'Science Fair Prep Workshop', NULL,
 'Guided support for students entering a school or regional science fair. Instructors help children develop an original hypothesis, design a controlled experiment, collect and analyze data, and build a presentation board. Past participants have placed in the Massachusetts State Science Fair.',
 NULL,
 '40000000-0000-0000-0000-000000000007', 10, 14, ARRAY['science','math']::category[], 'events/science_fair_prep.jpg'),

-- ── New England Dance Academy ────────────────────────────────────────────────
('60000000-0000-0000-0000-000000000013',
 'Ballet for Beginners', NULL,
 'A nurturing introduction to classical ballet for young dancers. Students learn barre exercises, basic center combinations, and simple choreography in an encouraging, non-competitive atmosphere. Class ends with a short performance sharing for families each session.',
 NULL,
 '40000000-0000-0000-0000-000000000008', 5, 10, ARRAY['art']::category[], 'events/ballet_beginners.jpg'),

('60000000-0000-0000-0000-000000000014',
 'Hip Hop Fundamentals', NULL,
 'High-energy hip hop dance classes covering foundational moves: toprock, grooves, freezes, and floor work. Instructors emphasize musicality, freestyle expression, and cypher etiquette. Each session ends with a freestyle cipher and a new routine to take home.',
 NULL,
 '40000000-0000-0000-0000-000000000008', 8, 15, ARRAY['art']::category[], 'events/hip_hop_fundamentals.jpg'),

('60000000-0000-0000-0000-000000000015',
 'Creative Movement & Rhythm', NULL,
 'Designed for the youngest dancers, this class uses storytelling, props, and music to develop coordination, spatial awareness, and rhythm. Each class has a theme — animals, seasons, space — and children create their own movements within it. A joyful, imaginative introduction to dance.',
 NULL,
 '40000000-0000-0000-0000-000000000008', 4, 7, ARRAY['art','music']::category[], 'events/creative_movement.jpg'),

-- ── Fenway Chess & Math Club ─────────────────────────────────────────────────
('60000000-0000-0000-0000-000000000016',
 'Chess for Young Minds', NULL,
 'Learn chess from an expert FIDE-rated instructor. Students master piece movement, basic tactics (forks, pins, skewers), opening principles, and endgame fundamentals. Weekly puzzles reinforce pattern recognition. Students are entered into monthly in-house tournaments.',
 NULL,
 '40000000-0000-0000-0000-000000000009', 6, 13, ARRAY['math']::category[], 'events/chess_kids.jpg'),

('60000000-0000-0000-0000-000000000017',
 'Math Olympiad Prep', NULL,
 'Intensive problem-solving preparation for the AMC 8, Math Olympiad, and MATHCOUNTS competitions. Topics include number theory, combinatorics, geometry, and algebraic thinking. Past students have qualified for state-level MATHCOUNTS and scored 25+ on the AMC 8.',
 NULL,
 '40000000-0000-0000-0000-000000000009', 9, 14, ARRAY['math']::category[], 'events/math_olympiad_prep.jpg');
