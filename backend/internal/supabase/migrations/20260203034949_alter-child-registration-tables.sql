ALTER TABLE registration
ALTER child_id DROP NOT NULL;

ALTER TABLE registration
ALTER guardian_id DROP NOT NULL;

ALTER TABLE child
DROP CONSTRAINT IF EXISTS child_guardian_id_fkey,
ADD CONSTRAINT child_guardian_id_fkey
    FOREIGN KEY (guardian_id) REFERENCES guardian(id) ON DELETE CASCADE;

ALTER TABLE registration
DROP CONSTRAINT IF EXISTS registration_guardian_id_fkey,
ADD CONSTRAINT registration_guardian_id_fkey
    FOREIGN KEY (guardian_id) REFERENCES guardian(id) ON DELETE SET NULL;

ALTER TABLE registration
DROP CONSTRAINT IF EXISTS registration_child_id_fkey,
ADD CONSTRAINT registration_child_id_fkey
    FOREIGN KEY (child_id) REFERENCES child(id) ON DELETE SET NULL;