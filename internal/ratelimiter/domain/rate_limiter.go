package domain

import "fmt"

type RateLimiter struct {
	MaxRequest     int
	CooldownPeriod int
}

func NewRateLimiter() *RateLimiter {
	return &RateLimiter{}
}

func (r *RateLimiter) setMaxRequest(MaxRequest int) error {

	if MaxRequest < 1 {
		return fmt.Errorf("should be greater 1")
	}

	r.MaxRequest = MaxRequest
	return nil
}

func (r *RateLimiter) setCooldownPeriod(period int) error {
	if period < 1 {
		return fmt.Errorf("should be greater 1")
	}

	r.CooldownPeriod = period
	return nil
}
