// Solution I : backtrack
func wordBreak(s string, wordDict []string) []string {
    size := len(s)
    preIdxs := make([][]int, size+1)
    preIdxs[0] = []int{0}
    dict := make(map[string]bool)
    for _, word:=range wordDict{
        dict[word]=true
    }
    for i:=1;i<=size;i++{
        for j:=0;j<i;j++{
            if len(preIdxs[j])>0 && dict[s[j:i]]{
                preIdxs[i] = append(preIdxs[i], j)
            }
        }
    }

    return genSolutions(s, size, preIdxs)
}

func genSolutions(s string, start int, preIdxs [][]int) []string{
    var res []string
    if start == 0{
        res = append(res, "")
        return res
    }
    
    for _, idx:= range preIdxs[start]{
        sublist := solution(s, idx, preIdxs)
        for _, item := range sublist{
            if item == ""{
                res = append(res, s[idx:start])
            }else{
                res = append(res, item + " " + s[idx:start])
            }
        }
    }
    return res
}

// Solution II : LTE using DFS and memory
// func wordBreak(s string, wordDict []string) []string {
//     mem := make(map[string][]string)
//     return DFS(s, wordDict, mem)
// }

// func DFS(s string, dict []string, mem map[string][]string) []string{
//     if mem[s] != nil {
//         return mem[s]
//     }
//     var res []string    
//     if s == ""{
//         res = append(res, "")
//         return res
//     }
    
//     for _,word:= range dict{
//         if strings.HasPrefix(s, word) {
//             sublist := DFS(s[len(word):], dict, mem)
//             for _, item := range sublist{
//                 if item  == ""{
//                     res = append(res, word)
//                 }else {
//                     res = append(res, word + " " + item)
//                 }
//             }
//         }
//     }
//     mem[s] = res
//     return res
// }
 
// solution III Forward tracking
// func wordBreak(s string, wordDict []string) []string {
//     size := len(s)
//     list := make([][]string, size+1)
//     list[0] = []string{" "}
//     dict := make(map[string]bool)
//     for _, word:=range wordDict{
//         dict[word]=true
//     }
//     for i:=1;i<=size;i++{
//         for j:=0;j<i;j++{
//             if len(list[j]) >0 && dict(s[j:i]){
//                 if j == 0{
//                     list[i] = append(list[i], s[j:i])
//                 }else{
//                     for _, item := range list[j]{
//                         list[i] = append(list[i], item+" "+s[j:i])
//                     }
//                 }
//             }
//         }
//     }
//     return list[size]
// }
