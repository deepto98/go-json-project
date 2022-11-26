package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

//Defining functions in common Interface, can be implemented by diff DBs
type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccounts() ([]*Account, error)
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

func (store *PostgresStore) Init() error {
	return store.CreateAccountTable()
}

func (store *PostgresStore) CreateAccountTable() error {
	query := CreateAccountTableQuery
	_, err := store.db.Exec(query)

	return err
}

//Implementing interface functions
func (store *PostgresStore) CreateAccount(acc *Account) error {
	resp, err := store.db.Query(CreateAccountQuery,
		acc.FirstName,
		acc.LastName,
		acc.Number,
		acc.Balance,
		acc.CreatedAt,
	)
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", resp)
	return nil
}

func (store *PostgresStore) DeleteAccount(id int) error {
	return nil
}
func (store *PostgresStore) UpdateAccount(acc *Account) error {
	return nil
}

func (store *PostgresStore) GetAccounts() ([]*Account, error) {
	accounts := []*Account{}

	rows, err := store.db.Query(GetAccountsQuery)
	if err != nil {
		return nil, err
	}
	// fmt.Printf("%+v\n", rows)

	for rows.Next() {
		account := new(Account)
		//Scan copies values from row into struct
		err := rows.Scan(&account.ID,
			&account.FirstName,
			&account.LastName,
			&account.Number,
			&account.Balance,
			&account.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}

	return accounts, nil
}

func (store *PostgresStore) GetAccountById(id int) (*Account, error) {
	return nil, nil
}
