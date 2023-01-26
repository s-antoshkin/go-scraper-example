package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const BASE_URL = "http://books.toscrape.com/"

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	pages := make(chan int, 50)
	results := make(chan []string)
	file, err := os.Create("books.csv")
	check(err)
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.Comma = ';'
	defer writer.Flush()

	start := time.Now()

	for i := 1; i <= cap(pages); i++ {
		go scraping(pages, results)
	}

	go func() {
		for i := 1; i <= cap(pages); i++ {
			pages <- i
		}
	}()

	for i := 1; i <= 1000; i++ {
		result := <-results

		if err := writer.Write(result); err != nil {
			fmt.Println(err)
		}
	}

	close(pages)
	close(results)

	fmt.Printf("End: %s", time.Since(start))

}

func scraping(pages chan int, result chan []string) {
	for i := range pages {
		res, err := http.Get(BASE_URL + fmt.Sprintf("catalogue/page-%d.html", i))
		check(err)

		if res.StatusCode > 400 {
			fmt.Println("Status code: ", res.StatusCode)
		}

		doc, err := goquery.NewDocumentFromReader(res.Body)
		check(err)

		s := "article.product_pod"
		doc.Find(s).Each(func(i int, s *goquery.Selection) {
			h3A := s.Find("h3 a")
			url, ok := h3A.Attr("href")
			if !ok {
				return
			}
			fullURL := BASE_URL + url

			title, ok := h3A.Attr("title")
			if !ok {
				return
			}
			stringPrice := strings.TrimSpace(s.Find("div.product_price p.price_color").Text())
			price := strings.Replace(stringPrice, "£", "", 1)

			available := strings.TrimSpace(s.Find("div.product_price p.instock").Text())

			book := []string{title, fullURL, price, available}
			result <- book

		})

		res.Body.Close()
	}

}
