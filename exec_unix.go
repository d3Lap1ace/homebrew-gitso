//go:build !windows

package main

import (
	"os"
	"syscall"
)

func executeGit(path string, args []string) error {
	return syscall.Exec(path, append([]string{"git"}, args...), os.Environ())
}
