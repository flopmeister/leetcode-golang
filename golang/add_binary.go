func addBinary(a string, b string) string {
    if len(b) > len(a) {
        return addBinary(b,a)
    }
    rA := []rune(a)
    rB := []rune(b)
    sA := len(rA)
    sB := len(rB)
    ints:= make([]int, len(a)+1)
    push := 0
    tmp := 0
    i:=0
    for ; i<sA; i++{
        if i<= sB-1{
            tmp = int(rB[sB-i-1] - '0') + int(rA[sA-1-i] - '0') + push
        }else{
            tmp = int(rA[sA-1-i] - '0') + push  
        }
        ints[sA-i] = tmp%2
        push = tmp/2
    }
    if push == 1 {
        ints[0] = '1' - '0'
        return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(ints)), ""), "[]")

    }
    return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(ints[1:])), ""), "[]")
}

func addBinary(a string, b string) string {
    i:=len(a)-1
    j:=len(b)-1
    c:=0
    s:=""
    for i>=0 || j>=0 || c==1 {
        if i>=0 {
            c+=int(a[i]-'0')
            i--
        }
        if j>=0 {
            c+=int(b[j]-'0')
            j--
        }
        s=strconv.Itoa(c%2)+s
        c/=2
    }
    return s
}
