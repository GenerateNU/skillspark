-- ============================================
-- 11. REVIEWS
-- ============================================
INSERT INTO review (
    id, registration_id, guardian_id, rating, description_en, description_th, categories
) VALUES

-- ══════════════════════════════════════════════════════════════════════════════
-- SARAH JOHNSON — Bangkok past events
-- ══════════════════════════════════════════════════════════════════════════════
('90000000-0000-0000-0000-000000000001', '80000000-0000-0000-0000-000000000001', '11111111-1111-1111-1111-111111111111', 5,
 'Emily absolutely loved this workshop. She came home beaming, immediately explaining how she programmed her robot to navigate a maze. The instructors were incredibly patient and matched their pace to each child. We signed her up for the next session the same evening.',
 'เอมิลี่ชอบเวิร์คช็อปนี้มาก เธอกลับบ้านด้วยรอยยิ้ม อธิบายทันทีว่าเธอเขียนโปรแกรมหุ่นยนต์ให้นำทางผ่านเขาวงกตได้อย่างไร ครูผู้สอนใจเย็นมากและปรับจังหวะให้เข้ากับเด็กแต่ละคน เราสมัครให้เธอเรียนรอบถัดไปในคืนเดียวกัน',
 ARRAY['informative', 'interesting', 'engaging']::review_category[]),

('90000000-0000-0000-0000-000000000002', '80000000-0000-0000-0000-000000000002', '11111111-1111-1111-1111-111111111111', 4,
 'Really well-organized class with age-appropriate experiments. Emily made slime and grew a copper sulfate crystal — she hasn''t stopped talking about it. My only minor note is that the session ran 10 minutes over. Will definitely be back.',
 'ชั้นเรียนที่จัดได้ดีมากพร้อมการทดลองที่เหมาะสมกับอายุ เอมิลี่ทำสไลม์และปลูกผลึกคอปเปอร์ซัลเฟต เธอไม่หยุดพูดถึงมัน จะกลับมาแน่นอน',
 ARRAY['fun', 'informative']::review_category[]),

('90000000-0000-0000-0000-000000000003', '80000000-0000-0000-0000-000000000003', '11111111-1111-1111-1111-111111111111', 5,
 'Emily built an animated story about a cat who goes to space — she was so proud. The instructor explained loops and conditionals in a way that clicked for a 9-year-old. The small class size made it easy for everyone to get personal help. Highly recommend.',
 'เอมิลี่สร้างเรื่องราวแอนิเมชั่นเกี่ยวกับแมวที่ไปอวกาศ เธอภูมิใจมาก ครูผู้สอนอธิบายลูปและเงื่อนไขในแบบที่เข้าใจได้จริงสำหรับเด็กอายุ 9 ปี แนะนำอย่างยิ่ง',
 ARRAY['engaging', 'interesting']::review_category[]),

('90000000-0000-0000-0000-000000000004', '80000000-0000-0000-0000-000000000004', '11111111-1111-1111-1111-111111111111', 4,
 'Alex is 7 and came in not knowing much beyond kicking the ball. He left knowing how to trap, turn, and pass with purpose. Coach Marcus has a gift for keeping young kids focused without being strict.',
 'อเล็กซ์อายุ 7 ปีและมาโดยไม่รู้อะไรมากนอกจากการเตะบอล เขาออกไปพร้อมกับรู้วิธีหยุดบอล หมุน และส่งบอลอย่างมีจุดประสงค์ โค้ชมาร์คัสมีพรสวรรค์ในการทำให้เด็กเล็กมีสมาธิ',
 ARRAY['fun', 'engaging']::review_category[]),

('90000000-0000-0000-0000-000000000005', '80000000-0000-0000-0000-000000000005', '11111111-1111-1111-1111-111111111111', 5,
 'Alex surprised me completely — I wasn''t sure he had the patience for piano, but he sat focused the entire hour. The teacher introduced notes through nursery rhymes Alex already knew, which was genius. He''s been playing "Hot Cross Buns" on our keyboard every day since.',
 'อเล็กซ์ทำให้ฉันประหลาดใจอย่างสิ้นเชิง เขานั่งโฟกัสตลอดทั้งชั่วโมง ครูแนะนำโน้ตผ่านเพลงที่อเล็กซ์รู้จักอยู่แล้ว ซึ่งเป็นแนวคิดที่ยอดเยี่ยม คุ้มค่ามาก',
 ARRAY['informative', 'engaging', 'fun']::review_category[]),

