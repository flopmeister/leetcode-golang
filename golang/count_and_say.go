func countAndSay(n int) string {
	if n <= 0 {
		return ""
	}
	nums := "1"
	for i := 1; i < n; i++ {
		var buffer bytes.Buffer
		count := 1
		for j := 0; j < len(nums); j++ {
			if j > 0 {
				if j > 0 && nums[j] == nums[j-1] {
					count++
				} else if nums[j] != nums[j-1] {
					buffer.WriteString(strconv.Itoa(count))
					buffer.WriteByte(nums[j-1])
					count = 1
				}
			}
		}
        buffer.WriteString(strconv.Itoa(count))
        buffer.WriteByte(nums[len(nums)-1])
		nums = buffer.String()
	}
	return nums
}
