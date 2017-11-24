// as long as n is 4n + 1 + 2 + 3
func canWinNim(n int) bool {
    if n % 4 != 0 {
        return true
    }
    return false
}