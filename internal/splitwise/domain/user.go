package domain

type User struct {
	ID      string
	Balance float64
	Debts   []User
}

func NewUser(ID string, InitialBalance float64) *User {
	return &User{
		ID:      ID,
		Balance: InitialBalance,
	}
}
