package main

import (
	"log"

	"github.com/jbpratt78/apis/postgres/models"
)

func main() {
	db, err := models.InitDB()
	if err != nil {
		log.Println(db)
	}
}
