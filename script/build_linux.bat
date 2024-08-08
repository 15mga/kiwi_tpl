@echo off

REM 获取当前脚本所在目录并切换到该目录
cd %~dp0
if %errorlevel% neq 0 exit /b %errorlevel%

REM 设置版本号
set VER=0.0.3

REM 运行 build.bat 脚本并传递参数
call build.bat game\game linux amd64 %VER%

pause