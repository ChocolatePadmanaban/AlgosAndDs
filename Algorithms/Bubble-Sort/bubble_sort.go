package bubble_sort

func bubble_sort(L []int) []int {
	swap := true
	for swap {
		swap = false
		for j := 1; j < len(L); j++ {
			if L[j] < L[j-1] {
				swap = true
				L[j], L[j-1] = L[j-1], L[j]
			}
		}
	}
	return L
}
