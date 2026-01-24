// Package dto contains Data Transfer Objects for use cases.
package dto

import "time"

// CreateQuestionRequest represents a question creation request.
type CreateQuestionRequest struct {
	LessonID string `json:"lesson_id" validate:"required"`
	Title    string `json:"title" validate:"required"`
	Body     string `json:"body" validate:"required"`
}

// UpdateQuestionRequest represents a question update request.
type UpdateQuestionRequest struct {
	Title *string `json:"title,omitempty"`
	Body  *string `json:"body,omitempty"`
}

// QuestionResponse represents a question in API responses.
type QuestionResponse struct {
	ID          string            `json:"id"`
	LessonID    string            `json:"lesson_id"`
	UserID      string            `json:"user_id"`
	User        *UserResponse     `json:"user,omitempty"`
	Title       string            `json:"title"`
	Body        string            `json:"body"`
	AnswerCount int               `json:"answer_count"`
	IsResolved  bool              `json:"is_resolved"`
	Answers     []*AnswerResponse `json:"answers,omitempty"`
	CreatedAt   time.Time         `json:"created_at"`
}

// CreateAnswerRequest represents an answer creation request.
type CreateAnswerRequest struct {
	QuestionID string `json:"question_id" validate:"required"`
	Body       string `json:"body" validate:"required"`
}

// UpdateAnswerRequest represents an answer update request.
type UpdateAnswerRequest struct {
	Body *string `json:"body,omitempty"`
}

// AnswerResponse represents an answer in API responses.
type AnswerResponse struct {
	ID         string        `json:"id"`
	QuestionID string        `json:"question_id"`
	UserID     string        `json:"user_id"`
	User       *UserResponse `json:"user,omitempty"`
	Body       string        `json:"body"`
	IsBest     bool          `json:"is_best"`
	CreatedAt  time.Time     `json:"created_at"`
}

// ListQuestionsRequest represents a list questions request.
type ListQuestionsRequest struct {
	Page  int `json:"page" validate:"min=1"`
	Limit int `json:"limit" validate:"min=1,max=100"`
}

// ListQuestionsResponse represents a list questions response.
type ListQuestionsResponse struct {
	Questions  []*QuestionResponse `json:"questions"`
	TotalCount int64               `json:"total_count"`
	Page       int                 `json:"page"`
	Limit      int                 `json:"limit"`
}
