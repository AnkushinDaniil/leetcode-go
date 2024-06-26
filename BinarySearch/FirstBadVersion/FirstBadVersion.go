package firstBadVersion

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

var isBadVersion func(int) bool

func firstBadVersion(n int) int {
	l, r := 0, n
	for l <= r {
		m := (l + r) / 2
		if isBadVersion(m) {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return l
}
