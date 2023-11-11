package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "modernc.org/sqlite"
)

var (
	lock = sync.Mutex{}
)

type CustomContext struct {
	echo.Context
}

func (c *CustomContext) Foo(test string) string {
	println(test)
	return test
}

func mid1(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("i was here")
		return next(c)
	}
}

func hand1(c echo.Context) error {
	fmt.Println("then here")
	cc := c.(*CustomContext)
	fmt.Println(cc.Foo("work?"), 1)
	return c.JSON(http.StatusOK, "ayoooooo")
}

func main17() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &CustomContext{c}
			return next(cc)
		}
	})

	// Routes
	e.GET("/", hand1, mid1)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