('90000000-0000-0000-0000-000000000006', '80000000-0000-0000-0000-000000000006', '11111111-1111-1111-1111-111111111111', 5,
 'The choir session was the highlight of our January. Alex sang his heart out. The director taught the kids a two-part harmony in 90 minutes. The mini-performance at the end for parents had me in tears. Already booked the next one.',
 'เซสชั่นคณะนักร้องประสานเสียงเป็นไฮไลต์ของเดือนมกราคมของเรา อเล็กซ์ร้องเต็มที่ ผู้อำนวยการสอนเด็กๆ การประสานสองส่วนภายใน 90 นาที การแสดงขนาดเล็กสำหรับผู้ปกครองทำให้ฉันน้ำตาไหล',
 ARRAY['fun', 'engaging', 'interesting']::review_category[]),

('90000000-0000-0000-0000-000000000007', '80000000-0000-0000-0000-000000000007', '11111111-1111-1111-1111-111111111111', 4,
 'Alex went in as a fearful swimmer who only wanted to hold the wall, and came out doing basic freestyle across the pool. The coach was calm and never pushed too hard — big deal for an anxious kid. Great program overall.',
 'อเล็กซ์เข้าไปในฐานะนักว่ายน้ำที่กลัวซึ่งต้องการจับกำแพงเท่านั้น และออกมาว่ายน้ำฟรีสไตล์พื้นฐานข้ามสระ โค้ชสงบและไม่เคยกดดันมากเกินไป โปรแกรมที่ยอดเยี่ยมโดยรวม',
 ARRAY['informative', 'engaging']::review_category[]),

-- ── Sarah: Boston past events ────────────────────────────────────────────────
('90000000-0000-0000-0000-000000000008', '80000000-0000-0000-0000-000000000009', '11111111-1111-1111-1111-111111111111', 5,
 'Boston STEM Lab exceeded every expectation. Emily came home with a working LED badge she built from scratch and a notebook full of circuit diagrams she drew herself. Jennifer and her team clearly love what they do — the energy in that room is infectious.',
 NULL, ARRAY['informative', 'interesting', 'engaging']::review_category[]),

('90000000-0000-0000-0000-000000000009', '80000000-0000-0000-0000-00000000000a', '11111111-1111-1111-1111-111111111111', 5,
 'Emily dissected a squid, looked at plankton under a microscope, and spent the whole drive home explaining the food web to me. The instructor is an actual marine biologist. You can hear the genuine passion for the ocean in every sentence.',
 NULL, ARRAY['informative', 'interesting']::review_category[]),

('90000000-0000-0000-0000-00000000000a', '80000000-0000-0000-0000-00000000000b', '11111111-1111-1111-1111-111111111111', 4,
 'Emily had never played chess and was a little intimidated. Patrick spent the first 20 minutes on the fun history of chess, which got the kids invested before learning a single rule. By the end she beat me at home. My only note: the chairs are a bit small for older kids.',
 NULL, ARRAY['informative', 'engaging']::review_category[]),

('90000000-0000-0000-0000-00000000000b', '80000000-0000-0000-0000-00000000000c', '11111111-1111-1111-1111-111111111111', 5,
 'We signed Emily up on a whim and she came home asking to do it every week. The problem-solving approach is completely different from school math — more creative, more playful, but absolutely rigorous. Patrick has a rare ability to make combinatorics feel like a puzzle game.',
 NULL, ARRAY['informative', 'interesting', 'engaging']::review_category[]),

('90000000-0000-0000-0000-00000000000c', '80000000-0000-0000-0000-00000000000d', '11111111-1111-1111-1111-111111111111', 5,
 'I was nervous about how Alex, an active 7-year-old boy, would take to ballet. Within 15 minutes he was completely absorbed. Maria has this rare ability to make discipline feel like play. Alex immediately asked "when''s the next one?".',
 NULL, ARRAY['fun', 'engaging', 'interesting']::review_category[]),

