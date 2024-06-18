def bisect_search(L,e):
    def bisect_search_helper(L,e,low,high, steps=0):
        steps+=1
        if high == low:
            return L[low] == e, steps # now high and low vale are same then , either we found the value or the value is not in the list
        mid = (high+low)//2 # get the mid value

        # this step is for us to under stand how bi-sectional search is proceeding uncomment the below for better understanding
        # print("This is step ", steps, "high mid and low values are", L[high],L[mid],L[low]) 
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
    print("This Following program will execute bisectional search on a sorted list or array: ")
    print("Output format of the function contains : (bool - weather the value is present or not, step-count - how much steps it took to find the output)")
    print("Test Case 1 : List- [10000-20000] search for 15001",bisect_search([i for i in range(10000,20001)], 15001) )
    print("Test Case 2 : List- [10000, 10002,10004 ....-30000] for 15002", bisect_search([i for i in range(10000,30001,2)], 15002))
    print("Test Case 3 : List- [10000, 10002,10004 ....-30000] for 15001", bisect_search([i for i in range(10000,30001,2)], 15001))
    print("Note: Linear search would have taken (5000-10000) steps to find these answers but Bisectional search finds in 14 steps max")