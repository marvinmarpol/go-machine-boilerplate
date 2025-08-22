package app

import (
	"bufio"
	"fmt"
	"go-machine-boilerplate/internal/splitwise/adapter/cli"
	"go-machine-boilerplate/internal/splitwise/service"
	"os"
	"strings"
)

const (
	exitCode     = "99"
	exitMessage  = "exit"
	debugMessage = "debug"
)

func Run() error {
	var splitWiseService = service.NewSplitwiseService()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := strings.TrimSpace(scanner.Text())
		if input == exitCode || strings.ToLower(input) == exitMessage {
			fmt.Println("\nsystem exited")
			break
		}

		if strings.ToLower(input) == debugMessage {
			fmt.Println(splitWiseService.Users, "\n")
		}

		command := cli.Parse(input)
		err := command.Dispatch(splitWiseService)
		if err != nil {
			fmt.Printf("\n%s\n\n", err.Error())
		}
	}

	return nil
}
