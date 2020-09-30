#!/bin/bash

echo "Testing go module algos"

for d in $(go list ./...); do
  go test "$d"
done
