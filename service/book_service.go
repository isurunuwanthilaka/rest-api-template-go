package service

import (
	"rest-api-template-go/dao"
	"rest-api-template-go/dto"
)

func AddBook(newBook dto.BookReqDto) (*dto.BookResDto, error) {

	var book = dao.Book{Title: newBook.Title, Summary: newBook.Summary, AuthorId: newBook.AuthorId, GenreId: newBook.GenreId, StatusId: 1}
	book.Save()

	var res = dto.BookResDto{
		Id:      book.ID,
		Summary: book.Summary,
		Title:   book.Title,
		Author:  book.Author.Name,
		Genre:   book.Genre.Name,
	}
	return &res, nil
}
