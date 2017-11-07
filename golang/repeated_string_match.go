// e.g. we say A_Final
// in order to find B in A_Final, len(A_Final) >=B
// and we also find, when A_Final <= len(A)+len(B)
func repeatedStringMatch(A string, B string) int {
    m:=1
    aFinal := A
    for len(aFinal) < len(B){
        aFinal += A
        m ++
    }
    if strings.Contains(aFinal, B){
        return m
    }
    if strings.Contains(aFinal+A, B){
        return m+1
    }
    return -1
}