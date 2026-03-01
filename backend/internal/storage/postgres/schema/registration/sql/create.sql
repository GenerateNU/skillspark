WITH inserted AS (
    INSERT INTO registration (
        child_id, 
        guardian_id, 
        event_occurrence_id, 
        status,
        stripe_payment_intent_id,
        stripe_customer_id,
        org_stripe_account_id,
        stripe_payment_method_id,
        total_amount,
        provider_amount,
        platform_fee_amount,
        currency,
        payment_intent_status
    )
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
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
    i.id,
    i.child_id,
    i.guardian_id,
    i.event_occurrence_id,
    i.status,
    i.created_at,
    i.updated_at,
    i.stripe_customer_id,
    i.org_stripe_account_id,
    i.currency,
    i.payment_intent_status,
    i.cancelled_at,
    i.stripe_payment_intent_id,
    i.total_amount,
    i.provider_amount,
    i.platform_fee_amount,
    i.paid_at,
    i.stripe_payment_method_id,
    e.title AS event_name,
    eo.start_time AS occurrence_start_time
FROM inserted i
JOIN event_occurrence eo ON i.event_occurrence_id = eo.id
JOIN event e ON eo.event_id = e.id;