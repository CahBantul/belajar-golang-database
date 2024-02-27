package belajargolangdatabase

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func TestEmpty(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	query := "INSERT INTO customer(id, name) VALUES ('nozami', 'Nozami')"
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}

func TestConnection(t *testing.T) {
	db, err := sql.Open("mysql", "kli:password@tcp(localhost:3306)/belajar_golang_database)")

	if err != nil {
		panic(err)
	}
	defer db.Close()
}

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "kali:password@tcp(localhost:3306)/belajar_golang_database?parseTime=true")

	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
