SELECT 
    r.id,
    r.child_id,
    r.guardian_id,
    r.event_occurrence_id,
    r.status,
    r.created_at,
    r.updated_at,
    r.stripe_customer_id,
    r.org_stripe_account_id,
    r.currency,
    r.payment_intent_status,
    r.cancelled_at,
    r.stripe_payment_intent_id,
    r.total_amount,
    r.provider_amount,
    r.platform_fee_amount,
    r.paid_at,
    r.stripe_payment_method_id,
    e.title AS event_name,
    eo.start_time AS occurrence_start_time
FROM registration r
JOIN event_occurrence eo ON r.event_occurrence_id = eo.id
JOIN event e ON eo.event_id = e.id
WHERE r.child_id = $1