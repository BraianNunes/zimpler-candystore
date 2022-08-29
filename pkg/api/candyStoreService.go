package api

import (
	"encoding/json"
	"sort"
)

func orderFavoriteSnacksByTotalSnacksDescending(favoriteSnacks []FavoriteSnack) {
	sort.SliceStable(favoriteSnacks, func(i, j int) bool {
		return favoriteSnacks[i].TotalSnacks > favoriteSnacks[j].TotalSnacks
	})
}

func getFavoriteSnack(groupedByName map[string][]CandyStore) []FavoriteSnack {
	var favoriteSnacks = make([]FavoriteSnack, 0)

	for name, candyStores := range groupedByName {
		var favoriteSnack = FavoriteSnack{
			Name:          name,
			FavoriteSnack: "",
			TotalSnacks:   0,
		}

		for _, c := range candyStores {
			if c.Eaten > favoriteSnack.TotalSnacks {
				favoriteSnack.FavoriteSnack = c.Candy
			}
			favoriteSnack.TotalSnacks += c.Eaten
		}

		favoriteSnacks = append(favoriteSnacks, favoriteSnack)
	}
	return favoriteSnacks
}

func groupCandyStoreByCandy(stores []CandyStore) map[string][]CandyStore {
	nameMap := make(map[string][]CandyStore)

	for _, candyStore := range stores {
		if _, ok := nameMap[candyStore.Name]; !ok {
			nameMap[candyStore.Name] = append(nameMap[candyStore.Name], candyStore)
		} else {
			var candyStores = nameMap[candyStore.Name]
			var isCandyInTheList = false

			for idx, c := range candyStores {
				if c.Candy == candyStore.Candy {
					candyStores[idx].Eaten += candyStore.Eaten
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

func GetTopCustomersByCandy(candyStore []CandyStore) []FavoriteSnack {
	// Group the candy stores by name
	var groupedByCandy = groupCandyStoreByCandy(candyStore)

	// Get the favorite snack for each person
	favoriteSnacks := getFavoriteSnack(groupedByCandy)

	// Order the favorite snacks by total snacks descending
	orderFavoriteSnacksByTotalSnacksDescending(favoriteSnacks)

	return favoriteSnacks
}

func ConvertToJson(data interface{}) string {
	json, _ := json.MarshalIndent(data, "", " ")
	return string(json)
}
