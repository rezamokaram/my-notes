package main

import (
	"log"

	"gorm.io/gorm"
	"github.com/labstack/echo/v4"
)

var db *gorm.DB
var e *echo.Echo

func main() {
	var err error
	db, err = NewConnection()
	if err != nil {
		log.Fatal(err.Error())
	}

	db.AutoMigrate(
		User{},
	)
	dataDump()

	e = echo.New()

	e.GET(
		"/",
		HandelFunc(),
	)

	log.Fatal(e.Start(":8080"))
}

func HandelFunc() echo.HandlerFunc {
	return func(c echo.Context) error {
		res := make([]User, 0)
		if err := db.Where("name = ?").Find(&res).Error; err != nil {
			return c.JSON(
				400,
				"NOK",
			)
		}

		return c.JSON(
			200,
			res,
		)
	}
}

type User struct {
	ID   uint   `gorm:"primaryKey,autoIncrement" json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}


func dataDump() {
	db.Save(&User{
		Name: "a1",
	})
	db.Save(&User{
		Name: "a2",
	})
	db.Save(&User{
		Name: "a3",
	})
	db.Save(&User{
		Name: "a4",
	})
	db.Save(&User{
		Name: "a5",
	})
}