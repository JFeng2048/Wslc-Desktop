<#
.SYNOPSIS
    打包便携版 ZIP。

.DESCRIPTION
    将指定架构的可执行文件复制到临时暂存目录并统一命名为 wslc-desktop.exe，
    再压缩为 <app>-<version>-<arch>.zip。这样压缩包内部文件名保持一致，
    架构与版本信息体现在压缩包文件名上。

.EXAMPLE
    powershell -File make-zip.ps1 -ExePath ..\..\bin\wslc-desktop-amd64.exe -ZipPath ..\..\bin\wslc-desktop-0.1.0-amd64.zip -ExeName wslc-desktop.exe
#>
param(
    [Parameter(Mandatory = $true)][string]$ExePath,
    [Parameter(Mandatory = $true)][string]$ZipPath,
    [Parameter(Mandatory = $true)][string]$ExeName
)

$ErrorActionPreference = 'Stop'

if (-not (Test-Path $ExePath)) {
    throw "Executable not found: $ExePath"
}

$zipDir = Split-Path -Parent $ZipPath
if ($zipDir -and -not (Test-Path $zipDir)) {
    New-Item -ItemType Directory -Force -Path $zipDir | Out-Null
}

$stage = Join-Path ([IO.Path]::GetTempPath()) ("wslc-zip-" + [Guid]::NewGuid().ToString('N'))
try {
    New-Item -ItemType Directory -Force -Path $stage | Out-Null
    Copy-Item $ExePath (Join-Path $stage $ExeName)

    if (Test-Path $ZipPath) {
        Remove-Item $ZipPath -Force
    }
    Compress-Archive -Path (Join-Path $stage '*') -DestinationPath $ZipPath -CompressionLevel Optimal
}
finally {
    if (Test-Path $stage) {
        Remove-Item -Recurse -Force $stage
    }
}

$info = Get-Item $ZipPath
Write-Host ("Created {0} ({1:N0} bytes, contains {2})" -f $info.FullName, $info.Length, $ExeName)
