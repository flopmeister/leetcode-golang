// as long as current day's price is bigger than yesterday, we sell
func maxProfit(prices []int) int {
    size := len(prices)
    profit := 0
    if size <= 1{
        return profit
    }
    for i:=1; i<size; i++ {
        if prices[i] > prices[i-1]{
            profit = profit + prices[i] - prices[i-1]
        }
    }
    return profit
}