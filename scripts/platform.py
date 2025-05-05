import sys
import compat

def node_os_to_goos(node_os: str) -> str:
    match node_os:
        case "win32":
            return "windows"
        case "linux" | "darwin":
            return node_os

    raise Exception(f"unknown os: {node_os}")

def node_cpu_to_goarch(node_cpu: str) -> str:
    match node_cpu:
        case "x64":
            return "amd64"
        case "ia32":
            return "386"
        case "arm" | "arm64":
            return node_cpu

    raise Exception(f"unknown cpu: {node_cpu}")

def platform_wildcard() -> str:
    return " ".join(compat.target_doubles())

def exe_suffix(target_double: str) -> str | None:
    return ".exe" if target_double.startswith("win32-") else None

if __name__ == "__main__":
    op = sys.argv[1]

    match op:
        case "--node-os-to-goos":
            try:
                print(node_os_to_goos(sys.argv[2]))
            except:
                pass
        case "--node-cpu-to-goarch":
            try:
                print(node_cpu_to_goarch(sys.argv[2]))
            except:
                pass
        case "--platform-wildcard":
            print(platform_wildcard())
        case "--exe-suffix":
            suffix = exe_suffix(sys.argv[2])
            if suffix != None:
                print(suffix)
        case _:
            raise Exception(f"unknown op: {op}")
