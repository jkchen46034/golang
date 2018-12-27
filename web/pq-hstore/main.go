//hstore:
// http://www.postgresqltutorial.com/postgresql-hstore/

package main

import (
	"database/sql"
	_ "fmt"
	"log"

	_ "github.com/lib/pq"
	_ "github.com/lib/pq/hstore"
)

var Db *sql.DB

func main() {
	var err error

	Db, err = sql.Open("postgres", "user=jk password=Tom00son dbname=testdb")

	defer Db.Close()

	if err != nil {
		log.Fatal(err)
	}

	err = Db.Ping()
	if err != nil {
		log.Panic(err)
	}
	log.Println("DB ping success!")

	sel := "select attr->'ISBN-13' AS isbn from books"

	var isbn []string
	rows, err := Db.Query(sel)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var isbnstr string
		if err := rows.Scan(&isbnstr); err != nil {
			log.Fatal(err)
		}
		isbn = append(isbn, isbnstr)
	}
	log.Println(isbn)
}

/*
2018/12/27 15:22:25 DB ping success!
2018/12/27 15:22:25 [978-1449370000 978-1449370001]
*/

/*
CREATE EXTENSION hstore;
CREATE TABLE books (
 id serial primary key,
 title VARCHAR (255),
 attr hstore
);
INSERT INTO books (title, attr)
VALUES
 (
 'PostgreSQL Tutorial',
 '"paperback" => "243",
    "publisher" => "postgresqltutorial.com",
    "language"  => "English",
    "ISBN-13"   => "978-1449370000",
 "weight"    => "11.2 ounces"'
 );
INSERT INTO books (title, attr)
VALUES
 (
 'PostgreSQL Cheat Sheet',
 '
"paperback" => "5",
"publisher" => "postgresqltutorial.com",
"language"  => "English",
"ISBN-13"   => "978-1449370001",
"weight"    => "1 ounces"'
 );
*/
