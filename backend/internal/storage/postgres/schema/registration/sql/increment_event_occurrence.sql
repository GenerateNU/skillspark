UPDATE event_occurrence
SET curr_enrolled = curr_enrolled + 1
WHERE id = $1
RETURNING curr_enrolled;