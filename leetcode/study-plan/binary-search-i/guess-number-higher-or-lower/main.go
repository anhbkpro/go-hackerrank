package guess_number_higher_or_lower

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

// GuessNumber My Solution
// 	Runtime: 0 ms, faster than 100.00% of Go online submissions for Guess Number Higher or Lower.
//	Memory Usage: 1.8 MB, less than 100.00% of Go online submissions for Guess Number Higher or Lower.
func GuessNumber(n int) int {
	lo, hi := 0, n
	for hi-lo > 1 {
		mid := (hi + lo) >> 1
		if guess(mid) > 0 { // assume that the guess is given
			lo = mid
		} else if guess(mid) < 0 {
			hi = mid
		} else {
			return mid
		}
	}
	return hi
}
