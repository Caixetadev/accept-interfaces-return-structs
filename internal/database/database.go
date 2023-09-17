package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect() *pgxpool.Pool {
	var psqlconn string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", "localhost", "root", "password", "teste")

	poolConfig, err := pgxpool.ParseConfig(psqlconn)
	if err != nil {
		panic(err)
	}

	db, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}

	err = db.Ping(context.Background())
	if err != nil {
		panic(err)
	}

	return db
}
