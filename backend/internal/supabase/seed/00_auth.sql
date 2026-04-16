-- ============================================
-- 0. auth.users
-- Default password for every seed user: Skillspark1!
-- ============================================
CREATE EXTENSION IF NOT EXISTS pgcrypto;

INSERT INTO auth.users (
  instance_id,
  id,
  aud,
  role,
  email,
  encrypted_password,
  email_confirmed_at,
  raw_app_meta_data,
  raw_user_meta_data,
  created_at,
  updated_at,
  confirmation_token,
  recovery_token,
  email_change,
  email_change_token_new
) VALUES
-- Guardians (original 8)
('00000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-00000000000a', 'authenticated', 'authenticated', 'sarah.johnson@email.com',    crypt('Skillspark1!', gen_salt('bf')), NOW(), '{"provider":"email","providers":["email"]}', '{}', NOW(), NOW(), '', '', '', ''),
('00000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-00000000000b', 'authenticated', 'authenticated', 'michael.chen@email.com',     crypt('Skillspark1!', gen_salt('bf')), NOW(), '{"provider":"email","providers":["email"]}', '{}', NOW(), NOW(), '', '', '', ''),
('00000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-00000000000c', 'authenticated', 'authenticated', 'priya.patel@email.com',      crypt('Skillspark1!', gen_salt('bf')), NOW(), '{"provider":"email","providers":["email"]}', '{}', NOW(), NOW(), '', '', '', ''),
('00000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-00000000000d', 'authenticated', 'authenticated', 'carlos.rodriguez@email.com', crypt('Skillspark1!', gen_salt('bf')), NOW(), '{"provider":"email","providers":["email"]}', '{}', NOW(), NOW(), '', '', '', ''),
('00000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-00000000000e', 'authenticated', 'authenticated', 'emma.thompson@email.com',    crypt('Skillspark1!', gen_salt('bf')), NOW(), '{"provider":"email","providers":["email"]}', '{}', NOW(), NOW(), '', '', '', ''),
('00000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-00000000000f', 'authenticated', 'authenticated', 'yuki.tanaka@email.com',      crypt('Skillspark1!', gen_salt('bf')), NOW(), '{"provider":"email","providers":["email"]}', '{}', NOW(), NOW(), '', '', '', ''),
('00000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000001', 'authenticated', 'authenticated', 'olivia.martinez@email.com',  crypt('Skillspark1!', gen_salt('bf')), NOW(), '{"provider":"email","providers":["email"]}', '{}', NOW(), NOW(), '', '', '', ''),
('00000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000002', 'authenticated', 'authenticated', 'james.wilson@email.com',     crypt('Skillspark1!', gen_salt('bf')), NOW(), '{"provider":"email","providers":["email"]}', '{}', NOW(), NOW(), '', '', '', ''),
-- Bangkok org managers
('00000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000003', 'authenticated', 'authenticated', 'amanda.lee@scienceacademy.com',    crypt('Skillspark1!', gen_salt('bf')), NOW(), '{"provider":"email","providers":["email"]}', '{}', NOW(), NOW(), '', '', '', ''),
('00000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000004', 'authenticated', 'authenticated', 'marcus.thompson@sportscenter.com', crypt('Skillspark1!', gen_salt('bf')), NOW(), '{"provider":"email","providers":["email"]}', '{}', NOW(), NOW(), '', '', '', ''),
('00000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000005', 'authenticated', 'authenticated', 'sofia.rossi@artsstudio.com',       crypt('Skillspark1!', gen_salt('bf')), NOW(), '{"provider":"email","providers":["email"]}', '{}', NOW(), NOW(), '', '', '', ''),
('00000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000006', 'authenticated', 'authenticated', 'david.kim@musicschool.com',        crypt('Skillspark1!', gen_salt('bf')), NOW(), '{"provider":"email","providers":["email"]}', '{}', NOW(), NOW(), '', '', '', ''),
-- Boston org managers
('00000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000007', 'authenticated', 'authenticated', 'jennifer.walsh@bostonstemlab.com', crypt('Skillspark1!', gen_salt('bf')), NOW(), '{"provider":"email","providers":["email"]}', '{}', NOW(), NOW(), '', '', '', ''),
('00000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000008', 'authenticated', 'authenticated', 'maria.fontaine@nedance.com',       crypt('Skillspark1!', gen_salt('bf')), NOW(), '{"provider":"email","providers":["email"]}', '{}', NOW(), NOW(), '', '', '', ''),
('00000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000009', 'authenticated', 'authenticated', 'patrick.obrien@fenwaychess.com',   crypt('Skillspark1!', gen_salt('bf')), NOW(), '{"provider":"email","providers":["email"]}', '{}', NOW(), NOW(), '', '', '', ''),
-- Additional guardians (Bangkok-based)
('00000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000010', 'authenticated', 'authenticated', 'nattaporn.chaiyasit@email.com',    crypt('Skillspark1!', gen_salt('bf')), NOW(), '{"provider":"email","providers":["email"]}', '{}', NOW(), NOW(), '', '', '', ''),
('00000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000012', 'authenticated', 'authenticated', 'ananya.krishnamurthy@email.com',   crypt('Skillspark1!', gen_salt('bf')), NOW(), '{"provider":"email","providers":["email"]}', '{}', NOW(), NOW(), '', '', '', ''),
('00000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000014', 'authenticated', 'authenticated', 'siriporn.wattanabe@email.com',     crypt('Skillspark1!', gen_salt('bf')), NOW(), '{"provider":"email","providers":["email"]}', '{}', NOW(), NOW(), '', '', '', ''),
('00000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000015', 'authenticated', 'authenticated', 'meiling.huang@email.com',          crypt('Skillspark1!', gen_salt('bf')), NOW(), '{"provider":"email","providers":["email"]}', '{}', NOW(), NOW(), '', '', '', ''),
('00000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000017', 'authenticated', 'authenticated', 'arjun.sharma@email.com',           crypt('Skillspark1!', gen_salt('bf')), NOW(), '{"provider":"email","providers":["email"]}', '{}', NOW(), NOW(), '', '', '', ''),
('00000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000019', 'authenticated', 'authenticated', 'somchai.thongprasert@email.com',   crypt('Skillspark1!', gen_salt('bf')), NOW(), '{"provider":"email","providers":["email"]}', '{}', NOW(), NOW(), '', '', '', ''),
-- Additional guardians (Boston-based)
('00000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000011', 'authenticated', 'authenticated', 'james.oconnor@email.com',          crypt('Skillspark1!', gen_salt('bf')), NOW(), '{"provider":"email","providers":["email"]}', '{}', NOW(), NOW(), '', '', '', ''),
('00000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000013', 'authenticated', 'authenticated', 'thomas.brennan@email.com',         crypt('Skillspark1!', gen_salt('bf')), NOW(), '{"provider":"email","providers":["email"]}', '{}', NOW(), NOW(), '', '', '', ''),
('00000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000016', 'authenticated', 'authenticated', 'rachel.kim@email.com',             crypt('Skillspark1!', gen_salt('bf')), NOW(), '{"provider":"email","providers":["email"]}', '{}', NOW(), NOW(), '', '', '', ''),
('00000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000018', 'authenticated', 'authenticated', 'patricia.walsh@email.com',         crypt('Skillspark1!', gen_salt('bf')), NOW(), '{"provider":"email","providers":["email"]}', '{}', NOW(), NOW(), '', '', '', '')
ON CONFLICT (id) DO NOTHING;
