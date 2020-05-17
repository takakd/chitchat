#!/usr/bin/env bash

if [[ $# -le 0 ]]; then
    echo "usage: sh commit comment"
    exit 0
fi

git add --all
git commit -m "${1}"
