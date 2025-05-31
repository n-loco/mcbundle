import os.path as path
import compat
import json
import sys
import os

def get_version() -> str:
    with open("assets/program_version.txt", "r") as program_version:
        content = program_version.read()
    return content.strip()

def target(target_double: str):
    target_as_list = target_double.split("-")

    node_os = target_as_list[0]
    node_cpu = target_as_list[1]

    package_dir = path.join("npm", "@mcbundle", target_double)
    package_path = path.join(package_dir, "package.json")

    debug_package_dir = path.join(package_dir, "debug")
    debug_package_path = path.join(debug_package_dir, "package.json")

    package_json_obj = {
        "name": f"@mcbundle/{target_double}",
        "description": f"{target_double} binary for mcbundle",
        "version": get_version(),
        "files": ["mcbundle.exe" if node_os == "win32" else "mcbundle"],
        "os": [node_os],
        "cpu": [node_cpu]
    }

    try:
        os.mkdir(package_dir)
    except:
        pass

    try:
        os.mkdir(debug_package_dir)
    except:
        pass

    package_json_obj["version"] = get_version()
    content = json.dumps(package_json_obj, indent="  ")

    with open(package_path, "wb") as package_json:
        package_json.write(bytes(content, "UTF-8"))

    with open(debug_package_path, "wb") as package_json:
        package_json.write(bytes(content, "UTF-8"))
    

def main_package():
    package_path = path.join("npm", "mcbundle", "package.json")
    with open(package_path, "r") as package_json:
        package_json_obj = json.load(package_json)
    with open(package_path, "wb") as package_json:
        package_json_obj["version"] = get_version()

        target_doubles = compat.target_doubles()

        optional_dependencies: dict = {}

        for target_double in target_doubles:
            optional_dependencies[f"@mcbundle/{target_double}"] = f"workspace:../@mcbundle/{target_double}"
        
        package_json_obj["optionalDependencies"] = optional_dependencies

        content = json.dumps(package_json_obj, indent="  ")
        package_json.write(bytes(content, "UTF-8"))

def create_package():
    package_path = path.join("npm", "create", "package.json")
    with open(package_path, "r") as package_json:
        package_json_obj = json.load(package_json)
    with open(package_path, "wb") as package_json:
        package_json_obj["version"] = get_version()

        content = json.dumps(package_json_obj, indent="  ")
        package_json.write(bytes(content, "UTF-8"))

if __name__ == "__main__":
    op = sys.argv[1]
    match op:
        case "--main-package":
            main_package()
        case "--create-package":
            create_package()
        case "--target":
            target(sys.argv[2])
        case _:
            raise Exception(f"unknown op: {op}")
    
