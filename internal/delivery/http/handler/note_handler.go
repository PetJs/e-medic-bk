// Package handler provides HTTP request handlers.
package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"emedic-bk/internal/application/dto"
	"emedic-bk/internal/application/usecase/notes"
)

// NoteHandler handles personal lesson-note requests.
type NoteHandler struct {
	createUC *notes.CreateNoteUseCase
	listUC   *notes.ListNotesUseCase
	deleteUC *notes.DeleteNoteUseCase
}

// NewNoteHandler creates a new NoteHandler.
func NewNoteHandler(createUC *notes.CreateNoteUseCase, listUC *notes.ListNotesUseCase, deleteUC *notes.DeleteNoteUseCase) *NoteHandler {
	return &NoteHandler{createUC: createUC, listUC: listUC, deleteUC: deleteUC}
}

// CreateNote creates a timestamped personal note on a lesson.
func (h *NoteHandler) CreateNote(c *gin.Context) {
	var req dto.CreateNoteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	userID := c.GetString("user_id")
	response, err := h.createUC.Execute(c.Request.Context(), c.Param("id"), userID, &req)
	if err != nil {
		if errors.Is(err, notes.ErrLessonNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Lesson not found"})
			return
		}
		internalError(c, "Failed to create note", err)
		return
	}
	c.JSON(http.StatusCreated, response)
}

// ListNotes lists the caller's own notes on a lesson.
func (h *NoteHandler) ListNotes(c *gin.Context) {
	userID := c.GetString("user_id")
	response, err := h.listUC.Execute(c.Request.Context(), userID, c.Param("id"))
	if err != nil {
		internalError(c, "Failed to list notes", err)
		return
	}
	c.JSON(http.StatusOK, response)
}

// DeleteNote deletes one of the caller's own notes.
func (h *NoteHandler) DeleteNote(c *gin.Context) {
	userID := c.GetString("user_id")
	if err := h.deleteUC.Execute(c.Request.Context(), userID, c.Param("id")); err != nil {
		if errors.Is(err, notes.ErrNoteNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
			return
		}
		internalError(c, "Failed to delete note", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Note deleted"})
}
