from builtins import tuple

COMPATIBILITY_MATRIX: dict[str, list[str]] = {
    "win32": ["x64", "ia32", "arm64"],
    "linux": ["x64", "ia32", "arm64"],
    "darwin": ["x64", "arm64"]
}

def target_double_tuples() -> list[tuple[str, str]]:
    doubles: list[tuple[str, str]] = []

    for os in COMPATIBILITY_MATRIX:
        for cpu in COMPATIBILITY_MATRIX[os]:
            doubles.append((os, cpu))

    return doubles

def target_doubles() -> list[str]:
    double_tuples = target_double_tuples()
    target_doubles: list[str] = []

    for double_tuple in double_tuples:
        (os, cpu) = double_tuple
        target_doubles.append(f"{os}-{cpu}")

    target_doubles.sort()
    return target_doubles
