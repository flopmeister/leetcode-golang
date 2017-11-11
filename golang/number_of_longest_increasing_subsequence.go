// 经典dp 入门题
// 来个cache dp[i], 记录两个值，到当前index 的LIS 长度，以及有多少种方式， default value is {1,1}
// 所以 dp[i] = max(dp[i-1] for j in [0 to i-1]+1)
type node struct{
    size int // record the LIS length to current idx
    count int // record how many ways can meet the size at current idx
}
func findNumberOfLIS(nums []int) int {
    if len(nums) == 0{
        return 0
    }
    dp := make([]node, len(nums))
    max:=1
    for i:=0; i<len(nums);i++{
        dp[i] = node{1,1}
        for j:=0;j<i;j++{
            if nums[i] > nums[j]{
                if dp[i].size < dp[j].size+1{
                    dp[i].size = dp[j].size+1
                    dp[i].count = dp[j].count
                }else if dp[i].size == dp[j].size+1{
                    dp[i].count += dp[j].count
                }
            }

        }
        if dp[i].size > max{
            max = dp[i].size
        }
    }
    sum:=0
    for _, item:= range dp{
        if item.size == max{
            sum+= item.count
        }
    }
    return sum
}


