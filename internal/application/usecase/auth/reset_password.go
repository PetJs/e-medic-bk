// Package auth contains authentication use cases.
package auth

import (
	"context"

	"emedic-bk/internal/application/dto"
)

// ResetPasswordUseCase handles password reset.
type ResetPasswordUseCase struct {
	// TODO: Add dependencies
}

// NewResetPasswordUseCase creates a new ResetPasswordUseCase.
func NewResetPasswordUseCase() *ResetPasswordUseCase {
	return &ResetPasswordUseCase{}
}

// RequestReset sends a password reset email.
func (uc *ResetPasswordUseCase) RequestReset(ctx context.Context, req *dto.ResetPasswordRequest) error {
	// TODO: Find user by email
	// TODO: Generate reset token
	// TODO: Send reset email
	return nil
}

// ConfirmReset resets the password using the reset token.
func (uc *ResetPasswordUseCase) ConfirmReset(ctx context.Context, req *dto.ConfirmResetRequest) error {
	// TODO: Validate reset token
	// TODO: Hash new password
	// TODO: Update user password
	return nil
}
