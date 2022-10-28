package maximize_the_confusion_of_an_exam

import "math"

func MaxConsecutiveAnswers(answerKey string, k int) int {
	l, ans, mostFreq := 0, 0, 0
	freqWindow := make(map[string]int)
	for r := 0; r < len(answerKey); r++ {
		cRight := string(answerKey[r])
		freqWindow[cRight]++
		mostFreq = int(math.Max(float64(mostFreq), float64(freqWindow[cRight])))

		for r-l+1-mostFreq > k {
			cLeft := string(answerKey[l])
			freqWindow[cLeft]--
			l++
		}
		ans = int(math.Max(float64(ans), float64(r-l+1)))
	}
	return ans
}
