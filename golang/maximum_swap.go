func maximumSwap(num int) int {
    str := strconv.Itoa(num)
    nums := []rune(str)
    dp := make([]int, len(nums))
    maxPos := len(nums)-1
    for i:=len(nums)-1; i>=0; i--{
        if nums[i] > nums[maxPos]{
            maxPos = i
        }
        dp[i] = maxPos
    }
    for j:=0; j<len(nums);j++{
        if nums[j] < nums[dp[j]]{
            nums[j], nums[dp[j]] = nums[dp[j]], nums[j]
            break
        }
    }
    res,_ := strconv.ParseInt(string(nums),10,32)
    return int(res)
}