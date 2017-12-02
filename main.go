package main

import (
	"strings"
)

func main() {
	a := []int{1, 2, 3, 4, 5}
	longestConsecutive(a)
}

// 个人感觉所有的palindrome问题都跟two pointers 有关

// string tokenization + stack
func simplifyPath(path string) string {
	paths := strings.Split(path, "/")
	stk := []string{}
	for _, s := range paths {
		if s == "." || s == "" {
			continue
		}
		// pop back stk encounter ".."
		if s == ".." && len(stk) > 0 {
			stk = stk[:len(stk)-1]
		} else if s != ".." {
			stk = append(stk, s)
		}
	}
	var res string
	for _, item := range stk {
		res += "/" + item
	}
	if res == "" {
		res = "/"
	}
	return res
}
