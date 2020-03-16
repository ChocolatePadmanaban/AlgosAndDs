import time

def isPrime(anumber):
    if anumber > 2:        
        for i in range(2,anumber):
            if anumber%i==0:                
                return False
        return True
    elif anumber<2:        
        return False
    else:        
        return True

def lessPrimeList(anumber):
    return [  i   for i in range(2,anumber+1)   if isPrime(i) ]

def lessPrimeList1(anumber):
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


if __name__ == "__main__":
    start = time.time()
    nonalgo=lessPrimeList(100003)
    end = time.time()
    print("Time 1  ", end - start)
    start = time.time()
    algo=lessPrimeList1(100003)
    end = time.time()
    print("Time 2  ",end - start)
    print(nonalgo==algo)