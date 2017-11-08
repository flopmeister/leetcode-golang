func minPathSum(grid [][]int) int {
    rows := len(grid)
    cols := len(grid[0])
    dp := make([][]int, rows)
    for idx,_:= range dp{
        dp[idx] = make([]int,cols )
    }
    
    dp[0][0] = grid[0][0]
    for i:=1; i< rows; i++{
        dp[i][0] = dp[i-1][0] + grid[i][0]
    }
    for j:=1; j<cols; j++{
        dp[0][j] = dp[0][j-1] + grid[0][j]
    }
    for i:=1; i<rows; i++{
        for j:=1; j<cols; j++{
            dp[i][j] = min(dp[i-1][j],dp[i][j-1]) + grid[i][j]
        }
    }
    return dp[rows-1][cols-1]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}