('90000000-0000-0000-0000-00000000000d', '80000000-0000-0000-0000-00000000000e', '11111111-1111-1111-1111-111111111111', 5,
 'This class is pure joy to watch. Alex learned an 8-count routine in the first session and has been performing it for anyone who will watch. The instructors keep the energy high without letting it devolve into chaos — impressive class management.',
 NULL, ARRAY['fun', 'engaging']::review_category[]),

-- ══════════════════════════════════════════════════════════════════════════════
-- MICHAEL CHEN — Sophie past Digital Art (Bangkok, Nov 2025)
-- ══════════════════════════════════════════════════════════════════════════════
('90000000-0000-0000-0000-00000000000e', '80000000-0000-0000-0000-00000000001d', '22222222-2222-2222-2222-222222222222', 5,
 'Sophie has always been artistic but working on a Wacom tablet was completely new territory. By the end of the 2-hour session she had a finished character illustration she was incredibly proud of. The instructor walked around constantly and gave very specific, technical feedback.',
 'โซฟี่ชอบวาดรูปมาเสมอ แต่การทำงานบนแท็บเล็ต Wacom เป็นดินแดนใหม่อย่างสิ้นเชิง เมื่อสิ้นสุดเซสชั่น 2 ชั่วโมงเธอมีภาพประกอบตัวละครที่เสร็จสมบูรณ์ที่เธอภูมิใจมาก',
 ARRAY['informative', 'interesting', 'engaging']::review_category[]),

-- ══════════════════════════════════════════════════════════════════════════════
-- CARLOS RODRIGUEZ — Lucas past Soccer (Bangkok, Oct 2025)
-- ══════════════════════════════════════════════════════════════════════════════
('90000000-0000-0000-0000-00000000000f', '80000000-0000-0000-0000-000000000026', '44444444-4444-4444-4444-444444444444', 5,
 'Lucas has been playing informal soccer in the park for years but this was his first structured training. The session was full — 20 kids — but the coaches managed the chaos brilliantly. Lucas came home and immediately asked to practice his first touch in the hallway.',
 'ลูคัสเล่นฟุตบอลแบบไม่เป็นทางการในสวนสาธารณะมาหลายปีแล้ว แต่นี่เป็นการฝึกซ้อมที่มีโครงสร้างครั้งแรกของเขา เซสชั่นเต็ม 20 คน แต่โค้ชจัดการความวุ่นวายได้อย่างยอดเยี่ยม',
 ARRAY['engaging', 'fun']::review_category[]),

-- ══════════════════════════════════════════════════════════════════════════════
-- YUKI TANAKA — Hana past Piano (Bangkok, Sep 2025)
-- ══════════════════════════════════════════════════════════════════════════════
('90000000-0000-0000-0000-000000000010', '80000000-0000-0000-0000-000000000031', '66666666-6666-6666-6666-666666666666', 5,
 'Hana has shown an interest in music since she was very small. This lesson confirmed it. David Kim is an exceptional teacher — patient, encouraging, and clearly skilled. Hana cried a little when it was time to leave, which says everything. We''ll be regulars.',
 'ฮานะแสดงความสนใจในดนตรีตั้งแต่ยังเล็กมาก บทเรียนนี้ยืนยันสิ่งนั้น David Kim เป็นครูที่ยอดเยี่ยม อดทน ให้กำลังใจ และมีทักษะอย่างชัดเจน ฮานะร้องไห้นิดหน่อยเมื่อถึงเวลาต้องออกไป',
 ARRAY['informative', 'engaging', 'interesting']::review_category[]),

-- ══════════════════════════════════════════════════════════════════════════════
-- OLIVIA MARTINEZ — Noah past Soccer (Bangkok, Oct 2025)
-- ══════════════════════════════════════════════════════════════════════════════
('90000000-0000-0000-0000-000000000011', '80000000-0000-0000-0000-000000000036', '77777777-7777-7777-7777-777777777777', 4,
 'Noah is 9 and has been begging for proper soccer training for months. This delivered. The warm-up games were clever and immediately got everyone engaged. The session was high intensity but well-paced. One small thing: it would be nice to send a parent summary of what was covered.',
 'โนอาห์อายุ 9 ปีและขอการฝึกฟุตบอลที่เหมาะสมมาหลายเดือนแล้ว สิ่งนี้ตอบสนองได้ดี เกมวอร์มอัพฉลาดและทำให้ทุกคนมีส่วนร่วมทันที',
 ARRAY['fun', 'engaging']::review_category[]),

