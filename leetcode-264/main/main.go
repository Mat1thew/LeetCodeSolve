package main

import (
	"LeetCodeSolve/leetcode-264/nthUglyNumber"
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	a := nthUglyNumber.NthUglyNumber(1000)
	fmt.Println(a)
	fmt.Println(time.Now())
}
