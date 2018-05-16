/*
	https://echo.labstack.com/guide
*/

package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello Word!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
