@echo off

setlocal enabledelayedexpansion

REM 获取当前脚本所在目录
set DIR=%~dp0

set NAME=server REM 项目go.mod中指定的模组名
set GOOGLEPBIDR=%GOPATH%\pb REM google/protobuf的根目录，包含any.proto,api.proto等
set PBDIR=%DIR%\..\proto\msg REM 项目protobuf文件根目录
set KIWIDIR=..\..\..\.. REM kiwi.proto存放根目录
set OUTDIR=%DIR%\.. REM 输出目录根目录
set PBOUTDIR=%OUTDIR%\proto REM *.pb.go输出根目录
set KIWIOUTDIR=%OUTDIR%\internal REM 生成文件根目录

echo %KIWIDIR%

echo complie kiwi

protoc --proto_path=%GOOGLEPBIDR% --proto_path=%KIWIDIR% --proto_path=%PBDIR% --go_out=%PBOUTDIR% --kiwi_out=-m=%NAME%,-r=guest_player,-c=cs,-db=mgo:%KIWIOUTDIR% %PBDIR%\model\*.proto %PBDIR%\fail\*.proto %PBDIR%\service\*.proto %PBDIR%\client\*.proto

echo kiwi finished

echo mgo bson
protoc-mgo-bson -d=%PBOUTDIR%\pb
echo mgo bson finished

endlocal

pause