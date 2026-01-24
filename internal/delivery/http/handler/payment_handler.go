// Package handler provides HTTP request handlers.
package handler

import "github.com/gin-gonic/gin"

// PaymentHandler handles payment requests.
type PaymentHandler struct{}

func NewPaymentHandler() *PaymentHandler { return &PaymentHandler{} }

func (h *PaymentHandler) InitiatePayment(c *gin.Context) {
	c.JSON(501, gin.H{"error": "not implemented"})
}
func (h *PaymentHandler) ListPayments(c *gin.Context) { c.JSON(501, gin.H{"error": "not implemented"}) }
func (h *PaymentHandler) StripeWebhook(c *gin.Context) {
	c.JSON(501, gin.H{"error": "not implemented"})
}
func (h *PaymentHandler) PaystackWebhook(c *gin.Context) {
	c.JSON(501, gin.H{"error": "not implemented"})
}
