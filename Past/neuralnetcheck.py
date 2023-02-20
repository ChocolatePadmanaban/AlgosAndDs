import numpy as np

def Z(w_0,w,x):
    return np.dot(w,x)+w_0

def forapoint(w_0,w,x):
    z=Z(w_0,w,x)
    print("function of z",5*z-2)
    print("RELU",[max(i,0) for i in z])
    print("tanh",np.tanh(z))
    print("z",z)


def foreachpoint(w_0,w):
    for i in [-1,1]:
        for j in [-1,1]:
            print("For the point", i,j,":")
            forapoint(w_0,w,np.array([i,j]))

def CheckLine(x,y,z):
    print((x[1]-y[1])/(z[1]-y[1]) == (x[0]-y[0])/(z[0]-y[0]) )

if __name__ =="__main__":
    foreachpoint(np.array([1,1]),np.array([[1,-1],[-1,1]]))
    CheckLine([0.76159416, 0.76159416], [-0.76159416 , 0.99505475],[ 0.99505475 ,-0.76159416]  )