-- ══════════════════════════════════════════════════════════════════════════════
-- JAMES WILSON — Liam past 3D Modeling (Bangkok, Jan 2026)
-- ══════════════════════════════════════════════════════════════════════════════
('90000000-0000-0000-0000-000000000012', '80000000-0000-0000-0000-00000000003d', '88888888-8888-8888-8888-888888888888', 5,
 'Liam designed a fully functional phone stand and had it printed in the same session. He came home and immediately started designing version two with improvements. The instructor struck a perfect balance between showing technique and letting students explore independently.',
 'เลียมออกแบบที่วางโทรศัพท์ที่ใช้งานได้จริงและพิมพ์ออกมาในเซสชั่นเดียวกัน เขากลับบ้านและเริ่มออกแบบเวอร์ชัน 2 พร้อมการปรับปรุงทันที',
 ARRAY['informative', 'interesting', 'engaging']::review_category[]),

-- ══════════════════════════════════════════════════════════════════════════════
-- NATTAPORN CHAIYASIT — Bangkok past events
-- ══════════════════════════════════════════════════════════════════════════════
('90000000-0000-0000-0000-000000000013', '80000000-0000-0000-0000-00000000003e', '99999999-9999-9999-9999-999999999999', 5,
 'เพียรมาถึงบ้านด้วยความตื่นเต้นมากจนนอนไม่หลับ เธอเล่าให้พ่อฟังทุกขั้นตอนที่เธอตั้งโปรแกรมหุ่นยนต์ ครูผู้สอนพูดภาษาอังกฤษและไทยสลับกัน ซึ่งช่วยให้เพียรสบายใจมากขึ้น',
 'เพียรมาถึงบ้านด้วยความตื่นเต้นมากจนนอนไม่หลับ เธอเล่าให้พ่อฟังทุกขั้นตอนที่เธอตั้งโปรแกรมหุ่นยนต์',
 ARRAY['informative', 'interesting', 'engaging']::review_category[]),

('90000000-0000-0000-0000-000000000014', '80000000-0000-0000-0000-00000000003f', '99999999-9999-9999-9999-999999999999', 4,
 'การทดลองเคมีสนุกมากและปลอดภัยดี เพียรทำสไลม์กลับบ้านและทำให้พ่อแม่ดูซ้ำหลายครั้ง เซสชั่นยาวนานกว่าที่คาดเล็กน้อย แต่เนื้อหาดีมาก',
 'การทดลองเคมีสนุกมากและปลอดภัยดี เพียรทำสไลม์กลับบ้านและทำให้พ่อแม่ดูซ้ำหลายครั้ง',
 ARRAY['fun', 'informative']::review_category[]),

('90000000-0000-0000-0000-000000000015', '80000000-0000-0000-0000-000000000040', '99999999-9999-9999-9999-999999999999', 5,
 'ฝ้ายวาดภาพสีน้ำสวยงามมากจนเราเอาไปติดที่ผนัง Sofia Rossi เป็นครูที่มีพลังงานมาก เธอพูดกับเด็กทุกคนเป็นรายบุคคลและให้คำแนะนำที่เหมาะสมกับระดับของแต่ละคน',
 'ฝ้ายวาดภาพสีน้ำสวยงามมากจนเราเอาไปติดที่ผนัง Sofia Rossi เป็นครูที่มีพลังงานมาก',
 ARRAY['interesting', 'engaging']::review_category[]),

('90000000-0000-0000-0000-000000000016', '80000000-0000-0000-0000-000000000041', '99999999-9999-9999-9999-999999999999', 4,
 'ฝ้ายชอบการวาดภาพดิจิทัลมาก เธอเรียนรู้วิธีใช้เลเยอร์และแปรงดิจิทัลได้เร็วมาก ครูอธิบายดี แต่บางครั้งรอเด็กที่เร็วกว่าน้อยไปหน่อย',
 'ฝ้ายชอบการวาดภาพดิจิทัลมาก เธอเรียนรู้วิธีใช้เลเยอร์และแปรงดิจิทัลได้เร็วมาก',
 ARRAY['interesting', 'informative']::review_category[]),

