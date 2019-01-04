package main

import (
	"log"
	"net/url"
	"os"

	"github.com/gocolly/colly"
)

type YouTubeUrl struct {
	Url string
}

func main() {
	if len(os.Args) != 2 {
		log.Println("missing url argument")
		os.Exit(1)
	}
	scrapeUrl := os.Args[1]
	if !isValidUrl(scrapeUrl) {
		log.Println("invalid url given")
		os.Exit(1)
	}
	//init collector
	c := colly.NewCollector(
		colly.AllowedDomains(scrapeUrl),
		colly.CacheDir("./colly_cache"),
	)

	//videos := make([]YouTubeUrl, 0, 999)

	loginUrl := scrapeUrl + "?login/login"
	log.Println("attempting to login at ", loginUrl)
	err := c.Post(loginUrl, map[string]string{"login": "xxx", "password": "xxx"})
	if err != nil {
		log.Println("there was an error")
		log.Fatal(err)
	} else {
		log.Println("i think it worked....")
	}
}

func isValidUrl(testString string) bool {
	_, err := url.ParseRequestURI(testString)
	if err != nil {
		return false
	} else {
		return true
	}

