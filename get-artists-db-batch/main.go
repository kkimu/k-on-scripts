package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	baseUrl := os.Getenv("BASE_URL")

	for page := 0; page <= 70; page++ {
		url := baseUrl + strconv.Itoa(page) + "/"
		fmt.Print(url)
		insertFromSite(url)
		time.Sleep(5 * time.Second)
		fmt.Println(" complete")
	}

}

func insertFromSite(url string) {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	//var artists []Artist

	doc.Find("div .anchor_box").Each(func(i int, doc2 *goquery.Selection) {
		suffix := doc2.Find("dt").Text()
		doc2.Find("li").Each(func(i int, doc3 *goquery.Selection) {
			now := time.Now()
			artist := &Artist{
				id:         uuid.New().String(),
				name:       doc3.Find("a").Text(),
				kanaPrefix: suffix,
				createdAt: now,
				updatedAt: now,
			}
			err := insert(*artist)
			if err != nil {
				log.Fatal(err)
			}
			//artists = append(artists, *artist)
			//fmt.Printf("%s,%s\n", artist.name, artist.kanaPrefix)
		})
	})
}
