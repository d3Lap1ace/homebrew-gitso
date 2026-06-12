<div align="right">
  <a href="#english"><img src="https://img.shields.io/badge/English-blue?style=flat-square" alt="English" /></a>
  &nbsp;
  <a href="#中文"><img src="https://img.shields.io/badge/中文-gray?style=flat-square" alt="中文" /></a>
</div>

<a id="english"></a>

# gitso

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
curl -fsSL https://raw.githubusercontent.com/d3Lap1ace/gitso/main/install.sh | sh
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

---

<br /><br />

<div align="right">
  <a href="#english"><img src="https://img.shields.io/badge/English-gray?style=flat-square" alt="English" /></a>
  &nbsp;
  <a href="#中文"><img src="https://img.shields.io/badge/中文-blue?style=flat-square" alt="中文" /></a>
</div>

<a id="中文"></a>

# gitso

> `gitso` 是一个 CLI 工具，可克隆任意 GitHub 仓库（HTTPS **或** SSH），并按作者名将其存放到可预测的目录树中：  
> `~/Code/GitHub.com/<owner>/<repo>`。  
> 如果宿主环境没有系统级 `git`，它会自动回退到内嵌的 **go-git** 引擎，因此即使在 scratch 容器和 FaaS 中也能正常工作。

---

## 使用示例

```bash
gitso git@github.com:owner/repo.git
gitso https://github.com/owner/repo.git

# 克隆到自定义目录
gitso -d ~/Projects git@github.com:owner/repo.git

# 克隆指定分支
gitso -b dev https://github.com/owner/repo.git
```

## 效果图
![img.png](img/render_ch.png)

## 项目动机

| 痛点 | gitso 的解决方案 |
|------|-----------------|
| 仓库散落在随机文件夹 | 统一的根目录结构将所有仓库集中管理 |
| 最小化容器 / CI 镜像缺少 `git` | 当系统缺少 `git` 时，自动切换到纯 Go 实现 |
| 保持仓库最新 | 检测已存在仓库 → 执行 `git pull --ff-only` |
| 每次工作前需手动 `cd && git pull` | 一条命令，无需手动导航 |

---

## 安装

**Linux / macOS — 一行命令**
```bash
curl -fsSL https://raw.githubusercontent.com/d3Lap1ace/gitso/main/install.sh | sh
```

**macOS — Homebrew**
```bash
brew tap d3Lap1ace/gitso
brew install gitso
```

**Linux — 手动安装**  
从 [Releases](https://github.com/d3Lap1ace/gitso/releases/latest) 页面下载对应架构的压缩包：
```bash
# amd64
tar -zxvf gitso_linux_amd64.tar.gz && sudo mv gitso /usr/local/bin/

# arm64
tar -zxvf gitso_linux_arm64.tar.gz && sudo mv gitso /usr/local/bin/
```

**Windows — amd64**  
从 [Releases](https://github.com/d3Lap1ace/gitso/releases/latest) 页面下载 `gitso_windows_amd64.zip`，将 `gitso.exe` 解压至 `%PATH%` 中的某个目录（例如 `C:\Tools\`），然后运行 `gitso --help`。

> **提示**  
> • 检测 CPU 架构：`uname -m`（Linux/macOS）/ `wmic os get osarchitecture`（Windows）  
> • 通过一行命令或 Homebrew 安装时，quarantine 标记会自动移除。  
> • 在 macOS 上手动安装时，需运行：`sudo xattr -d com.apple.quarantine /usr/local/bin/gitso`

---

## 默认目录

按以下优先级依次解析：

| 优先级 | 来源 |
|--------|------|
| 1 | `-d` / `--dest` flag |
| 2 | `GITSO_DEST` 环境变量 |
| 3 | `~/.gitso_config`（由 `gitso config --dest` 写入） |
| 4 | `~/Code/Github.com`（内置默认值） |

```bash
gitso config                      # 查看当前默认目录
gitso config --dest ~/Projects    # 永久修改默认目录
```
