package common

import (
	"os"
	"path/filepath"
	"strings"
)

// ExpandHome Convert ~/ to an absolute path
func ExpandHome(p string) string {
	if strings.HasPrefix(p, "~/") {
		if home, err := os.UserHomeDir(); err == nil {
			return filepath.Join(home, strings.TrimPrefix(p, "~/"))
		}
	}
	return p
}
