package main

func main() {
	a := []int{1, 2, 3, 4, 5}
	longestConsecutive(a)
}

// 个人感觉所有的palindrome问题都跟two pointers 有关

// string tokenization + stack
func setZeroes(matrix [][]int) {
	// top-down
	// down-top
	m := len(matrix)
	n := len(matrix[0])
	col0 := 1
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			col0 = 0
		}
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 1; j-- {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
		if col0 == 0 {
			matrix[i][0] = 0
		}
	}
}
