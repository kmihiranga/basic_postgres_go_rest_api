package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccountByID(int) (*Account, error)
}

type PostgresStore struct {
	db *pgx.Conn
}

func NewPostgressStore() (*PostgresStore, error) {
	ctx := context.Background()
	connStr := fmt.Sprintf("postgres://%s:%s@localhost:5432/postgres_rest", "postgres", "1234")
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		log.Fatalf("Error connecting to the database. %v", err)
		return nil, err
	}
	defer conn.Close(ctx)

	return &PostgresStore{
		db: conn,
	}, nil
}

func (s *PostgresStore) CreateAccount(*Account) error {
	return nil
}

func (s *PostgresStore) UpdateAccount(*Account) error {
	return nil
}

func (s *PostgresStore) DeleteAccount(id int) error {
	return nil
}

func (s *PostgresStore) GetAccountByID(id int) (*Account, error) {
	return nil, nil
}
