package maximize_the_confusion_of_an_exam

import "math"

func MaxConsecutiveAnswers(answerKey string, k int) int {
	l, r, max := 0, 0, 0
	freqWindow := make(map[string]int)
	for r < len(answerKey) {
		cRight := string(answerKey[r])
		if _, ok := freqWindow[cRight]; ok {
			freqWindow[cRight]++
		} else {
			freqWindow[cRight] = 1
		}

		if r-l+1-maxOf(freqWindow) <= k {
			max = int(math.Max(float64(max), float64(r-l+1)))
			r++
		} else {
			cLeft := string(answerKey[l])
			freqWindow[cLeft]--
			r++
			l++
		}
	}
	return max
}

func maxOf(m map[string]int) int {
	var res int
	for _, v := range m {
		if res < v {
			res = v
		}
	}

	return res
}
