import os.path as path
import compat
import shutil
import sys
import os

def rm_unused_builds():
    builds_dir = path.join("npm", "@bpbuild")

    builds = filter(lambda b: path.isdir(path.join(builds_dir, b)), os.listdir(builds_dir))

    used_builds = set(compat.target_doubles())

    for build in builds:
        if build not in used_builds:
            build_dir = path.join(builds_dir, build)
            shutil.rmtree(build_dir, True)

if __name__ == "__main__":
    op_or_fd = sys.argv[1]

    match op_or_fd:
        case "--unused-builds":
            rm_unused_builds()
        case _:
            shutil.rmtree(op_or_fd, True)
