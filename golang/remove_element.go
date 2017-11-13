func removeElement(nums []int, val int) int {
    res := 0
    for _, num := range nums {
        if num != val{
            nums[res] = num 
            res++
        }
    }
    return res
}