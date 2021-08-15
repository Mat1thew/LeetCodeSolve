package main

import (
	"sort"
	"strconv"
)

func checkNum(x, y, u, v int) bool {
	if friendsPreferences[x][u] > friendsPreferences[x][y] {
		return false
	}
	if friendsPreferences[u][x] >= friendsPreferences[u][v] {
		return false
	}
	return true
}

var friendsPreferences = [501][501]byte{}

func unhappyFriends(n int, preferences [][]int, pairs [][]int) int {
	result := 0
	// 选用二维数组存储朋友关系友好度
	// 题目中最多为500，预申请长度为500*500的二维数组
	// 下标表示伙伴，数组中存储关系顺序，0为关系最好，数组长度表示关系最差
	friendsPreferences = [501][501]byte{}
	for i := 0; i < len(preferences); i++ {
		for j := 0; j < len(preferences[i]); j++ {
			friend := preferences[i][j]
			friendsPreferences[i][friend] = byte(j)
		}
	}
	isVisit := [501]int{}
	for i := 0; i < len(pairs); i++ {
		for j := 0; j < len(pairs); j++ {
			if i == j {
				continue
			}
			isSad1 := checkNum(pairs[i][0], pairs[i][1], pairs[j][0], pairs[j][1])
			isSad2 := checkNum(pairs[i][0], pairs[i][1], pairs[j][1], pairs[j][0])
			isSad3 := checkNum(pairs[i][1], pairs[i][0], pairs[j][0], pairs[j][1])
			isSad4 := checkNum(pairs[i][1], pairs[i][0], pairs[j][1], pairs[j][0])
			if isSad1 || isSad2 {
				if isVisit[pairs[i][0]] == 0 {
					result++
				}
				isVisit[pairs[i][0]] = 1
			}
			if isSad3 || isSad4 {
				if isVisit[pairs[i][1]] == 0 {
					result++
				}
				isVisit[pairs[i][1]] = 1
			}
		}

	}
	return result
}

func main() {
	strconv.Itoa()
	sort.Ints()
}
