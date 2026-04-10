# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What this project is

`sugit` is a Go CLI tool that clones GitHub repositories (HTTPS or SSH) into a predictable directory tree: `~/Documents/GitHub.com/<owner>/<repo>`. It prefers the system `git` binary and falls back to the embedded `go-git` engine when `git` is not on `$PATH`. If the target repo already exists locally, it runs `git pull --ff-only` instead of cloning.

## Commands

```bash
# Build
go build -o sugit .

# Run locally
go run . <repo-url>
go run . -d ~/Code -b dev git@github.com:owner/repo.git

# Tidy dependencies
go mod tidy

# Run tests
go test ./...
```

## Architecture

The entry point is `main.go` → `cmd.Execute()` (Cobra root command).

**`cmd/`** — Cobra command definitions
- `root.go`: Root command. Accepts one positional arg (repo URL) and flags `-d`/`--dest` and `-b`/`--branch`. Resolves the destination via: CLI flag → `SUGIT_DEST` env var → hardcoded `~/Documents/Github.com`. Then calls `downloader.Clone`.
- `config.go`: `sugit config --dest <path>` subcommand. Writes `~/.sugit_dest` to persist the default destination (note: this file is not currently read back by `loadDefaultDest()`).

**`internal/downloader/downloader.go`** — Core logic
- `Clone()`: Parses the URL, builds the local path as `<destBase>/<owner>/<repo>`, detects existing repos (pulls) vs new ones (clones).
- `cloneRepo()`: Prefers system `git clone --depth 1`; falls back to `go-git` (`gogit.PlainClone`) when `git` is not on `$PATH`.
- `gitPull()`: Runs `git -C <path> pull --ff-only`; falls back to `go-git` (`wt.Pull`) when `git` is absent.
- `detectAuth()`: Used only by the go-git fallback. HTTPS uses `GITHUB_TOKEN` env var; SSH reads `~/.ssh/id_rsa`.
- URL parsing: Two regexes — `sshLike` for `user@host:owner/repo.git` and `httpLike` for `https://host/owner/repo.git`.

**`internal/common/pathutil.go`** — `ExpandHome()` expands `~/` to the absolute home directory path.

## Release process

Releases are triggered by pushing a `v*.*.*` tag. GitHub Actions runs GoReleaser (`.goreleaser.yaml`), which builds cross-platform binaries (darwin/linux/windows, amd64/arm64) with `CGO_ENABLED=0` and publishes them as GitHub Release assets. The Homebrew formula lives in `Formula/sugit.rb`.
