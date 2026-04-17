SELECT
    r.id,
    r.child_id,
    r.guardian_id,
    r.event_occurrence_id,
    r.status,
    r.created_at,
    r.updated_at,
    COALESCE(p.stripe_customer_id, '') AS stripe_customer_id,
    COALESCE(p.org_stripe_account_id, '') AS org_stripe_account_id,
    COALESCE(p.currency, '') AS currency,
    COALESCE(p.payment_intent_status::text, '') AS payment_intent_status,
    r.cancelled_at,
    COALESCE(p.stripe_payment_intent_id, '') AS stripe_payment_intent_id,
    COALESCE(p.total_amount, 0) AS total_amount,
    COALESCE(p.provider_amount, 0) AS provider_amount,
    COALESCE(p.platform_fee_amount, 0) AS platform_fee_amount,
    p.paid_at,
    COALESCE(p.stripe_payment_method_id, '') AS stripe_payment_method_id,
    e.title_en,
    e.title_th,
    eo.start_time AS occurrence_start_time
FROM registration r
LEFT JOIN payment p ON p.registration_id = r.id
JOIN event_occurrence eo ON r.event_occurrence_id = eo.id
JOIN event e ON eo.event_id = e.id
WHERE r.event_occurrence_id = $1
