---
layout: default
---

```
git clone https://github.com/ChocolatePadmanaban/AlgosAndDs.git
export ALGOS_DS_REPO=`pwd`/AlgosAndDs
cd $ALGOS_DS_REPO/Algorithms/Bisectional-Search
```

# Bisectional Search Python

- Lets say you have a Sorted List you have to find whether the an element is present it the list
- If you do do a Linear search (for or while loop) it would take `n` times for the algorithm to find the answer
- But if you do Bisectional Search it would take only `log(n) to the base 2` times to find the value
- why `log(n) to the base 2` ? every time we are dividing the desion list by 2 so 
```
2 ^ (no of steps ) =  Length of the Sorted Array

= > no of steps = log(n) to the base 2

where n = Length of the Sorted Array
```
- note : if we get no of sets as float value then we have to take next integers
```
example:
>>> math.log(10000,2)
13.28771237954945

so if we have a list of length 10000 then it would take 14(at most) to check if the value is present in the list
where as in normal linear check it would take 10000(at most) and 5000(mean) steps to check if its present in the list
```


## Python Code 

```
def bisect_search(L,e):
    def bisect_search_helper(L,e,low,high, steps=0):
        steps+=1
        if high == low:
            return L[low] == e, steps # now high and low vale are same then , either we found the value or the value is not in the list
        mid = (high+low)//2 # get the mid value
        print("This is step ", steps, "high mid and low values are", L[high],L[mid],L[low]) # this step is for us to under stand how bi-sectional search is proceeding
        if L[mid]==e:
            return True, steps # if the mid value is the expected value then it should stop and return the number of steps
        elif L[mid] > e: 
            if low == mid: # there is nothing Left to search in the List 
                return False, steps 
            else:
                return bisect_search_helper(L,e,low, mid-1,steps) # the guessed value is in lower half of the list search there 
        else:
            return bisect_search_helper(L,e,mid+1,high,steps) # the guessed value is in higher half of the list search there 
    if len(L)==0:
        return False , 0
    else:
        return bisect_search_helper(L, e, 0, len(L) - 1)


if __name__ == "__main__":
    print(bisect_search([i for i in range(10000,20001)], 15001))

```
**Explaination:**
- we are searching for 15001 in a list of [10000 to 20000] it took only 13 steps to find the answer
- where as in Linear search it would have taken 5001 steps to find the answer


**Output:**
```
$ python bisectional-search.py 
This is step  1 high mid and low values are 20000 15000 10000
This is step  2 high mid and low values are 20000 17500 15001
This is step  3 high mid and low values are 17499 16250 15001
This is step  4 high mid and low values are 16249 15625 15001
This is step  5 high mid and low values are 15624 15312 15001
This is step  6 high mid and low values are 15311 15156 15001
This is step  7 high mid and low values are 15155 15078 15001
This is step  8 high mid and low values are 15077 15039 15001
This is step  9 high mid and low values are 15038 15019 15001
This is step  10 high mid and low values are 15018 15009 15001
This is step  11 high mid and low values are 15008 15004 15001
This is step  12 high mid and low values are 15003 15002 15001
(True, 13)
```

### Negative Case:

- Suppose if The Searchign value is not in the list 
- The Algorithm should exit gracefully, stating the value not present

**Code:**
```
print(bisect_search([i for i in range(10000,20001)], 35001))
```

**Output:**
```
$ python bisectional-search.py 
This is step  1 high mid and low values are 20000 15000 10000
This is step  2 high mid and low values are 20000 17500 15001
This is step  3 high mid and low values are 20000 18750 17501
This is step  4 high mid and low values are 20000 19375 18751
This is step  5 high mid and low values are 20000 19688 19376
This is step  6 high mid and low values are 20000 19844 19689
This is step  7 high mid and low values are 20000 19922 19845
This is step  8 high mid and low values are 20000 19961 19923
This is step  9 high mid and low values are 20000 19981 19962
This is step  10 high mid and low values are 20000 19991 19982
This is step  11 high mid and low values are 20000 19996 19992
This is step  12 high mid and low values are 20000 19998 19997
This is step  13 high mid and low values are 20000 19999 19999
(False, 14)
```

**Explaination:**
- It took only 14 steps to find that the value is not in the list
- the Algorthim exited gracefully, not in failure