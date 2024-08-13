def selection_sort(L):
    suffixSt = 0 
    while suffixSt != len(L):
        for i in range(suffixSt, len(L)):
            if L[i] < L[suffixSt]:
                L[i], L[suffixSt] = L[suffixSt],L[i]
        suffixSt += 1
    return L  

if __name__ == "__main__":
    print(selection_sort([i for i in range(10, 0 , -1)]))