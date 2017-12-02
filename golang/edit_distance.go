
// DP, 
// 关键是要能理解：
// e.g. 假设我们已知 word[i-2] 到word[j-2]的 转换，
// 到 word[i-1], word[j-1] 只需要考虑
// Replace word1[i - 1] by word2[j - 1] (dp[i][j] = dp[i - 1][j - 1] + 1 (for replacement));
// Delete word1[i - 1] and word1[0..i - 2] = word2[0..j - 1] (dp[i][j] = dp[i - 1][j] + 1 (for deletion));
// Insert word2[j - 1] to word1[0..i - 1] and word1[0..i - 1] + word2[j - 1] = word2[0..j - 1] (dp[i][j] = dp[i][j - 1] + 1 (for // insertion)).
// Ref : ttps://discuss.leetcode.com/topic/17639/20ms-detailed-explained-c-solutions-o-n-space
func minDistance(word1 string, word2 string) int {
    dp := make([][]int, len(word1)+1)
    for i:=range dp {
        dp[i] = make([]int, len(word2)+1)
    }
    
    for i:=1; i<=len(word1);i++{
        dp[i][0] = i
    }
    for j:=1;j<=len(word2);j++{
        dp[0][j] = j
    }
    for i:=1; i<=len(word1);i++{
        for j:=1;j<=len(word2);j++{
            if word1[i-1] == word2[j-1]{
                dp[i][j] = dp[i-1][j-1]
            }else{
                dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1
            }
        }
    }
    return dp[len(word1)][len(word2)]
}
func min(a, b int) int{
    if a <b{
        return a
    }
    return b
}