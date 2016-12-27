package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", postgresPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal("Ping error: ", err)
	}

	var (
		id   int
		name string
	)

	rows, err := db.Query("select id, name from test")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal("Row scanning error: ", err)
		}
		log.Println(id, name)

	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

}
