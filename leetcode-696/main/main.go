package main

import (
	"fmt"
	"math/bits"
	"strconv"
)

func countBinarySubstrings(s string) int {
	result := 0
	index := 0
	tempCnt := make([]int, 0)
	for index < len(s) {
		cnt := 0
		c := s[index]
		for index < len(s) && s[index] == c {
			cnt++
			index++
		}
		tempCnt = append(tempCnt, cnt)
	}
	for i := 1; i < len(tempCnt); i++ {
		if tempCnt[i] < tempCnt[i-1] {
			result += tempCnt[i]
		} else {
			result += tempCnt[i-1]
		}
	}
	return result
}

func hammingDistance(x int, y int) int {
	xString := strconv.FormatInt(int64(x), 2)
	yString := strconv.FormatInt(int64(y), 2)
	xPointer := len(xString) - 1
	yPointer := len(yString) - 1
	result := 0
	bits.OnesCount(uint(x ^ y))
	for xPointer >= 0 || yPointer >= 0 {
		if xString[xPointer] != yString[yPointer] {
			result++
		}
		xPointer--
		yPointer--
	}
	result = result + xPointer + 1 + yPointer + 1
	return result
}

func main() {
	//a := countBinarySubstrings("1010001")
	//fmt.Println(a)
	//fmt.Println(strconv.FormatInt(4, 2))
	//fmt.Println(strconv.FormatInt(14, 2))

	fmt.Println(hammingDistance(1, 10))

}
