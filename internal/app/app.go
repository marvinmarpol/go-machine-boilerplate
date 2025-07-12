package app

import (
	"fmt"
	"go-machine-boilerplate/internal/ratelimiter/service"
	"time"
)

func Run() error {

	rateLimiterService := service.NewRateLimiterService()

	for i := 0; i < 15; i++ {
		fmt.Println(time.Now(), rateLimiterService.Passthrough())
		time.Sleep(1 * time.Second)
	}

	return nil
}
