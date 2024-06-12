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
    print(bisect_search([i for i in range(10000,20001)], 35001))
