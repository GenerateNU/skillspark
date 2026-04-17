WITH cancelled_reg AS (
    UPDATE registration
    SET
        status       = COALESCE($2::registration_status, 'cancelled'),
        cancelled_at = NOW(),
        updated_at   = NOW()
    WHERE id = $1
    RETURNING id, child_id, guardian_id, event_occurrence_id, status, cancelled_at, created_at, updated_at
),
updated_payment AS (
    UPDATE payment p
    SET
        payment_intent_status = COALESCE($3::payment_intent_status, p.payment_intent_status),
        updated_at            = NOW()
    FROM cancelled_reg cr
    WHERE p.registration_id = cr.id
    RETURNING
        p.registration_id,
        p.stripe_payment_intent_id,
        p.stripe_customer_id,
        p.org_stripe_account_id,
        p.stripe_payment_method_id,
        p.total_amount,
        p.provider_amount,
        p.platform_fee_amount,
        p.currency,
        p.payment_intent_status,
        p.paid_at
)
SELECT
    cr.id,
    cr.child_id,
    cr.guardian_id,
    cr.event_occurrence_id,
    cr.status,
    up.stripe_payment_intent_id,
    up.stripe_customer_id,
    up.org_stripe_account_id,
    up.stripe_payment_method_id,
    up.total_amount,
    up.provider_amount,
    up.platform_fee_amount,
    up.currency,
    up.payment_intent_status,
    up.paid_at,
    cr.cancelled_at,
    cr.created_at,
    cr.updated_at,
    e.title_en,
    e.title_th,
    eo.start_time AS occurrence_start_time
FROM cancelled_reg cr
LEFT JOIN updated_payment up ON up.registration_id = cr.id
JOIN event_occurrence eo ON eo.id = cr.event_occurrence_id
JOIN event e ON e.id = eo.event_id;
