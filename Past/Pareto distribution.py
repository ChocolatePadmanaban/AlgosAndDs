#!/usr/bin/env python
# coding: utf-8

# # I would try to Create a working Pareto Distribution

#              Lets say there are ten people a = ['Newton', 'Maxwell', 'Gauss', 'Pythogarus', 'Leibniz', 'Reimann','Taylor', 'Fibonnaci', 'Napier', 'Bernoulli']
#              they all start with 10 rupees each iteration wat will happen at nth iteration

# In[97]:


import random


# In[98]:


people = ['Newton', 'Maxwell', 'Gauss', 'Pythogarus', 'Leibniz', 'Reimann','Taylor', 'Fibonnaci', 'Napier', 'Bernoulli']


# In[99]:


peoplewithmoney= {person:10  for person in people}
print("Before Iteration................................")
print(peoplewithmoney)


# In[106]:


def Trade(Peoplewithmoney, therepeat=0):
    
    for i in range(len(Peoplewithmoney)-1):
        people=list(Peoplewithmoney.keys())
        person1=random.choice(people)
        people.remove(person1)
        person2=random.choice(people)
        if Peoplewithmoney[person2]>0:
            Peoplewithmoney[person1]+=1
            Peoplewithmoney[person2]-=1
        else:
            del Peoplewithmoney[person2]
            
        
    if therepeat>0:
        return Trade(Peoplewithmoney, therepeat-1)
    else:
        return Peoplewithmoney
        


print("After Iteration................................")
print(Trade(peoplewithmoney, 100))


# In[ ]:




