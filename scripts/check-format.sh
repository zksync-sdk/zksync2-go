#!/bin/bash

fmt_output=$(gofmt -l ../.)

if [ -n "$fmt_output" ]; then
  echo "Code not formatted correctly:"
  echo "$fmt_output"
  exit 1
else
  echo "Code is formatted correctly"
  exit 0
fi