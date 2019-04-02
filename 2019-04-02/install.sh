#!/usr/bin/env bash

set -euo pipefail

mkdir -p ~/Documents/workshops

git clone git@github.com:kamilsk/workshops.git ~/Documents/workshops

cd ~/Documents/workshops/2019-04-02

make deps format
