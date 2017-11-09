func hammingDistance(x int, y int) int {
    xor := x^y
    bres := strconv.FormatInt(int64(xor), 2)
    fmt.Println(bres)
    res := 0
    for _,char := range bres{
        if char == '1' {
            res ++
        }
    }
    return res
}