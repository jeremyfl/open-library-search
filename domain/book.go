package domain

// Book is the main model of book
type Book struct {
	ID          int    `json:"id"`
	Title       string `json:"name"`
	Description string `json:"description"`
	AuthorName  string `json:"author_name"`
	Thumbnail   string `json:"thumbnail"`
}

// BookService represent the book service or business logic
type BookService interface {
	//
}

// BookRepository represent the book repository
type BookRepository interface {
	//
}
