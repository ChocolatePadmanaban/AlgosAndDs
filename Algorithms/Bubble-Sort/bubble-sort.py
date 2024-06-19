def bubble_sort(L):
    '''
    bubble sort is bringing the lowest item forward
    '''
    swap = False
    while not swap:
        swap = True
        for j in range(1,len(L)):
            if L[j-1] > L[j]:
                swap = False
                temp = L[j]
                L[j]= L[j-1]
                L[j-1] = temp
    return L

if __name__ == "__main__":
    print(bubble_sort([5,7,7,8,4,5,56,9,1,2,3,4,6,6,6,6,]))
