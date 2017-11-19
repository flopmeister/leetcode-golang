func judgeCircle(moves string) bool {
    runes := []rune(moves)
    i, j:=0,0
    for _, val := range runes{
        switch val{
            case 'U':
                j++
            case 'R':
                i++
            case 'L':
                i--
            case 'D':
                j--
        }
    }
    if i == 0 && j == 0 {
        return true
    }
    return false
}