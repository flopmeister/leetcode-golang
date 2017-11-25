package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	longestConsecutive(a)
}

func longestConsecutive(nums []int) int {
	set := make(map[int]bool)
	for _, val := range nums {
		if set[val] == false {
			set[val] = true
		}
		fmt.Println(set)
	}
	return 1
}
