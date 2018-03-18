package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/vimsucks/blogsucks/model"
	"github.com/vimsucks/blogsucks/router"
)

func main() {
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		e.Logger.Fatal(err)
	}
	defer db.Close()
	model.InitDb(db)

	router.SetupRouters(e)
	e.Logger.Fatal(e.Start(":4000"))
}
