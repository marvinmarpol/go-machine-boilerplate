package app

import (
	"fmt"
	"go-machine-boilerplate/internal/ratelimiter/service"
	"time"
)

func Run() error {

	rateLimiterService := service.NewRateLimiterService()

	for i := 0; i < 5; i++ {
		result := rateLimiterService.Passthrough()
		fmt.Println("incoming request status:", result)
		if !result {
			time.Sleep(time.Duration(rateLimiterService.GetCooldownPeriod()))
		}
	}

	return nil
}
