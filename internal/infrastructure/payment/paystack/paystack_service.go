// Package paystack provides Paystack payment gateway implementation.
package paystack

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"emedic-bk/internal/domain/service"
)

const baseURL = "https://api.paystack.co"

// PaymentService implements service.PaymentService using Paystack.
type PaymentService struct {
	secretKey string
	client    *http.Client
}

// NewPaymentService creates a new Paystack payment service.
func NewPaymentService(secretKey string) service.PaymentService {
	return &PaymentService{
		secretKey: secretKey,
		client:    &http.Client{Timeout: 30 * time.Second},
	}
}

type initializeRequest struct {
	Email       string `json:"email"`
	Amount      int64  `json:"amount"` // kobo
	Currency    string `json:"currency,omitempty"`
	Reference   string `json:"reference"`
	CallbackURL string `json:"callback_url,omitempty"`
}

type initializeResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    struct {
		AuthorizationURL string `json:"authorization_url"`
		AccessCode       string `json:"access_code"`
		Reference        string `json:"reference"`
	} `json:"data"`
}

type verifyResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Status    string     `json:"status"`
		Reference string     `json:"reference"`
		Amount    int64      `json:"amount"`
		Currency  string     `json:"currency"`
		PaidAt    *time.Time `json:"paid_at"`
	} `json:"data"`
}

func (s *PaymentService) InitializeTransaction(ctx context.Context, email string, amount int64, currency, reference, callbackURL string) (*service.CheckoutSession, error) {
	body, err := json.Marshal(initializeRequest{
		Email:       email,
		Amount:      amount,
		Currency:    currency,
		Reference:   reference,
		CallbackURL: callbackURL,
	})
	if err != nil {
		return nil, err
	}

	var res initializeResponse
	if err := s.do(ctx, http.MethodPost, "/transaction/initialize", bytes.NewReader(body), &res); err != nil {
		return nil, err
	}
	if !res.Status {
		return nil, fmt.Errorf("paystack initialize failed: %s", res.Message)
	}

	return &service.CheckoutSession{
		AuthorizationURL: res.Data.AuthorizationURL,
		AccessCode:       res.Data.AccessCode,
		Reference:        res.Data.Reference,
	}, nil
}

func (s *PaymentService) VerifyTransaction(ctx context.Context, reference string) (*service.TransactionStatus, error) {
	var res verifyResponse
	if err := s.do(ctx, http.MethodGet, "/transaction/verify/"+reference, nil, &res); err != nil {
		return nil, err
	}
	if !res.Status {
		return nil, fmt.Errorf("paystack verify failed: %s", res.Message)
	}

	status := &service.TransactionStatus{
		Reference: res.Data.Reference,
		Status:    res.Data.Status,
		Amount:    res.Data.Amount,
		Currency:  res.Data.Currency,
	}
	if res.Data.PaidAt != nil {
		status.PaidAt = *res.Data.PaidAt
	}
	return status, nil
}

// VerifyWebhookSignature validates Paystack's x-paystack-signature header
// (HMAC SHA512 of the raw body with the secret key).
func (s *PaymentService) VerifyWebhookSignature(payload []byte, signature string) bool {
	mac := hmac.New(sha512.New, []byte(s.secretKey))
	mac.Write(payload)
	expected := hex.EncodeToString(mac.Sum(nil))
	return hmac.Equal([]byte(expected), []byte(signature))
}

func (s *PaymentService) do(ctx context.Context, method, path string, body io.Reader, out any) error {
	req, err := http.NewRequestWithContext(ctx, method, baseURL+path, body)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+s.secretKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(io.LimitReader(resp.Body, 1<<20))
	if err != nil {
		return err
	}
	if resp.StatusCode >= 500 {
		return fmt.Errorf("paystack API error: %s", resp.Status)
	}
	return json.Unmarshal(data, out)
}
