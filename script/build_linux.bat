@echo off

cd %~dp0

set VER=0.0.3

call build.bat game\game linux amd64 %VER%

pause