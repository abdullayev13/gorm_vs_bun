package config

import "fmt"

type DB struct {
	Username  string
	Password  string
	Name      string
	DebugMode bool
}

func (db *DB) Dsn() (dsn string) {
	dsn = fmt.Sprintf("postgres://%s:%s@localhost:5432/%s?sslmode=disable",
		db.Username, db.Password, db.Name)
	return dsn
}
