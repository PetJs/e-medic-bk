// Package subscription contains subscription use cases.
package subscription

import "context"

// CheckAccessUseCase handles checking if a user has premium access.
type CheckAccessUseCase struct{}

func NewCheckAccessUseCase() *CheckAccessUseCase { return &CheckAccessUseCase{} }

func (uc *CheckAccessUseCase) Execute(ctx context.Context, userID string) (bool, error) {
	// TODO: Check if user has active subscription
	return false, nil
}