-- ══════════════════════════════════════════════════════════════════════════════
-- JAMES O'CONNOR — Boston past events
-- ══════════════════════════════════════════════════════════════════════════════
('90000000-0000-0000-0000-000000000017', '80000000-0000-0000-0000-000000000044', 'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa', 5,
 'Connor has always been competitive but he''d never played chess. Patrick O''Brien has an incredible way of making every child feel like they have a real shot at becoming great. Connor came home and challenged every adult in the house. He''s completely hooked.',
 NULL, ARRAY['informative', 'interesting', 'engaging']::review_category[]),

('90000000-0000-0000-0000-000000000018', '80000000-0000-0000-0000-000000000045', 'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa', 5,
 'Fiona is 7 and I wasn''t sure ballet was the right fit. But watching the class through the window I saw her completely transformed — totally focused, trying so hard, and beaming when she got a step right. Maria creates a magical atmosphere. We''re signing up for the full term.',
 NULL, ARRAY['fun', 'engaging']::review_category[]),

('90000000-0000-0000-0000-000000000019', '80000000-0000-0000-0000-000000000046', 'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa', 5,
 'The Math Olympiad prep class completely changed how Connor thinks about math. He came home talking about pigeonhole problems and "elegant" solutions. Patrick has a way of presenting hard problems as adventures rather than tests. Connor asked if we could do private lessons.',
 NULL, ARRAY['informative', 'interesting']::review_category[]),

-- ══════════════════════════════════════════════════════════════════════════════
-- ANANYA KRISHNAMURTHY — Bangkok past events
-- ══════════════════════════════════════════════════════════════════════════════
('90000000-0000-0000-0000-00000000001a', '80000000-0000-0000-0000-00000000004a', 'bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb', 5,
 'Aryan is 11 and has been interested in robotics for years but only through YouTube videos. Actually building and programming a LEGO robot was completely different. Dr. Lee''s team is exceptional. The bilingual delivery (English and Thai) meant no child felt left behind.',
 'อาร์ยานอายุ 11 ปีและสนใจหุ่นยนต์มาหลายปีแต่ผ่านวิดีโอ YouTube เท่านั้น การสร้างและเขียนโปรแกรมหุ่นยนต์ LEGO จริงๆ แตกต่างกันอย่างสิ้นเชิง ทีมของ Dr. Lee ยอดเยี่ยมมาก',
 ARRAY['informative', 'interesting', 'engaging']::review_category[]),

('90000000-0000-0000-0000-00000000001b', '80000000-0000-0000-0000-00000000004b', 'bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb', 4,
 'Aryan learned the core concepts of Scratch loops and events and built a working game in one session — impressive pacing. My one note is that the older kids (Aryan is nearly 12) could use a slightly more advanced track. But the fundamentals they covered were solid.',
 'อาร์ยานเรียนรู้แนวคิดหลักของลูปและเหตุการณ์ใน Scratch และสร้างเกมที่ใช้งานได้ในเซสชั่นเดียว น่าประทับใจมาก',
 ARRAY['informative', 'engaging']::review_category[]),

('90000000-0000-0000-0000-00000000001c', '80000000-0000-0000-0000-00000000004c', 'bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb', 5,
 'Priya is 8 and was completely new to piano. David Kim has a gift for making young children feel capable. He found a song Priya already loved and used it as the entry point, which was brilliant. She practiced every evening for a week afterward.',
 'ปรียาอายุ 8 ปีและไม่รู้เรื่องเปียโนเลย David Kim มีพรสวรรค์ในการทำให้เด็กเล็กรู้สึกว่าตัวเองทำได้ เขาพบเพลงที่ปรียาชื่นชอบอยู่แล้วและใช้เป็นจุดเข้า ซึ่งยอดเยี่ยมมาก',
 ARRAY['informative', 'engaging', 'fun']::review_category[]),

