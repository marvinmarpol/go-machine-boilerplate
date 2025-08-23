package service

import (
	"fmt"
	"go-machine-boilerplate/internal/splitwise/domain"
	"strconv"
)

type SplitWiseService struct {
	Users map[string]*domain.User
}

func NewSplitwiseService() *SplitWiseService {
	return &SplitWiseService{
		Users: make(map[string]*domain.User),
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

func (s *SplitWiseService) ShowBalanceCLI(ID string) {
	currentUser, ok := s.Users[ID]
	if !ok {
		fmt.Println("User not found")
		return
	}

	if len(currentUser.Debts) < 1 && len(currentUser.Receivables) < 1 {
		fmt.Println("No Balances")
		return
	}

	for k, debt := range currentUser.Debts {
		fmt.Printf("%s owes %s: %v\n", currentUser.ID, k, debt)
	}

	for k, receivable := range currentUser.Receivables {
		fmt.Printf("%s owes %s: %v\n", k, currentUser.ID, receivable)
	}
}

func (s *SplitWiseService) PrintUsers() {
	for _, user := range s.Users {
		fmt.Println("\nid:", user.ID, "balance:", user.Balance)

		if len(user.Debts) > 0 {
			fmt.Println("list of debts:")
			for id, total := range user.Debts {
				fmt.Println("debtorID:", id, "debtAmount:", total)
			}
		}

		if len(user.Receivables) > 0 {
			fmt.Println("list of receivables:")
			for id, total := range user.Receivables {
				fmt.Println("receivableID:", id, "recieveableAmount:", total)
			}
		}
	}
}

func (s *SplitWiseService) ExpenseCLI(payerID, amountString string, args []string) error {
	userLength, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	offset := userLength + 1
	if offset > len(args) {
		return err
	}

	userIDs := args[1 : userLength+1]

	amount, err := strconv.Atoi(amountString)
	if err != nil {
		return err
	}
	amountF := float64(amount)

	expenseType := args[offset]
	switch expenseType {
	case "EQUAL":
		s.SplitEqual(payerID, amountF, userLength, userIDs)
	}

	return nil
}

func (s *SplitWiseService) SplitEqual(payerID string, amountF float64, numDivide int, userIDs []string) {
	splitAmount := amountF / float64(numDivide)
	payerUser, ok := s.Users[payerID]
	if !ok {
		return
	}

	payerUser.Balance = payerUser.Balance - amountF + splitAmount

	for _, userID := range userIDs {
		if userID == payerID {
			continue
		}

		debtUser, ok := s.Users[userID]
		if !ok {
			continue
		}
		debtUser.Balance -= splitAmount
		debtUser.Debts[payerID] += splitAmount
		payerUser.Receivables[userID] += splitAmount
	}
}

func (s *SplitWiseService) addUser(user domain.User) {
	s.Users[user.ID] = &user
}
