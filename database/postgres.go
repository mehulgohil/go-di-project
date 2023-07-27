package database

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
)

type Client struct {
	db *sqlx.DB
}

func (p *Client) Execute(query string, args ...interface{}) (sql.Result, error) {
	return p.db.Exec(query, args...)
}

func (p *Client) Query(query string, args ...interface{}) (*sqlx.Rows, error) {
	return p.db.Queryx(query, args...)
}

func (p *Client) Get(dest interface{}, query string, args ...interface{}) error {
	return p.db.Get(dest, query, args...)
}

func (p *Client) Select(dest interface{}, query string, args ...interface{}) error {
	return p.db.Select(dest, query, args...)
}

func NewPGClient(pgConnectionString string) (*Client, error) {
	db, err := sqlx.Connect("postgres", pgConnectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Client{db: db}, nil
}
