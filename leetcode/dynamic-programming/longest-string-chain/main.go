package longest_string_chain

import (
	"math"
	"sort"
)

func LongestStrChain(words []string) int {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	dp := make(map[string]int)
	res := 0
	for _, w := range words {
		best := 0
		for i := 0; i < len(w); i++ {
			pre := w[0:i] + w[i+1:]
			best = int(math.Max(float64(best), float64(dp[pre]+1)))
		}
		dp[w] = best
		res = int(math.Max(float64(res), float64(best)))
	}
	return res
}
