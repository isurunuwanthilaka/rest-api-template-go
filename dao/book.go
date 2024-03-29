package dao

import (
	"os"
	"rest-api-template-go/datasource/mysql"
	"rest-api-template-go/utils/log"
	"strconv"

	"gorm.io/gorm/clause"

	"github.com/joho/godotenv"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func init() {

	e := godotenv.Load("dev.env") //Load .env file
	if e != nil {
		log.Error("%s", e)
	}

	dbCreate, _ := strconv.ParseBool(os.Getenv("db_create"))
	db = mysql.GetDB()

	if dbCreate {
		if err := db.AutoMigrate(&Status{}, &Gender{}, &Genre{}, &Author{}, &Book{}); err != nil {
			return
		}

		db.Create(&Gender{Name: "male"})
		db.Create(&Gender{Name: "female"})

		db.Create(&Status{Name: "active"})
		db.Create(&Status{Name: "deleted"})

		db.Create(&Genre{Name: "adventure"})

		db.Create(&Author{Name: "Isuru Nuwanthilaka", Age: 26, GenderId: 1, StatusId: 1})

		db.Create(&Book{Title: "Programming", AuthorId: 1, GenreId: 1, Summary: "History of programming", StatusId: 1})

	}
}

type Book struct {
	gorm.Model
	Title    string
	AuthorId uint
	Author   Author `gorm:"foreignKey:AuthorId"`
	GenreId  uint
	Genre    Genre `gorm:"foreignKey:GenreId"`
	Summary  string
	StatusId uint
	Status   Status `gorm:"foreignKey:StatusId"`
}

func (book *Book) Save() {
	db.Create(&book)
	db.Preload(clause.Associations).Find(&book, book.ID)
}

func UpdateBook(newBook Book, id uint) Book {
	var book Book
	book.ID = id
	db.Model(&book).Updates(Book{Title: newBook.Title, Summary: newBook.Summary})
	db.Preload(clause.Associations).Find(&book, book.ID)
	return book
}

func GetBookById(id uint) Book {
	var book Book
	db.Preload(clause.Associations).Find(&book, id)
	return book
}

func GetBooks() []Book {
	var books []Book
	db.Preload(clause.Associations).Find(&books)
	return books
}

func DeleteBookById(id uint) error {
	var book Book
	book.ID = id
	result := db.Delete(&book)
	return result.Error
}
