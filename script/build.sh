#!/bin/bash

DIR=`dirname $0`

BINDIR=$DIR/../bin
CMDDIR=$DIR/../cmd

# 引用 ver.sh 文件以获取当前版本号
source ./ver.sh

if [[ $VER == "" ]]; then
  VER="0.0.0"
fi

# 分割版本号为主版本号、次版本号和补丁版本号
IFS='.' read -r -a version_parts <<< "$VER"

# 递增补丁版本号
version_parts[2]=$((version_parts[2] + 1))

# 组合新的版本号
new_version="${version_parts[0]}.${version_parts[1]}.${version_parts[2]}"

# 保存新的版本号回 ver.sh 文件
echo "VER=$new_version" > ./ver.sh

# 将新的版本号导出为环境变量
export VER=$new_version

# 打印版本号
echo "Version is $VER"

FILE=$1
if [[ $FILE == "" ]]; then
  FILE="game/game"
fi

OS=$2
if [[ $OS == "" ]]; then
  OS="linux"
fi

ARCH=$3
if [[ $ARCH == "" ]]; then
  if [[ $OS == "darwin" ]]; then
    ARCH="arm64"
  el
    ARCH="amd64"
  fi
fi

export VER_NAME="${version_parts[0]}_${version_parts[1]}_${version_parts[2]}"

OUTPUT=$FILE"_"$OS"_"$NAME"_"$VER_NAME
if [[ $OS == "windows" ]]; then
  OUTPUT=$OUTPUT.exe
fi

echo build $SVC $CMDDIR/$FILE.go -o $BINDIR/$OUTPUT
GOOS=$OS  GOARCH=$ARCH go build -ldflags "-X main.version=$VER" -o $BINDIR/$OUTPUT $CMDDIR/$FILE.go

echo finished