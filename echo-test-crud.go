package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "modernc.org/sqlite"
)

var (
	lock2 = sync.Mutex{}
)

func main5() {
	e := echo.New()

	db, err1 := sql.Open("sqlite", "./DB_1.db")
	if err1 != nil {
		fmt.Println(err1)
	}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", func(c echo.Context) error {

		rows, err := db.Query("SELECT id, words FROM Test")
		if err != nil {
			return err
		}
		defer rows.Close()
		rows.Next()
		number := 0
		words := ""
		rows.Scan(&number, &words)
		fmt.Println(number, words)
		return c.JSON(http.StatusOK, 4)
	})
	e.POST("/words", func(c echo.Context) error {

		stmt, err := db.Prepare("INSERT INTO Test (id, words) VALUES (?, ?)")
		stmt.Exec(nil, "words")
		if err != nil {
			return err
		}
		defer stmt.Close()
		return c.JSON(http.StatusOK, "Added person")
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
