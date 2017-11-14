func nextPermutation(nums []int)  { 
    j:= len(nums) -2
    for ;j>=0 && nums[j+1] <= nums[j]; j--{} // find the number break the condition
    if j >= 0 {
        for i:=len(nums)-1; i>j;i--{
            if nums[i] > nums[j]{
                nums[i],nums[j] = nums[j], nums[i]
                break
            }
        }
    }
    fmt.Println(nums)
    left := j+1
    right := len(nums)-1
    for left<right{
        nums[left], nums[right] = nums[right], nums[left]
        left++
        right--
    }
}