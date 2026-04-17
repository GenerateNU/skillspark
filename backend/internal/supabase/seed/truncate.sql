-- ============================================
-- TRUNCATE ALL SEED DATA
-- Clears every table and resets sequences.
-- RESTART IDENTITY resets serial/sequence counters.
-- CASCADE handles all foreign key dependencies automatically.
-- ============================================
TRUNCATE TABLE
    payment,
    review,
    saved,
    emergency_contacts,
    scheduled_notification,
    registration,
    event_occurrence,
    event,
    manager,
    child,
    organization,
    school,
    guardian,
    "user",
    location
RESTART IDENTITY CASCADE;
