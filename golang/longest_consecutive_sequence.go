func longestConsecutive(nums []int) int {
	set := make(map[int]bool)
	for _, val := range nums {
		if set[val] == false {
			set[val] = true
		}
	}
    max := 0
    for key := range set {
        if set[key-1] == false { // the start of a new streak
            next := key + 1
            for set[next] == true{
                next++
            }
            if max < next-key {
                max = next-key
            }
        }
    }
    return max
}
