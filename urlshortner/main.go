package main

import (
	"crypto/md5"
	"fmt"
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
	fmt.Println(hasher)
	return ""
}

func main() {
	fmt.Println("Url shortener")
	fmt.Println(genrateShortURL("heyhey"))
}
