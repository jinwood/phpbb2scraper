package main

import (
	"os"
	"log"
	"net/url"

	//"github.com/gocolly/colly"
)

type YouTubeUrl struct {
	Url string
}

func main() {
	if len(os.Args) != 2 {
		log.Println("missing url argument")
		os.Exit(1)
	}
	if !isValidUrl(os.Args[1]) {
		log.Println("invalid url given")
		os.Exit(1)
	}
	//init collector
}

func isValidUrl(testString string) bool {
	_, err := url.ParseRequestURI(testString)
	if err != nil {
		return false
	} else {
		return true
	}
}

