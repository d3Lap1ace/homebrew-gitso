# gitso

[English](./README.md) | 中文

`gitso` 是系统 Git 的透明代理。除了一种情况外，所有命令和参数都会原样传给 Git：执行没有指定目标目录的 `clone` 时，GitHub 仓库会自动存入：

```text
~/Code/Github.com/<owner>/<repo>
```

## 效果图

![按 GitHub owner 分类存储仓库](./img/render_ch.png)

## 使用方法

按照正常 Git 语法，把 `git` 换成 `gitso`：

```bash
gitso clone https://github.com/owner/repo.git
gitso clone -b dev --depth 1 git@github.com:owner/repo.git
gitso status --short
gitso pull --ff-only
```

第一条命令等价于：

```bash
git clone https://github.com/owner/repo.git ~/Code/Github.com/owner/repo
```

如果明确提供了目标目录，或者仓库不在 GitHub，gitso 不会修改命令：

```bash
gitso clone https://github.com/owner/repo.git ./custom-dir
gitso clone https://gitlab.com/owner/repo.git
```

可以通过 `GITSO_DEST` 修改仓库根目录：

```bash
GITSO_DEST=~/Projects gitso clone https://github.com/owner/repo.git
```

参数应遵循 Git 文档中的 clone 语法，放在仓库 URL 之前。

## 实现原则

- 依赖系统中的 `git` 命令。
- 认证、凭据、SSH、代理、LFS、子模块及所有 Git 行为都由 Git 自己处理。
- Unix 下 gitso 会用 Git 替换自身，完整保留终端交互、信号和退出码。
- Windows 下会转发标准输入、输出、错误和 Git 退出码。

## 安装

GitHub Releases 是唯一的发布渠道。安装脚本会自动识别操作系统与 CPU 架构，下载对应的预编译二进制，并在安装前校验 SHA-256。

### macOS / Linux

```bash
curl -fsSL https://raw.githubusercontent.com/d3Lap1ace/gitso/master/install.sh | sh
```

默认安装到 `/usr/local/bin`。如需修改安装目录：

```bash
curl -fsSL https://raw.githubusercontent.com/d3Lap1ace/gitso/master/install.sh | GITSO_INSTALL_DIR="$HOME/.local/bin" sh
```

### Windows PowerShell

```powershell
irm https://raw.githubusercontent.com/d3Lap1ace/gitso/master/install.ps1 | iex
```

Windows 默认安装到 `%LOCALAPPDATA%\Programs\gitso`，并自动加入用户 `PATH`。所有二进制和 `checksums.txt` 都发布在 [GitHub Releases](https://github.com/d3Lap1ace/gitso/releases/latest)。
