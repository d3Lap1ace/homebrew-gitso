# gitso

> **Clone → Organize → Done**  
> `gitso` is a zero-friction CLI that clones any GitHub repository (HTTPS **or** SSH) and files it under a predictable tree  
> `~/Documents/GitHub.com/<owner>/<repo>`.  
> If the host doesn’t have a system-wide `git`, it silently falls back to an embedded **go-git** engine, so it also works inside scratch containers and FaaS.

---

## Project motivation

| Pain point | gitso’s answer |
|------------|----------------|
| Repos scattered across random folders | Opinionated layout keeps everything under one root |
| Minimal containers / CI images lack `git` | Falls back to pure-Go implementation when `git` isn’t present |
| Keeping repos up-to-date | Detects existing repo → runs `git pull --ff-only` |
| Tedious `cd && git pull` before work | One command, no manual navigation |
| Re-entering the same options | Global `~/.gitso.yaml` or `GITSO_*` envs provide defaults |

---

## Installation

| Platform | Quick start |
|----------|-------------|
| **macOS / Linux** | **Binary** (current)\*: download latest `gitso_<os>_<arch>.tar.gz` →<br>`tar -xf … && sudo mv gitso /usr/local/bin` |
| **Windows** | **Binary**: download `gitso_<ver>_windows_amd64.zip` → unzip and put `gitso.exe` in `%PATH%` |
| *(coming soon)* | **Homebrew** `brew install gitso` &nbsp;•&nbsp; **Scoop** `scoop install gitso` |

\* Pre-compiled binaries are linked on every GitHub Release.


## Global config

Modify `gitso/.gitso.yaml` (respects `GITSO_*` env too):

```yaml
# Global defaults
dest: "~/Documents/GitHub.com"  # Base directory
branch: "main"                  # Default branch