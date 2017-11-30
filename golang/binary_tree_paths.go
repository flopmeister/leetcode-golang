/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 标准都dfs

// non pointer 版本
func binaryTreePaths(root *TreeNode)[]string{
    return solution("", root)
}
func solution(path string, root *TreeNode) []string{
    if root == nil {
        return []string{}
    }
    if path == ""{
        path = strconv.Itoa(root.Val)
    }else{
        path = path + "->" + strconv.Itoa(root.Val)
    }
    if root.Left == nil && root.Right == nil {
        return []string{path}
    }

    l := solution(path, root.Left)
    r := solution(path, root.Right)
    return append(l, r...)
}
// pointer 版本
// func binaryTreePaths(root *TreeNode) []string {
//     res := []string{}
//     if root == nil {
//         return res
//     }
// 	var cur = strconv.Itoa(root.Val)
// 	solution(&res, &cur, root)
// 	return res
// }

// func solution(res *[]string, cur *string, root *TreeNode) {
// 	// reach the left
// 	if root.Left == nil && root.Right == nil {
// 		// we need to copy current path
// 		tmp := cur
// 		*res = append(*res, *tmp)
// 	}
//     if root.Left != nil {
// 		*cur = *cur + "->" + strconv.Itoa(root.Left.Val)
// 		solution(res, cur, root.Left)
// 		*cur = strings.TrimSuffix(*cur, "->"+strconv.Itoa(root.Left.Val))
// 	}
//     if root.Right != nil {
// 		*cur = *cur + "->" + strconv.Itoa(root.Right.Val)
// 		solution(res, cur, root.Right)
// 		*cur = strings.TrimSuffix(*cur, "->"+strconv.Itoa(root.Right.Val))
// 	}
// }
