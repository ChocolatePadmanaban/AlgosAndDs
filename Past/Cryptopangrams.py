import time

def lessPrimeList(anumber):
    PrimeList=[2]
    for i in range(3,anumber+1):
        isPrimeB = True
        for Prime in PrimeList:
            if i % Prime == 0:
                isPrimeB = False
                break
        if isPrimeB:
            PrimeList.append(i)                        
    return PrimeList

def CryptoExec(TestcaseNo,N,L,MultiPrime):
    PrimeList=lessPrimeList(N)
    start, nextone=0,0
    for i in PrimeList:
        if MultiPrime[0] % i ==0:
            start=int(i)
            nextone=int(MultiPrime[0]/i)
            break
    breakcount=0
    if MultiPrime[1] % nextone ==0:
        PrimeList=[start,nextone,int(MultiPrime[1]/nextone)]        
    elif MultiPrime[1] % start ==0:
        PrimeList=[nextone,start,int(MultiPrime[1]/start)]
    else:
        breakcount+=1
    for i in MultiPrime[2:]:
        if i%PrimeList[-1]==0:
            PrimeList.append(int(i/PrimeList[-1]))
        else:
            breakcount+=1
    
    
    
    PrimeListSort=PrimeList.copy()
    PrimeListSort.sort()
    PrimeListSort=list(set(PrimeListSort))
    alphabets=list('abcdefghijklmnopqrstuvwxyz'.upper())
    NuericKey = { PrimeListSort[i]:alphabets[i] for i in range(len(PrimeListSort))}
    return "Case #"+str(TestcaseNo)+': '+''.join([NuericKey[i] for i in PrimeList])





if __name__ == "__main__":
    numberOfTestCase=int(input())
    ciphertext=[]
    for i in range(numberOfTestCase):
        N,L=input().split(' ')
        MultiPrime=input().split(' ')
        ciphertext.append(CryptoExec(i,int(N),int(L), [int(j) for j in MultiPrime]))
    for i in ciphertext:
        print(i)