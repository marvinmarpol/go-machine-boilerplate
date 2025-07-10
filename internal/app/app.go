package app

import (
	"bufio"
	"go-machine-boilerplate/internal/parking/adapter/cli"
	"go-machine-boilerplate/internal/parking/service"
	"os"
	"strings"
)

const (
	exitCommand = "exit"
	exitCode    = "99"
)

func Run() error {
	parkingLotService := service.NewParkingLotService()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		if input == exitCode || strings.ToLower(input) == exitCommand {
			break
		}

		cmd := cli.ParseInput(input)
		cmd.Dispatch(parkingLotService)
	}

	return nil
}
