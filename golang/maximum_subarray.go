// Thoughts:
// Convert this problem to what is the maximum subarray ending at i?
// To find the the maximum subarray ending at i: dp[i] = nums[i] + dp[i-1]>0?dp[i-1]:0
func maxSubArray(nums []int) int {
    size := len(nums)
    dp := make([]int, size)
    dp[0] = nums[0]
    result := dp[0]
    for i:=1; i<size;i++{
        dp[i] = nums[i] + max(dp[i-1], 0)
        result = max(result,dp[i])
    }
    return result
}

func max(a, b int) int{
    if a > b {
        return a
    }
    return b
}