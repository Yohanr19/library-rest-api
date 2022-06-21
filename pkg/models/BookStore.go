package models

import "gorm.io/gorm"
import "gorm.io/driver/postgres"

type BookStore struct {
	DB *gorm.DB
}

func(bs *BookStore)InitDB()error{
	dsn := `host=localhost port=5432 dbname=library user=yohan password=yohan1234`
	db , err := gorm.Open(postgres.Open(dsn))
	if err!=nil {
		return err
	}
	db.AutoMigrate(&Book{},&Page{})
	bs.DB = db
	return nil
}
func (bs *BookStore)AddBook(book Book, raw []string, html[]string){
	for i, page := range raw {
		rawPage := Page{}
		rawPage.Type = "raw"
		rawPage.PageNumber = i+1
		rawPage.Text = page
		book.Pages = append(book.Pages, rawPage)
	}
	for i, page := range html {
		htmlPage := Page{}
		htmlPage.Type = "html"
		htmlPage.PageNumber = i+1
		htmlPage.Text = page
		book.Pages = append(book.Pages, htmlPage)
	}
	bs.DB.Create(&book)
}
func (bs *BookStore)RemoveBooks(ids ...int){
	for _ ,v := range ids {
		bs.DB.Where(`book_id =?`,v).Delete(&Page{})
		bs.DB.Where(`id = ?`,v).Delete(&Book{})
	}
}
func (bs *BookStore)GetBooks()[]*Book{
	Books := []*Book{}
	bs.DB.Find(&Books)
	for _ , book := range Books{
		bs.DB.Find(&book.Pages,`book_id = ?`,book.ID)
	}
	return Books
}