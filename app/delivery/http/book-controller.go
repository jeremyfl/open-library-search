package http

import (
	"open-library-search/domain"

	"github.com/labstack/echo/v4"
)

// BookController represent http handler for book service
type BookController struct {
	Service domain.BookService
}

func (bc *BookController) SearchBook(c echo.Context) error {
	book := &domain.Book{
		ID:          1,
		Title:       "Leadership 101",
		Description: "The Maxwell Daily Reader",
		AuthorName:  "John Maxwell",
		Thumbnail:   "leadership-john-1.png",
	}

	return GenerateResponse(c, 200, "Fetch Book success", book)
}
