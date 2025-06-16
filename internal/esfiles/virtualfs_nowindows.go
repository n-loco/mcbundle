//go:build !windows

package esfiles

func normalizeZipPath(path string) string {
	return path
}
