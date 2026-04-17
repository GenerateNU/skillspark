-- ============================================
-- 22. REGISTRATIONS + PAYMENTS — Demo Guardian Accounts
--
-- Jennifer Park  (cc000001-...001)  cus_jenniferpark
--   children: Lily (ad000001-...001), Max (ad000001-...002)
-- Nattaya Srisuk (cc000001-...002)  cus_nattayath
--   children: Ploy (ad000001-...003), Ton  (ad000001-...004)
-- Marcus Webb    (cc000001-...003)  cus_marcuswebb
--   children: Zoe  (ad000001-...005)
-- ============================================

INSERT INTO registration (id, child_id, guardian_id, event_occurrence_id, status) VALUES

-- ── Jennifer Park — COMPLETED (past) ────────────────────────────────────────
('ae000001-0000-0000-0000-000000000001', 'ad000001-0000-0000-0000-000000000001', 'cc000001-0000-0000-0000-000000000001', '7b000000-0000-0000-0000-000000000001', 'registered'), -- Lily → Young Engineers (past)
('ae000001-0000-0000-0000-000000000002', 'ad000001-0000-0000-0000-000000000001', 'cc000001-0000-0000-0000-000000000001', '7b000000-0000-0000-0000-000000000002', 'registered'), -- Lily → Science Explorers (past)
('ae000001-0000-0000-0000-000000000003', 'ad000001-0000-0000-0000-000000000001', 'cc000001-0000-0000-0000-000000000001', '7b000000-0000-0000-0000-000000000003', 'registered'), -- Lily → App Inventor (past)
('ae000001-0000-0000-0000-000000000004', 'ad000001-0000-0000-0000-000000000001', 'cc000001-0000-0000-0000-000000000001', '7b000000-0000-0000-0000-00000000000d', 'registered'), -- Lily → Web Design (past)
('ae000001-0000-0000-0000-000000000005', 'ad000001-0000-0000-0000-000000000001', 'cc000001-0000-0000-0000-000000000001', '7b000000-0000-0000-0000-00000000000e', 'registered'), -- Lily → Game Design (past)
('ae000001-0000-0000-0000-000000000006', 'ad000001-0000-0000-0000-000000000002', 'cc000001-0000-0000-0000-000000000001', '7b000000-0000-0000-0000-000000000004', 'registered'), -- Max  → Youth Soccer (past)
('ae000001-0000-0000-0000-000000000007', 'ad000001-0000-0000-0000-000000000002', 'cc000001-0000-0000-0000-000000000001', '7b000000-0000-0000-0000-000000000005', 'registered'), -- Max  → Basketball (past)
('ae000001-0000-0000-0000-000000000008', 'ad000001-0000-0000-0000-000000000002', 'cc000001-0000-0000-0000-000000000001', '7b000000-0000-0000-0000-00000000000a', 'registered'), -- Max  → Watercolor (past)

-- ── Jennifer Park — UPCOMING ─────────────────────────────────────────────────
('ae000001-0000-0000-0000-000000000009', 'ad000001-0000-0000-0000-000000000001', 'cc000001-0000-0000-0000-000000000001', '7b000000-0000-0000-0000-000000000015', 'registered'), -- Lily → App Inventor (upcoming)
('ae000001-0000-0000-0000-00000000000a', 'ad000001-0000-0000-0000-000000000001', 'cc000001-0000-0000-0000-000000000001', '7b000000-0000-0000-0000-000000000013', 'registered'), -- Lily → Young Engineers (upcoming)
('ae000001-0000-0000-0000-00000000000b', 'ad000001-0000-0000-0000-000000000002', 'cc000001-0000-0000-0000-000000000001', '7b000000-0000-0000-0000-000000000018', 'registered'), -- Max  → Soccer (upcoming)
('ae000001-0000-0000-0000-00000000000c', 'ad000001-0000-0000-0000-000000000002', 'cc000001-0000-0000-0000-000000000001', '7b000000-0000-0000-0000-000000000022', 'registered'), -- Max  → Comic Art (upcoming)

