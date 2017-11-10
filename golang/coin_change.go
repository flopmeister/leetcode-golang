func coinChange(coins []int, amount int) int {
	
		dp := make([]int, amount+1)
	
		dp[0] = 0
	
		for j:=1;j<=amount;j++{
			dp[j] = amount +1
			for i:=0; i<len(coins);i++{
				if coins[i] <= j{
					if dp[j-coins[i]] < dp[j]{
						dp[j] = dp[j-coins[i]] + 1
					}
				}
			}
		}
		if dp[amount] > amount{
			return -1
		}
		return dp[amount]
	}
	
	