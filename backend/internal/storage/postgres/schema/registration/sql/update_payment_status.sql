WITH updated AS (
    UPDATE registration
    SET
        payment_intent_status = $2::payment_intent_status,
        paid_at = CASE 
            WHEN $2 = 'succeeded' THEN NOW() 
            ELSE paid_at 
        END,
        updated_at = NOW()
    WHERE id = $1
    RETURNING 
        id, 
        child_id, 
        guardian_id, 
        event_occurrence_id, 
        status, 
        created_at, 
        updated_at,
        stripe_customer_id,
        org_stripe_account_id,
        currency,
        payment_intent_status,
        cancelled_at,
        stripe_payment_intent_id,
        total_amount,
        provider_amount,
        platform_fee_amount,
        paid_at,
        stripe_payment_method_id
)
SELECT 
    u.id,
    u.child_id,
    u.guardian_id,
    u.event_occurrence_id,
    u.status,
    u.created_at,
    u.updated_at,
    u.stripe_customer_id,
    u.org_stripe_account_id,
    u.currency,
    u.payment_intent_status,
    u.cancelled_at,
    u.stripe_payment_intent_id,
    u.total_amount,
    u.provider_amount,
    u.platform_fee_amount,
    u.paid_at,
    u.stripe_payment_method_id,
    e.title AS event_name,
    eo.start_time AS occurrence_start_time
FROM updated u
JOIN event_occurrence eo ON u.event_occurrence_id = eo.id
JOIN event e ON eo.event_id = e.id;