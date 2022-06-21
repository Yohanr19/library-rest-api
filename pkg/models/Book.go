package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title  string
	Author string
	Year   int
	Pages  []Page
}
type Page struct {
	gorm.Model
	Text       string
	BookID     uint //Foreign Key
	PageNumber int
	Type       string
}
