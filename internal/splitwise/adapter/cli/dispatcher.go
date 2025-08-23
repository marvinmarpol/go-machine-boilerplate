package cli

import (
	"fmt"
	"go-machine-boilerplate/internal/splitwise/service"
)

type Command struct {
	Name string
	Args []string
}

const (
	RegisterCMD = "REGISTER"
	ShowCMD     = "SHOW"
	ExpenseCMD  = "EXPENSE"
	DebugCMD    = "DEBUG"
)

func (cmd *Command) validate() bool {
	switch cmd.Name {
	case RegisterCMD:
		if len(cmd.Args) < 2 || len(cmd.Args)%2 != 0 {
			return false
		}
	case ShowCMD:
		if len(cmd.Args) != 1 {
			return false
		}
	case ExpenseCMD:
		if len(cmd.Args) < 3 {
			return false
		}
	}

	return true
}

func (cmd *Command) Dispatch(s *service.SplitWiseService) error {
	if !cmd.validate() {
		return fmt.Errorf(`bad format for "%s"`, cmd.Name)
	}

	switch cmd.Name {
	case RegisterCMD:
		for i := 0; i < len(cmd.Args); i += 2 {
			err := s.RegisterUserCLI(cmd.Args[i], cmd.Args[i+1])
			if err != nil {
				return err
			}
		}

	case ShowCMD:
		s.ShowBalanceCLI(cmd.Args[0])

	case ExpenseCMD:
		s.ExpenseCLI(cmd.Args[0], cmd.Args[1], cmd.Args[2:])

	case DebugCMD:
		s.PrintUsers()
	}

	return nil
}
