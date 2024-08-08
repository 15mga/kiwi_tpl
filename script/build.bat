@echo off

REM 获取当前脚本所在目录
set DIR=%~dp0

REM 设置目录变量
set BINDIR=%DIR%\..\bin
set CMDDIR=%DIR%\..\cmd

REM 获取输入参数
set FILE=%1
if "%FILE%"=="" (
  set FILE=game\game
)

set OS=%2
if "%OS%"=="" (
  set OS=linux
)

set ARCH=%3
if "%ARCH%"=="" (
  if "%OS%"=="darwin" (
    set ARCH=arm64
  ) else (
    set ARCH=amd64
  )
)

set VER=%4
if "%VER%"=="" (
  set VER=0.0.1
)

REM 设置输出文件名
set OUTPUT=%FILE%_%OS%_%NAME%_%VER%
if "%OS%"=="windows" (
  set OUTPUT=%OUTPUT%.exe
)

REM 打印构建信息
echo build %SVC% %CMDDIR%\%FILE%.go -o %BINDIR%\%OUTPUT%

REM 执行构建命令
set GOOS=%OS%
set GOARCH=%ARCH%
go build -ldflags "-X main.version=%VER%" -o %BINDIR%\%OUTPUT% %CMDDIR%\%FILE%.go

REM 打印完成信息
echo finished

pause