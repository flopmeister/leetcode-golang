// let's define n(i) the number of distinct ways to reach stairs i
// n(i) = n(i-1) + n(i-2)
func climbStairs(n int) int {
    ways := make([]int, n+1)
    ways[0], ways[1], ways[2] = 0, 1, 2
    for i:= 3; i<=n; i++{
        ways[i]= ways[i-2] + ways[i-1]
    }
    return ways[n]
}