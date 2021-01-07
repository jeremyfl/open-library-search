package service

import "open-library-search/domain"

type bookService struct {
	repo domain.BookRepository
}

// SearchBook search book from repo and return
// list of domain book
func (bs *bookService) SearchBook(keyword *string) *[]domain.Book {
	return &[]domain.Book{
		{
			ID:          1,
			Title:       "Leadership 101",
			Description: "The Maxwell Daily Reader",
			AuthorName:  "John Maxwell",
			Thumbnail:   "leadership-john-1.png",
		},
		{
			ID:          2,
			Title:       "Clean Architecture: A Craftsman's Guide to Software Structure",
			Description: "Practical Software Architecture Solutions from the Legendary Robert C. Martin (“Uncle Bob”) By applying universal rules of software architecture, you can dramatically improve developer productivity throughout the life of any software system.o",
			AuthorName:  "Uncle Bob",
			Thumbnail:   "clean-architecture-uncle-bob-1.png",
		},
	}
}
