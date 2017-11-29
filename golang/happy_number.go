func isHappy(n int) bool {
    dict := make(map[int]bool)
    dict[n] = true
    for {
        if n == 1 {
            return true
        }
        n = numSquare(n)
        if dict[n] == true {
            return false
        }
        dict[n] = true
    }
}


func numSquare(n int) (output int){
    for ; n!=0; n/=10{
        d:=n%10
        output += d*d
    }
    return
}
