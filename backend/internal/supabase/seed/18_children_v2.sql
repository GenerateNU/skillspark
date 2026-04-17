-- ============================================
-- 18. NEW CHILDREN — Demo Guardian Kids
-- ============================================
INSERT INTO child (id, name, school_id, birth_month, birth_year, interests, guardian_id) VALUES

-- ── Jennifer Park's children ─────────────────────────────────────────────────
-- Lily Park: 8 years old, science/tech/math kid in Newton
('ad000001-0000-0000-0000-000000000001', 'Lily Park',  'dd000001-0000-0000-0000-000000000001', 6,  2017, ARRAY['science','technology','math','art']::category[],      'cc000001-0000-0000-0000-000000000001'),
-- Max Park: 5 years old, active kid who loves sports and making things
('ad000001-0000-0000-0000-000000000002', 'Max Park',   'dd000001-0000-0000-0000-000000000001', 3,  2020, ARRAY['sports','art']::category[],                           'cc000001-0000-0000-0000-000000000001'),

-- ── Nattaya Srisuk's children ────────────────────────────────────────────────
-- Ploy Srisuk: 10 years old, passionate about dance and languages
('ad000001-0000-0000-0000-000000000003', 'Ploy Srisuk', 'dd000001-0000-0000-0000-000000000004', 8,  2015, ARRAY['music','art','language','science']::category[],       'cc000001-0000-0000-0000-000000000002'),
-- Ton Srisuk: 7 years old, muay thai fan, loves technology
('ad000001-0000-0000-0000-000000000004', 'Ton Srisuk',  '20000000-0000-0000-0000-000000000001', 1,  2018, ARRAY['sports','technology']::category[],                    'cc000001-0000-0000-0000-000000000002'),

-- ── Marcus Webb's children ───────────────────────────────────────────────────
-- Zoe Webb: 11 years old, multi-talented tech + art kid at Boston Latin
('ad000001-0000-0000-0000-000000000005', 'Zoe Webb',   'dd000001-0000-0000-0000-000000000002', 11, 2014, ARRAY['technology','math','science','art']::category[],       'cc000001-0000-0000-0000-000000000003');
