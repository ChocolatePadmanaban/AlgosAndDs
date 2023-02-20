def TheDigitor(anumdigit):
    intdigit= int(anumdigit)
    intdigit+=1
    intdigit=str(intdigit)
    if len(anumdigit)>len(intdigit):
        intdigit = (len(anumdigit)-len(intdigit))*"0"+ intdigit
        return intdigit
    else:
        return intdigit
    

def increment_string(astring):
    checkFor =0
    for i in range(len(astring)):
        if astring[i].isdigit():
            checkFor=i
            break
    if checkFor==0:
        return astring + "1"
    else:
        return astring.replace(astring[checkFor:], TheDigitor(astring[checkFor:]))



if __name__ == '__main__':
    print(increment_string("foo")== "foo1",increment_string("foo"))
    print(increment_string("foobar001")== "foobar002",increment_string("foobar001"))
    print(increment_string("foobar1")== "foobar2",increment_string("foobar1"))
    print(increment_string("foobar00")== "foobar01",increment_string("foobar00"))
    print(increment_string("foobar99")== "foobar100",increment_string("foobar99"))
    print(increment_string("foobar099")== "foobar100",increment_string("foobar099"))
    print(increment_string("")== "1",increment_string(""))
    print(increment_string("1"))
