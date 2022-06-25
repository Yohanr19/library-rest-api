package main

import (
	"fmt"
	"log"
	"net/http"

	//"github.com/yohanr19/library-rest-api/pkg/models"
	"github.com/yohanr19/library-rest-api/pkg/controlers"
)

//var rawPages []string = []string{"fine rawpage number 1","fine rawpage number 2","fine rawpage number 3","fine rawpage number 4"}
//var htmlPages []string = []string{"fine html page number 1","fine html page number 2","fine html page number 3"}

func main() {

	BookControler := controlers.BookControler{}
	err := BookControler.Init()
	if err != nil {
		fmt.Println(err)
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/books", BookControler.GetAllBooks)
	mux.HandleFunc("/book", BookControler.GetBook)
	mux.HandleFunc("/page", BookControler.GetPage)
	mux.HandleFunc("/create",BookControler.CreateBook)

	log.Fatal(http.ListenAndServe("localhost:3001", mux))
	/*
		BookStore.AddBook(models.Book{
			Title: "FineBook",
			Author: "Not R19",
			Year:2022,
		}, rawPages, htmlPages)
		//BookStore.RemoveBooks(1,2)

		books := BookStore.GetBooks()
		for _ , book := range books {
			for _, page := range book.Pages {
				if page.Type == "html"{
					fmt.Println(page.Text)
				}
			}
		} */

}
