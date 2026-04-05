Insert into emergency_contacts(name, guardian_id, phone_number)
VALUES ($1, $2, $3)
RETURNING id, name, guardian_id, phone_number, created_at, updated_at