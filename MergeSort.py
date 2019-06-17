def mergeSort(alist):
    if len(alist)>1:
        midcount=len(alist)//2
        LeftA=alist[:midcount]
        RightA=alist[midcount:]

        mergeSort(LeftA)
        mergeSort(RightA)

        i=j=k=0

        while i<len(LeftA) and j<len(RightA):
            if LeftA[i]<RightA[j]:
                alist[k]=LeftA[i]
                i+=1
            else:
                alist[k]=RightA[j]
                j+=1
            k+=1
        
        while i<len(LeftA):
            alist[k]=LeftA[i]
            i+=1
            k+=1
        while j<len(RightA):
            alist[k]=RightA[j]
            j+=1
            k+=1

if __name__== '__main__':
    Alist=[9,9,99,8,7,6,5,4,3,2,1,9,9,9,9]
    print(Alist)
    mergeSort(Alist)
    print(Alist)