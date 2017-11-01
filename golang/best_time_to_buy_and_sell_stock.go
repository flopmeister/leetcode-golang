func maxProfit(prices []int) int {
    // because buying must happen before selling
    days := len(prices)
    max := 0
    if days < 1 {
        return max
    }
    
    highest := prices[days - 1]
    for i:= days -2; i>=0 ; i-- {
        if prices[i] >= highest {
            highest = prices[i]
        }else if highest - prices[i] > max {
            max = highest - prices[i] 
        }
    }
    return max
}