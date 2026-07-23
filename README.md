# gitso

English | [中文](./README.zh-CN.md)

`gitso` is a transparent proxy for the system Git. It forwards every command and argument unchanged, except that `clone` without an explicit destination stores GitHub repositories under:

```text
~/Code/Github.com/<owner>/<repo>
```

## Preview

![GitHub repositories organized by owner](./img/render_en.png)

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

gitso is not distributed through package managers or GitHub Releases. Build it from source:

```bash
git clone https://github.com/d3Lap1ace/gitso.git
cd gitso
go build -o gitso .
```

Install the resulting binary somewhere on your `PATH`:

```bash
sudo install -m 0755 gitso /usr/local/bin/gitso
```

On Windows PowerShell:

```powershell
git clone https://github.com/d3Lap1ace/gitso.git
Set-Location gitso
go build -o gitso.exe .
.\gitso.exe --version
```

Move `gitso.exe` to a directory already on your `PATH` to use it globally.
