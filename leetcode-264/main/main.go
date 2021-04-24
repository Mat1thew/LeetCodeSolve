package main

import (
	"LeetCodeSolve/leetcode-264/nthUglyNumber_three_pointer"
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	//a := nthUglyNumber.NthUglyNumber(8)
	a := nthUglyNumber_three_pointer.NthUglyNumber(1000)
	fmt.Println(a)
	fmt.Println(time.Now())
}
