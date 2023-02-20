def insertionSort(Alist):
    if len(Alist)>1:
        for i in range(1,len(Alist)):
            tempVar=Alist[i]
            j=i-1
            while j>=0 and Alist[j]>tempVar:
                Alist[j+1]=Alist[j]
                j-=1
            Alist[j+1]=tempVar


if __name__== '__main__':
    Alist=[9,9,99,8,7,6,5,4,3,2,1,9,9,9,9]
    print(Alist)
    insertionSort(Alist)
    print(Alist)