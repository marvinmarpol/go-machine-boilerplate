package cli

import "strings"

func ParseInput(input string) Command {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return Command{Name: "", Args: []string{}}
	}

	return Command{
		Name: strings.ToLower(parts[0]),
		Args: parts[1:],
	}
}
