WITH updated_payment AS (
    UPDATE payment
    SET
        payment_intent_status = $2::payment_intent_status,
        paid_at = CASE
            WHEN $2 = 'succeeded' THEN NOW()
            ELSE paid_at
        END,
        updated_at = NOW()
    WHERE registration_id = $1
    RETURNING
        registration_id,
        stripe_customer_id,
        org_stripe_account_id,
        currency,
        payment_intent_status,
        stripe_payment_intent_id,
        total_amount,
        provider_amount,
        platform_fee_amount,
        paid_at,
        stripe_payment_method_id
)
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
JOIN updated_payment p ON p.registration_id = r.id
JOIN event_occurrence eo ON r.event_occurrence_id = eo.id
JOIN event e ON eo.event_id = e.id;
