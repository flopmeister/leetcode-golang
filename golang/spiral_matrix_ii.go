func generateMatrix(n int) [][]int {
	// initialize the 2D array as all 0s
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}
	m := 1
	level := 0
	for n >= 1 {
		i, j := 0, 0
		// base case 1 , if there is only one point left
		if n == 1 {
			for ; j <= n-1; j++ {
			    res[i+level][j+level] = m
			}
			break
		}
		// top
		for ; j < n-1; j++ {
			res[i+level][j+level] = m
			m++
		}
		// right
		for ; i < n-1; i++ {
			res[i+level][j+level] = m
			m++
		}
		// btm
		for ; j > 0; j-- {
			res[i+level][j+level] = m
			m++
		}
		// left
		for ; i > 0; i-- {
			res[i+level][j+level] = m
			m++
		}
		level++
		n = n - 2
	}
	return res
}