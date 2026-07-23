# gitso

English | [中文](./README.zh-CN.md)

`gitso` is a transparent proxy for the system Git. It forwards every command and argument unchanged, except that `clone` without an explicit destination stores GitHub repositories under:

```text
~/Code/Github.com/<owner>/<repo>
```

## Usage

Use normal Git syntax, replacing `git` with `gitso`:

```bash
gitso clone https://github.com/owner/repo.git
gitso clone -b dev --depth 1 git@github.com:owner/repo.git
gitso status --short
gitso pull --ff-only
```

The first command runs the equivalent of:

```bash
git clone https://github.com/owner/repo.git ~/Code/Github.com/owner/repo
```

If a clone destination is supplied, or the remote is not hosted on GitHub, gitso does not modify the command:

```bash
gitso clone https://github.com/owner/repo.git ./custom-dir
gitso clone https://gitlab.com/owner/repo.git
```

Set `GITSO_DEST` to change the organization root:

```bash
GITSO_DEST=~/Projects gitso clone https://github.com/owner/repo.git
```

Options should follow Git's documented clone syntax and appear before the repository URL.

## Design

- Requires the system `git` executable.
- Git handles authentication, credentials, SSH, proxies, LFS, submodules and all command behavior.
- On Unix, gitso replaces itself with Git, preserving terminal interaction, signals and exit codes.
- On Windows, standard input, output, error and the Git exit code are forwarded.

## Installation

GitHub Releases is the only distribution channel. The installer detects your OS and CPU architecture, downloads the matching prebuilt binary, and verifies its SHA-256 checksum before installing it.

### macOS / Linux

```bash
curl -fsSL https://raw.githubusercontent.com/d3Lap1ace/gitso/master/install.sh | sh
```

The default install location is `/usr/local/bin`. Override it when needed:

```bash
curl -fsSL https://raw.githubusercontent.com/d3Lap1ace/gitso/master/install.sh | GITSO_INSTALL_DIR="$HOME/.local/bin" sh
```

### Windows PowerShell

```powershell
irm https://raw.githubusercontent.com/d3Lap1ace/gitso/master/install.ps1 | iex
```

Windows installs to `%LOCALAPPDATA%\Programs\gitso` and adds it to the user `PATH`. All binaries and `checksums.txt` are published on [GitHub Releases](https://github.com/d3Lap1ace/gitso/releases/latest).
