package main

import (
	"simple_api/module/db"
	"simple_api/router"

	"github.com/labstack/echo"
)

func main() {
	err := db.InitMongo()
	if err != nil {
		panic(err)
	}
	start()
}
func start() {
	e := echo.New()
	e.HideBanner = true
	router.Router(e)
	e.Logger.Fatal(e.Start(":8080"))
}
