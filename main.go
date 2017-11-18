package main

import "sort"

func main() {

}

// 题目如果给的不是set， 需要加一个移除duplicates
// 第一反应，sort 然后back track

func permuteUnique(nums []int) [][]int {
	// 先 sort， 然后直到
	sort.Ints(nums)
	out := [][]int{}
	buf := []int{}
	solution(nums, &out, &buf)
	return out
}

func solution(nums []int, out *[][]int, buf *[]int) {
	if len(nums) == 0 {
		tmp := make([]int, len(*buf))
		copy(tmp, *buf)
		*out = append(*out, tmp)
	}
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		*buf = append(*buf, nums[i])
		next := make([]int, len(nums))
		copy(next, nums)
		next = append(next[0:i], next[i+1:]...)
		solution(next, out, buf)
		*buf = (*buf)[:len(*buf)-1]
	}
}
