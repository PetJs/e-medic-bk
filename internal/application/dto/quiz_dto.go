// Package dto contains Data Transfer Objects for use cases.
package dto

import "time"

// QuizCreateQuestionRequest represents a quiz question creation request.
// The lesson is identified by the URL path, not the body. Type is validated
// manually in the usecase (not by a struct tag — see quiz.ErrInvalidQuestionType).
type QuizCreateQuestionRequest struct {
	Type    string                     `json:"type" binding:"required"` // "multiple_choice" | "free_text"
	Prompt  string                     `json:"prompt" binding:"required"`
	Order   int                        `json:"order"`
	Options []QuizCreateOptionRequest  `json:"options,omitempty"` // required (>=2, exactly one is_correct) iff type=multiple_choice
}

// QuizCreateOptionRequest represents one option on a multiple_choice question.
type QuizCreateOptionRequest struct {
	Text      string `json:"text" binding:"required"`
	IsCorrect bool   `json:"is_correct"`
}

// QuizUpdateQuestionRequest represents a quiz question update request.
// Options are not editable after creation in this slice — delete and
// recreate the question if the options need to change.
type QuizUpdateQuestionRequest struct {
	Prompt *string `json:"prompt,omitempty"`
	Order  *int    `json:"order,omitempty"`
}

// QuizOptionResponse is the student-facing option shape — no is_correct,
// so a student can't see the answer in the network response before answering.
type QuizOptionResponse struct {
	ID    string `json:"id"`
	Text  string `json:"text"`
	Order int    `json:"order"`
}

// QuizOptionAdminResponse is the admin-facing option shape, including is_correct.
type QuizOptionAdminResponse struct {
	ID        string `json:"id"`
	Text      string `json:"text"`
	IsCorrect bool   `json:"is_correct"`
	Order     int    `json:"order"`
}

// QuizQuestionResponse is the student-facing question shape.
type QuizQuestionResponse struct {
	ID       string                `json:"id"`
	LessonID string                `json:"lesson_id"`
	Type     string                `json:"type"`
	Prompt   string                `json:"prompt"`
	Order    int                   `json:"order"`
	Options  []*QuizOptionResponse `json:"options,omitempty"`
}

// QuizQuestionAdminResponse is the admin-facing question shape, including is_correct on options.
type QuizQuestionAdminResponse struct {
	ID       string                     `json:"id"`
	LessonID string                     `json:"lesson_id"`
	Type     string                     `json:"type"`
	Prompt   string                     `json:"prompt"`
	Order    int                        `json:"order"`
	Options  []*QuizOptionAdminResponse `json:"options,omitempty"`
}

// QuizSubmitAnswerRequest represents a student's answer submission.
// The question is identified by the URL path.
type QuizSubmitAnswerRequest struct {
	SelectedOptionID *string `json:"selected_option_id,omitempty"` // multiple_choice
	FreeTextBody     *string `json:"free_text_body,omitempty"`     // free_text
}

// QuizAnswerResponse represents a submitted quiz answer in API responses.
// Author is populated only on the admin review listing.
type QuizAnswerResponse struct {
	ID               string        `json:"id"`
	QuestionID       string        `json:"question_id"`
	UserID           string        `json:"user_id"`
	SelectedOptionID *string       `json:"selected_option_id,omitempty"`
	FreeTextBody     *string       `json:"free_text_body,omitempty"`
	IsCorrect        *bool         `json:"is_correct,omitempty"`
	Author           *UserResponse `json:"author,omitempty"`
	CreatedAt        time.Time     `json:"created_at"`
}

// QuizListAnswersResponse represents a list of quiz answers (own answers, or
// an admin's review of every student's answer to one question).
type QuizListAnswersResponse struct {
	Answers []*QuizAnswerResponse `json:"answers"`
}
