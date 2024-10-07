package domain

import "errors"

type User struct {
	ID       int64
	Username string
	Balance  float64
}

func (u *User) ReplenishBalance(amount float64) error {
	if amount <= 0 {
		return errors.New("invalid amount: must be greater than 0")
	}
	u.Balance += amount
	return nil
}
