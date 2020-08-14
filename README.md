# Proverbs 
Proverb application spreading wisdom to the masses


[Swagger doc](http://localhost:8081)

## API Documentation
- [Swagger doc](http://localhost:8081)
- [Project Structure](#structure)
- [Web Server](#server)
- [Open API](#Open API)
- [Documentation] (#documentation)

## structure
https://github.com/golang-standards/project-layout

## server
go get -u github.com/labstack/echo/...

run via cmd/proverb-echo/main.go

##  Open API
oapi-codegen --generate types,server --package server ./docs/swagger/proverb.yaml > ./api/server/server.go

