package main

import (
	"github.com/stevedesilva/my-proverb-app/api/server"

	"github.com/labstack/echo/v4"
)

const msg = "My proverb-echo"

func main() {
	startServer()
}

func startServer() {
	e := echo.New()
	// open api codegen used to generate client and server stubs
	impl := NewProverbsService()
	server.RegisterHandlers(e, impl)

	// serve static content
	e.Static("/swaggerui", "api/swaggerui")
	e.Static("/redoc", "api/redoc")

	e.Logger.Fatal(e.Start(":8081"))
}
