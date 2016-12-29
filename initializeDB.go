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
													revenue numeric,
													PRIMARY KEY (id)
									);
										
									CREATE TABLE "public"."expense" (
													id serial,
													item text NOT NULL,
													cost numeric NOT NULL,
													reimbursed bool NOT NULL DEFAULT 'false',
													purchase_date date NOT NULL DEFAULT now(),
													event_id integer REFERENCES event (id),
													who text NOT NULL,
													PRIMARY KEY ("id")
									);

									INSERT INTO event (type, event_date, description, estimated_mouths, revenue)
									VALUES
										('Recurring Multiple Events', '2016-12-01', '{"notes":"These are recurring expenses that span multiple events"}', -1, -1),
										('Grow Lunch', '2016-12-15', '{"notes": "1/2 lb orders for lunch"}', 30, 228),
										('Christmas 2016', '2016-12-25', '{"notes": "Whole brisket orders for families"}', 4, 295)
										;

									INSERT INTO expense (item, cost, reimbursed, purchase_date, event_id, who) 
									VALUES
										('Curing Salt', 10.99, TRUE, '2016-12-11', 1, 'Jimmy'),
										('Brisket', 49.74, TRUE, '2016-12-13', 2, 'Jimmy'),
										('Brisket', 67.68, TRUE, '2016-12-14', 2, 'Joey'),
										('White Bread', 10.25, TRUE, '2016-12-14', 2, 'Joey'),
										('BBQ Sauce', 17.48, TRUE, '2016-12-14', 2, 'Joey'),
										('Brisket', 114.62, TRUE, '2016-12-20', 3, 'Jimmy'),
										('Salt + Pepper', 11.10, TRUE, '2016-12-21', 3, 'Jimmy'),
										('Brisket', 43.04, TRUE, '2016-12-20', 3, 'Joey')
										;

									INSERT INTO brisket (price, weight_pre_trim, weight_post_trim, weight_post_cook, rub, sous_vide_time, sous_vide_temp, smoker_time, smoker_temp, final_temp, meta)
									VALUES
										(2.69, 12.8, NULL, NULL, 'Salt, Pepper, Liquid Smoke', 36, 155, 2.33, 300, 195, '{"notes":"Bark was not great. Pepper not ground enough"}'),
										(2.69, 18, 13.6, NULL, 'Salt, Pepper, Liquid Smoke, Curing Salt', 36, 155, 3.5, 300, 160, '{"notes":"Bark was much better. Did an ice bath so we could smoke it longer"}'),
										(2.99, 9.47, 7.15, NULL, 'Salt, Pepper, Liquid Smoke, Curing Salt, INSERT BEEF RUB NAME HERE', 24, 155, 3, 300, 175, '{"notes":"The flat was sous vided for 24 hours and we could not tell a difference really. Had it for Christmas Dinner and it was great. A little too much curing salt though. Also used a sweet beef rub that Karissa really liked, but should have used more pepper in conjunction with it"}'),
										(2.99, 13.88, 10.28, 6.1, 'Salt, Pepper, Liquid Smoke, Curing Salt', 36, 155, 3, 300, 175, '{"notes":""}'),
										(2.99, 11.62, 9.32, NULL, 'Salt, Pepper, Liquid Smoke, Curing Salt', 36, 155, 3, 300, 175, '{"notes":""}'),
										(2.99, 11.77, 9.55, NULL, 'Salt, Pepper, Liquid Smoke, Curing Salt', 36, 155, 3, 300, 175, '{"notes":""}'),
										(2.69, 14.08, 10.28, 6.6, 'Salt, Pepper, Liquid Smoke, Curing Salt', 36, 155, 3, 200, 175, '{"notes":""}')
										;
										
										`)

	fmt.Println("finished exec")
	if err != nil {
		log.Fatal(err)
	}
	return
}
