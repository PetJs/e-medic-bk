// Package entity contains the core domain entities.
package entity

// QuizOption is one choice on a multiple_choice QuizQuestion.
type QuizOption struct {
	ID         string
	QuestionID string
	Text       string
	IsCorrect  bool
	Order      int
}
