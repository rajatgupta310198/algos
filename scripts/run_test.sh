#!/bin/bash

echo "Testing go module algos"

time for d in $(go list ./...); do
  go test "$d"
done
