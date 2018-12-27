// testdb=# create table posts ( title text not null primary key, tags text[]);
// testdb=# insert into posts (title, tags) values('pq-array with golang', '{"postgres", "golang"}');

package main

import (
	"database/sql"
	"fmt"
	"log"

	pq "github.com/lib/pq"
)

var Db *sql.DB

func main() {
	//db, err := sql.Open("postgres", "host=172.16.2.100 dbname=test")
	var err error

	Db, err = sql.Open("postgres", "user=jk password=xxxxxx dbname=testdb")

	defer Db.Close()

	if err != nil {
		log.Fatal(err)
	}

	err = Db.Ping()
	if err != nil {
		log.Panic(err)
	}
	log.Println("DB ping success!")

	ins := "INSERT INTO posts (title, tags) VALUES ($1, $2)"

	tags := []string{"go", "goroutines", "queues"}

	title := "Job Queues in Go"
	_, err = Db.Exec(ins, title, pq.Array(tags))

	if err != nil {
		fmt.Println(err)
	}

	sel := "SELECT tags FROM posts WHERE title=$1"

	if err := Db.QueryRow(sel, title).Scan(pq.Array(&tags)); err != nil {
		log.Fatal(err)
	}
	fmt.Println(tags)
}

/*
2018/12/27 12:14:14 DB ping success!
pq: duplicate key value violates unique constraint "posts_pkey"
[go goroutines queues]
*/
