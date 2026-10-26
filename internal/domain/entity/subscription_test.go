package entity

import (
	"testing"
	"time"
)

func TestSubscriptionGraceWindow(t *testing.T) {
	periodEnd := time.Date(2026, 8, 11, 12, 0, 0, 0, time.UTC)
	graceEnd := periodEnd.Add(SubscriptionGracePeriod)

	active := &Subscription{
		Status:           SubscriptionStatusActive,
		CurrentPeriodEnd: periodEnd,
	}

	cases := []struct {
		name    string
		now     time.Time
		access  bool
		inGrace bool
	}{
		{"inside paid period", periodEnd.Add(-time.Hour), true, false},
		{"first moment of grace", periodEnd, true, true},
		{"one day into grace", periodEnd.Add(24 * time.Hour), true, true},
		{"last moment of grace", graceEnd.Add(-time.Nanosecond), true, true},
		{"grace elapsed", graceEnd, false, false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := active.HasAccess(tc.now); got != tc.access {
				t.Errorf("HasAccess(%v) = %v, want %v", tc.now, got, tc.access)
			}
			if got := active.IsInGrace(tc.now); got != tc.inGrace {
				t.Errorf("IsInGrace(%v) = %v, want %v", tc.now, got, tc.inGrace)
			}
		})
	}

	if got := active.AccessEndsAt(); !got.Equal(graceEnd) {
		t.Errorf("AccessEndsAt() = %v, want %v", got, graceEnd)
	}

	expired := &Subscription{
		Status:           SubscriptionStatusExpired,
		CurrentPeriodEnd: periodEnd,
	}
	if expired.HasAccess(periodEnd.Add(-time.Hour)) {
		t.Error("expired subscription must not grant access")
	}
}
