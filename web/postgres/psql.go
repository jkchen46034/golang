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
	password = "xxxxxxxx"
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

	sqlStatement := `
		insert into userinfo(username, department, created)
		values ($1,$2,$3)
		returning uid`

	id := -1

	err = db.QueryRow(sqlStatement, "username_abc", "HR", "12-8-2015").Scan(&id)

	checkError(err)

	fmt.Println("Successfully inserted, last inserted id", id)

	sqlStatement = `
		update userinfo
		SET username = $2
		Where uid = $1;`
	res, err1 := db.Exec(sqlStatement, id, "username_def")
	checkError(err1)
	//fmt.Println("Successfully updated, id:", id, ", username_def")

	count, err2 := res.RowsAffected()
	checkError(err2)
	fmt.Println(count)

	sqlStatement = `
		delete from userinfo
		where uid = $1;`
	_, err = db.Exec(sqlStatement, id)
	checkError(err)
	fmt.Println("Successfully deleted, id:", id)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

/*
$ go run psql.go
Successfully connected!
Successfully inserted, last inserted id 13
1
Successfully deleted, id: 13

CREATE TABLE userinfo
    (
        uid serial NOT NULL,
        username character varying(100) NOT NULL,
        departname character varying(500) NOT NULL,
        Created date,
        CONSTRAINT userinfo_pkey PRIMARY KEY (uid)
    )
    WITH (OIDS=FALSE);

testdb=# \d userinfo
                                      Table "public.userinfo"
   Column   |          Type          | Collation | Nullable |                Default
------------+------------------------+-----------+----------+---------------------------------------
 uid        | integer                |           | not null | nextval('userinfo_uid_seq'::regclass)
 username   | character varying(100) |           | not null |
 department | character varying(500) |           | not null |
 created    | date                   |           |          |
Indexes:
    "userinfo_pkey" PRIMARY KEY, btree (uid)

*/
