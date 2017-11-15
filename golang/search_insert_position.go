// binary search ,
// if found, return index, 
// if not found, which means left==right, compare the two values, if bigger than nums[left],
// pos = left + 1, else pos --
func searchInsert(nums []int, target int) int {
    // empty array always insert into the 1st position
    if len(nums) == 0{
        return 0
    }
    
    left, right := 0, len(nums)-1
    // binary search, left < right
    for {
        if left > right {
            return left
        }
        mid := (left + right) /2 // mid is always smaller than right
        if nums[mid] == target{
            return mid
        }
        if nums[mid] < target{
            left =  mid + 1
        }else{
            right = mid -1
        }
    }
}