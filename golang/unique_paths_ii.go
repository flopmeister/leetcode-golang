func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    // if current grid is obstacle then ,the value is 0
    if obstacleGrid[0][0] == 1{
        return 0
    }
    m := len(obstacleGrid)
    n := len(obstacleGrid[0])
    dp := make([][]int, m)
    for i:= range dp{
        dp[i] = make([]int, n)
        for j:= range dp[i] {
            dp[i][j] = 1
        }
    }
    for i:=0; i<m; i++ {
        for j:=0; j<n; j++{
            if obstacleGrid[i][j] == 1{
                dp[i][j] = 0
            }else {
                if i==0 && j==0{
                    continue
                }else if i == 0{
                    dp[0][j] = dp[0][j-1]
                }else if j==0{
                   dp[i][0] = dp[i-1][0]  
                }else{
                    dp[i][j] = dp[i-1][j] + dp[i][j-1]
                }
            }
        }
    }
    return dp[m-1][n-1]
}