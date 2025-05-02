import os.path as path
import compat
import json
import sys
import os

def get_version() -> str:
    with open("assets/program_version.txt", "r") as program_version:
        content = program_version.read()
    return content.strip()

def executable(node_os: str, node_cpu: str):
    target_double = f"{node_os}-{node_cpu}"
    package_dir = path.join("npm", "@bpbuild", target_double)
    package_path = path.join(package_dir, "package.json")

    package_json_obj = {
        "name": f"@bpbuild/{target_double}",
        "description": f"{target_double} binary for bpbuild",
        "version": get_version(),
        "files": ["bpbuild.exe" if node_os == "win32" else "bpbuild"],
        "os": [node_os],
        "cpu": [node_cpu]
    }

    try:
        os.mkdir(package_dir)
    except:
        pass

    with open(package_path, "w") as package_json:
        package_json_obj["version"] = get_version()
        json.dump(package_json_obj, package_json, indent="  ")
    

def main_package():
    package_path = path.join("npm", "bpbuild", "package.json")
    with open(package_path, "r") as package_json:
        package_json_obj = json.load(package_json)
    with open(package_path, "w") as package_json:
        package_json_obj["version"] = get_version()

        target_doubles = compat.target_doubles()

        optional_dependencies: dict = {}

        for target_double in target_doubles:
            optional_dependencies[f"@bpbuild/{target_double}"] = f"workspace:../@bpbuild/{target_double}"
        
        package_json_obj["optionalDependencies"] = optional_dependencies

        json.dump(package_json_obj, package_json, indent="  ")

def create_package():
    package_path = path.join("npm", "create", "package.json")
    with open(package_path, "r") as package_json:
        package_json_obj = json.load(package_json)
    with open(package_path, "w") as package_json:
        package_json_obj["version"] = get_version()

        json.dump(package_json_obj, package_json, indent="  ")

if __name__ == "__main__":
    op = sys.argv[1]
    match op:
        case "--main-package":
            main_package()
        case "--create-package":
            create_package()
        case "--executable":
            executable(sys.argv[2], sys.argv[3])
        case _:
            raise Exception(f"unknown op: {op}")
    
