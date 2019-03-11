package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type User struct {
	gorm.Model
	Orders []Order
	Data   string `sql:"type:JSONB NOT NULL DEFAULT '{}'::JSONB" json:"-"`
}

type Order struct {
	gorm.Model
	User User
	Data string `sql:"type:JSONB NOT NULL DEFAULT '{}'::JSONB"`
}

// GORM creates plural table names, surpress with:
func (User) TableName() string {
	return "user"
}

func (Order) TableName() string {
	return "order"
}

func InitDB() (*gorm.DB, error) {
	var err error
	db, err := gorm.Open("postgres", "postgres://majora:autumn@127.0.21.2/mydb?sslmode=disable")
	if err != nil {
		return nil, err
	} else {
		// if !db.HasTable("user") { db.CreateTable(&User{}) }
		db.AutoMigrate(&User{}, &Order{})
		return db, nil
	}
}
