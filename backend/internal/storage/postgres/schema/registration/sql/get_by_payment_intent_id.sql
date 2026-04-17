SELECT
    r.id,
    r.child_id,
    r.guardian_id,
    r.event_occurrence_id,
    r.status,
    p.stripe_payment_intent_id,
    p.stripe_customer_id,
    p.org_stripe_account_id,
    p.stripe_payment_method_id,
    p.total_amount,
    p.provider_amount,
    p.platform_fee_amount,
    p.currency,
    p.payment_intent_status,
    p.paid_at,
    r.cancelled_at,
    r.created_at,
    r.updated_at,
    e.title_en,
    e.title_th,
    eo.start_time AS occurrence_start_time
FROM registration r
JOIN payment p ON p.registration_id = r.id
JOIN event_occurrence eo ON eo.id = r.event_occurrence_id
JOIN event e ON e.id = eo.event_id
WHERE p.stripe_payment_intent_id = $1;
