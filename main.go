package main

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"sort"
	"strconv"
)

type CandyStore struct {
	Name  string
	Candy string
	Eaten int
}

type FavoriteSnack struct {
	Name          string `json:"name"`
	FavoriteSnack string `json:"favoriteSnack"`
	TotalSnacks   int    `json:"totalSnacks"`
}

func OrderFavoriteSnacksByTotalSnacksDescending(favoriteSnacks []FavoriteSnack) []FavoriteSnack {
	sort.SliceStable(favoriteSnacks, func(i, j int) bool {
		return favoriteSnacks[i].TotalSnacks > favoriteSnacks[j].TotalSnacks
	})
	return favoriteSnacks
}

func ScrapeCandyStorePage() []CandyStore {
	const URL = "https://candystore.zimpler.net"

	res, err := http.Get(URL)
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

	if len(candyStores) == 0 {
		log.Fatal("No candy stores found")
	}

	return candyStores
}

func GetFavoriteSnack(groupedByName map[string][]CandyStore) []FavoriteSnack {
	var favoriteSnacks = make([]FavoriteSnack, 0)

	for name, candyStores := range groupedByName {
		var favoriteSnack = FavoriteSnack{
			Name:          name,
			FavoriteSnack: "",
			TotalSnacks:   0,
		}

		for _, candyStore := range candyStores {
			if candyStore.Eaten > favoriteSnack.TotalSnacks {
				favoriteSnack.FavoriteSnack = candyStore.Candy
			}
			favoriteSnack.TotalSnacks += candyStore.Eaten
		}

		favoriteSnacks = append(favoriteSnacks, favoriteSnack)
	}
	return favoriteSnacks
}

func GroupCandyStoreByCandy(stores []CandyStore) map[string][]CandyStore {
	nameMap := make(map[string][]CandyStore)

	for _, candyStore := range stores {
		if _, ok := nameMap[candyStore.Name]; !ok {
			nameMap[candyStore.Name] = append(nameMap[candyStore.Name], candyStore)
		} else {
			var list = nameMap[candyStore.Name]

			var isCandyInTheList = false
			for idx, candy := range list {
				if candy.Candy == candyStore.Candy {
					list[idx].Eaten += candyStore.Eaten
					isCandyInTheList = true
					break
				}
			}
			if !isCandyInTheList {
				nameMap[candyStore.Name] = append(nameMap[candyStore.Name], candyStore)
			}
		}
	}

	return nameMap
}

func main() {
	// Scrape the candy store page
	candyStores := ScrapeCandyStorePage()

	// Group the candy stores by name
	var groupedByCandy = GroupCandyStoreByCandy(candyStores)

	// Get the favorite snack for each person
	favoriteSnacks := GetFavoriteSnack(groupedByCandy)

	// Order the favorite snacks by total snacks descending
	OrderFavoriteSnacksByTotalSnacksDescending(favoriteSnacks)

	// Convert the favorite snacks to json
	data, _ := json.MarshalIndent(favoriteSnacks, "", " ")

	// Print the favorite snacks
	fmt.Println(string(data))
}
