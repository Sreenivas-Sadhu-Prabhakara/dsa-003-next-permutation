package main

import (
	"reflect"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	cases := []struct {
		in, want []int
	}{
		{[]int{1, 2, 3}, []int{1, 3, 2}},
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{1, 1, 5}, []int{1, 5, 1}},
		{[]int{1, 3, 2}, []int{2, 1, 3}},
		{[]int{5}, []int{5}},
		{[]int{}, []int{}},
	}
	for _, c := range cases {
		in := append([]int{}, c.in...)
		NextPermutation(in)
		if !reflect.DeepEqual(in, c.want) {
			t.Errorf("NextPermutation(%v) = %v, want %v", c.in, in, c.want)
		}
	}
}
