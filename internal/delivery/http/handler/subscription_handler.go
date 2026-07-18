// Package handler provides HTTP request handlers.
package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"emedic-bk/internal/application/usecase/subscription"
)

// SubscriptionHandler handles subscription requests.
type SubscriptionHandler struct {
	listUC   *subscription.ListSubscriptionsUseCase
	cancelUC *subscription.CancelSubscriptionUseCase
}

// NewSubscriptionHandler creates a new SubscriptionHandler.
func NewSubscriptionHandler(
	listUC *subscription.ListSubscriptionsUseCase,
	cancelUC *subscription.CancelSubscriptionUseCase,
) *SubscriptionHandler {
	return &SubscriptionHandler{listUC: listUC, cancelUC: cancelUC}
}

// List returns the authenticated user's subscriptions.
func (h *SubscriptionHandler) List(c *gin.Context) {
	response, err := h.listUC.Execute(c.Request.Context(), c.GetString("user_id"))
	if err != nil {
		internalError(c, "Failed to list subscriptions", err)
		return
	}
	c.JSON(http.StatusOK, response)
}

// Cancel cancels one of the user's subscriptions.
func (h *SubscriptionHandler) Cancel(c *gin.Context) {
	err := h.cancelUC.Execute(c.Request.Context(), c.GetString("user_id"), c.Param("id"))
	if err != nil {
		if errors.Is(err, subscription.ErrSubscriptionNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Subscription not found"})
			return
		}
		internalError(c, "Failed to cancel subscription", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Subscription canceled"})
}
