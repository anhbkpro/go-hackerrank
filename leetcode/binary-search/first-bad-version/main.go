package first_bad_version

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func FirstBadVersion(n int) int {
	lo, hi := 0, n
	for hi-lo > 1 {
		mid := lo + (hi-lo)/2
		// isBadVersion is given :)
		// an API bool isBadVersion(version) which returns whether version is bad.
		if isBadVersion(mid) {
			hi = mid
		} else {
			lo = mid
		}
	}
	return hi
}
