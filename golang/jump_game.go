// dp
// dp[i-1] can reach i? if yes, compare dp[i-1] + nums[i]
func canJump(nums []int) bool {
	dp := make([]int, len(nums))
	if len(nums) < 1 {
		return false
	}
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
        if dp[i-1] < i {
            return false
        }
		if dp[i-1] >= i {
			if i+nums[i] > dp[i-1] {
				dp[i] = i + nums[i]
			} else {
				dp[i] = dp[i-1]
			}
		}
	}
    return true
}
