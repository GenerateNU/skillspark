CREATE TABLE IF NOT EXISTS guardian_payment_methods (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    guardian_id UUID NOT NULL REFERENCES guardian(id) ON DELETE CASCADE,                 
    stripe_payment_method_id VARCHAR(255) NOT NULL UNIQUE,                            
    
    card_brand VARCHAR(50),                  
    card_last4 VARCHAR(4),                  
    card_exp_month INTEGER,                   
    card_exp_year INTEGER,           
    
    is_default BOOLEAN DEFAULT false NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW() NOT NULL,  
    updated_at TIMESTAMPTZ DEFAULT NOW() NOT NULL, 
    
    CONSTRAINT unique_guardian_payment_method UNIQUE(guardian_id, stripe_payment_method_id)
);

CREATE INDEX IF NOT EXISTS idx_guardian_payment_methods ON guardian_payment_methods(guardian_id);
CREATE INDEX IF NOT EXISTS idx_default_payment_method ON guardian_payment_methods(guardian_id, is_default) WHERE is_default = true;

ALTER TABLE registration
ADD COLUMN IF NOT EXISTS stripe_customer_id VARCHAR(255),
ADD COLUMN IF NOT EXISTS org_stripe_account_id VARCHAR(255),
ADD COLUMN IF NOT EXISTS currency VARCHAR(3) DEFAULT 'thb',
ADD COLUMN IF NOT EXISTS payment_intent_status VARCHAR(50),
ADD COLUMN IF NOT EXISTS cancelled_at TIMESTAMPTZ,
ADD COLUMN IF NOT EXISTS stripe_payment_intent_id VARCHAR(255) UNIQUE,
ADD COLUMN IF NOT EXISTS total_amount INTEGER,
ADD COLUMN IF NOT EXISTS provider_amount INTEGER,
ADD COLUMN IF NOT EXISTS platform_fee_amount INTEGER,
ADD COLUMN IF NOT EXISTS paid_at TIMESTAMPTZ,
ADD COLUMN IF NOT EXISTS stripe_payment_method_id VARCHAR(255);

-- Indexes
CREATE INDEX IF NOT EXISTS idx_registration_payment_intent ON registration(stripe_payment_intent_id);
CREATE INDEX IF NOT EXISTS idx_registration_customer ON registration(stripe_customer_id);
CREATE INDEX IF NOT EXISTS idx_registration_provider_account ON registration(org_stripe_account_id);
CREATE INDEX IF NOT EXISTS idx_registration_payment_status ON registration(payment_intent_status);
CREATE INDEX IF NOT EXISTS idx_registration_guardian ON registration(guardian_id);
CREATE INDEX IF NOT EXISTS idx_registration_event ON registration(event_occurrence_id);

ALTER TABLE organization 
ADD COLUMN IF NOT EXISTS stripe_account_id VARCHAR(255) UNIQUE,
ADD COLUMN IF NOT EXISTS stripe_account_activated BOOLEAN NOT NULL DEFAULT false;

ALTER TABLE guardian
ADD COLUMN IF NOT EXISTS stripe_customer_id VARCHAR(255) UNIQUE;


ALTER TABLE event_occurrence
ADD COLUMN IF NOT EXISTS price INTEGER NOT NULL DEFAULT 0; -- Price in bhat

CREATE INDEX IF NOT EXISTS idx_event_occurrence_price ON event_occurrence(price);

