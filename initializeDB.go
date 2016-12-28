package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

//Initialize is so good
func Initialize() {
	db, err := sql.Open("postgres", postgresPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal("Ping error: ", err)
	} else {
		fmt.Println("Ping worked")
	}
	var count int
	err = db.QueryRow(`SELECT COUNT(*) FROM pg_catalog.pg_tables where tableowner = 'cozza';`).Scan(&count)

	if err != nil {
		log.Fatal(err)
	}

	if count > 0 {
		fmt.Println("Not building tables because they already exist")
		return
	}
	_, err = db.Exec(`CREATE TABLE public.brisket (
													id serial,
													price numeric,
													weight_pre_trim numeric,
													weight_post_trim numeric,
													weight_post_cook numeric,
													rub text,
													sous_vide_time numeric,
													sous_vide_temp int,
													smoker_time numeric,
													smoker_temp int,
													final_temp int,
													meta jsonb,
													PRIMARY KEY (id)
										);
										
									CREATE TABLE public.event (
													id serial,
													type text,
													event_date date,
													description jsonb,
													estimated_mouths int,
													PRIMARY KEY (id)
									);
										
									CREATE TABLE "public"."expense" (
													id serial,
													item text NOT NULL,
													cost numeric NOT NULL,
													reimbursed bool NOT NULL DEFAULT 'false',
													purchase_date date NOT NULL DEFAULT now(),
													event_id integer REFERENCES event (id),
													PRIMARY KEY ("id")
									);
									
										`)

	fmt.Println("finished exec")
	if err != nil {
		log.Fatal(err)
	}
	return
}
