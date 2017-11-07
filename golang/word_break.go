func wordBreak(s string, wordDict []string) bool {
    size := len(s)
    dp := make([]bool, size+1)
    dp[0] = true
    for i:=1; i<=size;i++{
        for j:=0; j<i;j ++{
            if dp[j] &&inDict(s[j:i], wordDict){
                dp[i]=true
            }
        }
    }
    return dp[size]
}


func inDict(s string, wordDict []string) bool{
    for _, item:= range wordDict{
        if item == s{
            return true
        }
    }
    return false
}