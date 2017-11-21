package main

func main() {

}

// 题目如果给的不是set， 需要加一个移除duplicates
// 第一反应，sort 然后back track

// reverse the 2D array, then flip [i,j] to [j,i]
// reverse the 2D array, then flip [i,j] to [j,i]

/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */

// 解法1 ： 根据merge interval 思路， 先插入，再merge

// 想法很简单, 找到第一个非空格，开始计数，
// 再持续找下去，直到第一个空格。 退出，返回计数
func lengthOfLastWord(s string) int {
	i, count := len(s)-1, 0
	for ; i >= 0 && s[i] == ' '; i-- {
	}
	for ; i >= 0 && s[i] != ' '; i-- {
		count++
	}
	return count
}
