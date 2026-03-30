SELECT 
    ec.id,
    ec.name,
    ec.guardian_id,
    ec.phone_number,
    ec.created_at,
    ec.updated_at
FROM emergency_contacts ec
WHERE ec.id = $1;