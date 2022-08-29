package api

import (
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"strconv"
)

func ScrapeCandyStorePage() []CandyStore {
	// Should be in a config file
	const url = "https://candystore.zimpler.net"

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(res.Body)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var candyStores = make([]CandyStore, 0)

	doc.Find("table[id='top.customers']").Each(func(i int, item *goquery.Selection) {
		item.Find("tbody tr").Each(func(index int, item *goquery.Selection) {
			name := item.Find("td:nth-child(1)").Text()
			candy := item.Find("td:nth-child(2)").Text()
			eaten, _ := strconv.Atoi(item.Find("td:nth-child(3)").Text())

			var candyStore = CandyStore{
				Name:  name,
				Candy: candy,
				Eaten: eaten,
			}
			candyStores = append(candyStores, candyStore)
		})
	})

	return candyStores
}
