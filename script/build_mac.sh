#!/bin/bash

cd "$(dirname "$0")" || exit

./build.sh game/game darwin arm64
