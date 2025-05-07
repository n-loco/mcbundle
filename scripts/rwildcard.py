from glob import glob
import sys

if __name__ == "__main__":
    print(" ".join(glob(sys.argv[1], recursive=True)))