-- ── Jennifer Park — CANCELLED ────────────────────────────────────────────────
('ae000001-0000-0000-0000-00000000000d', 'ad000001-0000-0000-0000-000000000002', 'cc000001-0000-0000-0000-000000000001', '7b000000-0000-0000-0000-00000000001a', 'cancelled'),  -- Max  → Swim Team (cancelled)

-- ── Nattaya Srisuk — COMPLETED (past) ───────────────────────────────────────
('ae000001-0000-0000-0000-00000000000e', 'ad000001-0000-0000-0000-000000000003', 'cc000001-0000-0000-0000-000000000002', '7b000000-0000-0000-0000-000000000011', 'registered'), -- Ploy → Ballet (past)
('ae000001-0000-0000-0000-00000000000f', 'ad000001-0000-0000-0000-000000000003', 'cc000001-0000-0000-0000-000000000002', '7b000000-0000-0000-0000-000000000012', 'registered'), -- Ploy → STEM Camp (past)
('ae000001-0000-0000-0000-000000000010', 'ad000001-0000-0000-0000-000000000003', 'cc000001-0000-0000-0000-000000000002', '7b000000-0000-0000-0000-000000000028', 'registered'), -- Ploy → Robotics Workshop (past)
('ae000001-0000-0000-0000-000000000011', 'ad000001-0000-0000-0000-000000000003', 'cc000001-0000-0000-0000-000000000002', '7b000000-0000-0000-0000-000000000029', 'registered'), -- Ploy → Chemistry (past)
('ae000001-0000-0000-0000-000000000012', 'ad000001-0000-0000-0000-000000000004', 'cc000001-0000-0000-0000-000000000002', '7b000000-0000-0000-0000-000000000010', 'registered'), -- Ton  → Muay Thai Beginner (past)
('ae000001-0000-0000-0000-000000000013', 'ad000001-0000-0000-0000-000000000004', 'cc000001-0000-0000-0000-000000000002', '7b000000-0000-0000-0000-00000000002a', 'registered'), -- Ton  → Soccer Skills (past)

-- ── Nattaya Srisuk — UPCOMING ────────────────────────────────────────────────
('ae000001-0000-0000-0000-000000000014', 'ad000001-0000-0000-0000-000000000003', 'cc000001-0000-0000-0000-000000000002', '7b000000-0000-0000-0000-000000000034', 'registered'), -- Ploy → Ballet (upcoming)
('ae000001-0000-0000-0000-000000000015', 'ad000001-0000-0000-0000-000000000003', 'cc000001-0000-0000-0000-000000000002', '7b000000-0000-0000-0000-000000000037', 'registered'), -- Ploy → STEM Camp (upcoming)
('ae000001-0000-0000-0000-000000000016', 'ad000001-0000-0000-0000-000000000004', 'cc000001-0000-0000-0000-000000000002', '7b000000-0000-0000-0000-000000000031', 'registered'), -- Ton  → Muay Thai (upcoming)
('ae000001-0000-0000-0000-000000000017', 'ad000001-0000-0000-0000-000000000004', 'cc000001-0000-0000-0000-000000000002', '7b000000-0000-0000-0000-00000000001b', 'registered'), -- Ton  → Soccer (upcoming, Boston!)

