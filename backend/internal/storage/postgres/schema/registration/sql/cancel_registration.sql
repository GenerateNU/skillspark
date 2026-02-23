WITH cancelled_registration AS (
    UPDATE registration r
    SET
        status                = COALESCE($2::registration_status, 'cancelled'),
        cancelled_at          = NOW(),
        payment_intent_status = COALESCE($3::payment_intent_status, r.payment_intent_status),
        updated_at            = NOW()
    WHERE r.id = $1
    RETURNING
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
        r.updated_at
)
SELECT
    cr.id,
    cr.child_id,
    cr.guardian_id,
    cr.event_occurrence_id,
    cr.status,
    cr.stripe_payment_intent_id,
    cr.stripe_customer_id,
    cr.org_stripe_account_id,
    cr.stripe_payment_method_id,
    cr.total_amount,
    cr.provider_amount,
    cr.platform_fee_amount,
    cr.currency,
    cr.payment_intent_status,
    cr.paid_at,
    cr.cancelled_at,
    cr.created_at,
    cr.updated_at,
    e.title       AS event_name,
    eo.start_time AS occurrence_start_time
FROM cancelled_registration cr
JOIN event_occurrence eo ON eo.id = cr.event_occurrence_id
JOIN event e ON e.id = eo.event_id;