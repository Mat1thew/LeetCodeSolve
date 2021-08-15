package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var stringByte []byte
var sliceIndex int

func dfs(index int, target byte) bool {
	sliceIndex = index
	if index < 0 {
		return false
	}
	if stringByte[index] == target {
		return true
	}
	return dfs(index-1, target)
}

func strangePrinter(s string) int {
	flagChar := byte(0)
	result := 0
	stringByte = []byte(s)
	resultByte := make([]byte, len(s), len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == flagChar {
			continue
		}
		flagChar = s[i]
		resultByte[i] = flagChar
		if dfs(len(s)-1, flagChar) {
			for j := i + 1; j <= sliceIndex; j++ {
				resultByte[j] = flagChar
			}
		}
		result++
		if string(resultByte) == s {
			break
		}
	}
	return result
}

func findMaxForm(strs []string, m int, n int) int {
	length := len(strs)
	dp := make([][][]int, length+1)
	for i := range dp {
		dp[i] = make([][]int, m+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, n+1)
		}
	}
	for i, v := range strs {
		zeros := strings.Count(v, "0")
		ones := len(v) - zeros
		for j := 0; j <= m; j++ {
			for k := 0; k <= n; k++ {
				if j >= zeros && k >= ones {
					if dp[i+1][j][k] < dp[i][j-zeros][k-ones]+1 {
						dp[i+1][j][k] = dp[i][j-zeros][k-ones] + 1
					}

				}
			}
		}
	}
	return dp[length][m][n]
}

func main() {
	result := strangePrinter("abcabcabc")
	fmt.Println(result)
	//dp := make([][]int, 5, 5)
	ccc := "aaabbbccc"
	fmt.Println(strings.Count(ccc, "a"))

	fmt.Println(strconv.FormatInt(10, 2))
	a := []int{2, 3, 4}
	sort.Ints(a)

}
