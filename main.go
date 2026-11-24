package main

import "fmt"

// NextPermutation rearranges nums into the lexicographically next permutation,
// wrapping around to the smallest (sorted) order when nums is already the largest.
func NextPermutation(nums []int) {
	n := len(nums)
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i >= 0 {
		j := n - 1
		for nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	for l, r := i+1, n-1; l < r; l, r = l+1, r-1 {
		nums[l], nums[r] = nums[r], nums[l]
	}
}

func main() {
	nums := []int{1, 2, 3}
	NextPermutation(nums)
	fmt.Println(nums)
}
