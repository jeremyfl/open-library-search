package http

import (
	"open-library-search/domain"

	"github.com/labstack/echo/v4"
)

// BookController represent http handler for book service
type BookController struct {
	Service domain.BookService
}

// SearchBook search book entities for HTTP call
func (bc *BookController) SearchBook(c echo.Context) error {
	keyword := c.QueryParam("title")
	data := bc.Service.SearchBook(keyword)

	return GenerateResponse(c, 200, "Fetch Book success", data)
}
