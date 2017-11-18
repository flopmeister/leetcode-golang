package main

import "bytes"

func main() {

}

// 题目如果给的不是set， 需要加一个移除duplicates
// 第一反应，sort 然后back track

func multiply(num1 string, num2 string) string {
	m, n := len(num1), len(num2)
	pos := make([]int, m+n) // maximum possible length

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			mul := int(num1[m]) * int(num2[n])
			p1, p2 := i+j, i+j+1
			sum := mul + pos[p2]
			pos[p1] += sum / 10
			pos[p2] = sum % 10
		}
	}
	var buf bytes.Buffer
	for _, val := range pos {
		if buf.Len() != 0 || val != 0 {
			buf.WriteByte(byte(val))
		}
	}
	return buf.String()
}
