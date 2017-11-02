import "math"
func maxProfit(k int, prices []int) int {
    size := len(prices)
    if size < 2 {
        return 0
    }
    if k <= 0 {
        return 0
    }
    if k >= size/2{
        profit := 0
        for i:=1; i<size; i++{
            profit = profit + max(prices[i]-prices[i-1],0)
        }
        return profit
    }
    
    buy := make([]int, k)
    for i, _ := range buy {
        buy[i] = math.MinInt32
    }
    sell:= make([]int, k)

    for i:=0; i<size; i ++{
        buy[0] = max(buy[0], 0-prices[i])
        sell[0] = max(sell[0], buy[0] + prices[i])
        for j:=1; j<k; j ++{
            buy[j] = max(buy[j], sell[j-1] - prices[i])
            sell[j] = max(sell[j], buy[j] + prices[i])
        }
    }
    return sell[k-1]
}

func max(a, b int) int{
    if a < b {
        return b
    }
    return a
}