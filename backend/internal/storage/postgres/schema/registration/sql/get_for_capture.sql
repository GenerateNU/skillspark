SELECT
    r.id,
    r.child_id,
    r.guardian_id,
    r.event_occurrence_id,
    r.status,
    r.created_at,
    r.updated_at,
    p.stripe_customer_id,
    p.org_stripe_account_id,
    p.currency,
    p.payment_intent_status,
    r.cancelled_at,
    p.stripe_payment_intent_id,
    p.total_amount,
    p.provider_amount,
    p.platform_fee_amount,
    p.paid_at,
    p.stripe_payment_method_id,
    e.title_en,
    e.title_th,
    eo.start_time AS occurrence_start_time
FROM registration r
JOIN payment p ON p.registration_id = r.id
JOIN event_occurrence eo ON r.event_occurrence_id = eo.id
JOIN event e ON eo.event_id = e.id
WHERE p.payment_intent_status = 'requires_capture'
  AND r.status = 'registered'
  AND eo.start_time BETWEEN $1 AND $2
ORDER BY eo.start_time ASC;
