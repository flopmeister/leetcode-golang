// 满足 cotinuous， increasing
// nums[i] > nums[i-1]

func findLengthOfLCIS(nums []int) int {
    if len(nums) == 0{
        return 0
    }
    current, max:=1,1
    for i:=1; i< len(nums); i++{
        if nums[i] > nums[i-1]{
            current ++
            if current > max{
                max = current
            }
        }else{
            current = 1
        }
    }
    return max
}