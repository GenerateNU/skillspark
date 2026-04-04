UPDATE emergency_contacts
SET name = $2, guardian_id = $3, phone_number = $4, updated_at = NOW()
WHERE id = $1
RETURNING id, name, guardian_id, phone_number, created_at, updated_at