// Package valueobject contains immutable value objects with validation.
package valueobject

import (
	"errors"
	"fmt"
)

// ErrNegativeAmount is returned when a money amount is negative.
var ErrNegativeAmount = errors.New("money amount cannot be negative")

// Money represents a monetary value with currency.
type Money struct {
	amount   int64  // in smallest currency unit (cents, kobo, etc.)
	currency string // ISO 4217 currency code
}

// NewMoney creates a new Money value object.
func NewMoney(amount int64, currency string) (Money, error) {
	if amount < 0 {
		return Money{}, ErrNegativeAmount
	}
	return Money{amount: amount, currency: currency}, nil
}

// Amount returns the amount in smallest currency unit.
func (m Money) Amount() int64 {
	return m.amount
}

// Currency returns the ISO 4217 currency code.
func (m Money) Currency() string {
	return m.currency
}

// String returns a formatted string representation.
func (m Money) String() string {
	return fmt.Sprintf("%d %s", m.amount, m.currency)
}

// Add adds two Money values of the same currency.
func (m Money) Add(other Money) (Money, error) {
	if m.currency != other.currency {
		return Money{}, errors.New("cannot add money with different currencies")
	}
	return Money{amount: m.amount + other.amount, currency: m.currency}, nil
}