-- ── Marcus Webb — COMPLETED (past) ──────────────────────────────────────────
('ae000001-0000-0000-0000-000000000018', 'ad000001-0000-0000-0000-000000000005', 'cc000001-0000-0000-0000-000000000003', '7b000000-0000-0000-0000-000000000001', 'registered'), -- Zoe → Young Engineers (past)
('ae000001-0000-0000-0000-000000000019', 'ad000001-0000-0000-0000-000000000005', 'cc000001-0000-0000-0000-000000000003', '7b000000-0000-0000-0000-00000000000f', 'registered'), -- Zoe → Robotics Lab (past)
('ae000001-0000-0000-0000-00000000001a', 'ad000001-0000-0000-0000-000000000005', 'cc000001-0000-0000-0000-000000000003', '7b000000-0000-0000-0000-00000000000e', 'registered'), -- Zoe → Game Design (past)
('ae000001-0000-0000-0000-00000000001b', 'ad000001-0000-0000-0000-000000000005', 'cc000001-0000-0000-0000-000000000003', '7b000000-0000-0000-0000-00000000000a', 'registered'), -- Zoe → Watercolor (past)
('ae000001-0000-0000-0000-00000000001c', 'ad000001-0000-0000-0000-000000000005', 'cc000001-0000-0000-0000-000000000003', '7b000000-0000-0000-0000-00000000000c', 'registered'), -- Zoe → Comic Art (past)
('ae000001-0000-0000-0000-00000000001d', 'ad000001-0000-0000-0000-000000000005', 'cc000001-0000-0000-0000-000000000003', '7b000000-0000-0000-0000-000000000028', 'registered'), -- Zoe → Robotics Workshop Bangkok (past)
('ae000001-0000-0000-0000-00000000001e', 'ad000001-0000-0000-0000-000000000005', 'cc000001-0000-0000-0000-000000000003', '7b000000-0000-0000-0000-000000000002', 'registered'), -- Zoe → Science Explorers (past)

-- ── Marcus Webb — UPCOMING ───────────────────────────────────────────────────
('ae000001-0000-0000-0000-00000000001f', 'ad000001-0000-0000-0000-000000000005', 'cc000001-0000-0000-0000-000000000003', '7b000000-0000-0000-0000-000000000027', 'registered'), -- Zoe → Robotics Lab (upcoming)
('ae000001-0000-0000-0000-000000000020', 'ad000001-0000-0000-0000-000000000005', 'cc000001-0000-0000-0000-000000000003', '7b000000-0000-0000-0000-000000000013', 'registered'), -- Zoe → Young Engineers (upcoming)
('ae000001-0000-0000-0000-000000000021', 'ad000001-0000-0000-0000-000000000005', 'cc000001-0000-0000-0000-000000000003', '7b000000-0000-0000-0000-000000000026', 'registered'), -- Zoe → Game Design (upcoming)
('ae000001-0000-0000-0000-000000000022', 'ad000001-0000-0000-0000-000000000005', 'cc000001-0000-0000-0000-000000000003', '7b000000-0000-0000-0000-000000000022', 'registered'), -- Zoe → Comic Art (upcoming)

-- ── Marcus Webb — CANCELLED ──────────────────────────────────────────────────
('ae000001-0000-0000-0000-000000000023', 'ad000001-0000-0000-0000-000000000005', 'cc000001-0000-0000-0000-000000000003', '7b000000-0000-0000-0000-000000000025', 'cancelled');  -- Zoe → Web Design (cancelled)


-- ============================================
-- 22b. PAYMENTS for the registrations above
-- ============================================
-- Org Stripe account mapping:
--   MIT Kids Lab        → acct_mitkidslab
--   Boston Athletic     → acct_bostonathletic
--   Boston Art Center   → acct_bostonartcenter
--   Code & Create       → acct_codecreate
--   Siam Muay Thai      → acct_siammuaythai
--   Bangkok Ballet      → acct_bangkokballet
--   Geniuses STEM       → acct_geniusstem
--   Science Academy BKK → acct_seed_123  (existing)
--   Champions Sports BKK→ acct_seed_123  (existing)

