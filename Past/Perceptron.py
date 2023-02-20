def perceptron(featuresl,labelsl):
    theta=[0,0]
    j=theta0=i=0
    #for k in range(T):
    while i<len(featuresl):
        if labelsl[i]*(theta[0]*featuresl[i][0]+theta[1]*featuresl[i][1])<= 0:
            theta = [theta[0]+labelsl[i]*featuresl[i][0],theta[1]+labelsl[i]*featuresl[i][1]]
            j+=1   
        else:
            i+=1
            print(j)
            print(theta)
    

if __name__ == "__main__":
    perceptron([[1,0],[-1,-1],[-1,10]],[-1,1,1])



