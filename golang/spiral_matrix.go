func spiralOrder(matrix [][]int) []int {
	res := []int{}
	if len(matrix) == 0 {
		return res
	}
	level := 0
	m := len(matrix)
	n := len(matrix[0])
	for m >= 1 && n >= 1 {
		i, j := 0, 0
		// special case, if there is only one line
		if m == 1 {
			for ; j <= n-1; j++ {
				res = append(res, matrix[i+level][j+level])
			}
			break
		}
		// special case, if there is only one column
		if n == 1 {
			for ; i <= m-1; i++ {
				res = append(res, matrix[i+level][j+level])
			}
			break
		}
		// top
		for ; j < n-1; j++ {
			res = append(res, matrix[i+level][j+level])
		}
		// right
		for ; i < m-1; i++ {
			res = append(res, matrix[i+level][j+level])
		}
		// btm
		for ; j > 0; j-- {
			res = append(res, matrix[i+level][j+level])
		}
		// left
		for ; i > 0; i-- {
			res = append(res, matrix[i+level][j+level])
		}
		level++
		m = m - 2
		n = n - 2
	}
	return res
}