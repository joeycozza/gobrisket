package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func initialize() {
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

	rows, err := db.Query(`CREATE TABLE "public"."expense" (
    "id" serial,
    "item" text NOT NULL,
    "cost" numeric NOT NULL,
    "reimbursed" bool NOT NULL DEFAULT 'false',
    "purchase_date" date NOT NULL DEFAULT now(),
    PRIMARY KEY ("id")
		);`)
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
