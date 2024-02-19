package request

import "favorit-book-tracker/domain/entity"

type RequestBookDTO struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func (book RequestBookDTO) ToBookEntity() entity.Book {
	return entity.Book{
		Title:  book.Title,
		Author: book.Author,
		Rating: book.Rating,
	}
}
