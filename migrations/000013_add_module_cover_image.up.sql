ALTER TABLE modules ADD COLUMN cover_image_key VARCHAR(500);
ALTER TABLE modules ADD COLUMN cover_image_status VARCHAR(20) NOT NULL DEFAULT 'pending';
