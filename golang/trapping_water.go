// two pointers, left right, one statck
func trap(height []int) int {
    // to trapping water, height size must bigger than 2
    if len(height) < 2{
        return 0
    }
    res := 0
    left, right := 0, 1
    for ; right <len(height); right++{
        // current value > last: but small then left; push into stack and change until
        if height[right] > height[right-1]{
            tmp:= height[right]
            if tmp > height[left] {
                tmp = height[left]
            }
            for t:=right-1; t>left; t--{
                for tmp > height[t]{
                    res += tmp - height[t]
                    height[t] = tmp
                }
            }
            if height[right] > height[left]{
                left= right
            }
        }
    }
    return res
}