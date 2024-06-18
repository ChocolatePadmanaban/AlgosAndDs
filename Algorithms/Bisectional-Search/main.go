package main

import "fmt"

func bisect_search_helper(L []int, e, low, high, steps int) (bool, int) {
	steps += 1
	if high == low {
		return L[low] == e, steps
	}
	mid := (high + low) / 2

	if L[mid] == e {
		return true, steps
	} else if L[mid] > e {
		if low == mid {
			return false, steps
		} else {
			return bisect_search_helper(L, e, low, mid-1, steps)
		}
	} else {
		return bisect_search_helper(L, e, mid+1, high, steps)
	}
}

func bisect_search(L []int, e int) (bool, int) {
	if len(L) == 0 {
		return false, 0
	} else {
		return bisect_search_helper(L, e, 0, len(L)-1, 0)
	}
}

func main() {
	fmt.Println("This Following program will execute bisectional search on a sorted list or array: ")
	fmt.Println("Output format of the function contains : (bool - weather the value is present or not, step-count - how much steps it took to find the output)")
	list10000_20000 := func() []int {
		nums := make([]int, 0, 10001)
		for i := 10000; i <= 20000; i++ {
			nums = append(nums, i)
		}
		return nums
	}()
	fmt.Println(bisect_search(list10000_20000, 15001))
}
