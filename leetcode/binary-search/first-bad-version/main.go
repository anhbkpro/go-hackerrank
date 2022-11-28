package first_bad_version

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

// https://leetcode.com/tag/binary-search/discuss/786126/Python-Powerful-Ultimate-Binary-Search-Template.-Solved-many-problems
//Runtime: 0 ms, faster than 100.00% of Go online submissions for First Bad Version.
//Memory Usage: 1.9 MB, less than 87.29% of Go online submissions for First Bad Version.
func firstBadVersion(n int) int {
	l, r := 0, n
	for l < r {
		m := (l + r) >> 1
		if isBadVersion(m) {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}
