package main

func main() {

}

// 题目如果给的不是set， 需要加一个移除duplicates
// 第一反应，sort 然后back track

// reverse the 2D array, then flip [i,j] to [j,i]
// reverse the 2D array, then flip [i,j] to [j,i]

func canJump(nums []int) bool {
	dp := make([]int, len(nums))
	if len(nums) < 1 {
		return 0
	}
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] >= i {
			if i+nums[i] > dp[i-1] {
				dp[i] = i + nums[i]
			} else {
				dp[i] = dp[i-1]
			}
		}
	}
}
