// dp： 定义状态 dp[i] 表示第i个ugly number
// 到达下一个状态有三个transition可以选择
// 建议：此题解体思路，多写几个数，找状态比较。
func nthUglyNumber(num int) int {
    m:= 0
    n:= 0
    k:= 0
    
    A:=make([]int, num)
    A[0] = 1
    for i:=1; i <num; i++{
        A[i] := min(2*A[m], 3*A[n],5*A[k])
        if 2*A[m] == A[i] {
            m++
        }
        if 3*A[n] == A[i] {
            n++
        }
        if 5*A[k] == A[i] {
            k++
        }
    }
    return A[num-1]
}

func min(a,b,c int) int{
    min:= a
    if b < a{
        min = b
    }
    if c < min{
        min = c
    }
    return min
}