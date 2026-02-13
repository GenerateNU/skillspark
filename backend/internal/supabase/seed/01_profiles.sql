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
