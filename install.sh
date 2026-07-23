#!/usr/bin/env sh
set -e

REPO="d3Lap1ace/gitso"
BIN="gitso"
INSTALL_DIR="/usr/local/bin"

# ── detect OS ────────────────────────────────────────────────────────────────
OS="$(uname -s)"
case "$OS" in
  Darwin) OS_LABEL="macOS" ;;
  Linux)  OS_LABEL="linux" ;;
  *)
    echo "Unsupported OS: $OS" >&2
    exit 1
    ;;
esac

# ── detect arch ──────────────────────────────────────────────────────────────
ARCH="$(uname -m)"
case "$ARCH" in
  x86_64)          ARCH_LABEL="amd64" ;;
  arm64 | aarch64) ARCH_LABEL="arm64" ;;
  *)
    echo "Unsupported architecture: $ARCH" >&2
    exit 1
    ;;
esac

# ── resolve latest version ───────────────────────────────────────────────────
VERSION="$(curl -fsSL "https://api.github.com/repos/${REPO}/releases/latest" \
  | grep '"tag_name"' \
  | sed 's/.*"tag_name": *"\(.*\)".*/\1/')"

if [ -z "$VERSION" ]; then
  echo "Failed to fetch latest release version." >&2
  exit 1
fi

ARCHIVE="${BIN}_${OS_LABEL}_${ARCH_LABEL}.tar.gz"
URL="https://github.com/${REPO}/releases/download/${VERSION}/${ARCHIVE}"

# ── download & install ───────────────────────────────────────────────────────
TMP="$(mktemp -d)"
trap 'rm -rf "$TMP"' EXIT

echo "Downloading gitso ${VERSION} (${OS_LABEL}/${ARCH_LABEL})..."
curl -fsSL "$URL" -o "${TMP}/${ARCHIVE}"
tar -xzf "${TMP}/${ARCHIVE}" -C "$TMP"

# ── move to PATH ─────────────────────────────────────────────────────────────
if [ ! -w "$INSTALL_DIR" ]; then
  echo "Installing to ${INSTALL_DIR} (requires sudo)..."
  sudo mv "${TMP}/${BIN}" "${INSTALL_DIR}/${BIN}"
  sudo chmod +x "${INSTALL_DIR}/${BIN}"
else
  mv "${TMP}/${BIN}" "${INSTALL_DIR}/${BIN}"
  chmod +x "${INSTALL_DIR}/${BIN}"
fi

# ── remove macOS quarantine ───────────────────────────────────────────────────
if [ "$OS_LABEL" = "macOS" ]; then
  xattr -d com.apple.quarantine "${INSTALL_DIR}/${BIN}" 2>/dev/null || true
fi

echo "Installed: ${INSTALL_DIR}/${BIN}"
