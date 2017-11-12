// 简单的思路就是，拿一个抠掉一个，大概O（n^2)
// 还有一种 就是o（M+N)
// hashMap delete ranSomeNote 就是o1
func canConstruct(ransomNote string, magazine string) bool {
    dict := make(map[rune]int)
    for _, char := range magazine{
        dict[char] ++
    }
    for _, char := range ransomNote{
        if dict[char] == 0{
            return false
        }
        dict[char]-- 
    }
    return true
}