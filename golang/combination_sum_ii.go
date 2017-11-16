import "sort"
func combinationSum2(candidates []int, target int) [][]int {
    sort.Ints(candidates)
    buf := []int{}
    out := [][]int{}
    // remove the duplications from solutions
    // buf， out 是两个block，
    // 传两个address 进去
	solution(candidates, target, 0, &buf, &out)
    return out
}

// DFS
func solution(candidates []int, target, start int, buf *[]int, out *[][]int) {
	// base scenario, target == 0, recursion is done, can quit
	if target == 0 {
        found := make([]int, len(*buf))
        copy(found, *buf)
		*out = append(*out, found)
    }else{
         for i:=start; i < len(candidates); i++ {
            if i> start && candidates[i] == candidates[i-1]{
                continue
            }
            value := candidates[i]
            if value <= target {
                (*buf) = append(*buf, value)  // push value
                solution(candidates, target-value, i+1, buf, out) // DFS search
                (*buf) = (*buf)[:len(*buf) - 1] // pop value
            }else{
                break
            }
         }
    }
}
