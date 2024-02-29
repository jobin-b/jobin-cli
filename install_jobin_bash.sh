#!/bin/bash

# Install jobin-cli
# Run this script again if you make updates to the repo/config in this directory

# Install jobin binary
go build -o jobin
sudo mv jobin /usr/local/bin/

# Config file
cp -a .jobin/ ~/

# Install bash completion
jobin completion bash > jobin-bash-completion.sh
sudo mv jobin-bash-completion.sh /etc/bash_completion.d/