type NumArray struct {
    Sums []int
}

func Constructor(nums []int) NumArray {
    sums := make([]int, len(nums) + 1)
    for i, _ := range nums {
        sums[i+1] = nums[i] + sums[i]
    }	
    return NumArray{sums}
}

func (this *NumArray) SumRange(i int, j int) int {
    return this.Sums[j+1] - this.Sums[i]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */