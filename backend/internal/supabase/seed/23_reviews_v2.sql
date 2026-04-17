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
