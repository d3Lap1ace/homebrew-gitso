#!/usr/bin/env sh
set -eu

REPO="d3Lap1ace/gitso"
BIN="gitso"
INSTALL_DIR="${GITSO_INSTALL_DIR:-/usr/local/bin}"
BASE_URL="https://github.com/${REPO}/releases/latest/download"

case "$(uname -s)" in
  Darwin) OS="macOS" ;;
  Linux)  OS="linux" ;;
  *) echo "gitso: unsupported operating system: $(uname -s)" >&2; exit 1 ;;
esac

case "$(uname -m)" in
  x86_64 | amd64)  ARCH="amd64" ;;
  arm64 | aarch64) ARCH="arm64" ;;
  *) echo "gitso: unsupported architecture: $(uname -m)" >&2; exit 1 ;;
esac

ARCHIVE="${BIN}_${OS}_${ARCH}.tar.gz"
TMP="$(mktemp -d)"
trap 'rm -rf "$TMP"' EXIT HUP INT TERM

echo "Downloading ${ARCHIVE}..."
curl -fsSL "${BASE_URL}/${ARCHIVE}" -o "${TMP}/${ARCHIVE}"
curl -fsSL "${BASE_URL}/checksums.txt" -o "${TMP}/checksums.txt"

EXPECTED="$(awk -v file="$ARCHIVE" '$2 == file { print $1; exit }' "${TMP}/checksums.txt")"
if [ -z "$EXPECTED" ]; then
  echo "gitso: ${ARCHIVE} is missing from checksums.txt" >&2
  exit 1
fi

if command -v sha256sum >/dev/null 2>&1; then
  ACTUAL="$(sha256sum "${TMP}/${ARCHIVE}" | awk '{ print $1 }')"
elif command -v shasum >/dev/null 2>&1; then
  ACTUAL="$(shasum -a 256 "${TMP}/${ARCHIVE}" | awk '{ print $1 }')"
else
  echo "gitso: sha256sum or shasum is required" >&2
  exit 1
fi

if [ "$ACTUAL" != "$EXPECTED" ]; then
  echo "gitso: checksum verification failed" >&2
  exit 1
fi
echo "Checksum verified."

tar -xzf "${TMP}/${ARCHIVE}" -C "$TMP"
if [ ! -f "${TMP}/${BIN}" ]; then
  echo "gitso: archive does not contain ${BIN}" >&2
  exit 1
fi

if mkdir -p "$INSTALL_DIR" 2>/dev/null && [ -w "$INSTALL_DIR" ]; then
  install -m 0755 "${TMP}/${BIN}" "${INSTALL_DIR}/${BIN}"
else
  echo "Installing to ${INSTALL_DIR} (sudo may ask for your password)..."
  sudo mkdir -p "$INSTALL_DIR"
  sudo install -m 0755 "${TMP}/${BIN}" "${INSTALL_DIR}/${BIN}"
fi

echo "Installed: ${INSTALL_DIR}/${BIN}"
