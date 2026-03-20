SELECT g.id, g.user_id, g.stripe_customer_id, g.expo_push_token, g.created_at, g.updated_at
FROM guardian g
INNER JOIN "user" u ON g.user_id = u.id
WHERE u.auth_id = $1;