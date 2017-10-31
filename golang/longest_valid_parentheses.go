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