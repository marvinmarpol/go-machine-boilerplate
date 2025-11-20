package app

import (
	"fmt"
	"go-machine-boilerplate/internal/lockmanager"
	"time"
)

func Run() error {
	im := lockmanager.NewLockManager()

	fmt.Println(im.Lock("payload", "marvin", 1*time.Second))
	fmt.Println(im.Lock("payload", "marvin", 1*time.Second))

	time.Sleep(3 * time.Second)
	fmt.Println(im.Lock("payload", "marvin", 1*time.Second))
	fmt.Println(im.Unlock("payload", "wrongClient"))
	fmt.Println(im.Unlock("payload", "marvin"))

	return nil
}
