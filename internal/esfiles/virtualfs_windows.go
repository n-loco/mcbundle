//go:build windows

package esfiles

import "strings"

func normalizeZipPath(path string) string {
	return strings.ReplaceAll(path, "\\", "/")
}
