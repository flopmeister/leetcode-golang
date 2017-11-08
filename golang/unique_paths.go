func uniquePaths(m int, n int) int {
    // create a dp to represent unique ways to current point
    dp := make([][]int, m) 
    for i, _ := range dp{
        dp[i] = make([]int, n)
    }
    if m==0 || n == 0 {
        return 0
    }
    dp[0][0] = 1
    for i:=1; i<m;i++{
        dp[i][0] = 1
    }
    for j:=1; j<n; j++{
        dp[0][j] = 1
    }
    
    for i:=1; i<m;i++{
        for j:=1; j<n; j++{
            dp[i][j] = dp[i-1][j] + dp[i][j-1]
        }
    }
    return dp[m-1][n-1]
}