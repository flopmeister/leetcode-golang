package main

import "sort"

func main() {

}

// 题目如果给的不是set， 需要加一个移除duplicates
// 第一反应，sort 然后back track

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return solution(candidates, target, 0)
}

func solution(candidates []int, target, start int) [][]int {
	res := [][]int{}
	// base scenario, target == 0, recursion is done, can quit
	if target == 0 {
		emp := []int{}
		res = append(res, emp)
	}
	for ; start < len(candidates); start++ {
		value := candidates[start]
		if value <= target {
			for _, val := range solution(candidates, target-value, start+1) {
				tmp := []int{value}
				tmp = append(tmp, val...)
				res = append(res, tmp)
			}
		}
	}
	return res
}
