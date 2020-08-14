package main

import (
	"github.com/stevedesilva/my-proverb-app/api/server"

	"github.com/labstack/echo/v4"
)

const msg = "My proverb-native"

func main() {
	startServer()
}

func startServer() {
	e := echo.New()
	impl := NewProverbsService()
	server.RegisterHandlers(e, impl)

	//e.GET("/", func(e echo.Context) error {
	//	return e.String(http.StatusOK, msg)
	//})
	e.Logger.Fatal(e.Start(":8081"))
}
