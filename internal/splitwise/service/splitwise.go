package service

import (
	"go-machine-boilerplate/internal/splitwise/domain"
	"strconv"
)

type SplitWiseService struct {
	Users []domain.User
}

func NewSplitwiseService() *SplitWiseService {
	return &SplitWiseService{}
}

func (s *SplitWiseService) RegisterUserCLI(IDArgs, BalanceArgs string) error {
	balance, err := strconv.Atoi(BalanceArgs)
	if err != nil {
		return err
	}

	s.addUser(*domain.NewUser(IDArgs, float64(balance)))
	return nil
}

func (s *SplitWiseService) addUser(user domain.User) {
	s.Users = append(s.Users, user)
}
