package main

import (
	"fmt"
	"zimpler-candystore/pkg/api"
)

func main() {
	// Scrape the candy store page
	candyStores := api.ScrapeCandyStorePage()

	// Get the favorite snack for each person
	favoriteSnacks := api.GetTopCustomersByCandy(candyStores)

	// Convert the favorite snacks to json
	json := api.ConvertToJson(favoriteSnacks)

	// Print the result
	fmt.Println(json)
}
