package service

import (
	"rest-api-template-go/dao"
	"rest-api-template-go/dto"
	"rest-api-template-go/utils/errors"
	"rest-api-template-go/utils/log"
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

func GetBook(id uint) (*dto.BookResDto, error) {

	book := dao.GetBookById(id)
	if book.ID == 0 {
		log.Error("Book not found for id : %x", id)
		return nil, errors.NewBadRequestError("Book not found.")
	}

	var res = dto.BookResDto{
		Id:      book.ID,
		Summary: book.Summary,
		Title:   book.Title,
		Author:  book.Author.Name,
		Genre:   book.Genre.Name,
	}
	return &res, nil
}
