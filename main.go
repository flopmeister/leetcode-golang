package main

import "sort"

func main() {

}

// 题目如果给的不是set， 需要加一个移除duplicates
// 第一反应，sort 然后back track

// reverse the 2D array, then flip [i,j] to [j,i]
// reverse the 2D array, then flip [i,j] to [j,i]
func groupAnagrams(strs []string) [][]string {
	result := [][]string{}
	for _, val := range strs {
		i := 0
		for ; i < len(result); i++ {
			if isPermutation(val, result[i][0]) {
				result[i] = append(result[i], val)
				break
			}
		}
		if i == len(result) {
			newGroup := []string{val}
			result = append(result, newGroup)
		}
	}
	return result
}

type RuneSlice []rune

func (p RuneSlice) Len() int           { return len(p) }
func (p RuneSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p RuneSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func isPermutation(str1, str2 string) bool {
	// if lengths are not equal - false
	if len(str1) != len(str2) {
		return false
	}

	// sort both strings/runes
	var rune1 RuneSlice = []rune(str1)
	var rune2 RuneSlice = []rune(str2)

	sort.Sort(rune1)
	sort.Sort(rune2)

	//fmt.Println(string(rune1[:]))
	//fmt.Println(string(rune2[:]))

	// compare rune1 and rune 2 by indexes
	for i := 0; i < len(rune1); i++ {
		if rune1[i] != rune2[i] {
			return false
		}
	}

	return true
}
