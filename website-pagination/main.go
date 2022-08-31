package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

func fetchItemsToDisplay(items [][]string, sortParameter, sortOrder, itemsPerPage, pageNumber int) []string {
	var result []string

	sort.Slice(items, func(i, j int) bool {
		ie, err1 := strconv.Atoi(items[i][sortParameter])
		je, err2 := strconv.Atoi(items[j][sortParameter])
		if err1 != nil || err2 != nil {
			return false
		}
		if sortOrder == 0 {
			return ie < je
		}
		return ie > je
	})

	fmt.Println(items)

	startIndex := pageNumber * itemsPerPage
	endIndex := (pageNumber + 1) * itemsPerPage
	for _, item := range items[startIndex:int(math.Min(float64(endIndex), float64(len(items))))] {
		result = append(result, item[0])
	}
	fmt.Println("Pagination result: ", result)
	return result
}

func main() {
	// Sample case 1
	items := [][]string{{"item1", "10", "15"}, {"item2", "3", "4"}, {"item3", "17", "8"}}
	sortParameter := 1
	sortOrder := 0
	itemsPerPage := 2
	pageNumber := 1
	fetchItemsToDisplay(items, sortParameter, sortOrder, itemsPerPage, pageNumber)

	// Sample case 2
	items = [][]string{
		{"owievtkuwv", "58584272", "62930912"},
		{"rpaggbixik", "9425650", "96088250"},
		{"dfbkasygcn", "37469674", "46363902"},
		{"virrisdfxe", "18666489", "88046739"},
	}
	sortParameter = 2
	sortOrder = 1
	itemsPerPage = 2
	pageNumber = 0
	fetchItemsToDisplay(items, sortParameter, sortOrder, itemsPerPage, pageNumber)
}
