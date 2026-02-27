# SkillSpark Payment Integration

SkillSpark uses [Stripe Connect](https://stripe.com/connect) to handle payments between guardians and activity providers (organizations). The platform acts as a marketplace — collecting payments from guardians and distributing funds to organizations, while retaining a platform fee.

---

## Overview

```
Guardian pays → SkillSpark Platform → Organization receives (minus platform fee & Stripe fee)
```

- **Currencies**: THB amd USD 
- **Capture strategy**: Manual capture 24 hours before the event
- **Platform fee (what we take)**: 10% of the total charge (may change in the future)
- **Stripe fee**: The amount Stripe takes after capture - since this is done after capture, it is currently not stored

---

## Key Concepts

### Stripe Connect
SkillSpark uses **destination charges**, meaning:
- The payment intent is created on the **platform account**
- Funds are automatically transferred to the **connected organization account**
- The platform retains `application_fee_amount` (10%)

### Manual Capture
Payment intents are created with `capture_method: manual`, meaning:
- The guardian's card is **authorized but not charged** at registration time
- The charge is **captured 24 hours before the event** via a cron job
- If the registration is cancelled before capture, the authorization is released at no cost to the customer

---

## Flows

### 1. Organization Onboarding
Before an organization can receive payments, they must connect a Stripe account.

```
POST /api/v1/stripe/orgaccount/{organization_id}
→ Creates a Stripe Express account for the organization

POST /api/v1/stripe/onboarding/{organization_id}
→ Generates a Stripe-hosted onboarding link
→ Organization fills in bank account details and other important information on Stripe

[Webhook] account.updated
→ Stripe notifies when charges_enabled and payouts_enabled are both true
→ Organization's stripe_account_activated is set to true in the DB

POST /api/v1/stripe/login/{organization_id}
→ Used by the organization to log in to their Stripe account when needed
```

### 2. Guardian Setup
Before a guardian can register for activities, they need a Stripe customer account and a saved payment method.

```
POST /api/v1/stripe/customer/{guardian_id}
→ Creates a Stripe Customer linked to the guardian

POST /api/v1/stripe/setup-intent/{guardian_id}
→ Creates a SetupIntent and returns client_secret to the frontend
→ Frontend uses Stripe.js to collect and save card details
→ Card is attached to the guardian's Stripe customer once Stripe.js component successfully collects card details

GET /api/v1/guardians/{guardian_id}/payment-methods
→ Returns all saved payment methods for a guardian
```

### 3. Registration & Payment Authorization
When a guardian registers a child for an activity:

```
POST /api/v1/registrations
{
  "child_id": "...",
  "guardian_id": "...",
  "event_occurrence_id": "...",
  "payment_method_id": "pm_...",
  "status": "registered"
}
```

The handler:
1. Validates the event occurrence (not started, not at capacity)
2. Validates the guardian has a Stripe customer ID
3. Validates the child belongs to the guardian
4. Fetches the organization's Stripe account ID
5. Creates a payment intent with:
   - `amount` and `currency` from the event occurrence
   - `customer` = guardian's Stripe customer ID
   - `payment_method` = provided payment method
   - `transfer_data.destination` = organization's Stripe account
   - `application_fee_amount` = 10% of the total cost of the event
   - `capture_method: manual`
   - `confirm: true`
6. Saves the registration with payment intent details

The payment intent status at this point is `requires_capture`.

### 4. Payment Capture (Cron Job)
A cron job runs every hour and captures payments for events starting in the next 24–25 hours.

```
Schedule: 0 * * * * (every hour)
Window: events starting between now+24h and now+25h
```

For each registration in the window:
1. Calls `CapturePaymentIntent` on Stripe
2. Updates the registration's `payment_intent_status` to `succeeded`

### 5. Cancellation & Refunds
```
POST /api/v1/registrations/{id}/cancel
```

Cancellation behavior depends on the payment intent status:

| Status | Action |
|--------|--------|
| `requires_capture` | Cancel the payment intent (no charge, authorization released) |
| `succeeded` + event > 24hrs away | Issue a full refund (minus Stripe fee) |
| `succeeded` + event within 24hrs | No refund |

### 6. Event Occurrence Cancellation
When an entire event occurrence is cancelled:

```
POST /api/v1/event-occurrences/{id}/cancel
```

1. Fetches all registrations for the event occurrence
2. Cancels the event occurrence in the DB (decrements `curr_enrolled` atomically)
3. For each registration, applies the same cancellation/refund logic as individual cancellations

### 7. Failed Payments (Webhook)
If a payment capture fails (e.g. card declined):

```
[Webhook] payment_intent.payment_failed
→ Registration is automatically cancelled
→ curr_enrolled is decremented
```

---

## Webhook Endpoints

| Endpoint | Events |
|----------|--------|
| `POST /api/v1/webhooks/stripe` | `payment_intent.payment_failed` |
| `POST /api/v1/webhooks/stripe/account` | `account.updated` |

Webhook signature verification is performed using `STRIPE_WEBHOOK_SECRET` and `STRIPE_ACCOUNT_WEBHOOK_SECRET`.

To test the webhooks, it's a little chopped icl.
You need to run stripe 
```stripe listen --forward-to localhost:8080/api/v1/webhooks/stripe```
Put the signing secret your receive in your .env file as STRIPE_WEBHOOK_SECRET


In a separate terminal, run stripe 
```stripe listen --forward-to localhost:8080/api/v1/webhooks/stripe/account```
Put the signing secret your receive in your .env file as STRIPE_ACCOUNT_WEBHOOK_SECRET
This listener is for the webhooks related to account, so it registers that the account has updated, so when you update anything about the account, it runs and updates whether the account can receive payments.

<b>You need to update the keys in your env every time you run the command</b>

---

## Fee Structure

For a ฿50,000 activity:

| Party | Amount |
|-------|--------|
| Guardian pays | ฿50,000 |
| Stripe fee (3.65% + ฿10) | ฿1,835 |
| Platform fee (10%) | ฿5,000 |
| Organization receives | ฿43,165 |

---

## Environment Variables

| Variable | Description |
|----------|-------------|
| `STRIPE_SECRET_KEY` | Stripe platform secret key |
| `STRIPE_WEBHOOK_SECRET` | Signing secret for platform webhooks |
| `STRIPE_ACCOUNT_WEBHOOK_SECRET` | Signing secret for Account webhooks |
