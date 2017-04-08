package main

import (
	"github.com/labstack/echo"
	"net/http"
	"fmt"
)

func mock(ctx echo.Context) error {
	fmt.Println(ctx.Request().Body)

	return ctx.NoContent(http.StatusOK, )
}

// RateCrypto return rate between currency -- JUST MOCK
func main() {
	e := echo.New()

	e.GET("/mock", mock)

	e.Logger.Fatal(e.Start("0.0.0.0:8081"))
}
