// 关键是只能o(1)
func firstMissingPositive(nums []int) int {
    size := len(nums)
    for i,_ := range nums{
        for nums[i] > 0 && nums[i] < size && nums[nums[i]-1]!= nums[i]{
            nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
        }
    }
    for i, val := range nums{
        if val != i+1{
            return i+1
        }
    }
    return size+1
}


//方法2
func firstMissingPositive(nums []int) int {
    ref := make([]int,len(nums)+1)
    for i:=0; i<len(nums);i++{
        if nums[i] >0 && nums[i] <= len(nums){
            ref[nums[i]] = 1
        }
    }
    fmt.Println(ref)
    for j:=1; j<len(ref);j++{
        if ref[j] == 0{
            return j
        }
    }
    return len(ref)
}


