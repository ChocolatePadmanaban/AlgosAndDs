def kuberCommand(ReadFile, WriteFile):
    readFile=open(ReadFile, 'r')
    writeFile=open(WriteFile,'w')

    # for line in readFile.readlines():
    #     #writeFile.write("kubectl get pods "+str(line).rstrip()+" -n sock-shop -o jsonpath='{.spec.containers[*].resources}'"+"\n\n")
    #     #writeFile.write("kubectl get pods "+str(line).rstrip()+" -n sock-shop -o jsonpath='{.spec.containers[*].name}'"+"\n\n")
    #     #writeFile.write("kubectl get crd "+str(line).rstrip()+" --all-namespaces --show-labels \n\n")
    #     writeFile.write("kubectl logs "+str(line).rstrip().split()[0]+" -n mesh7-system > logfilesinit/"+str(line).rstrip().split()[0]+".txt \n\n")
    #     #writeFile.write( "kubectl get pods  "+str(line).rstrip()+" -n mesh7-system -o jsonpath='{.spec.containers[*].resources}'"+"\n\n")

    data = readFile.read()
    data = data.strip().split()

    for i in data:
        writeFile.write("kubectl config delete-context   "+ str(i)+"\n\n")
        print("kubectl config delete-context   "+ str(i))

    readFile.close()
    writeFile.close()

if __name__ == "__main__":
    kuberCommand("getcontext.txt","commands.txt")