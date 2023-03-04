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

func NewPostgressStore(ctx context.Context) (*PostgresStore, error) {
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

func (s *PostgresStore) Init(ctx context.Context) error {
	return s.CreateAccountTable(ctx)
}

func (s *PostgresStore) CreateAccountTable(ctx context.Context) error {
	queryStr := `create table if not exists accounts (
		id varchar(255) primary key,
		first_name varchar(50),
		last_name varchar(50),
		number serial,
		balance serial,
		created_at timestamp
	)`

	_, err := s.db.Exec(ctx, queryStr)
	return err
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
