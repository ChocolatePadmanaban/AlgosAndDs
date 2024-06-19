def insertionSort(L):
    suffixSt=0
    while suffixSt != len(L):
        for i in range(suffixSt, len(L)):
            if L[i] < L[suffixSt]:
                L[suffixSt] , L[i]= L[i], L[suffixSt]
        suffixSt += 1
    return L

if __name__ == "__main__":
    print(insertionSort([5,7,7,8,4,5,56,9,1,2,3,4,6,6,6,6,]))