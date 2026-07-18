// Package handler provides HTTP request handlers.
package handler

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"emedic-bk/internal/application/dto"
	"emedic-bk/internal/application/usecase/payment"
	"emedic-bk/internal/domain/service"
)

// PaymentHandler handles payment requests.
type PaymentHandler struct {
	initiateUC *payment.InitiatePaymentUseCase
	verifyUC   *payment.VerifyPaymentUseCase
	listUC     *payment.ListPaymentsUseCase
	gateway    service.PaymentService
	plan       payment.PlanDetails
}

// NewPaymentHandler creates a new PaymentHandler.
func NewPaymentHandler(
	initiateUC *payment.InitiatePaymentUseCase,
	verifyUC *payment.VerifyPaymentUseCase,
	listUC *payment.ListPaymentsUseCase,
	gateway service.PaymentService,
	plan payment.PlanDetails,
) *PaymentHandler {
	return &PaymentHandler{
		initiateUC: initiateUC,
		verifyUC:   verifyUC,
		listUC:     listUC,
		gateway:    gateway,
		plan:       plan,
	}
}

// ListPlans returns the purchasable plans.
func (h *PaymentHandler) ListPlans(c *gin.Context) {
	c.JSON(http.StatusOK, []dto.PlanResponse{
		{
			ID:       payment.PlanMonthly,
			Name:     "Jerome's Monthly Access",
			Amount:   h.plan.Amount,
			Currency: h.plan.Currency,
			Interval: "month",
			Features: []string{
				"All Video Tutorials",
				"All Study Modules",
				"Community Access",
				"Expert Q&A Sessions",
			},
		},
	})
}

// InitiatePayment starts a Paystack hosted checkout.
func (h *PaymentHandler) InitiatePayment(c *gin.Context) {
	var req dto.InitiatePaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	response, err := h.initiateUC.Execute(c.Request.Context(), c.GetString("user_id"), &req)
	if err != nil {
		if errors.Is(err, payment.ErrUnknownPlan) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Unknown plan"})
			return
		}
		c.JSON(http.StatusBadGateway, gin.H{"error": "Failed to start payment"})
		return
	}
	c.JSON(http.StatusOK, response)
}

// VerifyPayment confirms a transaction after the checkout redirect.
func (h *PaymentHandler) VerifyPayment(c *gin.Context) {
	var req dto.VerifyPaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	response, err := h.verifyUC.Execute(c.Request.Context(), req.Reference)
	if err != nil {
		if errors.Is(err, payment.ErrPaymentNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
			return
		}
		c.JSON(http.StatusBadGateway, gin.H{"error": "Failed to verify payment"})
		return
	}
	c.JSON(http.StatusOK, response)
}

// ListPayments returns the authenticated user's payment history.
func (h *PaymentHandler) ListPayments(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	response, err := h.listUC.Execute(c.Request.Context(), c.GetString("user_id"), &dto.ListPaymentsRequest{
		Page:  page,
		Limit: limit,
	})
	if err != nil {
		internalError(c, "Failed to list payments", err)
		return
	}
	c.JSON(http.StatusOK, response)
}

// PaystackWebhook processes Paystack webhook events (charge.success).
func (h *PaymentHandler) PaystackWebhook(c *gin.Context) {
	body, err := io.ReadAll(io.LimitReader(c.Request.Body, 1<<20))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid body"})
		return
	}

	signature := c.GetHeader("x-paystack-signature")
	if !h.gateway.VerifyWebhookSignature(body, signature) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid signature"})
		return
	}

	var event struct {
		Event string `json:"event"`
		Data  struct {
			Reference string `json:"reference"`
		} `json:"data"`
	}
	if err := json.Unmarshal(body, &event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}

	if event.Event == "charge.success" && event.Data.Reference != "" {
		// VerifyPayment re-checks with Paystack and activates idempotently.
		if _, err := h.verifyUC.Execute(c.Request.Context(), event.Data.Reference); err != nil &&
			!errors.Is(err, payment.ErrPaymentNotFound) {
			internalError(c, "Processing failed", err)
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
