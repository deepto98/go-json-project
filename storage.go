package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

//Defining functions in common Interface, can be implemented by diff DBs
type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccountById(int) (*Account, error)
}

type PostgresStore struct {
	db *sql.DB
}

//Postgres setup
func NewPostgresStore() (*PostgresStore, error) {
	//Connect to postgres
	connStr := "user=postgres dbname=postgres password=abcd1234 sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{db: db}, nil
}

func (store *PostgresStore) CreateAccount(acc *Account) error {
	return nil
}

func (store *PostgresStore) DeleteAccount(id int) error {
	return nil
}
func (store *PostgresStore) UpdateAccount(acc *Account) error {
	return nil
}
func (store *PostgresStore) GetAccountById(id int) (*Account, error) {
	return nil, nil
}
