ALTER TABLE review
ALTER COLUMN guardian_id DROP NOT NULL;

-- keep reviews even if the guardian is deleted
ALTER TABLE review
DROP CONSTRAINT IF EXISTS review_guardian_id_fkey,
ADD CONSTRAINT review_guardian_id_fkey
    FOREIGN KEY (guardian_id) REFERENCES guardian(id) ON DELETE SET NULL;