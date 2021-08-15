package main

import "fmt"

func findChar(char byte, list []byte) bool {
	for _, v := range list {
		if char == v {
			return true
		}
	}
	return false
}

func lengthOfLongestSubstring(s string) int {
	result := 0
	for i := 0; i < len(s); i++ {
		tempSlice := make([]byte, 0)
		tempCnt := 0
		tempSlice = append(tempSlice, s[i])
		tempCnt++
		for j := i + 1; j < len(s); j++ {
			if !findChar(s[j], tempSlice) {
				tempSlice = append(tempSlice, s[j])
				tempCnt++
			} else {
				break
			}
		}
		if tempCnt > result {
			result = tempCnt
		}
	}
	return result
}

func main() {
	a := lengthOfLongestSubstring("abcda")
	fmt.Println(a)
}
