def decipher(N,L,MultiPrime):
    MultiPrime=[int(i) for i in MultiPrime]
    decipherPrime=[]
    for i in range(len(MultiPrime)):
        if i==0:
            for j in range(2,N):
                if MultiPrime[0]%j==0:
                    decipherPrime.append([j,MultiPrime[0]/j])
                    break
        else:
            for j in decipherPrime[i-1]:
                if  MultiPrime[i]%j==0:
                    decipherPrime.append([j,MultiPrime[i]/j])
                    break
    DecipherPrime=[]
    for i in decipherPrime:
        DecipherPrime=DecipherPrime+i
    if DecipherPrime[0]==DecipherPrime[2]:
        DecipherPrime[0],DecipherPrime[1]=DecipherPrime[1],DecipherPrime[0]
    decipherPrime1=[]
    for i in range(len(DecipherPrime)):
        if (i < len(DecipherPrime)-1) and (DecipherPrime[i]==DecipherPrime[i+1]):
            pass
        else:
            decipherPrime1.append(DecipherPrime[i])
    DecipherPrime2= list(dict.fromkeys(DecipherPrime))
    DecipherPrime2.sort()
    alphabets=list('abcdefghijklmnopqrstuvwxyz'.upper())
    PairAlpha=dict(zip(DecipherPrime2,alphabets))
    return ''.join([PairAlpha[i] for i in decipherPrime1])

def decipherList(ciphertext):
    for i in range(len(ciphertext)):
        A = decipher(ciphertext[i][0],ciphertext[i][1],ciphertext[i][2])
        print("Case #"+str(i+1)+': '+A)



numberOfTestCase=int(input())
ciphertext=[]
for i in range(numberOfTestCase):
    N,L=input().split(' ')
    MultiPrime=input().split(' ')
    ciphertext.append([int(N),int(L), MultiPrime])
decipherList(ciphertext)