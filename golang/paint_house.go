func minCost(costs [][]int) int {
	const (
		IntMax = 1<<31 - 1
	)
	dp := make([][]int, len(costs))
	m := len(costs)
	if m == 0 {
		return 0
	}
	n := len(costs[0])
	if n == 0 {
		return 0
	}
	for i := range dp {
		dp[i] = make([]int, len(costs[0]))
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 {
				dp[i][j] = costs[i][j]
			} else {
				min := IntMax
				for x := 0; x < n; x++ {
					if x == j {
						continue
					}
					if dp[i-1][x] < min {
						min = dp[i-1][x]
					}
				}
				dp[i][j] = min + costs[i][j]
			}
		}
	}
	min := IntMax
	for x := 0; x < n; x++ {
		if dp[m-1][x] < min {
			min = dp[m-1][x]
		}
	}
	return min
}