package main

import (
	"fmt"
	"go-machine-boilerplate/internal/app"
)

func main() {
	err := app.Run()
	if err != nil {
		fmt.Println(err)
	}
}
