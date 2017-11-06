// 解题思路
// 首先罗马数字符号： I:1, V:5, X:10, L:50, C:100, D:500, M:1000
// 注意： 在罗马数字中，4，9，40，90，400，900 的表现行为 是：IV,IX,XL,XC,CD,CM
// 考虑到 个位，十位，百位，千位，每个位置就9种情况，所以用枚举不失为一个好办法

func intToRoman(num int) string {
    var M = []string{"", "M","MM","MMM"}
    var C = []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
    var X = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
    var I = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
    return M[num/1000] + C[(num%1000)/100] + X[(num%100)/10] + I[(num%10)]
}