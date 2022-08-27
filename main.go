package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"strconv"
)

// CandyStore Create a struct that represents a candy store
type CandyStore struct {
	Name  string
	Candy string
	Eaten int
	Total int
}

type FavoriteSnack struct {
	Name          string
	FavoriteSnack string
	totalSnacks   int
}

func ExampleScrape() {
	res, err := http.Get("https://candystore.zimpler.net/")
	if err != nil {
		log.Fatal(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(res.Body)

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Create a list to hold the candy stores
	var candyStores = make([]CandyStore, 0)

	// Find all links and process them with the function
	// defined earlier
	// Find a table with id "top.customers"
	doc.Find("table[id='top.customers']").Each(func(i int, item *goquery.Selection) {
		item.Find("tbody tr").Each(func(index int, item *goquery.Selection) {
			name := item.Find("td:nth-child(1)").Text()
			candy := item.Find("td:nth-child(2)").Text()
			eaten, _ := strconv.Atoi(item.Find("td:nth-child(3)").Text())

			// Create a new instance of the struct
			var candyStore = CandyStore{
				Name:  name,
				Candy: candy,
				Eaten: eaten,
			}

			// Append the new instance to the list
			candyStores = append(candyStores, candyStore)
		})
	})

	// Print the list of candy stores
	//PrintCandyStoreList(candyStores)

	GroupCandyStoreByCandy(candyStores)

	//GroupCandyStoreByName(candyStores)
}

func GroupCandyStoreByName(stores []CandyStore) {
	// Group the candy stores by name
	nameMap := make(map[string][]CandyStore)
	for _, candyStore := range stores {
		nameMap[candyStore.Name] = append(nameMap[candyStore.Name], candyStore)
	}

	// Print the candy stores grouped by name
	for name, candyStores := range nameMap {
		fmt.Println("Name: ", name)
		for _, candyStore := range candyStores {
			// Printf with a new line
			fmt.Printf("Name: %s, Eaten: %d, Candy: %s, \n", candyStore.Name, candyStore.Eaten, candyStore.Candy)
		}
	}
}

func PrintCandyStoreList(candyStores []CandyStore) {
	for _, candyStore := range candyStores {
		fmt.Printf("Name: %s, Eaten: %d, Candy: %s \n", candyStore.Name, candyStore.Eaten, candyStore.Candy)
	}
}

func GroupCandyStoreByCandy(stores []CandyStore) {
	// Group the candy stores by candy
	nameMap := make(map[string][]CandyStore)
	//map2 := make(map[string][]CandyStore)

	// Calculate the total of totalSnacks by candy

	for _, candyStore := range stores {
		// Check if the Name is already in the map
		if _, ok := nameMap[candyStore.Name]; !ok {
			nameMap[candyStore.Name] = append(nameMap[candyStore.Name], candyStore)
		} else {
			// change the value of Name in nameMap[candyStore.Name]
			var list = nameMap[candyStore.Name]

			var foundNameInList = false
			for idx, candy := range list {
				// ge the pointer to the candyStore
				if candy.Candy == candyStore.Candy {
					list[idx].Eaten += candyStore.Eaten

					fmt.Println("Found candy in list")
					foundNameInList = true
					break
				}
			}

			if !foundNameInList {
				nameMap[candyStore.Name] = append(nameMap[candyStore.Name], candyStore)
			}
		}
	}

	//candyMap := make(map[string][]CandyStore)
	//Do a for in nameMap
	//for _, candyStores := range nameMap {
	//	Group the candy stores by candy
	//for _, candyStore := range candyStores {
	//	candyMap[candyStore.Candy] = append(candyMap[candyStore.Candy], candyStore)
	//}
	//}

	// Print the candy stores grouped by candy
	for name, candyStores := range nameMap {
		fmt.Println("Name: ", name)
		for _, candyStore := range candyStores {
			// Printf with a new line
			fmt.Printf("Name: %s, Candy: %s, Eaten: %d\n", candyStore.Name, candyStore.Candy, candyStore.Eaten)
		}
	}

}

func main() {
	ExampleScrape()
	//TestPointers()
}

func TestPointers() {
	// Create a new instance of the struct
	var candyStore = CandyStore{
		Name:  "Zimpler",
		Candy: "Gummy Bears",
		Eaten: 0,
		Total: 0,
	}

	// Print the struct
	fmt.Println(candyStore)

	// Add the struct to a list
	var candyStores = make([]CandyStore, 0)
	candyStores = append(candyStores, candyStore)

	fmt.Println(candyStores)

	// Edit the struct
	var candyStore2 = &candyStores[0]
	candyStore2.Eaten = 1
	candyStore2.Total = 1

	fmt.Println(candyStores)

}
