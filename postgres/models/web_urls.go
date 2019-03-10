package models

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func InitDB() (*sql.DB, error) {
	var err error
	db, err := sql.Open("postgres", "postgres://majora:autumn@127.0.21.2/mydb?sslmode=disable")
	if err != nil {
		return nil, err
	} else {
		// create model for URL service
		stmt, err := db.Prepare("CREATE TABLE WEB_URL(ID SERIAL PRIMARY KEY, URL TEXT NOT NULL);")
		if err != nil {
			log.Println(err)
			return nil, err
		}
		res, err := stmt.Exec()
		log.Println(res)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		return db, nil
	}
}
