package main
import (
	"fmt"
	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/types"
	"log"
	"encoding/json"
	"net/http"
)

type Entity struct {
	Id          int         `db:"id" json:"id,omitempty"`
	Name        string      `db:"name" json:"Name,omitempty"`
	Description string      `db:"description" json:"Description,omitempty"`
	Properties  types.JSONText `db:"perperties" json:"Properties,omitempty"`
}

func main() {
	http.HandleFunc("/insert", insert)
	log.Fatal(http.ListenAndServe(":8099", nil))
}

func insert(w http.ResponseWriter, r *http.Request) {

	Db, err := sqlx.Connect("postgres", "postgres:///tachyon_dev?host=/var/run/postgresql&sslmode=disable&binary_parameters=yes")

	err = Db.Ping()
	if err != nil {
		log.Panic(err.Error())
	}
	fmt.Println("ping database success")

    var f Entity
    decoder := json.NewDecoder(r.Body)
    decoder.Decode(&f)
	fmt.Println(f)

	query := `INSERT INTO entity(id, name, description, properties) 
          VALUES($1, $2, $3, $4)`

	_, err = Db.Queryx(query, f.Id, f.Name, f.Description, f.Properties)
	if err != nil {
	 log.Fatalln(err)
	 return
	}

	e := Entity{Id:f.Id }

	err = Db.QueryRow("SELECT name, description, properties FROM entity WHERE id = $1", 
              e.Id).Scan(&e.Name, &e.Description, &e.Properties)

    if err == nil {
	  fmt.Println("writing output")
      w.Header().Set("Content-Type", "application/json")
      w.WriteHeader(http.StatusOK)
      json.NewEncoder(w).Encode(e)
    }
}

/*
postgres
\d entity;
                               Table "public.entity"
   Column    |  Type   | Collation | Nullable |              Default               
-------------+---------+-----------+----------+------------------------------------
 id          | integer |           | not null | nextval('entity_id_seq'::regclass)
 name        | text    |           |          | 
 description | text    |           |          | 
 properties  | json    |           |          | 
Indexes:
    "entity_pkey" PRIMARY KEY, btree (id)
*/

/*
curl  -X POST --verbose http://localhost:8099/insert -d '{"Id":1300, "Name":"cyy", "Description":"at china", "Properties":{"Price": 9.99, "Title": "chinese101", "Published": false, "Year":2019}}'

select * from entity;
  id  | name | description |                               properties                                
------+------+-------------+-------------------------------------------------------------------------
 1300 | cyy  | at china    | {"Price": 9.99, "Title": "chinese101", "Published": false, "Year":2019}

$ curl  http://localhost:8099/insert -d '{"Id":1300, "Name":"cyy", "Description":"at china", "Properties":{"Price": 9.99, "Title": "chinese101", "Published": false, "Year":2019}}'
{"id":1300,"Name":"cyy","Description":"at china","Properties":{"Price":9.99,"Title":"chinese101","Published":false,"Year":2019}}
*/
