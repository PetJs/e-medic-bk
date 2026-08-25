CREATE TABLE IF NOT EXISTS lesson_quiz_questions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    lesson_id UUID NOT NULL REFERENCES lessons(id) ON DELETE CASCADE,
    type VARCHAR(20) NOT NULL, -- 'multiple_choice' | 'free_text'
    prompt TEXT NOT NULL,
    "order" INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_lesson_quiz_questions_lesson_id ON lesson_quiz_questions(lesson_id);
