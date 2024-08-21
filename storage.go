package main

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Storage interface{
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccountByID(int) (*Account, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "host=172.17.0.2 user=postgres dbname=postgres password=gobank sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}	

	return &PostgresStore{
		db: db,
	}, nil
}

func(s *PostgresStore) Init() error{
	return s.createAccountTable()
}

func(s *PostgresStore) createAccountTable() error{
	query := `CREATE TABLE IF NOT EXISTS account (
			id SERIAL PRIMARY KEY, 
			first_name VARCHAR(50), 
			last_name VARCHAR(50),
			number SERIAL,
			balance SERIAL,
			created_at TIMESTAMP
		)`

	_, err := s.db.Exec(query)
	return err
}

func(s *PostgresStore) CreateAccount(*Account) error {
	return nil
}	

func(s *PostgresStore) UpdateAccount(*Account) error {
	return nil
}	

func(s *PostgresStore) DeleteAccount(id int) error {
	return nil
}	

func(s *PostgresStore) GetAccountByID(id int) (*Account, error) {
	return nil, nil
}	