-- ══════════════════════════════════════════════════════════════════════════════
-- THOMAS BRENNAN — Boston past events
-- ══════════════════════════════════════════════════════════════════════════════
('90000000-0000-0000-0000-00000000001d', '80000000-0000-0000-0000-00000000004f', 'cccccccc-cccc-cccc-cccc-cccccccccccc', 4,
 'Sean is a math kid who wanted to try something new. Chess was a natural fit. Patrick''s teaching style — competitive but encouraging — suits Sean perfectly. The in-house tournament at the end was a brilliant touch that motivated even the newer players.',
 NULL, ARRAY['informative', 'engaging']::review_category[]),

('90000000-0000-0000-0000-00000000001e', '80000000-0000-0000-0000-000000000050', 'cccccccc-cccc-cccc-cccc-cccccccccccc', 5,
 'The circuits class was exactly what Sean needed. He''s a hands-on learner and sitting down for theory lessons at school bores him. Here he built three working circuits in two hours and was completely absorbed. Jennifer Walsh is a fantastic educator.',
 NULL, ARRAY['informative', 'interesting', 'engaging']::review_category[]),

('90000000-0000-0000-0000-00000000001f', '80000000-0000-0000-0000-000000000051', 'cccccccc-cccc-cccc-cccc-cccccccccccc', 5,
 'Aoife is dramatic by nature and ballet is her perfect outlet. The class was joyful and disciplined in equal measure. The teacher moved between Thai and English fluidly — actually this is Boston so purely English — and created a very inclusive, supportive environment.',
 NULL, ARRAY['fun', 'engaging', 'interesting']::review_category[]),

('90000000-0000-0000-0000-000000000020', '80000000-0000-0000-0000-000000000052', 'cccccccc-cccc-cccc-cccc-cccccccccccc', 5,
 'Aoife was intimidated walking in — all the other kids seemed to already know the moves. Within 20 minutes she was keeping up and laughing. The hip hop instructor has an extraordinary energy and manages to make every child feel seen. Aoife has watched the video of her routine probably 50 times.',
 NULL, ARRAY['fun', 'engaging']::review_category[]),

-- ══════════════════════════════════════════════════════════════════════════════
-- SIRIPORN WATTANABE — Bangkok past events
-- ══════════════════════════════════════════════════════════════════════════════
('90000000-0000-0000-0000-000000000021', '80000000-0000-0000-0000-000000000055', 'dddddddd-dddd-dddd-dddd-dddddddddddd', 4,
 'กฤตชอบงานหุ่นยนต์ แต่เขาทำงานช้ากว่าเด็กบางคนในกลุ่ม ครูผู้สอนช่วยเขาอย่างอดทนและทำให้เขาทันกลุ่มได้ กฤตบอกว่าอยากมาอีก',
 'กฤตชอบงานหุ่นยนต์ แต่เขาทำงานช้ากว่าเด็กบางคนในกลุ่ม ครูผู้สอนช่วยเขาอย่างอดทนและทำให้เขาทันกลุ่มได้',
 ARRAY['informative', 'engaging']::review_category[]),

('90000000-0000-0000-0000-000000000022', '80000000-0000-0000-0000-000000000056', 'dddddddd-dddd-dddd-dddd-dddddddddddd', 5,
 'กฤตสร้างเกมบน Scratch ได้เองในเซสชั่นแรก ครูผู้สอนอธิบายแนวคิดการเขียนโปรแกรมได้ดีมากสำหรับเด็ก กฤตไม่หยุดคุยเรื่องเกมที่เขาสร้างตลอดทางกลับบ้าน',
 'กฤตสร้างเกมบน Scratch ได้เองในเซสชั่นแรก ครูผู้สอนอธิบายแนวคิดการเขียนโปรแกรมได้ดีมากสำหรับเด็ก',
 ARRAY['informative', 'interesting', 'engaging']::review_category[]),

