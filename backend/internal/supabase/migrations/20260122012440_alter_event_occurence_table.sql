ALTER TABLE event_occurrence
DROP CONSTRAINT event_occurrence_manager_id_fkey,
ADD CONSTRAINT event_occurrence_manager_id_fkey 
    FOREIGN KEY (manager_id) REFERENCES manager(id) ON DELETE SET NULL;