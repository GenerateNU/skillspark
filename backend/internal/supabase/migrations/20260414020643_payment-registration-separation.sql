-- Create payment table to separate payment concerns from registration
CREATE TABLE IF NOT EXISTS payment (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    registration_id UUID NOT NULL REFERENCES registration(id) ON DELETE CASCADE,
    stripe_payment_intent_id VARCHAR(255) UNIQUE,
    stripe_customer_id VARCHAR(255),
    stripe_payment_method_id VARCHAR(255),
    org_stripe_account_id VARCHAR(255),
    total_amount INTEGER,
    provider_amount INTEGER,
    platform_fee_amount INTEGER,
    currency VARCHAR(3) DEFAULT 'thb',
    payment_intent_status payment_intent_status,
    paid_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Migrate existing payment data only if columns haven't already been dropped
-- (guards against re-running a partially applied migration)
DO $$
BEGIN
    IF EXISTS (
        SELECT 1 FROM information_schema.columns
        WHERE table_name = 'registration'
        AND column_name = 'stripe_payment_intent_id'
    ) THEN
        INSERT INTO payment (
            registration_id,
            stripe_payment_intent_id,
            stripe_customer_id,
            stripe_payment_method_id,
            org_stripe_account_id,
            total_amount,
            provider_amount,
            platform_fee_amount,
            currency,
            payment_intent_status,
            paid_at
        )
        SELECT
            id,
            stripe_payment_intent_id,
            stripe_customer_id,
            stripe_payment_method_id,
            org_stripe_account_id,
            total_amount,
            provider_amount,
            platform_fee_amount,
            currency,
            payment_intent_status,
            paid_at
        FROM registration
        WHERE stripe_payment_intent_id IS NOT NULL
           OR stripe_customer_id IS NOT NULL
           OR total_amount IS NOT NULL;
    END IF;
END $$;

-- Drop old indexes on payment columns in registration
DROP INDEX IF EXISTS idx_registration_payment_intent;
DROP INDEX IF EXISTS idx_registration_customer;
DROP INDEX IF EXISTS idx_registration_provider_account;
DROP INDEX IF EXISTS idx_registration_payment_status;

-- Remove payment columns from registration table
-- cancelled_at stays on registration as it tracks registration lifecycle, not payment
ALTER TABLE registration
DROP COLUMN IF EXISTS stripe_payment_intent_id,
DROP COLUMN IF EXISTS stripe_customer_id,
DROP COLUMN IF EXISTS stripe_payment_method_id,
DROP COLUMN IF EXISTS org_stripe_account_id,
DROP COLUMN IF EXISTS total_amount,
DROP COLUMN IF EXISTS provider_amount,
DROP COLUMN IF EXISTS platform_fee_amount,
DROP COLUMN IF EXISTS currency,
DROP COLUMN IF EXISTS payment_intent_status,
DROP COLUMN IF EXISTS paid_at;

-- Indexes on payment table
CREATE INDEX IF NOT EXISTS idx_payment_registration ON payment(registration_id);
CREATE INDEX IF NOT EXISTS idx_payment_intent ON payment(stripe_payment_intent_id);
CREATE INDEX IF NOT EXISTS idx_payment_customer ON payment(stripe_customer_id);
CREATE INDEX IF NOT EXISTS idx_payment_provider_account ON payment(org_stripe_account_id);
CREATE INDEX IF NOT EXISTS idx_payment_status ON payment(payment_intent_status);

-- updated_at trigger for payment table
CREATE TRIGGER update_payment_updated_at
BEFORE UPDATE ON payment
FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_column();
