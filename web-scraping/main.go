package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gocolly/colly"
)

const userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"
const domain = "quotes.toscrape.com"
const url = "https://quotes.toscrape.com"

func main() {

	//setup
	c := colly.NewCollector(
		colly.AllowedDomains(domain),
	)

	c.SetRequestTimeout(30 * time.Second)

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", userAgent)
		fmt.Printf("Visiting %v\n", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Printf("Response StatusCode: %v", r.StatusCode)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("Error: %v", err)
	})
	
	// extract data
	var quoteList []Quote
	c.OnHTML(".quote", func (h *colly.HTMLElement)  {
		dom := h.DOM
		title := dom.Find(".text").Text()
		author := dom.Find(".author").Text()
		
		q := Quote {
			Title: title,
			Author: author,
		}
		quoteList = append(quoteList, q)
	})
	
	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
		js, err := json.MarshalIndent(quoteList, "", "    ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Writing data to file")
		if err := os.WriteFile("data.json", js, 0664); err == nil {
			fmt.Println("Data written to file successfully")
		}
 
	})

	c.Visit(url)
}

type Quote struct {
	Title string
	Author string
}

