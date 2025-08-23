package domain

type User struct {
	ID          string
	Balance     float64
	Debts       map[string]float64
	Receivables map[string]float64
}

func NewUser(ID string, InitialBalance float64) *User {
	return &User{
		ID:          ID,
		Balance:     InitialBalance,
		Debts:       make(map[string]float64),
		Receivables: make(map[string]float64),
	}
}
