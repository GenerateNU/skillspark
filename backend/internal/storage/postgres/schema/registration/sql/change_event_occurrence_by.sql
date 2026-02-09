UPDATE event_occurrence
SET curr_enrolled = curr_enrolled + $2
WHERE id = $1
RETURNING curr_enrolled;