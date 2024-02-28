#!/bin/bash

go build -o jobin
sudo mv jobin /usr/local/bin/
jobin completion bash > jobin-bash-completion.sh
sudo mv jobin-bash-completion.sh /etc/bash_completion.d/