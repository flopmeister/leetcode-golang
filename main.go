package main

// 思路，刚从ransom note过来，
// 先对target 进行一个按字典排序，得到每个char 有多少的一个set
// 对word里面所有的单词进行字典排序，
// 举例： target = “thehat” -> {'a':1,'e':1,'h':2,'t'=2}
//      ['with','example'] - > [{'i':1,'h':1,'t':1,'w':1}, {'a':1,'e':2,'l':1,'m':1,'p':1,'x':1}]
// 这样就转换成了硬币题，更复杂一些罢了。// dp+backtrack
func minStickers(stickers []string, target string) int {

	// 把sticker 字典化，并且计数。
	words := make([][]int, len(stickers))
	for i, word := range stickers {
		words[i] = dictNize(word)
	}

	//提前判别找不到
	var wordsStr string
	for _, sticker := range stickers {
		wordsStr += sticker
	}
	targetDict := dictNize(target)
	wordsDict := dictNize(wordsStr)
	for i, val := range targetDict {
		if val > 0 && wordsDict[i] == 0 {
			return -1
		}
	}

	return findMin(words, targetDict)
}

func findMin(words [][]int, target []int) int {
	for _, word := range words {
		// 待实现ransom
		if ranSom(target, word) {
			return 1
		}
	}
	res := 1<<31 - 1

	for _, word := range words {
		subTarget, ok := deleteWord(target, word)
		if res < findMin(words, subTarget)+1 {
			res = findMin(words, subTarget) + 1
		}
	}
	return res
}

func ranSom(target, word []int) bool {
	for i, val := range word {
		target[i] -= val
		if target[i] > 0 {
			return false
		}
	}
	return true
}

func deleteWord(target, word []int) []int {
	for i, val := range word {
		target[i] -= val
	}
}

func dictNize(word string) []int {
	res := make([]int, 26)
	for _, char := range word {
		res[char-'a']++
	}
	return res
}
