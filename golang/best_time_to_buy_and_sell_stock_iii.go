func maxProfit(prices []int) int {
    // two transverse to get left best profit and right best profit
    size := len(prices)
    profit := 0
    if size <= 1{
        return 0;
    }
    
    left := make([]int, size)
    right := make([]int, size)
    
    min:= prices[0]
    for i:= 1; i<size;i++{
        if prices[i] < min {
            min = prices[i]
        }
        if prices[i] - min > left[i-1]{
            left[i] = prices[i] - min
        }else{
            left[i] = left[i-1]
        }
    }
    
    max:=prices[size-1]
    for j:=size-2; j>=0; j --{
        if prices[j] >max{
            max = prices[j]
        }
        if max - prices[j] > right[j+1]{
            right[j] = max - prices[j]
        }else{
            right[j] = right[j+1]
        }
    }
    
    for m:=1; m <size; m++{
        if left[m] + right[m] > profit {
            profit = left[m] + right[m]
        }
    }
    return profit
}