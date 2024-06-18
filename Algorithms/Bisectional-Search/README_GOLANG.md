---
layout: default
---

```
git clone https://github.com/ChocolatePadmanaban/AlgosAndDs.git
export ALGOS_DS_REPO=`pwd`/AlgosAndDs
cd $ALGOS_DS_REPO/Algorithms/Bisectional-Search
```

# Bisectional Search golang

- Lets say you have a Sorted List you have to find whether the an element is present it the list
- If you do do a Linear search (for or while loop) it would take `n` times for the algorithm to find the answer
- But if you do Bisectional Search it would take only `log(n) to the base 2` times to find the value
- why `log(n) to the base 2` ? every time we are dividing the desion list by 2 so 


```
2 ^ (no of steps ) =  Length of the Sorted Array

= > no of steps = log(n) to the base 2

where n = Length of the Sorted Array
```


- note : if we get no of sets as float value then we have to take next integers


```
example:
>>> math.log(10000,2)
13.28771237954945

so if we have a list of length 10000 then it would take 14(at most) to check if the value is present in the list
where as in normal linear check it would take 10000(at most) and 5000(mean) steps to check if its present in the list
```


## Golang Code 



```
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

```


**Explaination:**
- we are searching for 15001 in a list of [10000 to 20000] it took only 13 steps to find the answer
- where as in Linear search it would have taken 5001 steps to find the answer


**Output:**


```
$ go run main.go 
This Following program will execute bisectional search on a sorted list or array: 
Output format of the function contains : (bool - weather the value is present or not, step-count - how much steps it took to find the output)
true 13
```
