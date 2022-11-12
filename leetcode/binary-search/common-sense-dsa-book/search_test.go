package common_sense_dsa_book

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	searchField := []int{-1, 0, 3, 5, 9, 12}

	for _, searchNumber := range searchField {
		fmt.Printf("Searching for %d in list %v\n", searchNumber, searchField)
		result, searchCount := Search(searchField, searchNumber)
		fmt.Printf("Your number was found in position %d after %d steps.\n\n", result, searchCount)
	}
}
