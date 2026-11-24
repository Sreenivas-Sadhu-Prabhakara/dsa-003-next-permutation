# Next Permutation

> Striver SDE Sheet 003 · Day 1 Arrays — find the rightmost ascent, swap it with its ceiling in the suffix, reverse the suffix. Go, stdlib only, tested.

## Problem

Rearrange an array of numbers into the lexicographically next greater permutation. If the array is already the largest possible arrangement (fully descending), wrap around to the smallest (fully ascending). Do it in place.

## Approach

Scan from the right for the first index i where `nums[i] < nums[i+1]` — everything right of i is a descending "ceiling" that cannot grow. To get the *next* permutation, bump position i as little as possible: swap it with the smallest suffix element still greater than it, then reverse the (still descending) suffix to make it the smallest ascending tail. If no ascent exists, the whole array is the maximum — reversing it gives the wrap-around.

**Complexity:** O(n) time, O(1) space.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
