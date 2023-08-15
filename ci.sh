#!/bin/bash

npx go-semantic-release \
  --dry \
  --no-ci \
  --provider git \
  --allow-initial-development-versions \
  --changelog tmp.md

if [ -e CHANGELOG.md ]; then
    cat CHANGELOG.md >> tmp.md
    mv -f tmp.md CHANGELOG.md
else
    mv tmp.md CHANGELOG.md
fi
