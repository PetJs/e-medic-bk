// Package handler provides HTTP request handlers.
package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"emedic-bk/internal/application/dto"
	"emedic-bk/internal/application/usecase/content"
)

// maxUploadSize caps admin content uploads (2 GiB).
const maxUploadSize = 2 << 30

// ContentHandler handles content requests.
type ContentHandler struct {
	uploadUC *content.UploadContentUseCase
	getURLUC *content.GetContentURLUseCase
	deleteUC *content.DeleteContentUseCase
}

// NewContentHandler creates a new ContentHandler.
func NewContentHandler(
	uploadUC *content.UploadContentUseCase,
	getURLUC *content.GetContentURLUseCase,
	deleteUC *content.DeleteContentUseCase,
) *ContentHandler {
	return &ContentHandler{uploadUC: uploadUC, getURLUC: getURLUC, deleteUC: deleteUC}
}

// UploadContent accepts a multipart upload (fields: lesson_id, title, file).
func (h *ContentHandler) UploadContent(c *gin.Context) {
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxUploadSize)

	lessonID := c.PostForm("lesson_id")
	title := c.PostForm("title")
	fileHeader, err := c.FormFile("file")
	if err != nil || lessonID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "lesson_id and file are required"})
		return
	}
	if title == "" {
		title = fileHeader.Filename
	}

	file, err := fileHeader.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not read uploaded file"})
		return
	}
	defer file.Close()

	response, err := h.uploadUC.Execute(c.Request.Context(), &dto.UploadContentRequest{
		LessonID:    lessonID,
		Title:       title,
		File:        file,
		FileName:    fileHeader.Filename,
		ContentType: fileHeader.Header.Get("Content-Type"),
		Size:        fileHeader.Size,
	})
	if err != nil {
		switch {
		case errors.Is(err, content.ErrLessonNotFound):
			c.JSON(http.StatusNotFound, gin.H{"error": "Lesson not found"})
		case errors.Is(err, content.ErrUnsupportedType):
			c.JSON(http.StatusBadRequest, gin.H{"error": "Only video, PDF, and image files are supported"})
		default:
			internalError(c, "Upload failed — check storage configuration", err)
		}
		return
	}
	c.JSON(http.StatusCreated, response)
}

// GetContentURL returns a short-lived signed streaming URL.
func (h *ContentHandler) GetContentURL(c *gin.Context) {
	userID := c.GetString("user_id")
	role := c.GetString("role")

	response, err := h.getURLUC.Execute(c.Request.Context(), c.Param("id"), userID, role)
	if err != nil {
		switch {
		case errors.Is(err, content.ErrContentNotFound):
			c.JSON(http.StatusNotFound, gin.H{"error": "Content not found"})
		case errors.Is(err, content.ErrSubscriptionRequired):
			c.JSON(http.StatusPaymentRequired, gin.H{"error": "Active subscription required"})
		default:
			internalError(c, "Failed to generate stream URL", err)
		}
		return
	}
	c.JSON(http.StatusOK, response)
}

// DeleteContent removes content and its stored object (admin).
func (h *ContentHandler) DeleteContent(c *gin.Context) {
	if err := h.deleteUC.Execute(c.Request.Context(), c.Param("id")); err != nil {
		if errors.Is(err, content.ErrContentNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Content not found"})
			return
		}
		internalError(c, "Failed to delete content", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Content deleted"})
}
