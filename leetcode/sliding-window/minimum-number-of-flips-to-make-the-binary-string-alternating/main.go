package minimum_number_of_flips_to_make_the_binary_string_alternating

import (
	"math"
	"strings"
)

// MinFlips Refer https://www.youtube.com/watch?v=MOeuK6gaC2A&list=PLot-Xpze53leOBgcVsJBEGrHPd_7x_koV&index=8
//	Runtime: 43 ms, faster than 37.04% of Go online submissions for Minimum Number of Flips to Make the Binary String Alternating.
//	Memory Usage: 8.4 MB, less than 11.11% of Go online submissions for Minimum Number of Flips to Make the Binary String Alternating.
func MinFlips(s string) int {
	n := len(s)
	s += s
	//start := time.Now()
	sb1 := strings.Builder{}
	sb2 := strings.Builder{}
	for i := 0; i < len(s); i++ {
		if i%2 == 0 {
			sb1.WriteString("0")
			sb2.WriteString("1")
		} else {
			sb1.WriteString("1")
			sb2.WriteString("0")
		}
	}
	s1 := sb1.String()
	s2 := sb2.String()
	//elapsed := time.Since(start)
	//log.Printf("Init 2 strings took %s", elapsed)

	ans1, ans2, ans := 0, 0, math.MaxInt
	//start = time.Now()
	for i := 0; i < len(s); i++ {
		if s[i] != s1[i] {
			ans1++
		}
		if s[i] != s2[i] {
			ans2++
		}
		if i >= n {
			if s[i-n] != s1[i-n] {
				ans1--
			}
			if s[i-n] != s2[i-n] {
				ans2--
			}
		}
		if i >= n-1 {
			ansi := int(math.Min(float64(ans1), float64(ans2)))
			ans = int(math.Min(float64(ans), float64(ansi)))
		}
	}
	//elapsed = time.Since(start)
	//log.Printf("Sliding window took %s", elapsed)
	return ans
}
