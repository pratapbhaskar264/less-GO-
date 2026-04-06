package main

import (
	"fmt"
	"time"
)

type URL struct {
	ID           string    `json:"id"`
	OriginalURL  string    `json:"original_url"`
	ShortURL     string    `json:"short_url"`
	CreationDate time.Time `json:"creation_date"`
}

// map will work as in memory db here so we will create one
//mapping : shortUrl -> struct

func main() {
	fmt.Println("Url shortener")
}
