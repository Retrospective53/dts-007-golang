package config

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type PostgresConfig struct {
	Port int
	Host string
	Username string
	Password string
	DBName string
	MaxIdleConnection int
	MaxOpenConnection int
	MaxIdleTime int
}

func NewPostgresConfig() (db *sql.DB) {
	pgConfig := PostgresConfig{
		Port: 5432,
		Host: "localhost",
		Username: "postgres",
		Password: "mysecretpassword",
		DBName: "books_management",
		MaxOpenConnection: 7,
		MaxIdleConnection: 5,
		MaxIdleTime: int(30 * time.Minute),
	}

	connString := fmt.Sprintf(`
		host=%v
		port=%v
		user=%v
		password=%v
		dbname=%v
		sslmode=disable
	`,
		pgConfig.Host,
		pgConfig.Port,
		pgConfig.Username,
		pgConfig.Password,
		pgConfig.DBName,
	)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(pgConfig.MaxOpenConnection)
	db.SetMaxIdleConns(pgConfig.MaxIdleConnection)
	db.SetConnMaxIdleTime(time.Duration(pgConfig.MaxIdleConnection))

	if err := db.Ping(); err != nil {
		panic(err)
	}

	return db
}