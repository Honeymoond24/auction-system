package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"sync"
)

type DBPool struct {
	conn *pgxpool.Pool
}

var (
	pgInstance *DBPool
	pgOnce     sync.Once
)

func NewDB(databaseURI string) *DBPool {
	pgOnce.Do(func() {
		db, err := pgxpool.New(context.Background(), databaseURI)
		if err != nil {
			log.Fatal("Failed to connect to database", err)
		}
		pgInstance = &DBPool{conn: db}
	})

	return pgInstance
}

func (pg DBPool) Ping(ctx context.Context) error {
	row := pg.conn.QueryRow(ctx, `SELECT 1;`)
	var num int
	err := row.Scan(&num)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (pg DBPool) Close() {
	pg.conn.Close()
}
