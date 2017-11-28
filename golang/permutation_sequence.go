func getPermutation(n int, k int) string {
    nums, res := make([]byte, n), make([]byte, n)
    for i:=range nums {nums[i] = byte(i) + 1}
    k--
    for i:=1; i<=n;i++{
        index := k/fact(n) // index to add
        k -= index*fact(n) // next n
        res[i-1] = nums[index] + '0'	//这个可以帮助取消最后得special char
        nums = append(nums[0:index], nums[index+1:]...) // remove the number has been used
    }
    return string(res)
}

func fact(n int) int { if n<=1 {return 1}; return n*fact(n-1)}