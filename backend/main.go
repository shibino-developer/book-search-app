package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Book represents a single book item
type Book struct {
	Title       string   `json:"title"`
	Authors     []string `json:"authors"`
	Description string   `json:"description"`
}

// VolumeInfo represents the volume information of a book
type VolumeInfo struct {
	Title       string   `json:"title"`
	Authors     []string `json:"authors"`
	Description string   `json:"description"`
}

// Item represents a single item in the API response
type Item struct {
	VolumeInfo VolumeInfo `json:"volumeInfo"`
}

// BooksResponse represents the API response structure
type BooksResponse struct {
	Items []Item `json:"items"`
}

// getBooksHandler handles the /books endpoint
func getBooksHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "Missing 'q' parameter", http.StatusBadRequest)
		return
	}

	url := "https://www.googleapis.com/books/v1/volumes?q=" + query + "&key=" + os.Getenv("AIzaSyC00Su4QXe_NvBuDBKCNktDqkQQKo8_Qwc")
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Failed to fetch data from Google Books API: %v", err)
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Failed to read response body: %v", err)
		http.Error(w, "Failed to read response body", http.StatusInternalServerError)
		return
	}

	log.Printf("Raw JSON response: %s", string(body))

	var booksResp BooksResponse
	if err := json.Unmarshal(body, &booksResp); err != nil {
		log.Printf("Failed to decode JSON response: %v", err)
		http.Error(w, "Failed to decode JSON response", http.StatusInternalServerError)
		return
	}

	var books []Book
	for _, item := range booksResp.Items {
		book := Book{
			Title:       item.VolumeInfo.Title,
			Authors:     item.VolumeInfo.Authors,
			Description: item.VolumeInfo.Description,
		}
		books = append(books, book)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func main() {
	http.HandleFunc("/books", getBooksHandler)
	log.Println("Server is listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
