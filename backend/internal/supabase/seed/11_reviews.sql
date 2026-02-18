INSERT INTO review (
    id,
    registration_id,
    guardian_id,
    description_en,
    description_th,
    categories
) VALUES

(
    '90000000-0000-0000-0000-000000000001',
    '80000000-0000-0000-0000-000000000001',
    '11111111-1111-1111-1111-111111111111',
    'Emily had a great time and came home excited to talk about everything she learned.',
    'เอมิลี่สนุกมากและกลับบ้านด้วยความตื่นเต้นที่จะเล่าทุกสิ่งที่เธอได้เรียนรู้',
    ARRAY['informative', 'interesting', 'engaging']::review_category[]
),
(
    '90000000-0000-0000-0000-000000000002',
    '80000000-0000-0000-0000-000000000002',
    '11111111-1111-1111-1111-111111111111',
    'The event was excellently organized. The activities were fun and educational.',
    'งานจัดได้อย่างดีเยี่ยม กิจกรรมสนุกและให้ความรู้',
    ARRAY['fun', 'informative']::review_category[]
),
(
    '90000000-0000-0000-0000-000000000003',
    '80000000-0000-0000-0000-000000000003',
    '11111111-1111-1111-1111-111111111111',
    'Engaging instructors and hands-on activities kept Emily interested the whole time.',
    'ครูผู้สอนที่น่าสนใจและกิจกรรมภาคปฏิบัติช่วยให้เอมิลี่สนใจตลอดเวลา',
    ARRAY['engaging', 'interesting']::review_category[]
),

(
    '90000000-0000-0000-0000-000000000004',
    '80000000-0000-0000-0000-000000000004',
    '11111111-1111-1111-1111-111111111111',
    'Alex really enjoyed the physical activities and music at the event.',
    'อเล็กซ์สนุกกับกิจกรรมทางกายภาพและดนตรีในงานเป็นอย่างมาก',
    ARRAY['fun', 'engaging']::review_category[]
);
