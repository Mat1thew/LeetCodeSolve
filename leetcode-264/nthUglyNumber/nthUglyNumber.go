package nthUglyNumber

import (
	"math"
)

var resultMap = make(map[int]bool, 0)

func isUgly(n int) bool {
	if n < 1 {
		return false
	}
	for n != 1 {
		if _, ok := resultMap[n]; ok {
			return resultMap[n]
		}
		if n%2 == 0 {
			n = n / 2
		} else if n%3 == 0 {
			n = n / 3
		} else if n%5 == 0 {
			n = n / 5
		} else {
			return false
		}
	}
	return true
}

func NthUglyNumber(n int) int {
	resultSlice := make([]int, 0)
	for i := 1; i <= math.MaxInt32; i++ {
		resultMap[i] = isUgly(i)
		if resultMap[i] {
			resultSlice = append(resultSlice, i)
		}
		if len(resultSlice) == n {
			break
		}
	}
	return resultSlice[len(resultSlice)-1]
}
