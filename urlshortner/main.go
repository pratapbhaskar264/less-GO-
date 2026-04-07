package main

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"time"
)

// struct in golang are same as DTOs in java

type URL struct {
	ID           string    `json:"id"`
	OriginalURL  string    `json:"original_url"`
	ShortURL     string    `json:"short_url"`
	CreationDate time.Time `json:"creation_date"`
}

// map will work as in memory db here so we will create one
//mapping : shortUrl -> struct

var urlDB = make(map[string]URL)

//same as Pair class in java and hashmap of string -> pair ..... ez

// function to generate short url from original url
func genrateShortURL(OriginalURL string) string {
	hasher := md5.New()               // hasher
	hasher.Write([]byte(OriginalURL)) // converted to byte
	// fmt.Println(hasher)
	data := hasher.Sum(nil)
	hash := hex.EncodeToString(data)
	return hash[:8]
}

// db function to save the url mapping
func createURL(OriginalURL string) string {
	shortURL := genrateShortURL(OriginalURL)

	id := shortURL

	urlDB[id] = URL{
		ID:           id,
		OriginalURL:  OriginalURL,
		ShortURL:     shortURL,
		CreationDate: time.Now(),
	}
	return shortURL
}

// genrate URL by id
func getURL(id string) (URL, error) {
	url, ok := urlDB[id]

	if !ok {
		return URL{}, errors.New("URL Not Found")
	}
	return url, nil
}

func main() {
	fmt.Println("Url shortener")

	// creating server in GoLang
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server", err)
	}

}
