func searchRange(nums []int, target int) []int {
	
		res := []int{-1,-1}
		if len(nums) == 0{
			return res
		}
		left, right := 0, len(nums)-1
	
		for left < right {
			mid := (left+right)/2
			if nums[mid] < target{
				left = mid+1
			}else{
				right = mid
			}
		}
		if nums[left]!=target{
			return res
		}else{
			res[0] = left
		}
		
		right = len(nums)-1 // we do not need to reset left to 0 the 2nd time
		for left < right {
			mid := (left+right)/2+1 // need to make mid biased to right
			if nums[mid] > target{
				right = mid-1
			}else{
				left= mid
			}
		}
		res[1] = right
		return res
	}