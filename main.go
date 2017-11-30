package main

import (
	"strconv"
	"strings"
)

func main() {
	a := []int{1, 2, 3, 4, 5}
	longestConsecutive(a)
}

// 个人感觉所有的palindrome问题都跟two pointers 有关

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	res := []string{}
	var cur = strconv.Itoa(root.Val)
	solution(&res, &cur, root)
	return res
}

func solution(res *[]string, cur *string, root *TreeNode) {
	// reach the left
	if root.Left == nil && root.Right == nil {
		// we need to copy current path
		tmp := *cur
		*res = append(*res, tmp)
	} else if root.Left != nil {
		*cur = *cur + "->" + strconv.Itoa(root.Val)
		solution(res, cur, root.Left)
		*cur = strings.TrimSuffix(*cur, "->"+strconv.Itoa(root.Val))
	} else if root.Right != nil {
		*cur = *cur + "->" + strconv.Itoa(root.Val)
		solution(res, cur, root.Right)
		*cur = strings.TrimSuffix(*cur, "->"+strconv.Itoa(root.Val))
	}

}
