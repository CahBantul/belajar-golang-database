package belajargolangdatabase

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	query := "INSERT INTO customer VALUES ('tamaz', 'tama', 'tama@gmail.com', 10000, 95.0, '1996-10-18', true)"
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}

func TestExecSqlParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	username := "nozami"
	password := "password"

	ctx := context.Background()
	query := "INSERT INTO user(username, password) VALUES (?, ?)"
	_, err := db.ExecContext(ctx, query, username, password)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	query := "SELECT id, name FROM customer"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}

		fmt.Println(id + ": " + name)
	}
}

func TestQuerySqlComplext(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	query := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var birth_date sql.NullTime
		var married bool
		var created_at time.Time
		err := rows.Scan(&id, &name, &email, &balance, &rating, &birth_date, &married, &created_at)
		if err != nil {
			panic(err)
		}
		// Print details of each record
		fmt.Println("=========================")
		fmt.Println("id", id)
		fmt.Println("name", name)
		if email.Valid {
			fmt.Println("email", email.String)
		}
		fmt.Println("balance", balance)
		fmt.Println("rating", rating)
		if birth_date.Valid {
			fmt.Println("birth_date", birth_date.Time)
		}
		fmt.Println("married", married)
		fmt.Println("created_at", created_at)
	}
}

func TestQuerySqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	username := "admin'; #"
	password := "password"
	query := "SELECT username FROM user WHERE username = '" + username + "' AND password = '" + password + "' LIMIT 1" // SQL
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	if rows.Next() {
		var username string

		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("sukses Login", username)
	} else {
		fmt.Println("gagal login")
	}
}

func TestQuerySqlInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	username := "admin"
	password := "password"
	query := "SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1" // SQL
	rows, err := db.QueryContext(ctx, query, username, password)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	if rows.Next() {
		var username string

		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("sukses Login", username)
	} else {
		fmt.Println("gagal login")
	}
}

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	email := "nozami.id@gmail.com"
	comment := "oohalao"

	ctx := context.Background()
	query := "INSERT INTO comments(email, comment) VALUES (?, ?)"
	result, err := db.ExecContext(ctx, query, email, comment)
	if err != nil {
		panic(err)
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new comment with id", insertId)
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "INSERT INTO comments(email, comment) VALUES (?,?)"
	statement, err := db.PrepareContext(ctx, script)
	if err != nil {
		panic(err)
	}

	defer statement.Close()

	for i := 0; i < 10; i++ {
		email := "fardan" + strconv.Itoa(i) + "@example.com"
		comment := "ini komen ke" + strconv.Itoa(i)
		result, err := statement.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("Comment id ", id)
	}
}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	// do transaction
	script := "INSERT INTO comments(email, comment) VALUES (?,?)"
	for i := 0; i < 10; i++ {
		email := "fardannozami" + strconv.Itoa(i) + "@example.com"
		comment := "ini komen ke" + strconv.Itoa(i)
		result, err := tx.ExecContext(ctx, script, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("Comment id ", id)
	}

	err = tx.Commit()
	if err != nil {
		panic(err)
	}
}
