package domain

import "fmt"

type RateLimiter struct {
	WindowSize int
	MaxRequest int
}

func NewRateLimiter(windowsSize, maxRequest int) *RateLimiter {
	return &RateLimiter{WindowSize: windowsSize, MaxRequest: maxRequest}
}

func (r *RateLimiter) setMaxRequest(MaxRequest int) error {

	if MaxRequest < 1 {
		return fmt.Errorf("should be greater 1")
	}

	r.MaxRequest = MaxRequest
	return nil
}

func (r *RateLimiter) setWindowsize(windowSize int) error {
	if windowSize < 1 {
		return fmt.Errorf("should be greater 1")
	}

	r.WindowSize = windowSize
	return nil
}
