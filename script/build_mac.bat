@echo off

cd %~dp0

set VER=0.0.3

call build.bat game\game darwin arm64 %VER%

pause