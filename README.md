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

## Data storage generation
//go:generate jet -source=MySQL -host=localhost -port=3306 -user=root -password=rootpassword -dbname=silvade -path=generated

## Skeema
Stored in skeema/tables
init creates sql from running db
> skeema init -h 127.0.0.1 -u root -p --schema silvade --dir tables

Check difference between local db and tables folder
> skeema diff --password=rootpassword

push changes to running db
skeema push --password=rootpassword