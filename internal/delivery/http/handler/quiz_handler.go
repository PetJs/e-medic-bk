// Package handler provides HTTP request handlers.
package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"emedic-bk/internal/application/dto"
	"emedic-bk/internal/application/usecase/quiz"
	"emedic-bk/internal/domain/repository"
)

// QuizHandler handles post-lesson quiz requests.
type QuizHandler struct {
	createQuestionUC     *quiz.CreateQuestionUseCase
	updateQuestionUC     *quiz.UpdateQuestionUseCase
	deleteQuestionUC     *quiz.DeleteQuestionUseCase
	listQuestionsUC      *quiz.ListQuestionsUseCase
	listQuestionsAdminUC *quiz.ListQuestionsAdminUseCase
	submitAnswerUC       *quiz.SubmitAnswerUseCase
	listMyAnswersUC      *quiz.ListMyAnswersUseCase
	listAnswersReviewUC  *quiz.ListAnswersForReviewUseCase
}

// NewQuizHandler creates a new QuizHandler.
func NewQuizHandler(
	createQuestionUC *quiz.CreateQuestionUseCase,
	updateQuestionUC *quiz.UpdateQuestionUseCase,
	deleteQuestionUC *quiz.DeleteQuestionUseCase,
	listQuestionsUC *quiz.ListQuestionsUseCase,
	listQuestionsAdminUC *quiz.ListQuestionsAdminUseCase,
	submitAnswerUC *quiz.SubmitAnswerUseCase,
	listMyAnswersUC *quiz.ListMyAnswersUseCase,
	listAnswersReviewUC *quiz.ListAnswersForReviewUseCase,
) *QuizHandler {
	return &QuizHandler{
		createQuestionUC:     createQuestionUC,
		updateQuestionUC:     updateQuestionUC,
		deleteQuestionUC:     deleteQuestionUC,
		listQuestionsUC:      listQuestionsUC,
		listQuestionsAdminUC: listQuestionsAdminUC,
		submitAnswerUC:       submitAnswerUC,
		listMyAnswersUC:      listMyAnswersUC,
		listAnswersReviewUC:  listAnswersReviewUC,
	}
}

// CreateQuestion creates a quiz question on a lesson (admin).
func (h *QuizHandler) CreateQuestion(c *gin.Context) {
	var req dto.QuizCreateQuestionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	response, err := h.createQuestionUC.Execute(c.Request.Context(), c.Param("id"), &req)
	if err != nil {
		switch {
		case errors.Is(err, quiz.ErrLessonNotFound):
			c.JSON(http.StatusNotFound, gin.H{"error": "Lesson not found"})
		case errors.Is(err, quiz.ErrInvalidQuestionType):
			c.JSON(http.StatusBadRequest, gin.H{"error": "type must be multiple_choice or free_text"})
		case errors.Is(err, quiz.ErrInvalidOptions):
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		default:
			internalError(c, "Failed to create question", err)
		}
		return
	}
	c.JSON(http.StatusCreated, response)
}

// UpdateQuestion updates a quiz question's prompt/order (admin).
func (h *QuizHandler) UpdateQuestion(c *gin.Context) {
	var req dto.QuizUpdateQuestionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	response, err := h.updateQuestionUC.Execute(c.Request.Context(), c.Param("id"), &req)
	if err != nil {
		if errors.Is(err, quiz.ErrQuestionNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Question not found"})
			return
		}
		internalError(c, "Failed to update question", err)
		return
	}
	c.JSON(http.StatusOK, response)
}

// DeleteQuestion deletes a quiz question (admin).
func (h *QuizHandler) DeleteQuestion(c *gin.Context) {
	if err := h.deleteQuestionUC.Execute(c.Request.Context(), c.Param("id")); err != nil {
		if errors.Is(err, quiz.ErrQuestionNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Question not found"})
			return
		}
		internalError(c, "Failed to delete question", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Question deleted"})
}

// ListQuestions lists a lesson's quiz questions (student-facing, no is_correct).
func (h *QuizHandler) ListQuestions(c *gin.Context) {
	response, err := h.listQuestionsUC.Execute(c.Request.Context(), c.Param("id"))
	if err != nil {
		if errors.Is(err, quiz.ErrLessonNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Lesson not found"})
			return
		}
		internalError(c, "Failed to list questions", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"questions": response})
}

// ListQuestionsAdmin lists a lesson's quiz questions (admin-facing, includes is_correct).
func (h *QuizHandler) ListQuestionsAdmin(c *gin.Context) {
	response, err := h.listQuestionsAdminUC.Execute(c.Request.Context(), c.Param("id"))
	if err != nil {
		if errors.Is(err, quiz.ErrLessonNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Lesson not found"})
			return
		}
		internalError(c, "Failed to list questions", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"questions": response})
}

// SubmitAnswer submits a student's answer to one question.
func (h *QuizHandler) SubmitAnswer(c *gin.Context) {
	var req dto.QuizSubmitAnswerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	userID := c.GetString("user_id")
	response, err := h.submitAnswerUC.Execute(c.Request.Context(), c.Param("id"), userID, &req)
	if err != nil {
		switch {
		case errors.Is(err, quiz.ErrQuestionNotFound):
			c.JSON(http.StatusNotFound, gin.H{"error": "Question not found"})
		case errors.Is(err, quiz.ErrOptionNotFound):
			c.JSON(http.StatusBadRequest, gin.H{"error": "Selected option not found on this question"})
		case errors.Is(err, quiz.ErrAnswerShapeMismatch):
			c.JSON(http.StatusBadRequest, gin.H{"error": "Answer does not match the question type"})
		case errors.Is(err, repository.ErrAlreadyAnswered):
			c.JSON(http.StatusConflict, gin.H{"error": "You've already answered this question"})
		default:
			internalError(c, "Failed to submit answer", err)
		}
		return
	}
	c.JSON(http.StatusCreated, response)
}

// ListMyAnswers returns the calling student's own answers for a lesson.
func (h *QuizHandler) ListMyAnswers(c *gin.Context) {
	response, err := h.listMyAnswersUC.Execute(c.Request.Context(), c.Param("id"), c.GetString("user_id"))
	if err != nil {
		if errors.Is(err, quiz.ErrLessonNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Lesson not found"})
			return
		}
		internalError(c, "Failed to list answers", err)
		return
	}
	c.JSON(http.StatusOK, response)
}

// ListAnswersForReview lists every student's answer to one question (admin).
func (h *QuizHandler) ListAnswersForReview(c *gin.Context) {
	response, err := h.listAnswersReviewUC.Execute(c.Request.Context(), c.Param("id"))
	if err != nil {
		if errors.Is(err, quiz.ErrQuestionNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Question not found"})
			return
		}
		internalError(c, "Failed to list answers", err)
		return
	}
	c.JSON(http.StatusOK, response)
}
