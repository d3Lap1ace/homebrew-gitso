package downloader

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	gitssh "github.com/go-git/go-git/v5/plumbing/transport/ssh"

	"gitso-cli/internal/common"
)

const defaultBase = "~/Documents/GitHub.com"

type Task struct {
	URL    string
	Branch string // Default branch
}

type repoMeta struct {
	Host  string
	Owner string
	Repo  string
	URL   string
}

var (
	sshLike  = regexp.MustCompile(`^([\w-]+)@([\w.\-]+):([\w.\-]+)/([\w.\-]+)(\.git)?$`)
	httpLike = regexp.MustCompile(`^https?://([\w.\-]+)/([\w.\-]+)/([\w.\-]+)(\.git)?$`)
)

// Clone the specified branch
func Clone(url, branch, destBase string) (string, error) {
	meta, err := parseURL(url)
	if err != nil {
		return "", err
	}
	if destBase == "" {
		destBase = defaultBase
	}
	destBase = common.ExpandHome(destBase)
	ownerDir := filepath.Join(destBase, meta.Owner)
	if err := os.MkdirAll(ownerDir, 0o755); err != nil {
		return "", err
	}
	repoPath := filepath.Join(ownerDir, meta.Repo)

	if _, err := os.Stat(filepath.Join(repoPath, ".git")); err == nil {
		return repoPath, gitPull(repoPath, branch)
	}

	return repoPath, cloneRepo(meta, branch, repoPath)
}

func parseURL(u string) (repoMeta, error) {
	switch {
	case sshLike.MatchString(u):
		m := sshLike.FindStringSubmatch(u)
		repo := strings.TrimSuffix(m[4], ".git")
		return repoMeta{Host: m[2], Owner: m[3], Repo: repo, URL: u}, nil
	case httpLike.MatchString(u):
		m := httpLike.FindStringSubmatch(u)
		repo := strings.TrimSuffix(m[4], ".git")
		return repoMeta{Host: m[1], Owner: m[2], Repo: repo, URL: u}, nil
	default:
		return repoMeta{}, fmt.Errorf("unsupported git url: %s", u)
	}
}

func cloneRepo(meta repoMeta, branch, dest string) error {

	if path, err := exec.LookPath("git"); err == nil {
		args := []string{"clone", "--depth", "1"}
		if branch != "" {
			args = append(args, "-b", branch)
		}
		args = append(args, meta.URL, dest)

		cmd := exec.Command(path, args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	}
	// --- fallback: go-git ----------------------------------
	auth, _ := detectAuth(meta.URL)

	opts := &git.CloneOptions{URL: meta.URL, Auth: auth, Depth: 1}
	if branch != "" {
		opts.ReferenceName = plumbing.NewBranchReferenceName(branch)
		opts.SingleBranch = true
	}
	_, err := git.PlainClone(dest, false, opts)
	return err
}

func gitPull(repoPath string, branch string) error {
	path, err := exec.LookPath("git")
	if err != nil {
		return nil
	}
	args := []string{"-C", repoPath, "pull", "--ff-only"}
	if branch != "" {
		args = append(args, "origin", branch)
	}
	cmd := exec.Command(path, args...)
	cmd.Stdout = io.Discard
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// detectAuth Only for go-git fallback
func detectAuth(repoURL string) (transport.AuthMethod, error) {
	switch {
	case strings.HasPrefix(repoURL, "https://"):
		if token := os.Getenv("GITHUB_TOKEN"); token != "" {
			return &http.BasicAuth{Username: "token", Password: token}, nil
		}
	case strings.HasPrefix(repoURL, "ssh://") || sshLike.MatchString(repoURL):
		key := filepath.Join(os.Getenv("HOME"), ".ssh", "id_rsa")
		if _, err := os.Stat(key); err != nil {
			return nil, fmt.Errorf("ssh key not found: %s", key)
		}
		return gitssh.NewPublicKeysFromFile("git", key, "")
	}
	return nil, nil
}
