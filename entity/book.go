package entity

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Id       int64
	Title    string
	AuthorId int64
	GenreId  int64
	Summary  string
	Status   int8
}
