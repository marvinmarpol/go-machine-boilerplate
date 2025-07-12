package service

import (
	"go-machine-boilerplate/internal/ratelimiter/domain"
	"time"
)

type RateLimiterService struct {
	domain         domain.RateLimiter
	requestCounter int
	currentTime    time.Time
}

const (
	defaultMaxRequest = 3
	defaultWindowSize = 5
)

func NewRateLimiterService() *RateLimiterService {
	rs := RateLimiterService{
		currentTime: time.Now(),
		domain: domain.RateLimiter{
			MaxRequest: defaultMaxRequest,
			WindowSize: defaultWindowSize,
		},
	}

	return &rs
}

func (s *RateLimiterService) Passthrough() bool {
	diff := time.Now().Sub(s.currentTime)

	diffInSecond := diff.Seconds()
	windowSizeSecond := time.Duration(s.domain.WindowSize) * time.Second

	if diffInSecond > windowSizeSecond.Seconds() {
		s.currentTime = time.Now()
		s.requestCounter = 0
	}

	if s.requestCounter >= defaultMaxRequest {
		return false
	}

	s.requestCounter++
	return true
}
