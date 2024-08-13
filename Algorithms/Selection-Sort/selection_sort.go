package selection_sort

func selection_sort(L []int) []int {
	suffixSt := 0
	for suffixSt < len(L) {
		for i := suffixSt; i < len(L); i++ {
			if L[i] < L[suffixSt] {
				L[i], L[suffixSt] = L[suffixSt], L[i]
			}
		}
		suffixSt += 1
	}
	return L
}
