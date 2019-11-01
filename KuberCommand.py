def kuberCommand(ReadFile, WriteFile):
    readFile=open(ReadFile, 'r')
    writeFile=open(WriteFile,'w')

    for line in readFile.readlines():
        writeFile.write("kubectl get pods "+str(line).rstrip()+" -n sock-shop -o jsonpath='{.spec.containers[*].resources}'"+"\n\n")
        #writeFile.write("kubectl get crd "+str(line).rstrip()+" --all-namespaces --show-labels \n\n")
        #writeFile.write("kubectl logs "+str(line).rstrip()+" -n kavach-system \n\n")
        #writeFile.write( "kubectl get pods  "+str(line).rstrip()+" -n mesh7-system -o jsonpath='{.spec.containers[*].resources}'"+"\n\n")    
    readFile.close()
    writeFile.close()

if __name__ == "__main__":
    kuberCommand("pods.txt","commands.txt")