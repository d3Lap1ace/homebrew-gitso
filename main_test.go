package main

import (
	"path/filepath"
	"reflect"
	"testing"
)

func TestRewriteClone(t *testing.T) {
	base := t.TempDir()
	t.Setenv("GITSO_DEST", base)
	args := []string{"clone", "--depth", "1", "-b", "dev", "git@github.com:owner/repo.git"}

	got, err := rewriteClone(args)
	if err != nil {
		t.Fatal(err)
	}
	want := append(args, filepath.Join(base, "owner", "repo"))
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("rewriteClone() = %q, want %q", got, want)
	}
}

func TestRewriteCloneLeavesOtherCommandsAndExplicitDestinationsAlone(t *testing.T) {
	tests := [][]string{
		{"status", "--short"},
		{"clone", "https://gitlab.com/owner/repo.git"},
		{"clone", "https://github.com/owner/repo.git", "./repo"},
	}
	for _, args := range tests {
		got, err := rewriteClone(args)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(got, args) {
			t.Errorf("rewriteClone(%q) = %q", args, got)
		}
	}
}

func TestGitHubRepo(t *testing.T) {
	tests := map[string]bool{
		"https://github.com/owner/repo.git":      true,
		"git@github.com:owner/repo.git":          true,
		"ssh://git@github.com/owner/repo.git":    true,
		"https://github.com/owner/repo":          true,
		"https://gitlab.com/owner/repo.git":      false,
		"https://github.com/owner/repo/tree/dev": false,
		"not-a-url":                              false,
	}
	for raw, want := range tests {
		_, _, got := githubRepo(raw)
		if got != want {
			t.Errorf("githubRepo(%q) ok = %v, want %v", raw, got, want)
		}
	}
}
