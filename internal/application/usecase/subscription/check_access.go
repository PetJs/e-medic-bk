// Package subscription contains subscription use cases.
package subscription

import (
	"context"

	"emedic-bk/internal/application/dto"
	"emedic-bk/internal/domain/repository"
)

// CheckAccessUseCase handles checking if a user has premium access.
type CheckAccessUseCase struct {
	subRepo repository.SubscriptionRepository
}

// NewCheckAccessUseCase creates a new CheckAccessUseCase.
func NewCheckAccessUseCase(subRepo repository.SubscriptionRepository) *CheckAccessUseCase {
	return &CheckAccessUseCase{subRepo: subRepo}
}

// Execute reports whether the user has an active subscription.
func (uc *CheckAccessUseCase) Execute(ctx context.Context, userID string) (bool, error) {
	sub, err := uc.subRepo.GetActiveByUser(ctx, userID)
	if err != nil {
		return false, err
	}
	return sub != nil, nil
}

// ListSubscriptionsUseCase lists a user's subscriptions.
type ListSubscriptionsUseCase struct {
	subRepo repository.SubscriptionRepository
}

// NewListSubscriptionsUseCase creates a new ListSubscriptionsUseCase.
func NewListSubscriptionsUseCase(subRepo repository.SubscriptionRepository) *ListSubscriptionsUseCase {
	return &ListSubscriptionsUseCase{subRepo: subRepo}
}

// Execute returns all subscriptions for a user, newest first.
func (uc *ListSubscriptionsUseCase) Execute(ctx context.Context, userID string) ([]*dto.SubscriptionResponse, error) {
	subs, err := uc.subRepo.ListByUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	responses := make([]*dto.SubscriptionResponse, 0, len(subs))
	for _, s := range subs {
		responses = append(responses, toSubscriptionResponse(s))
	}
	return responses, nil
}
