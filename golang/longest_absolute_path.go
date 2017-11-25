
func lengthLongestPath(input string) int {
    maxLen:= 0
    pathLen:= make(map[int]int)
    for _,s := range strings.Split(input, "\n") {
        name := strings.Trim(s, "\t")
        depth := len(s) - len(name)
        if strings.Contains(s, ".") {
            if pathLen[depth] + len(name) > maxLen {
                maxLen = pathLen[depth] + len(name)
            }
        }else{
            pathLen[depth+1] = pathLen[depth] + len(name) + 1 // +1 for '/'
        }
    }
    return maxLen
}