#!/bin/bash

userProfilePath="$HOME"
newPath="$userProfilePath/.yock/mnt"

if [[ ":$PATH:" != *":$newPath:"* ]]; then
  export PATH="$newPath:$PATH"
fi

echo "Install success!"
