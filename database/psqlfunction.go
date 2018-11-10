package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	_ "time"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "xxxxxx"
	dbname   = "testdb"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	checkError(err)

	defer db.Close()

	err = db.Ping()

	checkError(err)

	fmt.Println("Successfully connected!")

	sqlStatement := `select insert_userinfo($1, $2);`

	_, err = db.Exec(sqlStatement, "username_ghi", "Executive")

	checkError(err)

	fmt.Println("Successfully inserted")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

/*
$ go run psqlfunction.go
Successfully connected!
Successfully inserted
*/
