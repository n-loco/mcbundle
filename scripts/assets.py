from collections.abc import Callable
from builtins import tuple
import os.path as path
import os
import sys

SOURCE_ASSETS = "assets"
OUT_ASSETS = path.join("internal", "assets")

def file_name_parts(file_name: str) -> tuple[str, str]:
    i = file_name.find(".")
    if i < 0:
        return (file_name, None)
    
    file_base_name = file_name[:i]
    file_ext = file_name[i:]

    return (file_base_name, file_ext)

def snake_to_pascal(s: str) -> str:
    return "".join(map(lambda w : w.capitalize(), s.split("_")))

def txt_importer(file_name: str) -> str:
    file_path = path.join(SOURCE_ASSETS, file_name)

    with open(file_path, "r") as asset:
        content = asset.read().strip()
    
    return f"`{content}`"


def binary_importer(file_name: str) -> str:
    file_path = path.join(SOURCE_ASSETS, file_name)

    with open(file_path, "rb") as asset:
        content = asset.read()

    byte_list: list[str] = []

    for byte in content:
        byte_list.append(f"0x{byte:02x}")

    array_lines: list[str] = []
    byte_list_si = (0, 16)
    while byte_list_si[0] < len(byte_list):
        byte_list_s = byte_list[byte_list_si[0]:byte_list_si[1]]
        array_line = "\t" + ", ".join(byte_list_s)
        array_lines.append(array_line)

        byte_list_si = (byte_list_si[1], byte_list_si[1] + 16)

    array_content = ",\n".join(array_lines)

    array_def = "[" + str(len(content)) + "]byte{\n" + array_content + ",\n}"

    return array_def

IMPORTER_MAP: dict[str, Callable[[str], str]] = {
    ".txt": txt_importer,
    "_default": binary_importer
}

def gen_go_source(file_name: str) -> None:
    (file_base_name, file_ext) = file_name_parts(file_name)

    importer = IMPORTER_MAP[file_ext] if file_ext in IMPORTER_MAP else IMPORTER_MAP["_default"]

    var_name = snake_to_pascal(file_base_name)
    var_content = importer(file_name)
    var_decl = "var {} = {}".format(var_name, var_content)

    go_source_content = f"// generated code\npackage assets\n\n{var_decl}\n"

    go_source_file_path = path.join(OUT_ASSETS, file_name + ".go")

    with open(go_source_file_path, "wb") as go_source_file:
        go_source_file.write(bytes(go_source_content, "UTF-8"))


if __name__ == "__main__":
    try:
        os.mkdir(OUT_ASSETS)
    except:
        pass

    gen_go_source(sys.argv[1])
