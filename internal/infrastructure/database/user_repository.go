package database

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Honeymoond24/auction-system/internal/application"
	"github.com/Honeymoond24/auction-system/internal/domain"
	"github.com/jackc/pgx/v5"
	"log"
)

type userRepository struct {
	db *DBPool
}

func NewUserRepository(db *DBPool) application.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) BeginTx(ctx context.Context) (pgx.Tx, error) {
	return r.db.conn.BeginTx(ctx, pgx.TxOptions{})
}

func (r *userRepository) CreateUser(ctx context.Context, user *domain.User) (int64, error) {
	if user == nil {
		return 0, errors.New("invalid user data")
	}

	query := `
		INSERT INTO users (username, balance)
		VALUES ($1, $2)
		RETURNING id
	`

	var userID int64
	err := r.db.conn.QueryRow(ctx, query, user.Username, user.Balance).Scan(&userID)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (r *userRepository) GetUserByID(ctx context.Context, userID int64) (*domain.User, error) {
	if userID <= 0 {
		return nil, errors.New("invalid user ID")
	}

	query := `
		SELECT id, username, balance
		FROM users
		WHERE id = $1
	`

	var user domain.User
	err := r.db.conn.QueryRow(ctx, query, userID).Scan(&user.ID, &user.Username, &user.Balance)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, user *domain.User) error {
	if user == nil || user.ID <= 0 {
		return errors.New("invalid user data")
	}

	query := `
		UPDATE users
		SET username = $1, balance = $2, updated_at = NOW()
		WHERE id = $3
	`

	_, err := r.db.conn.Exec(ctx, query, user.Username, user.Balance, user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) ReplenishBalanceTx(ctx context.Context, userID int64, amount float64) error {
	tx, err := r.db.conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer func(tx pgx.Tx, ctx context.Context) {
		err := tx.Rollback(ctx)
		if err != nil {
			if errors.Is(err, pgx.ErrTxClosed) {
				return
			}
			log.Println(err)
		}
	}(tx, ctx)

	var user domain.User
	query := `SELECT id, username, balance FROM users WHERE id = $1`
	err = tx.QueryRow(ctx, query, userID).Scan(&user.ID, &user.Username, &user.Balance)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return errors.New("user not found")
		}
		return err
	}

	user.Balance += amount
	updateQuery := `UPDATE users SET balance = $1, updated_at = NOW() WHERE id = $2`
	_, err = tx.Exec(ctx, updateQuery, user.Balance, user.ID)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}
