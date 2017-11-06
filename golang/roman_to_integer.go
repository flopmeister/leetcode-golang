
//解体思路，从左往右，
// 首先罗马数字符号： I:1, V:5, X:10, L:50, C:100, D:500, M:1000
// 注意： 在罗马数字中，4，9，40，90，400，900 的表现行为 是：IV,IX,XL,XC,CD,CM
func romanToInt(s string) int {
    sum := 0
    for i:=0; i<len(s) -1; i++{
        if roi(s[i]) < roi(s[i+1]){
            sum -= roi(s[i])
        }else{
            sum += roi(s[i])
        }
    }
    return sum + roi(s[len(s)-1])
}

func roi(char byte)int{
    switch char{
        case 'M':
            return 1000
        case 'D':
            return 500
        case 'C':
            return 100
        case 'L':
            return 50
        case 'X':
            return 10
        case 'V':
            return 5
        case 'I':
            return 1
        default: 
            return 0
    }
}