INSERT INTO payment (
    registration_id, stripe_payment_intent_id, stripe_customer_id, org_stripe_account_id,
    stripe_payment_method_id, total_amount, provider_amount, platform_fee_amount,
    currency, payment_intent_status
) VALUES
-- Jennifer — past (requires_capture = completed/held)
('ae000001-0000-0000-0000-000000000001', 'pi_demo_jen_001', 'cus_jenniferpark', 'acct_mitkidslab',     'pm_demo_jen', 7500,  6375,  1125, 'usd', 'requires_capture'),
('ae000001-0000-0000-0000-000000000002', 'pi_demo_jen_002', 'cus_jenniferpark', 'acct_mitkidslab',     'pm_demo_jen', 7500,  6375,  1125, 'usd', 'requires_capture'),
('ae000001-0000-0000-0000-000000000003', 'pi_demo_jen_003', 'cus_jenniferpark', 'acct_mitkidslab',     'pm_demo_jen', 7500,  6375,  1125, 'usd', 'requires_capture'),
('ae000001-0000-0000-0000-000000000004', 'pi_demo_jen_004', 'cus_jenniferpark', 'acct_codecreate',     'pm_demo_jen', 7500,  6375,  1125, 'usd', 'requires_capture'),
('ae000001-0000-0000-0000-000000000005', 'pi_demo_jen_005', 'cus_jenniferpark', 'acct_codecreate',     'pm_demo_jen', 7500,  6375,  1125, 'usd', 'requires_capture'),
('ae000001-0000-0000-0000-000000000006', 'pi_demo_jen_006', 'cus_jenniferpark', 'acct_bostonathletic', 'pm_demo_jen', 5500,  4675,  825,  'usd', 'requires_capture'),
('ae000001-0000-0000-0000-000000000007', 'pi_demo_jen_007', 'cus_jenniferpark', 'acct_bostonathletic', 'pm_demo_jen', 5500,  4675,  825,  'usd', 'requires_capture'),
('ae000001-0000-0000-0000-000000000008', 'pi_demo_jen_008', 'cus_jenniferpark', 'acct_bostonartcenter','pm_demo_jen', 4500,  3825,  675,  'usd', 'requires_capture'),
-- Jennifer — upcoming
('ae000001-0000-0000-0000-000000000009', 'pi_demo_jen_009', 'cus_jenniferpark', 'acct_mitkidslab',     'pm_demo_jen', 7500,  6375,  1125, 'usd', 'requires_capture'),
('ae000001-0000-0000-0000-00000000000a', 'pi_demo_jen_00a', 'cus_jenniferpark', 'acct_mitkidslab',     'pm_demo_jen', 7500,  6375,  1125, 'usd', 'requires_capture'),
('ae000001-0000-0000-0000-00000000000b', 'pi_demo_jen_00b', 'cus_jenniferpark', 'acct_bostonathletic', 'pm_demo_jen', 5500,  4675,  825,  'usd', 'requires_capture'),
('ae000001-0000-0000-0000-00000000000c', 'pi_demo_jen_00c', 'cus_jenniferpark', 'acct_bostonartcenter','pm_demo_jen', 4500,  3825,  675,  'usd', 'requires_capture'),
-- Jennifer — cancelled (succeeded = refunded)
('ae000001-0000-0000-0000-00000000000d', 'pi_demo_jen_00d', 'cus_jenniferpark', 'acct_bostonathletic', 'pm_demo_jen', 5500,  4675,  825,  'usd', 'succeeded'),

-- Nattaya — past
('ae000001-0000-0000-0000-00000000000e', 'pi_demo_nat_001', 'cus_nattayath', 'acct_bangkokballet',  'pm_demo_nat', 40000, 34000, 6000, 'thb', 'requires_capture'),
('ae000001-0000-0000-0000-00000000000f', 'pi_demo_nat_002', 'cus_nattayath', 'acct_geniusstem',     'pm_demo_nat', 45000, 38250, 6750, 'thb', 'requires_capture'),
('ae000001-0000-0000-0000-000000000010', 'pi_demo_nat_003', 'cus_nattayath', 'acct_seed_123',       'pm_demo_nat', 50000, 42500, 7500, 'thb', 'requires_capture'),
('ae000001-0000-0000-0000-000000000011', 'pi_demo_nat_004', 'cus_nattayath', 'acct_seed_123',       'pm_demo_nat', 40000, 34000, 6000, 'thb', 'requires_capture'),
('ae000001-0000-0000-0000-000000000012', 'pi_demo_nat_005', 'cus_nattayath', 'acct_siammuaythai',   'pm_demo_nat', 35000, 29750, 5250, 'thb', 'requires_capture'),
('ae000001-0000-0000-0000-000000000013', 'pi_demo_nat_006', 'cus_nattayath', 'acct_seed_123',       'pm_demo_nat', 30000, 25500, 4500, 'thb', 'requires_capture'),
-- Nattaya — upcoming
('ae000001-0000-0000-0000-000000000014', 'pi_demo_nat_007', 'cus_nattayath', 'acct_bangkokballet',  'pm_demo_nat', 40000, 34000, 6000, 'thb', 'requires_capture'),
('ae000001-0000-0000-0000-000000000015', 'pi_demo_nat_008', 'cus_nattayath', 'acct_geniusstem',     'pm_demo_nat', 45000, 38250, 6750, 'thb', 'requires_capture'),
('ae000001-0000-0000-0000-000000000016', 'pi_demo_nat_009', 'cus_nattayath', 'acct_siammuaythai',   'pm_demo_nat', 35000, 29750, 5250, 'thb', 'requires_capture'),
('ae000001-0000-0000-0000-000000000017', 'pi_demo_nat_00a', 'cus_nattayath', 'acct_bostonathletic', 'pm_demo_nat', 5500,  4675,  825,  'usd', 'requires_capture'),