('90000000-0000-0000-0000-000000000023', '80000000-0000-0000-0000-000000000057', 'dddddddd-dddd-dddd-dddd-dddddddddddd', 5,
 'พลอยร้องเพลงในคณะนักร้องประสานเสียงตลอดทางกลับบ้าน ครูผู้อำนวยการยอดเยี่ยมมาก เขาสอนให้เด็กๆ ร้องประสานสองส่วนในครั้งเดียว ซึ่งน่าทึ่งมากสำหรับเด็กอายุ 12 ปี',
 'พลอยร้องเพลงในคณะนักร้องประสานเสียงตลอดทางกลับบ้าน ครูผู้อำนวยการยอดเยี่ยมมาก',
 ARRAY['fun', 'engaging', 'interesting']::review_category[]),

-- ══════════════════════════════════════════════════════════════════════════════
-- MEI-LING HUANG — Bangkok past events
-- ══════════════════════════════════════════════════════════════════════════════
('90000000-0000-0000-0000-000000000024', '80000000-0000-0000-0000-00000000005b', 'eeeeeeee-eeee-eeee-eeee-eeeeeeeeeeee', 5,
 'We chose the evening astronomy session hoping the telescope observation would be the highlight — and it was even better than expected. Wei identified three constellations on his own. The instructor''s bilingual delivery was seamless. We stayed 20 minutes late asking questions.',
 'เราเลือกเซสชั่นดาราศาสตร์ตอนเย็นโดยหวังว่าการสังเกตการณ์กล้องโทรทรรศน์จะเป็นไฮไลต์ และมันดีกว่าที่คาดไว้มาก เหวยระบุดาวสามดวงได้ด้วยตัวเอง',
 ARRAY['informative', 'interesting']::review_category[]),

('90000000-0000-0000-0000-000000000025', '80000000-0000-0000-0000-00000000005c', 'eeeeeeee-eeee-eeee-eeee-eeeeeeeeeeee', 4,
 'Wei already knew some programming concepts from online tutorials, and the Scratch instructor recognized this quickly and gave him additional challenges. That kind of differentiated instruction is rare and we really appreciated it. The class size was also ideal.',
 'เหวยรู้แนวคิดการเขียนโปรแกรมบางอย่างจากบทเรียนออนไลน์แล้ว และครู Scratch รู้จักสิ่งนี้อย่างรวดเร็วและให้ความท้าทายเพิ่มเติมแก่เขา',
 ARRAY['informative', 'engaging']::review_category[]),

('90000000-0000-0000-0000-000000000026', '80000000-0000-0000-0000-00000000005d', 'eeeeeeee-eeee-eeee-eeee-eeeeeeeeeeee', 5,
 'Lily has always drawn in notebooks but this was her first formal art class. She came home with a watercolor piece she''s declared "her best work ever." Sofia Rossi is a wonderful teacher — she encouraged every child''s individual voice rather than pushing a single style.',
 'ลิลี่วาดรูปในสมุดมาตลอด แต่นี่เป็นชั้นเรียนศิลปะอย่างเป็นทางการครั้งแรกของเธอ เธอกลับบ้านพร้อมกับภาพสีน้ำที่เธอประกาศว่า "เป็นผลงานที่ดีที่สุดเท่าที่เคยทำ"',
 ARRAY['interesting', 'engaging', 'fun']::review_category[]),

-- ══════════════════════════════════════════════════════════════════════════════
-- RACHEL KIM — Boston past events
-- ══════════════════════════════════════════════════════════════════════════════
('90000000-0000-0000-0000-000000000027', '80000000-0000-0000-0000-000000000067', 'ffffffff-ffff-ffff-ffff-ffffffffffff', 5,
 'Hannah is the kid who asks "but why?" about every math problem at school and never gets a satisfying answer. This class is built for kids like her. Patrick structures every problem as a discovery, not a procedure. She came home with a notebook full of her own proofs.',
 NULL, ARRAY['informative', 'interesting', 'engaging']::review_category[]),

('90000000-0000-0000-0000-000000000028', '80000000-0000-0000-0000-000000000068', 'ffffffff-ffff-ffff-ffff-ffffffffffff', 5,
 'Hannah built a working LED circuit and a moisture sensor in two hours. She''s been asking to set up a "lab" in her bedroom ever since. Jennifer Walsh explains everything at exactly the right level — not too simple, not overwhelming. Boston STEM Lab is genuinely special.',
 NULL, ARRAY['informative', 'interesting', 'engaging']::review_category[]),

