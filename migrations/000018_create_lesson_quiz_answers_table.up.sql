CREATE TABLE IF NOT EXISTS lesson_quiz_answers (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    question_id UUID NOT NULL REFERENCES lesson_quiz_questions(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    selected_option_id UUID REFERENCES lesson_quiz_options(id) ON DELETE CASCADE,
    free_text_body TEXT,
    is_correct BOOLEAN, -- true/false for multiple_choice, null for free_text (ungraded)
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    UNIQUE(question_id, user_id)
);

CREATE INDEX idx_lesson_quiz_answers_user_id ON lesson_quiz_answers(user_id);
