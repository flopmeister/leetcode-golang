// using hash
func singleNumber(nums []int) int {
    res := 0 
    for _,val := range nums{
        res^=val
    }
    return res
}