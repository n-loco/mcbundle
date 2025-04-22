from builtins import tuple

COMPATIBILITY_MATRIX: dict[str, list[str]] = {
    "win32": ["x64", "ia32", "arm64"],
    "linux": ["x64", "ia32", "arm64"],
    "darwin": ["x64", "arm64"]
}

def target_double_tuples() -> list[tuple[str, str]]:
    pairs: list[tuple[str, str]] = []

    for os in COMPATIBILITY_MATRIX:
        for cpu in COMPATIBILITY_MATRIX[os]:
            pairs.append((os, cpu))

    return pairs

def target_doubles() -> list[str]:
    pairs = target_double_tuples()
    pairs_names: list[str] = []

    for pair in pairs:
        (os, cpu) = pair
        pairs_names.append(f"{os}-{cpu}")

    return pairs_names
