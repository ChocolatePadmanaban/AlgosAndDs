from itertools import permutations


def pictword(word, wordCount):
    with open('words_alpha.txt') as f:
        lines = f.read().splitlines()
    lines = set(lines)
    
    word=[i for i in word]
    word.sort()
    perm= set(permutations(word, wordCount))
    answer=[]
    for i in perm:
        
        word1=''.join(i)
        if word1 in lines:
            answer.append(word1)
            print(word1)
    answer=list(set(answer))
    answer.sort()
    print(answer)
    f.close()

if __name__ == "__main__":
    pictword("gbhotubillwu",9)
