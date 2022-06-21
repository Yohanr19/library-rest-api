package controlers

import (
	"github.com/yohanr19/library-rest-api/pkg/models"

	"encoding/json"
	"net/http"
)

type BookData struct {
	ID     uint       `json:"id"`
	Title  string     `json:"title"`
	Author string     `json:"author"`
	Year   int        `json:"year"`
	Pages  []PageData `json:"pages"`
}
type PageData struct {
	Text       string `json:"text"`
	BookID     uint   `json:"book_id"`
	PageNumber int    `json:"page_number"`
	Type       string `json:"type"`
}

type BookControler struct {
	Store *models.BookStore
}

func (bc *BookControler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var response []BookData

	books := bc.Store.GetBooks()
	for _, book := range books {
		var responseBook BookData
		responseBook.ID = book.ID
		responseBook.Author = book.Author
		responseBook.Title = book.Title
		responseBook.Year = book.Year
		for _, page := range book.Pages {
			var responsePage PageData
			responsePage.Text = page.Text
			responsePage.Type = page.Type
			responsePage.BookID = page.BookID
			responsePage.PageNumber = page.PageNumber
			responseBook.Pages = append(responseBook.Pages, responsePage)
		}
		response = append(response, responseBook)
	}
	enconder := json.NewEncoder(w)
	err := enconder.Encode(&response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}

func (bc *BookControler) Init() error {
	bc.Store = &models.BookStore{}
	return bc.Store.InitDB()
}
