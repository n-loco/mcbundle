from glob import glob
import os.path as path
import sys

if __name__ == "__main__":
    root_dir: str = None
    if len(sys.argv) >= 3:
        root_dir = sys.argv[2]
        files = glob(sys.argv[1], recursive=True, root_dir=root_dir)
        print(" ".join(map(lambda f : path.normpath(f"{root_dir}/{f}"), files)))
    else:
        print(" ".join(glob(sys.argv[1], recursive=True)))
