CREATE TABLE IF NOT EXISTS lesson_quiz_options (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    question_id UUID NOT NULL REFERENCES lesson_quiz_questions(id) ON DELETE CASCADE,
    text VARCHAR(500) NOT NULL,
    is_correct BOOLEAN NOT NULL DEFAULT false,
    "order" INTEGER NOT NULL DEFAULT 0
);

CREATE INDEX idx_lesson_quiz_options_question_id ON lesson_quiz_options(question_id);
