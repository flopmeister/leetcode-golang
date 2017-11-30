func countSubstrings(s string) int {
	if len(s) == 0 {
		return 0
	}
	count := 0
	for i := 0; i < len(s); i++ {
		count += extendPalindrome(s, i, i)   // calculate the palindromes which center is i
		count += extendPalindrome(s, i, i+1) // calculate the palindromes the center is i and i+1
	}
	return count
}

func extendPalindrome(s string, left, right int) int {
	res := 0
	for left >= 0 && right < len(s) && s[left] == s[right] {
		res++
		left--
		right++
	}
	return res
}

// DP solution
// func countSubstrings(s string) int {
// 	dp := make([]int, len(s))
//     dp[0] = 1
// 	for i := 1; i < len(s); i++ {
// 		dp[i] = dp[i-1] + 1 + sum(s, i)
// 	}
// 	return dp[len(s)-1]
// }

// func sum(s string, pos int) int {
// 	res := 0
// 	for i := 0; i < pos; i++ {
// 		if s[i] == s[pos] && isPalindrom(s, i, pos) {
// 			res++
// 		}
// 	}
// 	return res
// }

// func isPalindrom(s string, start, end int) bool {
// 	l := start
// 	r := end
// 	for l <= r {
// 		if s[l] != s[r] {
// 			return false
// 		}
// 		l++
// 		r--
// 	}
// 	return true
// }
