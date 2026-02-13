INSERT INTO review (
    id,
    registration_id,
    guardian_id,
    description,
    categories
) VALUES

(
    '90000000-0000-0000-0000-000000000001',
    '80000000-0000-0000-0000-000000000001',
    '11111111-1111-1111-1111-111111111111',
    'Emily had a great time and came home excited to talk about everything she learned.',
    ARRAY['informative', 'interesting', 'engaging']::review_category[]
),
(
    '90000000-0000-0000-0000-000000000002',
    '80000000-0000-0000-0000-000000000002',
    '11111111-1111-1111-1111-111111111111',
    'Very well organized event. The activities were fun and educational.',
    ARRAY['fun', 'informative']::review_category[]
),
(
    '90000000-0000-0000-0000-000000000003',
    '80000000-0000-0000-0000-000000000003',
    '11111111-1111-1111-1111-111111111111',
    'Engaging instructors and hands-on activities kept Emily interested the whole time.',
    ARRAY['engaging', 'interesting']::review_category[]
),

(
    '90000000-0000-0000-0000-000000000004',
    '80000000-0000-0000-0000-000000000004',
    '11111111-1111-1111-1111-111111111111',
    'Alex really enjoyed the physical activities and music portions of the event.',
    ARRAY['fun', 'engaging']::review_category[]
);