-- Marcus — past
('ae000001-0000-0000-0000-000000000018', 'pi_demo_mar_001', 'cus_marcuswebb', 'acct_mitkidslab',     'pm_demo_mar', 7500,  6375,  1125, 'usd', 'requires_capture'),
('ae000001-0000-0000-0000-000000000019', 'pi_demo_mar_002', 'cus_marcuswebb', 'acct_codecreate',     'pm_demo_mar', 7500,  6375,  1125, 'usd', 'requires_capture'),
('ae000001-0000-0000-0000-00000000001a', 'pi_demo_mar_003', 'cus_marcuswebb', 'acct_codecreate',     'pm_demo_mar', 7500,  6375,  1125, 'usd', 'requires_capture'),
('ae000001-0000-0000-0000-00000000001b', 'pi_demo_mar_004', 'cus_marcuswebb', 'acct_bostonartcenter','pm_demo_mar', 4500,  3825,  675,  'usd', 'requires_capture'),
('ae000001-0000-0000-0000-00000000001c', 'pi_demo_mar_005', 'cus_marcuswebb', 'acct_bostonartcenter','pm_demo_mar', 4500,  3825,  675,  'usd', 'requires_capture'),
('ae000001-0000-0000-0000-00000000001d', 'pi_demo_mar_006', 'cus_marcuswebb', 'acct_seed_123',       'pm_demo_mar', 50000, 42500, 7500, 'thb', 'requires_capture'),
('ae000001-0000-0000-0000-00000000001e', 'pi_demo_mar_007', 'cus_marcuswebb', 'acct_mitkidslab',     'pm_demo_mar', 7500,  6375,  1125, 'usd', 'requires_capture'),
-- Marcus — upcoming
('ae000001-0000-0000-0000-00000000001f', 'pi_demo_mar_008', 'cus_marcuswebb', 'acct_codecreate',     'pm_demo_mar', 7500,  6375,  1125, 'usd', 'requires_capture'),
('ae000001-0000-0000-0000-000000000020', 'pi_demo_mar_009', 'cus_marcuswebb', 'acct_mitkidslab',     'pm_demo_mar', 7500,  6375,  1125, 'usd', 'requires_capture'),
('ae000001-0000-0000-0000-000000000021', 'pi_demo_mar_00a', 'cus_marcuswebb', 'acct_codecreate',     'pm_demo_mar', 7500,  6375,  1125, 'usd', 'requires_capture'),
('ae000001-0000-0000-0000-000000000022', 'pi_demo_mar_00b', 'cus_marcuswebb', 'acct_bostonartcenter','pm_demo_mar', 4500,  3825,  675,  'usd', 'requires_capture'),
-- Marcus — cancelled
('ae000001-0000-0000-0000-000000000023', 'pi_demo_mar_00c', 'cus_marcuswebb', 'acct_codecreate',     'pm_demo_mar', 7500,  6375,  1125, 'usd', 'succeeded');
