package application

import (
	"context"
	"github.com/Honeymoond24/auction-system/internal/domain"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) (int64, error)
	GetUserByID(ctx context.Context, userID int64) (*domain.User, error)
	UpdateUser(ctx context.Context, user *domain.User) error
	ReplenishBalanceTx(ctx context.Context, userID int64, amount float64) error
}
