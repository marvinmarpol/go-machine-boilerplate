package app

import (
	"bufio"
	"fmt"
	"go-machine-boilerplate/internal/splitwise/adapter/cli"
	"go-machine-boilerplate/internal/splitwise/service"
	"os"
	"strings"
)

const exitCode = "99"
const exitMessage = "exit"

func Run() error {
	var splitWiseService = service.SplitWiseService{}
	fmt.Print("input mode active\n\n")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := strings.TrimSpace(scanner.Text())
		if input == exitCode || strings.ToLower(input) == exitMessage {
			fmt.Println("system exited")
			break
		}

		command := cli.Parse(input)
		command.Dispatch(splitWiseService)

	}

	return nil
}
