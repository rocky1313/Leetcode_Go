package main

import (
	"fmt"
	"testing"
)

type question1 struct {
	para1
	ans1
}
type para1 struct {
	nums   []int
	target int
}
type ans1 struct {
	one []int
}

func Test_TwoSum(t *testing.T) {
	qs := []question1{
		{
			para1{[]int{2, 7, 11, 15}, 9},
			ans1{[]int{1, 2}},
		},
	}
	for _, q := range qs {
		_, p := q.ans1, q.para1
		fmt.Printf("input:%v\n\toutput:%v", p.nums, p.target)
	}
}
