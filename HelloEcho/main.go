package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type M map[string]interface{}

var ActionAbout = echo.WrapHandler(
	http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello world, this is about page"))
		},
	),
)

func main() {
	r := echo.New()

	r.GET("/", func(ctx echo.Context) error {
		data := "Helo World from /index"
		return ctx.String(http.StatusOK, data)
	})

	r.GET("/page1", func(ctx echo.Context) error {
		name := ctx.QueryParam("name")
		data := fmt.Sprintf("Hello %s", name)

		return ctx.String(http.StatusOK, data)
	})

	r.GET("/page2/:name", func(ctx echo.Context) error {
		name := ctx.Param("name")
		data := fmt.Sprintf("Hello %s in page2", name)

		return ctx.String(http.StatusOK, data)
	})

	r.GET("/about", ActionAbout)

	r.Start(":9000")
}
