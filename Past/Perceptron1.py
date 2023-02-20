def perceptron(featuresl,labelsl,T):
    theta=[-3,3]
    j=theta0=i=0
    theta0=-3
    
    for k in range(T):
        while i<(len(featuresl)):
            if labelsl[i]*(theta[0]*featuresl[i][0]+theta[1]*featuresl[i][1]+theta0)<= 0:
                theta = [theta[0]+labelsl[i]*featuresl[i][0],theta[1]+labelsl[i]*featuresl[i][1]]
                theta0=theta0+labelsl[i]
                j+=1   
            else:
                i+=1
                print(j)
                print(theta)
                print(theta0)
                j=0
        i=0
    

if __name__ == "__main__":
    perceptron([[-4,2],[-2,1],[-1,-1],[2,2],[1,-2]],[1,1,-1,-1,-1],1)


