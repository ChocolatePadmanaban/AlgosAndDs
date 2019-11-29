import time
import threading



def Hello100(theword):
    for i in range(10):
        print("Hello world by", theword,time.time())
        time.sleep(.01)


t = time.time()
people = ['Newton', 'Maxwell', 'Gauss', 'Pythogarus', 'Leibniz', 'Reimann','Taylor', 'Fibonnaci', 'Napier', 'Bernoulli']
threads=[]
for i in people:
    thread=threading.Thread(target=Hello100,args=i[0])
    threads.append(thread)
    thread.start()

print("done in : ",time.time()-t)

