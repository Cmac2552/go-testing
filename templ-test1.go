package main

import (
	"context"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
func main() {

	server := echo.New()
	// t := &Template{
	// 	templates: template.Must(template.ParseGlob("public/views/*.html")),
	// }
	// server.Renderer = t
	server.Use(middleware.Logger())
	server.Static("/assets", "assets")
	component := hello("John")
	component2 := bye()
	// fmt.Println(component)

	server.GET("/", func(c echo.Context) error {
		return component.Render(context.Background(), c.Response().Writer)
	})

	server.GET("/replace", func(c echo.Context) error {
		return component2.Render(context.Background(), c.Response().Writer)
	})

	server.Logger.Fatal(server.Start(":1323"))
}
