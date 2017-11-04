
func maxProfit(prices []int) int {
    if len(prices) < 2 {
        return 0
    }
    buy, sell, last_buy, last_sell := -prices[0], 0, 0, 0
    
    for _, price := range prices {
        last_buy = buy
        buy = max(last_sell - price, last_buy)
        last_sell = sell
        sell = max(last_buy + price, last_sell)
    }
    return sell
}

func max(a, b int) int {
    if a < b{
        return b
    }
    return a
}