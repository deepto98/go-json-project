package main

const CreateAccountTableQuery = `CREATE TABLE IF NOT EXISTS account  (
	id serial PRIMARY KEY,
	first_name VARCHAR ( 50 ) NOT NULL,
	last_name VARCHAR ( 50 ) NOT NULL,
	number int,
	balance numeric,
	created_at TIMESTAMP NOT NULL
	)`
const CreateAccountQuery = `INSERT INTO account
	(first_name,last_name,number,balance,created_at)
	VALUES
	($1,$2,$3,$4,$5)`
const GetAccountsQuery = `SELECT * FROM account`
const GetAccountByQuery = `SELECT * FROM account WHERE id = $1`
const DeleteAccountQuery = `DELETE  FROM account WHERE id = $1`
