// solution using a stack
func longestValidParentheses(s string) int {
    runes := []rune(s)
    var stack Stack
    for i:=0; i< len(runes); i++{
        if runes[i] == ')'&& len(stack) > 0 && runes[stack.Top()] == '('{
            // we find a match pair
            stack.Pop()
        }else{
            stack.Push(i)
        }
    }
    // the whole string are matched
    if len(stack) == 0 {
        return len(runes)
    }
    // there is no matched pair
    if len(stack) == len(runes) {
        return 0
    }
    // compare the head and tail interval and get the max
    max := stack[0] - 0
    if max < len(runes) - stack.Top() - 1 {
        max = len(runes) - stack.Top() - 1
    }
    
    for j:=1; j<len(stack); j++{
        if max < stack[j] - stack[j-1] - 1{
            max = stack[j] - stack[j-1] - 1
        }
    }
    return max
}

type Stack []int

func (s *Stack) Top() int {
    return (*s)[len(*s)-1]
}

func (s *Stack) Pop(){
    *s = (*s)[:len(*s)-1]
}

func (s *Stack) Push(item int){
    *s = append(*s, item)
}

// solution using DP
func longestValidParentheses(s string) int {
    runes := []rune(s)
    size := len(runes)
    // in case size small than 2, return 0
    if size < 2 {
        return 0
    }
    // create a array to record the maximum length of parentheses to current position
    // all positions now have 0 initial value
    dp := make([]int, size)
    max:= 0
    for i:=1; i<size; i++{
        if runes[i] == ')'{
            if runes[i-1] == '('{ // a new pair formed
                if i == 1 { // special case
                    dp[i] = 2
                }else{
                    dp[i] = dp[i-2] + 2
                }
            }
            // cause every position whose value is ')' records maximum length of parentheses to current position
            if runes[i-1] ==')' { 
                if i >= dp[i-1] + 1  && runes[i-1-dp[i-1]] =='('{ // new pair formed
                    if i == dp[i-1] + 1 {
                        dp[i] = dp[i-1] + 2
                    }else{
                        dp[i] = dp[i-1] + 2 + dp[i-2-dp[i-1]] 
                    }
                }
            } 
            if max < dp[i] {
                max = dp[i]
            }
        }
    }
    return max
}

// solution using two transversal
func longestValidParentheses(s string) int {
    runes := []rune(s)
    size := len(runes)
    max := 0 
    left := 0
    right := 0
    for i:=0; i< size; i ++{
        if runes[i] == '('{
            left ++
        }
        if runes [i] ==')'{
            right ++
        } 
        if left==right{
            if max < (right)*2 {
                max = (right)*2
            }
        }else if right > left{
            left = 0
            right = 0
        }
    }

    left = 0
    right = 0
    for j:= size -1; j>0; j--{
        if runes[j] == '('{
            left ++
        }
        if runes [j] ==')'{
            right ++
        } 
        if left==right{
            if max < (left)*2 {
                max = (left)*2
            }
        }else if left > right{
            left = 0
            right = 0
        }
    }
    return max
}
