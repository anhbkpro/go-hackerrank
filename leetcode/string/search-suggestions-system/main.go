package search_suggestions_system

import (
	"fmt"
	"math"
	"sort"
)

func SuggestedProducts(products []string, searchWord string) [][]string {
	ans := make([][]string, 0)
	l, r := 0, len(products)-1
	sort.Slice(products, func(i, j int) bool {
		return products[i] < products[j]
	})
	fmt.Println(products)
	for i := 0; i < len(searchWord); i++ {
		c := searchWord[i]

		for l <= r && (len(products[l]) <= i || products[l][i] != c) {
			l++
		}
		for l <= r && (len(products[r]) <= i || products[r][i] != c) {
			r--
		}
		remaining := r - l + 1
		ans = append(ans, products[l:l+int(math.Min(float64(3), float64(remaining)))])
	}
	fmt.Println(ans)
	return ans
}
