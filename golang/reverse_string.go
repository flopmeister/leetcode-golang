func reverseString(s string) string {
    // golang to reverse a string, must use []rune
    runeSlice := []rune(s)
    size := len(s)
    for i:=0;  i<size/2; i++{
        runeSlice[i], runeSlice[size-i-1] = runeSlice[size-i-1], runeSlice[i]
    }
    return string(runeSlice)
}