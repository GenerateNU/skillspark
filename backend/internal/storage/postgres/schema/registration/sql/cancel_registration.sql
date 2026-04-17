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
    COALESCE(up.stripe_payment_intent_id, '') AS stripe_payment_intent_id,
    COALESCE(up.stripe_customer_id, '') AS stripe_customer_id,
    COALESCE(up.org_stripe_account_id, '') AS org_stripe_account_id,
    COALESCE(up.stripe_payment_method_id, '') AS stripe_payment_method_id,
    COALESCE(up.total_amount, 0) AS total_amount,
    COALESCE(up.provider_amount, 0) AS provider_amount,
    COALESCE(up.platform_fee_amount, 0) AS platform_fee_amount,
    COALESCE(up.currency, '') AS currency,
    COALESCE(up.payment_intent_status::text, '') AS payment_intent_status,
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
