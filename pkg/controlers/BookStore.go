package controlers

import (
	"fmt"
	"strconv"

	"github.com/yohanr19/library-rest-api/pkg/models"
	"gorm.io/gorm"
	"github.com/yohanr19/library-rest-api/pkg/views"

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
	store *models.BookStore
}

func (bc *BookControler) GetBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var responseBook BookData
	r.ParseForm()
	query := r.URL.Query().Get("id")
	id, err := strconv.Atoi(query)
	if err != nil {
		http.Error(w, "Bad Query", http.StatusBadRequest)
		return
	}
	book, err := bc.store.GetBookByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}
	copyBook(&responseBook, book)
	enconder := json.NewEncoder(w)
	err = enconder.Encode(&responseBook)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
func (bc *BookControler) GetPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	r.ParseForm()
	bookid, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Bad Query", http.StatusBadRequest)
		return
	}
	pageNumber, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "Bad Query", http.StatusBadRequest)
		return
	}
	pageType := r.URL.Query().Get("type")
	book, err := bc.store.GetBookByID(bookid)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}
	if len(book.Pages) == 0 {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	for i, v := range book.Pages {
		if v.PageNumber == pageNumber && v.Type == pageType {
			fmt.Fprintf(w, "%s", v.Text)
			return
		}
		if i == len(book.Pages)-1 {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}
	}
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (bc *BookControler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var response []BookData

	books := bc.store.GetBooks()
	for _, book := range books {
		var responseBook BookData
		copyBook(&responseBook, book)
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
	bc.store = &models.BookStore{}
	return bc.store.InitDB()
}

func copyBook(responseStruct *BookData, book *models.Book) {
	responseStruct.ID = book.ID
	responseStruct.Author = book.Author
	responseStruct.Title = book.Title
	responseStruct.Year = book.Year
	for _, page := range book.Pages {
		var responsePage PageData
		responsePage.Text = page.Text
		responsePage.Type = page.Type
		responsePage.BookID = page.BookID
		responsePage.PageNumber = page.PageNumber
		responseStruct.Pages = append(responseStruct.Pages, responsePage)
	}
}
func (bs *BookControler)CreateBook(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodGet{
		err := views.RenderForm(w, nil)
		if err!=nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			fmt.Print(err)
		}
		return 
	}
	if r.Method == http.MethodPost{
		fmt.Fprint(w,"You got the post method")
		return
	}
	

}