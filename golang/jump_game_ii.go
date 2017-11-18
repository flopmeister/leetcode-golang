func jump(nums []int) int {
	// BFS
	 if len(nums) < 2 {
		 return 0
	 }
	 currentMax, nextMax, level, i := 0, 0, 0, 0
	 
	 for ; i <len(nums)-1; i++ {
		 level ++
		 for j:=i; j<=currentMax; j++{
			 nextMax = max(nums[j]+j, nextMax)
			 if nextMax >=len(nums) -1{
				 return level
			 }
		 }
		 currentMax = nextMax
	 }
	 return 0
 }
 
 func max(a, b int) int{
	 if a < b {
		 return b
	 }
	 return a
 }