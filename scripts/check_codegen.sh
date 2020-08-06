#!/usr/bin/env bash

cd "$(dirname "$0")/.." || exit 1

rake codegen:all

status=$(git status --porcelain=v1)

if [[ -z "$status" ]]
then
  echo "nothing changed after running codegen:all"
else
  echo "git status --porcelain returned non-empty output, this means you should re-run rake codegen:all and check in the results"
  echo "$status"
fi
