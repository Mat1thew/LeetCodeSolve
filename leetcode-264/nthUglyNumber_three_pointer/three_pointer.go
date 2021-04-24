package nthUglyNumber_three_pointer

func getMin(i, j, k int) int {
	result := i
	if result > j {
		result = j
	}
	if result > k {
		result = k
	}
	return result
}

func NthUglyNumber(n int) int {
	resultSlice := make([]int, 0)
	resultSlice = append(resultSlice, 1)
	// three pointer for 2, 3, 5
	i, j, k := 0, 0, 0
	for index := 0; index < n; index++ {
		twoResult := resultSlice[i] * 2
		threeResult := resultSlice[j] * 3
		fiveResult := resultSlice[k] * 5
		minResult := getMin(twoResult, threeResult, fiveResult)
		if minResult == twoResult {
			i++
		}
		if minResult == threeResult {
			j++
		}
		if minResult == fiveResult {
			k++
		}
		resultSlice = append(resultSlice, minResult)
	}
	return resultSlice[n-1]
}
