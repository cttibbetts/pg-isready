package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/jackc/pgx/stdlib"
)

const sleep = 3

func env(variable string, defaultValue string) string {
	out := os.Getenv(variable)
	if out == "" {
		out = defaultValue
	}
	return out
}

func main() {
	host := env("POSTGRES_HOST", "localhost")
	port := env("POSTGRES_PORT", "5432")
	user := env("POSTGRES_USER", "postgres")
	pass := env("POSTGRES_PASS", "password")
	db := env("POSTGRES_DB", "postgres")

	dbinfo := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		user, pass, host, port, db)
	fmt.Printf("Connecting to: %s:%s/%s\n", host, port, db)

	connected := false
	for !connected {
		db, err := sql.Open("pgx", dbinfo)
		if err != nil {
			fmt.Println("Can't connect to database...")
			time.Sleep(sleep * time.Second)
			continue
		}

		err = db.Ping()
		if err != nil {
			fmt.Println("Can't ping database...")
			time.Sleep(sleep * time.Second)
			continue
		}
		connected = true
	}
	fmt.Println("Connected")
}
