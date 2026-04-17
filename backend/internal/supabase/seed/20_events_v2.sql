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
