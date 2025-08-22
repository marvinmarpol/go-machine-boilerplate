package cli

import "strings"

func Parse(input string) Command {
	cmd := Command{}
	args := strings.Fields(input)

	if len(args) < 1 {
		return cmd
	}

	cmd.Name = args[0]
	cmd.Args = args[1:]
	return cmd
}
