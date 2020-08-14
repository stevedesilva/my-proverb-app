package main

import (
	"fmt"

	"github.com/stevedesilva/my-proverb-app/api/server"

	"github.com/labstack/echo/v4"
)

/*
// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /info)
	GetInformation(ctx echo.Context) error

	// (GET /proverbs)
	FindProverbs(ctx echo.Context, params FindProverbsParams) error

	// (POST /proverbs)
	AddProverb(ctx echo.Context) error

	// (DELETE /proverbs/{id})
	DeleteProverb(ctx echo.Context, id int64) error

	// (GET /proverbs/{id})
	FindProverbById(ctx echo.Context, id int64) error
}
*/
type ProverbsService struct{}

func NewProverbsService() *ProverbsService {
	return &ProverbsService{}
}
func (s *ProverbsService) FindProverbs(ctx echo.Context, params server.FindProverbsParams) error {
	fmt.Println("FindProverbs")
	return nil
}

func (s *ProverbsService) AddProverb(ctx echo.Context) error {
	fmt.Println("AddProverb")
	return nil
}

func (s *ProverbsService) DeleteProverb(ctx echo.Context, id int64) error {
	fmt.Println("DeleteProverb")
	return nil
}

func (s *ProverbsService) FindProverbById(ctx echo.Context, id int64) error {
	fmt.Println("FindProverbById")
	return nil
}

func (s *ProverbsService) GetInformation(ctx echo.Context) error {
	fmt.Println("GetInformation")
	return nil
}
