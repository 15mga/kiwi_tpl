@echo off

set DIR=%~dp0

set BINDIR=%DIR%\..\bin
set CMDDIR=%DIR%\..\cmd

set FILE=%1
if %FILE%=="" (
  set FILE=game\game
)

set OS=%2
if %OS%=="" (
  set OS=linux
)

set ARCH=%3
if %ARCH%=="" (
  if %OS%=="darwin" (
    set ARCH=arm64
  ) else (
    set ARCH=amd64
  )
)

set VER=%4
if %VER%=="" (
  set VER=0.0.1
)

set OUTPUT=%FILE%_%OS%_%NAME%_%VER%
if %OS%=="windows" (
  set OUTPUT=%OUTPUT%.exe
)

echo build %SVC% %CMDDIR%\%FILE%.go -o %BINDIR%\%OUTPUT%

set GOOS=%OS%
set GOARCH=%ARCH%
go build -ldflags "-X main.version=%VER%" -o %BINDIR%\%OUTPUT% %CMDDIR%\%FILE%.go

echo finished

pause