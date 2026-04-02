ALTER TABLE emergency_contacts
    DROP CONSTRAINT emergency_contacts_guardian_id_fkey,
    ADD CONSTRAINT emergency_contacts_guardian_id_fkey
        FOREIGN KEY (guardian_id) REFERENCES guardian(id) ON DELETE CASCADE;
