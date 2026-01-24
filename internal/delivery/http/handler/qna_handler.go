// Package handler provides HTTP request handlers.
package handler

import "github.com/gin-gonic/gin"

// QnAHandler handles Q&A requests.
type QnAHandler struct{}

func NewQnAHandler() *QnAHandler { return &QnAHandler{} }

func (h *QnAHandler) CreateQuestion(c *gin.Context) { c.JSON(501, gin.H{"error": "not implemented"}) }
func (h *QnAHandler) UpdateQuestion(c *gin.Context) { c.JSON(501, gin.H{"error": "not implemented"}) }
func (h *QnAHandler) DeleteQuestion(c *gin.Context) { c.JSON(501, gin.H{"error": "not implemented"}) }
func (h *QnAHandler) ListQuestions(c *gin.Context)  { c.JSON(501, gin.H{"error": "not implemented"}) }
func (h *QnAHandler) CreateAnswer(c *gin.Context)   { c.JSON(501, gin.H{"error": "not implemented"}) }
func (h *QnAHandler) UpdateAnswer(c *gin.Context)   { c.JSON(501, gin.H{"error": "not implemented"}) }
func (h *QnAHandler) DeleteAnswer(c *gin.Context)   { c.JSON(501, gin.H{"error": "not implemented"}) }
func (h *QnAHandler) MarkBestAnswer(c *gin.Context) { c.JSON(501, gin.H{"error": "not implemented"}) }
