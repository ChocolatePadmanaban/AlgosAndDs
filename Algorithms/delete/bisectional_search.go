package bisectional_search

func bisect_search_helper(L []int, e, low, high int) bool {
	if low == high {
		return L[low] == e
	}
	mid := (high + low) / 2
	if L[mid] == e {
		return true
	} else if L[mid] > e {
		if mid == low {
			return false
		} else {
			return bisect_search_helper(L, e, low, mid-1)
		}
	} else {
		return bisect_search_helper(L, e, mid+1, high)

	}

}

func bisectional_search(L []int, e int) bool {
	if len(L) == 0 {
		return false
	} else {
		return bisect_search_helper(L, e, 0, len(L)-1)
	}
}
