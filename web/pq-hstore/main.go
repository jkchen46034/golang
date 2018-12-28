//hstore:
// http://www.postgresqltutorial.com/postgresql-hstore/

package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	hstore "github.com/lib/pq/hstore"
)

type Book struct {
	Id    int
	Title string
	Attr  hstore.Hstore
}

var Db *sql.DB

func main() {
	var err error

	Db, err = sql.Open("postgres", "user=jk password=xxxxx dbname=testdb")

	defer Db.Close()

	if err != nil {
		log.Fatal(err)
	}

	err = Db.Ping()
	if err != nil {
		log.Panic(err)
	}
	log.Println("DB ping success!")

	sel := "select id, title,  attr from books"

	var books []Book

	rows, err := Db.Query(sel)

	defer rows.Close()

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var book Book
		if err := rows.Scan(&book.Id, &book.Title, &book.Attr); err != nil {
			log.Fatal(err)
		}
		books = append(books, book)
	}

	for i := 0; i < len(books); i++ {
		log.Println(books[i].Attr)
	}

	ins := "insert into books(title,  attr) values ($1,$2)"

	title := "Job Queues in Go"
	_, err = Db.Exec(ins, title, books[0].Attr)

	if err != nil {
		fmt.Println(err)
	}
}

/*
2018/12/27 16:09:21 DB ping success!
2018/12/27 16:09:21 [{1 PostgreSQL Tutorial {map[publisher:{postgresqltutorial.com true} freeshipping:{yes true} weight:{11.2 ounces true} ISBN-13:{978-1449370000 true} language:{English true} paperback:{243 true}]}} {2 PostgreSQL Cheat Sheet {map[language:{English true} paperback:{5 true} publisher:{postgresqltutorial.com true} weight:{1 ounces true} ISBN-13:{978-1449370001 true}]}}]
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
 '
  "paperback" => "243",
  "publisher" => "postgresqltutorial.com",
  "language"  => "English",
  "ISBN-13"   => "978-1449370000",
  "weight"    => "11.2 ounces"
 '
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
	"weight"    => "1 ounces"
 '
 );
*/
