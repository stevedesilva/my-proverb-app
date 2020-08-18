# Proverbs 
Proverb application spreading wisdom to the masses


[Swagger doc](http://localhost:8081)

## API Documentation
- [Swagger doc](http://localhost:8081/swaggerui/index.html)
- [ReDoc](http://localhost:8081/redoc/index.html)
- [Project Structure](#structure)
- [Web Server](#server)
- [Open API](#Open API)
- [structure](https://github.com/golang-standards/project-layout)
- [Documentation] (#documentation)

## structure
https://github.com/golang-standards/project-layout

## server
go get -u github.com/labstack/echo/...

run via cmd/proverb-echo/main.go

##  Open API
https://github.com/deepmap/oapi-codegen

oapi-codegen --generate types,server --package server ./api/swagger/proverb.yaml > ./api/server/server.go

## ReDoc
https://github.com/Redocly/redoc