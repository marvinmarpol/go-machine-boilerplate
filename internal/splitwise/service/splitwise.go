package service

import (
	"fmt"
	"go-machine-boilerplate/internal/splitwise/domain"
	"strconv"
)

type SplitWiseService struct {
	Users map[string]domain.User
}

func NewSplitwiseService() *SplitWiseService {
	return &SplitWiseService{
		Users: make(map[string]domain.User),
	}
}

func (s *SplitWiseService) RegisterUserCLI(IDArgs, BalanceArgs string) error {
	balance, err := strconv.Atoi(BalanceArgs)
	if err != nil {
		return err
	}

	s.addUser(*domain.NewUser(IDArgs, float64(balance)))
	return nil
}

func (s *SplitWiseService) PrintUsers() {
	for _, user := range s.Users {
		fmt.Println("\nid:", user.ID, "balance:", user.Balance)
		fmt.Println("list of debts:")
		for _, debt := range user.Debts {
			fmt.Println("debtorID:", debt.ID, "debtAmount:", debt.Balance)
		}
	}
}

func (s *SplitWiseService) addUser(user domain.User) {
	s.Users[user.ID] = user
}
