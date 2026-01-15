-- ============================================
-- 8. EVENTS
-- ============================================
INSERT INTO event (id, title, description, organization_id, age_range_min, age_range_max, category, header_image_s3_key) VALUES
-- Science Academy Events
('60000000-0000-0000-0000-000000000001', 'Junior Robotics Workshop', 'Learn the basics of robotics with hands-on LEGO Mindstorms projects. Build and program your own robots!', '40000000-0000-0000-0000-000000000001', 8, 12, ARRAY['science','technology']::category[], 'events/robotics_workshop.jpg'),
('60000000-0000-0000-0000-000000000002', 'Chemistry for Kids', 'Exciting chemistry experiments that are safe and fun. Discover reactions, make slime, and learn about molecules!', '40000000-0000-0000-0000-000000000001', 7, 10, ARRAY['science']::category[], 'events/chemistry_kids.jpg'),
('60000000-0000-0000-0000-000000000003', 'Astronomy Club', 'Explore the wonders of space! Learn about planets, stars, and galaxies. Includes telescope observation sessions.', '40000000-0000-0000-0000-000000000001', 9, 14, ARRAY['science']::category[], 'events/astronomy.jpg'),
-- Sports Center Events
('60000000-0000-0000-0000-000000000004', 'Soccer Skills Training', 'Develop fundamental soccer skills including dribbling, passing, and teamwork in a fun environment.', '40000000-0000-0000-0000-000000000002', 6, 12, ARRAY['sports']::category[], 'events/soccer_training.jpg'),
('60000000-0000-0000-0000-000000000005', 'Basketball Basics', 'Learn basketball fundamentals: shooting, dribbling, defense, and game strategy.', '40000000-0000-0000-0000-000000000002', 7, 13, ARRAY['sports']::category[], NULL),
('60000000-0000-0000-0000-000000000006', 'Swimming Lessons', 'Professional swimming instruction for beginners to intermediate levels. Focus on technique and water safety.', '40000000-0000-0000-0000-000000000002', 5, 15, ARRAY['sports']::category[], 'events/swimming.jpg'),
-- Arts Studio Events
('60000000-0000-0000-0000-000000000007', 'Painting & Drawing Workshop', 'Explore various art techniques including watercolor, acrylic, and sketching. All materials provided!', '40000000-0000-0000-0000-000000000003', 6, 14, ARRAY['art']::category[], 'events/painting_workshop.jpg'),
('60000000-0000-0000-0000-000000000008', 'Pottery for Beginners', 'Learn to work with clay! Create bowls, cups, and sculptures using hand-building and wheel techniques.', '40000000-0000-0000-0000-000000000003', 8, 15, ARRAY['art']::category[], NULL),
('60000000-0000-0000-0000-000000000009', 'Digital Art & Design', 'Introduction to digital illustration using tablets. Learn basic design principles and digital tools.', '40000000-0000-0000-0000-000000000003', 10, 16, ARRAY['art','technology']::category[], 'events/digital_art.jpg'),
-- Music School Events
('60000000-0000-0000-0000-00000000000a', 'Piano for Beginners', 'Start your musical journey with piano! Learn to read music and play simple songs.', '40000000-0000-0000-0000-000000000004', 6, 12, ARRAY['music']::category[], 'events/piano_lessons.jpg'),
('60000000-0000-0000-0000-00000000000b', 'Guitar Fundamentals', 'Learn basic chords, strumming patterns, and your first songs on acoustic guitar.', '40000000-0000-0000-0000-000000000004', 8, 15, ARRAY['music']::category[], NULL),
('60000000-0000-0000-0000-00000000000c', 'Kids Choir', 'Join our fun choir! Learn harmony, vocal techniques, and perform in recitals.', '40000000-0000-0000-0000-000000000004', 7, 13, ARRAY['music']::category[], 'events/kids_choir.jpg'),
-- Tech Workshop Events
('60000000-0000-0000-0000-00000000000d', 'Coding with Scratch', 'Learn programming basics through Scratch! Create games and animations with visual coding blocks.', '40000000-0000-0000-0000-000000000005', 7, 11, ARRAY['technology']::category[], 'events/scratch_coding.jpg'),
('60000000-0000-0000-0000-00000000000e', 'Python for Kids', 'Introduction to Python programming. Build simple programs and games while learning core concepts.', '40000000-0000-0000-0000-000000000005', 10, 15, ARRAY['technology','math']::category[], NULL),
('60000000-0000-0000-0000-00000000000f', '3D Modeling Workshop', 'Design 3D objects using Tinkercad! Learn basics of 3D design and prepare models for 3D printing.', '40000000-0000-0000-0000-000000000005', 9, 14, ARRAY['technology','art']::category[], 'events/3d_modeling.jpg');
