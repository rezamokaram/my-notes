package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func NewConnection() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("./sample.sqlite"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic("failed to connect database")
	}
	return db, nil
}
