#!/usr/bin/env bash

set -euo pipefail

mkdir -p ~/Documents/service-declarative-definition-workshop

git clone \
          -b service-declarative-definition ssh://git@stash.msk.avito.ru:7999/swat/workshops.git \
          ~/Documents/service-declarative-definition-workshop

cd ~/Documents/service-declarative-definition-workshop/2019-04-02

make deps format
