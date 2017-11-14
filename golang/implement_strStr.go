// 相当于implement golang strings.Index(str, substr)
// 最蠢的实现方法
func strStr(haystack string, needle string) int {
    if needle == "" {
        return 0
    }
    if len(needle) > len(haystack) {
        return -1
    }
    
    for i:= 0; i< len(haystack);i++{
        j:=0
        for ;; j++{
            if j == len(needle) {
                return i
            }
            if i+j == len(haystack) {
                return -1
            }
            if haystack[i+j] != needle[j]{
                break
            }
        }
    }
    return -1
}
// 方法2 kmp
