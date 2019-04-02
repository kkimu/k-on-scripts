package main

import (
	"database/sql"
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
		prefix := doc2.Find("dt").Text()
		doc2.Find("li").Each(func(i int, doc3 *goquery.Selection) {
			artistName := doc3.Find("p.name > a").Text()
			artist, err := getArtistByName(artistName)
			if err != nil {
				if err != sql.ErrNoRows {
					log.Fatal(err)
					return
				}
			}

			if artist.id != "" {
				return
			}


			log.Println(artistName)
			now := time.Now()
			artist2 := &Artist{
				id:         uuid.New().String(),
				name:       artistName,
				kanaPrefix: prefix,
				createdAt:  now,
				updatedAt:  now,
			}

			if err := insert(*artist2); err != nil {
				log.Fatal(err)
			}
		})
	})
}
