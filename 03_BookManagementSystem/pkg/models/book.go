package models

import (
	"github.com/jinzhu/gorm"
	"github.com/learn-go/03_Book/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	//main함수가 실행되기 전에 실행된다.

	//db와 연결한다.
	config.Connect()
	//db를 가져온다.
	db = config.GetDB()
	//db에 테이블을 생성한다.
	db.AutoMigrate(&Book{})
}

// Book 리시버가 있다.
func (b *Book) CreateBook() *Book {
	//Primary Key가 비어있는지 아닌지를 판단한다.
	//boolean을 반환한다.
	//여기서는 if문으로 묶거나 하지도 않았네?
	db.NewRecord(b)

	//새 Row를 생성한다.
	db.Create(b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	//조건없이 Find니까 모든 Row를 가져온다.
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	//조건에 맞는 Row를 찾는다.
	//체인이 가능하다.
	db := db.Where("ID=?", Id).Find(&getBook) //here db is local var
	return &getBook, db
}

func DeleteBook(Id int64) Book {
	var book Book
	//조건에 맞는 Row를 삭제한다.
	db.Where("ID=?", Id).Delete(book)
	return book
}
