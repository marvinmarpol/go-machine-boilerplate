package service

import "go-machine-boilerplate/internal/ratelimiter/domain"

type RateLimiterService struct {
	domain         domain.RateLimiter
	requestCounter int
}

const (
	defaultMaxRequest     = 3
	defaultCooldownPeriod = 5
)

func NewRateLimiterService() *RateLimiterService {
	return &RateLimiterService{domain: domain.RateLimiter{
		MaxRequest:     defaultMaxRequest,
		CooldownPeriod: defaultCooldownPeriod,
	}}
}

func (s *RateLimiterService) Passthrough() bool {
	s.requestCounter++

	if s.requestCounter > s.domain.MaxRequest {
		return false
	}

	s.requestCounter = 0
	return true
}

func (s *RateLimiterService) GetCooldownPeriod() int {
	return s.domain.CooldownPeriod
}
