package downloader

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"gitso/internal/common"

	gogit "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport"
	gogithttp "github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
)

const defaultBase = "~/Code/Github.com"

type repoMeta struct {
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
		return repoMeta{Owner: m[3], Repo: repo, URL: u}, nil
	case httpLike.MatchString(u):
		m := httpLike.FindStringSubmatch(u)
		repo := strings.TrimSuffix(m[3], ".git")
		return repoMeta{Owner: m[2], Repo: repo, URL: u}, nil
	default:
		return repoMeta{}, fmt.Errorf("unsupported git url: %s", u)
	}
}

func cloneRepo(meta repoMeta, branch, dest string) error {
	path, err := exec.LookPath("git")
	if err == nil {
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

	// Fallback: use embedded go-git when system git is unavailable.
	auth, err := detectAuth(meta)
	if err != nil {
		return fmt.Errorf("go-git auth: %w", err)
	}
	opts := &gogit.CloneOptions{
		URL:   meta.URL,
		Auth:  auth,
		Depth: 1,
	}
	if branch != "" {
		opts.ReferenceName = plumbing.NewBranchReferenceName(branch)
		opts.SingleBranch = true
	}
	_, err = gogit.PlainClone(dest, false, opts)
	return err
}

func gitPull(repoPath string, branch string) error {
	path, err := exec.LookPath("git")
	if err == nil {
		args := []string{"-C", repoPath, "pull", "--ff-only"}
		if branch != "" {
			args = append(args, "origin", branch)
		}
		cmd := exec.Command(path, args...)
		cmd.Stdout = io.Discard
		cmd.Stderr = os.Stderr
		return cmd.Run()
	}

	// Fallback: use embedded go-git.
	repo, err := gogit.PlainOpen(repoPath)
	if err != nil {
		return err
	}
	wt, err := repo.Worktree()
	if err != nil {
		return err
	}
	opts := &gogit.PullOptions{RemoteName: "origin", Auth: nil}
	if branch != "" {
		opts.ReferenceName = plumbing.NewBranchReferenceName(branch)
	}
	err = wt.Pull(opts)
	if err == gogit.NoErrAlreadyUpToDate {
		return nil
	}
	return err
}

func detectAuth(meta repoMeta) (transport.AuthMethod, error) {
	if strings.HasPrefix(meta.URL, "https") {
		if token := os.Getenv("GITHUB_TOKEN"); token != "" {
			return &gogithttp.BasicAuth{Username: "x-access-token", Password: token}, nil
		}
		return nil, nil
	}
	// SSH
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	keyPath := filepath.Join(home, ".ssh", "id_rsa")
	return ssh.NewPublicKeysFromFile("git", keyPath, "")
}
