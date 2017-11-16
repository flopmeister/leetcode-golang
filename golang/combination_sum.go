func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates) // which is O(nlgn)
	return solution(candidates, target)
}

func solution(candidates []int, target int) [][]int {
	res := [][]int{}
	// basic scenario
	if target == 0 {
        emp:= []int{}
		res = append(res, emp)
	}
	for idx, value := range candidates {
		if value <= target {
            for _, sub := range solution(candidates[idx:], target-value) {
				tmp := []int{value}
				tmp = append(tmp, sub...)
				res = append(res, tmp)
			}
		}
	}
	return res
}
