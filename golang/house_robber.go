func rob(nums []int) int {
    rob:= 0 
    letItGo := 0 
    for _, money := range nums{
        prevLetItGo:=letItGo
        letItGo = max(letItGo, rob)
        rob = prevLetItGo + money
    }
    return max(rob, letItGo)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}