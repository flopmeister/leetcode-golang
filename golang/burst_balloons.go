// divide and conquer,
// 以last burst 气球，分左右两个sub section
func maxCoins(nums []int) int {
    iNums := []int{}
    // add 1 to the start and end of nums
    iNums = append(iNums, 1)
    iNums = append(iNums, nums...)
    iNums = append(iNums, 1)
    
    size := len(iNums)
    // create a 2D array
    dp := make([][]int, size)
    for i,_:=range dp{
        dp[i] = make([]int,size)
    }
    
    for m:=2; m<size;m++{ // 定义slide windows的长度
        for left:=0; left<size-m;left++{ // 找出按lide window长度，不同的left，right里最大的burst值。
            right := left+m
            for n:=left+1;n<right;n++{
                dp[left][right]=max(dp[left][right], 
                                    iNums[left]*iNums[n]*iNums[right]+dp[left][n]+dp[n][right])
            }
        }
    } 
    return dp[0][size-1]
}

func max(a, b int) int{
    if a < b{
        return b
    }
    return a
}