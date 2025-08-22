package cli

import (
	"errors"
	"fmt"
	"go-machine-boilerplate/internal/splitwise/domain"
	"go-machine-boilerplate/internal/splitwise/service"
	"net/http"
	"strconv"
)

type Command struct {
	Name string
	Args []string
}

func (cmd *Command) validate() bool {
	switch cmd.Name {
	case "REGISTER":
		if len(cmd.Args) < 2 || len(cmd.Args)%2 != 0 {
			return false
		}
	case "SHOW":
		if len(cmd.Args) != 1 {
			return false
		}
	case "EXPENSE":
		if len(cmd.Args) < 3 {
			return false
		}
	}

	return true
}

func (cmd *Command) Dispatch(s service.SplitWiseService) error {
	if !cmd.validate() {
		return errors.New(http.StatusText(http.StatusBadRequest))
	}

	switch cmd.Name {
	case "REGISTER":
		for i := 0; i < len(cmd.Args); i += 2 {
			balance, err := strconv.Atoi(cmd.Args[i+1])
			if err != nil {
				return err
			}
			s.Users = append(s.Users, *domain.NewUser(cmd.Args[i], float64(balance)))
		}
	}

	fmt.Println(s.Users)
	return nil
}