('90000000-0000-0000-0000-000000000029', '80000000-0000-0000-0000-000000000069', 'ffffffff-ffff-ffff-ffff-ffffffffffff', 4,
 'Jake went in skeptical about chess ("it''s for old people, Mom") and came out having won his first game against a classmate. Patrick has a remarkable ability to make the game feel urgent and exciting. Jake is now practicing daily tactics puzzles on his own.',
 NULL, ARRAY['informative', 'engaging']::review_category[]),

('90000000-0000-0000-0000-00000000002a', '80000000-0000-0000-0000-00000000006a', 'ffffffff-ffff-ffff-ffff-ffffffffffff', 5,
 'Jake is a natural mover and hip hop was an instant fit. He learned a full 16-count routine in 90 minutes. The instructor builds confidence alongside skill — Jake left walking taller. The end-of-class freestyle session was a great idea, giving every kid a moment to shine.',
 NULL, ARRAY['fun', 'engaging', 'interesting']::review_category[]),

-- ══════════════════════════════════════════════════════════════════════════════
-- PATRICIA WALSH — Boston past events
-- ══════════════════════════════════════════════════════════════════════════════
('90000000-0000-0000-0000-00000000002b', '80000000-0000-0000-0000-00000000006d', '13131313-1313-1313-1313-131313131313', 5,
 'Lena has been asking for ballet for two years. This exceeded all expectations. Maria Fontaine has a gift for communicating with young children through movement — she barely uses words. Lena was humming the music from class for days afterward. Exceptional.',
 NULL, ARRAY['fun', 'engaging', 'interesting']::review_category[]),

('90000000-0000-0000-0000-00000000002c', '80000000-0000-0000-0000-00000000006e', '13131313-1313-1313-1313-131313131313', 4,
 'Declan is 11 and very competitive. Chess gave him a great outlet for that. He was frustrated when he lost his first game but the instructor handled it beautifully — talked about how grandmasters analyze their losses. Declan now wants to study openings. Mission accomplished.',
 NULL, ARRAY['informative', 'engaging']::review_category[]),

('90000000-0000-0000-0000-00000000002d', '80000000-0000-0000-0000-00000000006f', '13131313-1313-1313-1313-131313131313', 5,
 'This was Declan''s first engineering-style class and it lit something up in him. He designed a circuit that controlled an LED with a button, then immediately started asking what else he could control. Jennifer Walsh''s enthusiasm is completely genuine and completely contagious.',
 NULL, ARRAY['informative', 'interesting', 'engaging']::review_category[]),

-- ══════════════════════════════════════════════════════════════════════════════
-- SOMCHAI THONGPRASERT — Bangkok past events
-- ══════════════════════════════════════════════════════════════════════════════
('90000000-0000-0000-0000-00000000002e', '80000000-0000-0000-0000-000000000073', '14141414-1414-1414-1414-141414141414', 5,
 'กามลได้รับการฝึกฟุตบอลที่เขาต้องการมาตลอด โค้ชสอนพื้นฐานได้ดีมาก กามลกลับบ้านและฝึกต่อกับพ่อในสนามหญ้าอีกหนึ่งชั่วโมง นั่นบอกทุกอย่างแล้ว',
 'กามลได้รับการฝึกฟุตบอลที่เขาต้องการมาตลอด โค้ชสอนพื้นฐานได้ดีมาก',
 ARRAY['fun', 'engaging']::review_category[]),

('90000000-0000-0000-0000-00000000002f', '80000000-0000-0000-0000-000000000074', '14141414-1414-1414-1414-141414141414', 4,
 'กามลไม่ค่อยกล้าลงน้ำในตอนแรก แต่โค้ชอดทนมากและไม่กดดัน ภายในหนึ่งชั่วโมงกามลว่ายน้ำได้ระยะสั้นๆ โดยไม่ต้องใช้แผ่นลอยน้ำ เราจะลงทะเบียนอีกแน่นอน',
 'กามลไม่ค่อยกล้าลงน้ำในตอนแรก แต่โค้ชอดทนมากและไม่กดดัน ภายในหนึ่งชั่วโมงกามลว่ายน้ำได้ระยะสั้นๆ',
 ARRAY['informative', 'engaging']::review_category[]);
