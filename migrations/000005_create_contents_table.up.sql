-- Create contents table
CREATE TABLE IF NOT EXISTS contents (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    lesson_id UUID NOT NULL REFERENCES lessons(id) ON DELETE CASCADE,
    type VARCHAR(20) NOT NULL, -- 'pdf' or 'video'
    title VARCHAR(255) NOT NULL,
    storage_key VARCHAR(500) NOT NULL,
    mime_type VARCHAR(100) NOT NULL,
    size BIGINT NOT NULL DEFAULT 0, -- in bytes
    duration INTEGER NOT NULL DEFAULT 0, -- for videos, in seconds
    "order" INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_contents_lesson_id ON contents(lesson_id);
CREATE INDEX idx_contents_type ON contents(type);
