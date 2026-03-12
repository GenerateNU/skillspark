SELECT
    r.id,
    r.child_id,
    r.guardian_id,
    r.event_occurrence_id,
    r.status,
    r.stripe_payment_intent_id,
    r.stripe_customer_id,
    r.org_stripe_account_id,
    r.stripe_payment_method_id,
    r.total_amount,
    r.provider_amount,
    r.platform_fee_amount,
    r.currency,
    r.payment_intent_status,
    r.paid_at,
    r.cancelled_at,
    r.created_at,
    r.updated_at,
    e.title_en,
    e.title_th,
    eo.start_time AS occurrence_start_time
FROM registration r
JOIN event_occurrence eo ON eo.id = r.event_occurrence_id
JOIN event e ON e.id = eo.event_id
WHERE r.stripe_payment_intent_id = $1;