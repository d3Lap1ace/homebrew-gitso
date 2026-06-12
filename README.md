# gitso

English | [中文](./README.zh-CN.md)

> `gitso` is a CLI tool that clones any GitHub repository (HTTPS **or** SSH) and files it under a predictable tree by owner name:  
> `~/Code/GitHub.com/<owner>/<repo>`.  
> If the host doesn't have a system-wide `git`, it silently falls back to an embedded **go-git** engine, so it even works inside scratch containers and FaaS.

---

## Usage

```bash
gitso git@github.com:owner/repo.git
gitso https://github.com/owner/repo.git

# clone into a custom directory
gitso -d ~/Projects git@github.com:owner/repo.git

# clone a specific branch
gitso -b dev https://github.com/owner/repo.git
```

## Rendering
![img.png](img/render_en.png)

## Motivation

| Pain point | gitso's answer |
|------------|----------------|
| Repos scattered across random folders | Opinionated layout keeps everything under one root |
| Minimal containers / CI images lack `git` | Falls back to pure-Go implementation when `git` isn't present |
| Keeping repos up-to-date | Detects existing repo → runs `git pull --ff-only` |
| Tedious `cd && git pull` before work | One command, no manual navigation |

---

## Installation

**Linux / macOS — one-liner**
```bash
curl -fsSL https://raw.githubusercontent.com/d3Lap1ace/gitso/master/install.sh | sh
```

**macOS — Homebrew**
```bash
brew tap d3Lap1ace/gitso
brew install gitso
```

**Linux — manual**  
Download the archive for your architecture from the [Releases](https://github.com/d3Lap1ace/gitso/releases/latest) page:
```bash
# amd64
tar -zxvf gitso_linux_amd64.tar.gz && sudo mv gitso /usr/local/bin/

# arm64
tar -zxvf gitso_linux_arm64.tar.gz && sudo mv gitso /usr/local/bin/
```

**Windows — amd64**  
Download `gitso_windows_amd64.zip` from the [Releases](https://github.com/d3Lap1ace/gitso/releases/latest) page, unzip `gitso.exe` into a directory on `%PATH%` (e.g. `C:\Tools\`), then run `gitso --help`.

> **Tips**  
> • Detect CPU arch: `uname -m` (Linux/macOS) / `wmic os get osarchitecture` (Windows)  
> • The one-liner and Homebrew remove the macOS quarantine flag automatically.  
> • For manual installs on macOS, run: `sudo xattr -d com.apple.quarantine /usr/local/bin/gitso`

---

## Default destination

Resolved in this order:

| Priority | Source |
|----------|--------|
| 1 | `-d` / `--dest` flag |
| 2 | `GITSO_DEST` environment variable |
| 3 | `~/.gitso_config` (set by `gitso config --dest`) |
| 4 | `~/Code/Github.com` (built-in default) |

```bash
gitso config                      # view current default
gitso config --dest ~/Projects    # persist a new default
```
