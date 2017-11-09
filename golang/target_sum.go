// 方法1:普通dfs, 2^n
// func findTargetSumWays(nums []int, S int) int {
//     return validSum(nums, S, len(nums))
// }

// func validSum(nums []int, s int, idx int) int{
//     var res = 0 
//     if idx == 0{
//         if s == 0{
//             res = 1
//         }
//     }else{
//         res = validSum(nums, s-nums[idx-1], idx-1) + validSum(nums, s+nums[idx-1], idx-1)
//     }
//     return res
// }

// 方法2： 因为对于数组里的每一个数，只能有正，负两种情况
// 举个例子，
// 数组， 【1，2，3】 全选时 sum = +1 + 2 + 3 =6， 我们不选3 时，其和 为 sum = +1 +2 -3 = 0
// 所以，假设，nums[i]没有被选择，那么 sum = max_sum - 2nums[i]，min_sum 为全
// 这样 (max_sum -S)%2 必须为0. 
// 那么我们将max_sum 与s 的差 除以2 得到的值 为target。
// 问题就转换成，给定target 求出nums里有多少个数字之和可以满足它们的和为target。

func findTargetSumWays(nums []int, s int) int {
    sum := 0
    for _, item := range nums{
        sum += item
    }
    if s > sum || (sum-s)%2==1{
        return 0
    }
      
    target := (sum-s)/2
    sort.Ints(nums)
    dp := make([]int, target + 1)
    dp[0] = 1
    for i:=0; i<len(nums);i++{
        if nums[i] > target{
            break
        }
        for j:=target;j>=nums[i];j--{
            dp[j]+=dp[j-nums[i]]
        }
    }
    return dp[target]
}