import sys

def main():
    repoDir = sys.argv[1]
    f = open(repoDir + "/script.txt","w+")
    f.write("Hello World")
    f.close()

if __name__ == "__main__":
    main()