$ErrorActionPreference = "Stop"

$Repo = "d3Lap1ace/gitso"
$Bin = "gitso.exe"
$BaseUrl = "https://github.com/$Repo/releases/latest/download"
$InstallDir = if ($env:GITSO_INSTALL_DIR) {
    $env:GITSO_INSTALL_DIR
} else {
    Join-Path $env:LOCALAPPDATA "Programs\gitso"
}

$Architecture = [System.Runtime.InteropServices.RuntimeInformation]::OSArchitecture.ToString()
switch ($Architecture) {
    "X64"   { $Arch = "amd64" }
    "Arm64" { $Arch = "arm64" }
    default  { throw "gitso: unsupported architecture: $Architecture" }
}

$Archive = "gitso_windows_$Arch.zip"
$TempDir = Join-Path ([System.IO.Path]::GetTempPath()) ("gitso-" + [guid]::NewGuid())

try {
    New-Item -ItemType Directory -Path $TempDir | Out-Null
    $ArchivePath = Join-Path $TempDir $Archive
    $ChecksumsPath = Join-Path $TempDir "checksums.txt"

    Write-Host "Downloading $Archive..."
    Invoke-WebRequest -UseBasicParsing "$BaseUrl/$Archive" -OutFile $ArchivePath
    Invoke-WebRequest -UseBasicParsing "$BaseUrl/checksums.txt" -OutFile $ChecksumsPath

    $ChecksumLine = Get-Content $ChecksumsPath | Where-Object {
        $Fields = $_ -split "\s+", 2
        $Fields.Count -eq 2 -and $Fields[1].TrimStart("*") -eq $Archive
    } | Select-Object -First 1
    if (-not $ChecksumLine) {
        throw "gitso: $Archive is missing from checksums.txt"
    }
    $Expected = ($ChecksumLine -split "\s+")[0].ToLowerInvariant()
    $Actual = (Get-FileHash -Algorithm SHA256 $ArchivePath).Hash.ToLowerInvariant()
    if ($Actual -ne $Expected) {
        throw "gitso: checksum verification failed"
    }
    Write-Host "Checksum verified."

    Expand-Archive -Path $ArchivePath -DestinationPath $TempDir -Force
    $BinaryPath = Join-Path $TempDir $Bin
    if (-not (Test-Path $BinaryPath -PathType Leaf)) {
        throw "gitso: archive does not contain $Bin"
    }

    New-Item -ItemType Directory -Force -Path $InstallDir | Out-Null
    Copy-Item -Force $BinaryPath (Join-Path $InstallDir $Bin)

    $UserPath = [Environment]::GetEnvironmentVariable("Path", "User")
    $PathEntries = @($UserPath -split ";" | Where-Object { $_ })
    if ($PathEntries -notcontains $InstallDir) {
        $NewUserPath = (@($PathEntries) + $InstallDir) -join ";"
        [Environment]::SetEnvironmentVariable("Path", $NewUserPath, "User")
        $env:Path = "$env:Path;$InstallDir"
        Write-Host "Added $InstallDir to your user PATH. Open a new terminal to use gitso."
    }

    Write-Host "Installed: $(Join-Path $InstallDir $Bin)"
} finally {
    if (Test-Path $TempDir) {
        Remove-Item -Recurse -Force $TempDir
    }
}
